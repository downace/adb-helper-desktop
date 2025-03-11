package main

import (
	"adb-helper-desktop/internal/adb"
	"adb-helper-desktop/internal/common"
	"context"
	_ "embed"
	"fmt"
	"fyne.io/systray"
	"github.com/betamos/zeroconf"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

//go:embed build/trayicon.png
var trayIcon []byte

var knownDevicesCsvUrl = "https://storage.googleapis.com/play_public/supported_devices.csv"

type App struct {
	ctx                   context.Context
	config                *AppConfig
	adbClient             *adb.Client
	ticker                *time.Ticker
	zeroconfClient        *zeroconf.Client
	stopConnectSearch     context.CancelFunc
	stopPairingSearch     context.CancelFunc
	services              map[string]*zeroconf.Service
	hidden                bool
	trayStart             func()
	trayEnd               func()
	shouldShutdownOnClose bool
}

func NewApp() *App {
	app := App{
		services:              make(map[string]*zeroconf.Service),
		config:                makeConfig(filepath.Join(common.Cwd(), "config.yaml")),
		hidden:                false,
		shouldShutdownOnClose: false,
	}

	app.trayStart, app.trayEnd = systray.RunWithExternalLoop(app.initTray(), nil)

	return &app
}

func (a *App) initTray() func() {
	return func() {
		systray.SetIcon(trayIcon)
		systray.SetTitle("ADB Helper")
		mToggle := systray.AddMenuItem("Toggle Window", "Toggle Window")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quit")

		go func() {
			for range mToggle.ClickedCh {
				if a.hidden {
					runtime.WindowShow(a.ctx)
				} else {
					runtime.WindowHide(a.ctx)
				}
				a.hidden = !a.hidden
			}
		}()

		go func() {
			for range mQuit.ClickedCh {
				a.shouldShutdownOnClose = true
				runtime.Quit(a.ctx)
			}
		}()
	}
}

var connectType = zeroconf.NewType("_adb-tls-connect._tcp")
var pairingType = zeroconf.NewType("_adb-tls-pairing._tcp")

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	sigintCh := make(chan os.Signal, 1)
	signal.Notify(sigintCh, os.Interrupt)
	go func() {
		for range sigintCh {
			a.shouldShutdownOnClose = true
		}
	}()

	a.trayStart()

	err := a.config.load()

	if err != nil {
		_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.WarningDialog,
			Title:   "Warning",
			Message: fmt.Sprintf("Unable to load settings, using defaults\n%s", err),
		})
	}

	err = a.initAdbClient()

	if err != nil {
		panic(err)
	}

	a.zeroconfClient = zeroconf.New()

	a.zeroconfClient.Browse(func(event zeroconf.Event) {
		switch event.Op {
		case zeroconf.OpAdded:
			a.services[event.Service.String()] = event.Service
			runtime.EventsEmit(a.ctx, "service-added", event.Service)
		case zeroconf.OpUpdated:
			a.services[event.Service.String()] = event.Service
			runtime.EventsEmit(a.ctx, "service-updated", event.Service)
		case zeroconf.OpRemoved:
			delete(a.services, event.Service.String())
			runtime.EventsEmit(a.ctx, "service-removed", event.Service)
		}
	}, connectType, pairingType)

	a.zeroconfClient.Open()
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if a.shouldShutdownOnClose {
		return false
	}
	runtime.WindowHide(ctx)
	a.hidden = true
	return true
}

func (a *App) shutdown(_ context.Context) {
	a.adbClient.Destroy()
	a.zeroconfClient.Close()
	a.trayEnd()
}

func (a *App) secondInstanceLaunch(_ options.SecondInstanceData) {
	runtime.WindowShow(a.ctx)
	runtime.WindowUnminimise(a.ctx)
}

func (a *App) RestartServicesSearch() {
	a.zeroconfClient.Reload()
}

func (a *App) GetServices() []*zeroconf.Service {
	res := make([]*zeroconf.Service, 0)
	for _, service := range a.services {
		res = append(res, service)
	}
	return res
}

func (a *App) PickAdbPath() (string, error) {
	return runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
}

func (a *App) GetAdbPath() string {
	return a.config.AdbPath
}

func (a *App) SetAdbPath(adbPath string) error {
	if a.config.AdbPath == adbPath {
		return nil
	}
	a.config.AdbPath = adbPath
	err := a.initAdbClient()

	if err != nil {
		return err
	}

	return a.config.save()
}

func (a *App) RefreshKnownDevices() error {
	filePath := filepath.Join(common.Cwd(), "supported_devices.csv")
	return downloadFile(filePath, knownDevicesCsvUrl)
}

func (a *App) initAdbClient() error {
	if a.adbClient != nil {
		a.adbClient.Destroy()
	}
	a.adbClient = adb.NewClient(a.config.AdbPath)
	return a.adbClient.TrackDevices(func(device adb.Device) {
		runtime.EventsEmit(a.ctx, "adb-device", device)
	})
}

func downloadFile(filePath string, fileUrl string) error {
	outFile, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer outFile.Close()

	resp, err := http.Get(fileUrl)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	defer resp.Body.Close()

	_, err = io.Copy(outFile, resp.Body)

	return nil
}
