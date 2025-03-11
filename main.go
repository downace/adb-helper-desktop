package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/linux"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var appIcon []byte

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:         "ADB Helper",
		Width:         400,
		Height:        640,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:     app.startup,
		OnShutdown:    app.shutdown,
		OnBeforeClose: app.beforeClose,
		//HideWindowOnClose: true,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "ed2ab3cd-d268-4c09-9df4-a514eea631d8",
			OnSecondInstanceLaunch: app.secondInstanceLaunch,
		},
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			allAdbDeviceStates,
		},
		Linux: &linux.Options{
			Icon: appIcon,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
