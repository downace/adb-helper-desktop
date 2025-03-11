package main

import (
	"github.com/downace/adb-helper-desktop/internal/adb"
	"net/netip"
)

var allAdbDeviceStates = []struct {
	Value  adb.DeviceState
	TSName string
}{
	{adb.StateOffline, "Offline"},
	{adb.StateDevice, "Device"},
	{adb.StateNoDevice, "NoDevice"},
	{adb.StateRecovery, "Recovery"},
	{adb.StateSideload, "Sideload"},
	{adb.StateUnauthorized, "Unauthorized"},
	{adb.StateNoPermissions, "NoPermissions"},
	{adb.StateHost, "Host"},
}

// Methods just for bindings generation

func (a *App) Noop(_ *netip.Addr) {}
