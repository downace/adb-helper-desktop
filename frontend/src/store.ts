import {
  ConnectDevice,
  DisconnectDevice,
  GetConnectedDevices,
  GetServices,
  PairDevice,
  RestartServicesSearch,
  WindowIsHidden,
} from "@/go/main/App";
import type { adb } from "@/go/models";
import type { zeroconf } from "@/go/models";
import { serviceAddrs, serviceId } from "@/helpers/mdns";
import { EventsOn } from "@/runtime";
import { useIdle, useIntervalFn } from "@vueuse/core";
import { defineStore } from "pinia";
import {
  onBeforeMount,
  onBeforeUnmount,
  shallowReactive,
  shallowReadonly,
  shallowRef,
  watch,
} from "vue";

export const useAppStore = defineStore("app", () => {
  const listenersCancelFns = shallowRef<(() => void)[]>([]);
  const services = shallowReactive(new Map<string, zeroconf.Service>());
  const connectedDevices = shallowReactive(new Map<string, adb.Device>());

  onBeforeMount(() => {
    listenersCancelFns.value = [
      EventsOn("service-added", (service) => {
        services.set(serviceId(service), service);
      }),
      EventsOn("service-updated", (service) => {
        services.set(serviceId(service), service);
      }),
      EventsOn("service-removed", (service) => {
        services.delete(serviceId(service));
      }),

      EventsOn("adb-device", () => {
        void refreshConnectedDevices();
      }),

      EventsOn("window-hidden", (hidden) => {
        windowHidden.value = hidden;
      }),
    ];
  });

  onBeforeUnmount(() => {
    for (const cancel of listenersCancelFns.value) {
      cancel();
    }
  });

  onBeforeMount(() => {
    void loadServices();
    void refreshConnectedDevices();
  });

  const windowHidden = shallowRef(false);

  onBeforeMount(async () => {
    windowHidden.value = await WindowIsHidden();
  });

  function pairDevice(host: string, codeOrPassword: string) {
    return PairDevice(host, codeOrPassword);
  }

  function getConnectedDevice(service: zeroconf.Service) {
    return (
      serviceAddrs(service)
        .map((h) => connectedDevices.get(`${h}:${service.port}`))
        .find((a) => a) ?? null
    );
  }

  async function connectDevice(host: string) {
    await ConnectDevice(host);
  }

  async function disconnectDevice(host: string) {
    await DisconnectDevice(host);
  }

  async function loadServices() {
    const newServices = await GetServices();
    services.clear();
    for (const val of newServices) {
      services.set(serviceId(val), val);
    }
  }

  async function refreshAll() {
    return Promise.all([refreshServices(), refreshConnectedDevices()]);
  }

  async function refreshServices() {
    await RestartServicesSearch();
  }

  async function refreshConnectedDevices() {
    const newDevices = await GetConnectedDevices();
    connectedDevices.clear();
    for (const val of newDevices) {
      connectedDevices.set(val.id, val);
    }
  }

  const { idle } = useIdle();

  watch(idle, (isIdle, wasIdle) => {
    if (wasIdle || !isIdle) {
      void refreshAll();
    }
  });

  useIntervalFn(
    () => {
      if (!windowHidden.value) {
        void refreshAll();
      }
    },
    () => 60_000,
  );

  return {
    services: shallowReadonly(services),
    connectedDevices: shallowReadonly(connectedDevices),
    pairDevice,
    connectDevice,
    disconnectDevice,
    getConnectedDevice,
  };
});
