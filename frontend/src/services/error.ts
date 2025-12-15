import axios from "axios";



export function handleError(e:any): Error {
    if (axios.isAxiosError(e)) {
        return new  Error(e.response?.data?.message)
    }

    return new Error("unknown error")
}