export namespace adb {
	
	export enum DeviceState {
	    Offline = "offline",
	    Device = "device",
	    NoDevice = "no device",
	    Recovery = "recovery",
	    Sideload = "sideload",
	    Unauthorized = "unauthorized",
	    NoPermissions = "no permissions",
	    Host = "host",
	}
	export interface DeviceDescription {
	    product: string;
	    model: string;
	    device: string;
	    transport_id: string;
	}
	export interface Device {
	    id: string;
	    state: DeviceState;
	    description: DeviceDescription;
	    brand: string;
	    model: string;
	}

}

export namespace netip {
	
	export interface Addr {
	
	}

}

export namespace zeroconf {
	
	export interface Type {
	    name: string;
	    domain: string;
	    subtypes?: string[];
	}
	export interface Service {
	    type: Type;
	    name: string;
	    port: number;
	    hostname: string;
	    addrs: netip.Addr[];
	    text: string[];
	}

}

