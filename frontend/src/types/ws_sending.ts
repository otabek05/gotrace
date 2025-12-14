export enum MessageType {
  StartCapturing = "start_capturing",
  StopCapturing = "stop_capturing",
  Reset = "reset",
  Restart = "restart",
  StartNmap = "start_nmap",
  StopNmap = "stop_nmap",
}

export enum TrafficOptions {
  Incoming = "incoming",
  Outgoing = "outgoing",
  Both = "both",
}

export enum NetworkLayer {
  IPv4 = "ipv4",
  IPv6 = "ipv6",
  ICMP = "icmp",
  Unknown = "unknown",
}

export enum TransportLayer {
  TCP = "tcp",
  UDP = "udp",
  Unknown = "unknown",
}

export enum  ApplicationLayer {
  Any = "any",
  WellKnown = "well-known",
  Custom = "custom"
}

export interface WSOutgoingMessage {
  type: MessageType;
  message?: any;

  trafficOptions?: TrafficOptions;
  networkLayer?: NetworkLayer;
  transport?: TransportLayer;
  services:string[];

  isOutgoing: true;
}

export interface WSIncomingMessage {
  type: string;
  message: any;

  trafficOptions: string;
  networkLayer: string;
  transport: string;
  isOutgoing?: false;
}
