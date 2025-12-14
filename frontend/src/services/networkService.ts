import { ApiResponse } from "src/types/api";
import { NetworkInterface } from "src/types/net_interface";
import { axiosWithToken } from "./api";

export const fetchNetworkInterfaces = async (): Promise<NetworkInterface[]> => {
  try {
    const response = await axiosWithToken.get<ApiResponse<NetworkInterface[]>>("/api/v1/network/interfaces");
    return response.data.data;
  } catch (error) {
    console.error("Failed to fetch network interfaces:", error);
    throw error;
  }
};
