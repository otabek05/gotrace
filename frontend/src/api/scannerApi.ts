import api from "./axios";

export interface NetworkInterface {
  name: string;
  mtu: number;
  hardwareAddr: string;
  addresses: string[];
  isUp: boolean;
}

export const getInterfaces = async () => {
  const res = await api.get("/api/interfaces");
  return res.data.data as NetworkInterface[];
};

export const startScan = async (payload: {
  interface: string;
  snaplen?: number;
  promisc?: boolean;
  filter?: string;
}) => {
  return api.post("/api/scan/start", payload);
};

export const stopScan = async () => {
  return api.post("/api/scan/stop");
};

export const getStatus = async () => {
  const res = await api.get("/api/status");
  return res.data.data;
};
