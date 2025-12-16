import { ApiResponse } from "../types/api";
import type { NetworkInterface } from "../types/netInterface";
import { axiosWithToken } from "./api";
import { handleError } from "./error";


export const getNetworkInterfaces = async(): Promise<ApiResponse<NetworkInterface[] | null>> =>{
    try {
        const response = await axiosWithToken.get("/api/v1/network/interfaces")
        console.log(response.data)
        return response.data
    }catch(e) {
        return  ApiResponse.error(handleError(e))
    }
}