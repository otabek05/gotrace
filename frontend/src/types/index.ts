export type ApiResponse<T> = {
  success: boolean;
  message: string;
  data?: T;
};

export type WSMessageType = "PACKET" | "SCAN_RESULT" | "STATUS" | "ERROR";

export type WSMessage = {
  type: WSMessageType;
  data: any;
};

export type NetworkInterface = {
  name: string;
  mtu: number;
  mac: string;
  addresses: string[];
  is_up: boolean;
};

export type CapturedPacket = {
  time: string; // ISO string from backend
  interface: string;
  src_ip: string;
  dst_ip: string;
  src_port?: string;
  dst_port?: string;
  network?: string;
  transport?: string;
  application?: string;
  length?: number;
  payload?: string;
};