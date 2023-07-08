import axios, {AxiosResponse, HttpStatusCode} from "axios";
import {NET_URL} from "../constants/netwrok";
import {useAuthStore} from "../stores/auth/authstore";

/**
 * axios instance that can be used with all configuration configured
 */
export const client = axios.create({  //TODO make this client global singleton so that auth header is not reset
    baseURL: NET_URL,
});

//config client with middleware to add auth and required configs to request
client.interceptors.request.use((config) => {
    console.log("auth  header is :" + config.headers.Authorization);
    config.headers.Authorization =
        config.headers.Authorization === "" ||
        config.headers.Authorization === undefined ?
            useAuthStore.getState().authToken : config.headers.Authorization;
    return config;
});


/**
 * send  creates as post  request with given routes
 * @param data data to send
 * @param route route to which send relative to base route
 * @returns data sent by the server
 */
export async function send<T>(data: unknown, route: string) {
    let result: AxiosResponse<T>;
    try {
        result = await client.post(route, data);
        if (result.status === HttpStatusCode.Ok) {
            return result.data satisfies T;
        }
    } catch (e) {
        console.log(e);
    }

    return {};
}


/**
 * get creates a get request to the given routes
 * @param route route to which send relative to base route
 * @returns data sent by the server
 */
export async function get<T>(route: string) {
    let result: AxiosResponse<T>;
    try {
        result = await client.get(route);
        if (result.status === HttpStatusCode.Ok) {
            return result.data satisfies T;
        }
    } catch (e) {
        console.log(e);
    }

    return {};
}
