import type { adb } from "@/go/models";
import type { zeroconf } from "@/go/models";

export type Device = {
  service: zeroconf.Service | null;
  adbDevice: adb.Device | null;
};
