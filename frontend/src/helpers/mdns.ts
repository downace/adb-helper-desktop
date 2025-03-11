import type { zeroconf } from "@/go/models";
import { isIPv4, isIPv6 } from "is-ip";
import { ascend } from "ramda";

export function serviceAddrs(service: zeroconf.Service) {
  return [service.hostname, ...(service.addrs ?? [])] as string[];
}

export function preferredServiceHost(service: zeroconf.Service) {
  const hostname = serviceAddrs(service).sort(
    // Prefer IPv4
    ascend((addr) => (isIPv4(addr) ? 0 : isIPv6(addr) ? 2 : 1)),
  )[0];
  return `${hostname}:${service.port}`;
}

// Conforms to github.com/betamos/zeroconf/service.go:117 `func (s *Service) String() string`
export function serviceId(service: zeroconf.Service) {
  return `${service.name}.${service.type.name}.${service.type.domain}`;
}
