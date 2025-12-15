export interface WSReceiving {
  type: string;
  packets?: ParsedPacket;
  internetSpeed?: InternetSpeed
}

export  interface InternetSpeed {
  bytesIn:string 
  bytesOut:string
}

export interface ParsedPacket {
  ethernet?: Ethernet;
  ipv4?: IPv4;
  ipv6?: IPv6;
  tcp?: TCP;
  udp?: UDP;
  icmp?: ICMP;
  arp?: ARP;
  dns?: DNS;
  http?: HTTP;
  https?: HTTPS;
  dhcp?: DHCP;
  application?: AppLayer;
  direction: string 
  timestamp:string
}

export interface Ethernet {
  src_mac: string;
  dst_mac: string;
  ethernet_type: string;
  length: number;
}

export interface IPv4 {
  version: number;
  ihl: number;
  tos: number;
  total_length: number;
  identification: number;
  flags: string;
  fragment_offset: number;
  ttl: number;
  protocol: string;
  checksum: number;
  src_ip: string;
  dst_ip: string;
}

export interface TCP {
  src_port: string;
  dst_port: string;
  seq: number;
  ack: number;
  data_offset: number;
  flags: string[];
  window: number;
  checksum: number;
  urgent: number;
  options: string[];
}

export interface UDP {
  src_port: number;
  dst_port: number;
  length: number;
  checksum: number;
}

export interface DNSQuestion {
  name: string;
  type: string;
  class: string;
}

export interface DNSAnswer {
  name: string;
  type: string;
  class: string;
  ttl: number;
  data: string;
}

export interface DNS {
  id: number;
  qr: string;
  opcode: string;
  aa: boolean;
  tc: boolean;
  rd: boolean;
  ra: boolean;
  z: number;
  response_code: string;
  questions: DNSQuestion[];
  answers: DNSAnswer[];
}

export interface IPv6 {
  traffic_class: number;
  flow_label: number;
  payload_length: number;
  next_header: number;
  hop_limit: number;
  src_ip: string;
  dst_ip: string;
}

export interface ICMP {
  type: number;
  code: number;
  checksum: number;
}

export interface ARP {
  operation: string;
  sender_mac: string;
  sender_ip: string;
  target_mac: string;
  target_ip: string;
}

export interface HTTP {
  method: string;
  url: string;
  version: string;
  status_code: number;
  status: string;
  headers: Record<string, string>;
  body: string;
}

export interface HTTPS {
  tls_version: string;
  cipher: string;
  server_name: string;
}

export interface DHCP {
  operation: string;
  transaction_id: string;
  client_ip: string;
  your_ip: string;
  server_ip: string;
  gateway_ip: string;
  hostname?: string;
  lease_time?: number;
}

export interface AppLayer {
  protocol: string;
  data: any;
}
