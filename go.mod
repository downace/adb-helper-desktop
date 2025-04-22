module github.com/downace/adb-helper-desktop

go 1.23.0

toolchain go1.24.0

require (
	// Commit 6539193 fixes https://github.com/fyne-io/systray/issues/62
	fyne.io/systray v1.11.1-0.20240701075416-6539193433ff
	github.com/betamos/zeroconf v0.1.8-0.20250208023331-d559d61612b7
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d
	github.com/wailsapp/wails/v2 v2.10.1
	golang.org/x/text v0.23.0
	gopkg.in/yaml.v3 v3.0.1
)

retract [v0.0.0, v0.1.0] // Contains invalid module name

require (
	github.com/bep/debounce v1.2.1 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jchv/go-winloader v0.0.0-20210711035445-715c2860da7e // indirect
	github.com/labstack/echo/v4 v4.13.3 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leaanthony/go-ansi-parser v1.6.1 // indirect
	github.com/leaanthony/gosod v1.0.4 // indirect
	github.com/leaanthony/slicer v1.6.0 // indirect
	github.com/leaanthony/u v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.63 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/samber/lo v1.49.1 // indirect
	github.com/tkrajina/go-reflector v0.5.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/wailsapp/go-webview2 v1.0.19 // indirect
	github.com/wailsapp/mimetype v1.4.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
)

// replace github.com/wailsapp/wails/v2 v2.10 => /home/jbot/go/pkg/mod
