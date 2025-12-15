import axios from "axios";
import Cookies from "js-cookie";


//const BASE_URL = import.meta.env.VITE_API_URL;

const BASE_URL = "http://localhost:8080";

console.log("BASE_URL: ", BASE_URL)

export const axiosWithoutToken = axios.create({
    baseURL: BASE_URL
})

export const axiosWithToken = axios.create({
    baseURL: BASE_URL
})

let isRefreshing = false;
let refreshSubscribers: ((token:string)=> void) [] =  [];

function onRefreshed(accessToken:string) {
    refreshSubscribers.map((callback) => callback(accessToken));
}

function addRefreshSubscriber(callback:(token:string)=> void) {
    refreshSubscribers.push(callback);
}

axiosWithToken?.interceptors.request.use(
    (config) => {
        const jwt = Cookies.get("access_token")
        if (jwt) {
            config.headers["Authorization"] = `Bearer ${jwt}`
            config.headers["Content-Type"] = "application/json"
        }
        return config
    },
    (error) =>  Promise.reject(error)
    
)

// Add an interceptor to handle Refresh Token request

axiosWithToken?.interceptors.response.use(
    (response) =>  response,
    async (error) => {
        const originalRequest = error.config;
        console.log(error)
        // Check if the error status is due to token expiration
        // refreshToken
        if (error.response && error.response.status === 403 && !originalRequest._retry) {
            if (!isRefreshing) {
                isRefreshing = true
                originalRequest._retry = true
                try {
                    const refresh_token = Cookies.get("refresh_token")
                    const response = await axiosWithoutToken.post("/api/v1/auth/refresh-token", {
                        refresh_token
                    }, {
                        headers: {
                            "Authorization": `Bearer ${refresh_token}`,
                            'Content-Type': 'application/json'
                        },
                    });
                    if (response.status === 200) {
                      
                        Cookies.set("access_token", response.data?.data?.access_token);
                        Cookies.set("refresh_token", response.data?.data?.refresh_token);
                        axiosWithToken.defaults.headers.common["Authorization"] = `Bearer ${response.data.access_token}`;
                        axiosWithToken.defaults.headers.common["Content-Type"] = "application/json";
                        isRefreshing = false;
                        onRefreshed(response.data.access_token)
                        refreshSubscribers = [];
                        return axiosWithToken(originalRequest);
                    }
                } catch (refreshError) {
                    isRefreshing = false
                    refreshSubscribers = [];
                
                    Cookies.remove("refresh_token");
                    Cookies.remove("access_token");
                    reloadPage();
                    return Promise.reject(refreshError);
                }
            }else {
                return new Promise((resolve )=> {
                    addRefreshSubscriber((accessToken:string)=>{
                        originalRequest.headers["Authorization"] = `Bearer ${accessToken}`;
                        resolve(axiosWithToken(originalRequest))
                    })
                })
            }
        }

        // logout
        if(error.response && error.response.status === 403 ) {
       
            Cookies.remove("access_token")
            Cookies.remove("refresh_token")
            reloadPage();
        }
        return Promise.reject(error);
    }
)


function reloadPage() {
    window.location.reload();
}