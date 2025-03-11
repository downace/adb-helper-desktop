package common

import (
	"encoding/csv"
	"fmt"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"iter"
	"os"
	"path/filepath"
	"slices"
)

type KnownDevice struct {
	RetailBranding string `json:"retailBranding"`
	MarketingName  string `json:"marketingName"`
	Device         string `json:"device"`
	Model          string `json:"model"`
}

func Cwd() string {
	executablePath, _ := os.Executable()
	return filepath.Dir(executablePath)
}

func DetectKnownDevice(device string) (KnownDevice, error) {
	filePath := filepath.Join(Cwd(), "supported_devices.csv")

	var retailBrandingCol uint
	var marketingNameCol uint
	var deviceCol uint
	var modelCol uint

	var knownDevice KnownDevice

	it, err := csvFileIterator(filePath)

	if err != nil {
		return knownDevice, err
	}

	for lineNumber, line := range it {
		if lineNumber == 1 {
			if i := slices.Index(line, "Retail Branding"); i >= 0 {
				retailBrandingCol = uint(i)
			}
			if i := slices.Index(line, "Marketing Name"); i >= 0 {
				marketingNameCol = uint(i)
			}
			if i := slices.Index(line, "Device"); i >= 0 {
				deviceCol = uint(i)
			}
			if i := slices.Index(line, "Model"); i >= 0 {
				modelCol = uint(i)
			}
		} else if line[deviceCol] == device {
			knownDevice = KnownDevice{
				RetailBranding: line[retailBrandingCol],
				MarketingName:  line[marketingNameCol],
				Device:         line[deviceCol],
				Model:          line[modelCol],
			}
			break
		}
	}

	return knownDevice, err
}

func csvFileIterator(filePath string) (iter.Seq2[uint, []string], error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	enc, err := detectFileEncoding(file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(transform.NewReader(file, enc.NewDecoder()))
	reader.ReuseRecord = true

	var line []string
	var lineNumber uint = 0

	return func(yield func(uint, []string) bool) {
		defer file.Close()

		for {
			line, err = reader.Read()
			lineNumber++
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}
			if !yield(lineNumber, line) {
				break
			}
		}
	}, nil
}

func detectFileEncoding(file *os.File) (encoding.Encoding, error) {
	bytes := make([]byte, 8)
	_, err := file.Read(bytes)

	if err != nil {
		return nil, err
	}

	det := chardet.NewTextDetector()

	res, err := det.DetectBest(bytes)

	if err != nil {
		return nil, err
	}

	switch res.Charset {
	case "UTF-8":
		return unicode.UTF8, nil
	case "UTF-16LE":
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM), nil
	case "UTF-16BE":
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM), nil
	}

	return nil, fmt.Errorf("unsupported charset: %s", res.Charset)
}
