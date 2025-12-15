import { ApiResponse } from "src/types/api";
import { NetworkInterface } from "src/types/net_interface";
import { axiosWithToken } from "./api";
import { handleError } from "./error";

export const getNetworkInterfaces = async (): Promise<NetworkInterface[] | Error> => {
  try {
    const response = await axiosWithToken.get<ApiResponse<NetworkInterface[]>>("/api/v1/interfaces");
    return response.data.data;
  } catch (error) {
    return handleError(error)
  }
};
