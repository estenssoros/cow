import axios, { AxiosResponse, AxiosRequestConfig } from 'axios'

const BASE_URL = 'http://localhost:3001'

const requestFactory = (
    method: string,
    url: string,
    headers: object,
    params: object,
    data: object,
    token: object,
) => {
    return new Promise(
        (resolve: Function): Promise<AxiosRequestConfig> => {
            return resolve({
                method,
                url,
                headers: {
                    Token: token ? token : null,
                    'Content-Type': 'application/json',
                    ...headers
                },
                params: params,
                data: data,
            })
        }
    )
}

const API = () => {
    return {
        get: (
            slug: string,
            headers: object = {},
            params: object = {}
        ): Promise<AxiosResponse<object>> => {
            return requestFactory(
                'GET',
                `${BASE_URL}/${slug}`,
                headers,
                params,
                {},
                {}
            ).then((REQUEST: AxiosRequestConfig) => {
                return axios(REQUEST)
            })
        },
        post: (
            slug: string,
            headers: object,
            params: object,
            data: object,
            token: {}
        ): Promise<AxiosResponse<object>> => {
            return requestFactory(
                'POST',
                `${BASE_URL}/${slug}`,
                headers,
                params,
                data,
                {},
            ).then((REQUEST: AxiosRequestConfig)=>{
                return axios(REQUEST)
            })
        },
        patch: (
            slug: string,
            headers: object,
            params: object,
            data: object
        ): Promise<AxiosResponse<object>> => {
            return requestFactory(
                'PATCH',
                `${BASE_URL}/${slug}`,
                headers,
                params,
                data,
                {}
            ).then((REQUEST: AxiosRequestConfig) => {
                return axios(REQUEST)
            })
        },
        put: (
            slug: string,
            headers: object,
            params: object,
            data: object
        ): Promise<AxiosResponse<object>> => {
            return requestFactory(
                'PUT',
                `${BASE_URL}/${slug}`,
                headers,
                params,
                data,
                {}
            ).then((REQUEST: AxiosRequestConfig) => {
                return axios(REQUEST)
            })
        },
        delete: (
            slug: string,
            headers: object,
            params: object,
            data: object
        ): Promise<AxiosResponse<object>> => {
            return requestFactory(
                'DELETE',
                `${BASE_URL}/${slug}`,
                headers,
                params,
                data,
                {}
            ).then(
                (REQUEST: AxiosRequestConfig): Promise<AxiosResponse<object>> => {
                    return axios(REQUEST)
                }
            )
        },
    }
}

export default API