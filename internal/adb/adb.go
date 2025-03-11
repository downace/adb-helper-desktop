package adb

import (
	"fmt"
	"github.com/downace/adb-helper-desktop/internal/common"
	"os"
	"os/exec"
	"strings"
)

type DeviceState string

const (
	// https://developer.android.com/tools/adb#devicestatus

	StateOffline  = "offline"
	StateDevice   = "device"
	StateNoDevice = "no device"

	// https://stackoverflow.com/a/31232709

	StateRecovery      = "recovery"
	StateSideload      = "sideload"
	StateUnauthorized  = "unauthorized"
	StateNoPermissions = "no permissions"
	StateHost          = "host"
)

type DeviceDescription struct {
	Product     string `json:"product"`
	Model       string `json:"model"`
	Device      string `json:"device"`
	TransportId string `json:"transport_id"`
}

type Device struct {
	Id          string            `json:"id"`
	State       DeviceState       `json:"state"`
	Description DeviceDescription `json:"description"`
	Brand       string            `json:"brand"`
	Model       string            `json:"model"`
}

type Client struct {
	trackDevicesProc *os.Process
	path             string
}

func NewClient(path string) *Client {
	return &Client{
		path: path,
	}
}

func (c *Client) Destroy() {
	if c.trackDevicesProc != nil {
		c.trackDevicesProc.Signal(os.Interrupt)
		c.trackDevicesProc.Wait()
	}
}

func (c *Client) exec(command ...string) (string, error) {
	cmd := exec.Command(c.path, command...)
	outputBytes, err := cmd.CombinedOutput()

	output := string(outputBytes)

	if err != nil && output != "" {
		err = fmt.Errorf("%s%s", output, err)
	}

	return output, err
}

func (c *Client) Version() (string, error) {
	return c.exec("version")
}

func (c *Client) Devices() ([]Device, error) {
	output, err := c.exec("devices", "-l")

	if err != nil {
		return nil, err
	}

	result := make([]Device, 0)

	for line := range strings.Lines(output) {
		if strings.Contains(line, "List of devices attached") {
			continue
		}
		device := parseDeviceString(line)

		if device != nil {
			result = append(result, *device)
		}
	}

	return result, nil
}

func (c *Client) Connect(host string) (string, error) {
	output, err := c.exec("connect", host)
	if err == nil && strings.Contains(output, "failed to connect") {
		err = fmt.Errorf(output)
	}
	return output, err
}

func (c *Client) Disconnect(host string) (string, error) {
	return c.exec("disconnect", host)
}

func (c *Client) Pair(host string, codeOrPassword string) (string, error) {
	output, err := c.exec("pair", host, codeOrPassword)
	if err == nil && strings.Contains(output, "Failed") {
		err = fmt.Errorf(output)
	}
	return output, err
}

func (c *Client) TrackDevices(onTrack func(device Device)) error {
	cmd := exec.Command(c.path, "track-devices", "-l")
	r, w, err := os.Pipe()

	if err != nil {
		return err
	}

	cmd.Stdout = w

	go func() {
		sizeBytes := make([]byte, 4)
		var message []byte
		for {
			n, err := r.Read(sizeBytes)

			if err != nil {
				// todo
				return
			}

			if n != 4 {
				// todo err
				return
			}

			var size int
			_, err = fmt.Sscanf(string(sizeBytes), "%04X", &size)

			if err != nil {
				// todo err
				return
			}

			if size == 0 {
				continue
			}

			message = make([]byte, size)

			n, err = r.Read(message)

			if err != nil {
				// todo
				return
			}

			device := parseDeviceString(string(message))
			if device != nil {
				onTrack(*device)
			}
		}
	}()

	err = cmd.Start()

	if err != nil {
		return err
	}

	c.trackDevicesProc = cmd.Process

	return nil
}

func parseDeviceString(line string) *Device {
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return nil
	}
	device := Device{
		Id:    fields[0],
		State: DeviceState(fields[1]),
	}
	for _, str := range fields[2:] {
		pair := strings.SplitN(str, ":", 2)
		switch pair[0] {
		case "product":
			device.Description.Product = pair[1]
		case "model":
			device.Description.Model = pair[1]
		case "device":
			device.Description.Device = pair[1]
			knownDevice, err := common.DetectKnownDevice(pair[1])

			if err == nil {
				device.Brand = knownDevice.RetailBranding
				device.Model = knownDevice.MarketingName
			}
		case "transport_id":
			device.Description.TransportId = pair[1]
		}
	}

	return &device
}
