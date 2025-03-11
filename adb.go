package main

import (
	"github.com/downace/adb-helper-desktop/internal/adb"
)

func (a *App) ConnectDevice(host string) (string, error) {
	return a.adbClient.Connect(host)
}

func (a *App) DisconnectDevice(host string) (string, error) {
	return a.adbClient.Disconnect(host)
}

func (a *App) PairDevice(host string, codeOrPassword string) (string, error) {
	return a.adbClient.Pair(host, codeOrPassword)
}

func (a *App) GetConnectedDevices() ([]adb.Device, error) {
	return a.adbClient.Devices()
}

func (a *App) AdbVersion() (string, error) {
	return a.adbClient.Version()
}
