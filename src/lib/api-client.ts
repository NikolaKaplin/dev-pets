import axios, { AxiosInstance, AxiosResponse, AxiosError, InternalAxiosRequestConfig, AxiosRequestConfig } from 'axios';

let onUnauthorizedCallback: (() => void) | null = null;

class ApiClient {
    private instance: AxiosInstance;

    constructor() {
        this.instance = axios.create({
            baseURL: 'http://localhost:8080',
            withCredentials: true,
            timeout: 10000,
        });

        this.setupInterceptors();
    }

    private setupInterceptors() {
        this.instance.interceptors.request.use(
            (config: InternalAxiosRequestConfig) => {
                console.log(`🚀 Making ${config.method?.toUpperCase()} request to: ${config.url}`);
                return config;
            },
            (error: AxiosError) => {
                return Promise.reject(error);
            }
        );

        this.instance.interceptors.response.use(
            (response: AxiosResponse) => {
                return response;
            },
            (error: AxiosError) => {
                this.handleError(error);
                return Promise.reject(error);
            }
        );
    }

    private handleError(error: AxiosError) {
        if (error.response?.status === 401) {
            console.warn('🔐 Session expired, triggering logout');

            if (onUnauthorizedCallback) {
                onUnauthorizedCallback();
            }

            this.clearFrontendAuth();
        }
    }

    private clearFrontendAuth() {
        localStorage.removeItem('auth_redirect');
        sessionStorage.removeItem('user_data');
        document.cookie.split(';').forEach(cookie => {
            const eqPos = cookie.indexOf('=');
            const name = eqPos > -1 ? cookie.substr(0, eqPos).trim() : cookie.trim();
            if (name.includes('auth') || name.includes('token')) {
                document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/`;
            }
        });
    }

    public get axiosInstance(): AxiosInstance {
        return this.instance;
    }

    public async get<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
        const response = await this.instance.get<T>(url, config);
        return response.data;
    }

    public async post<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
        const response = await this.instance.post<T>(url, data, config);
        return response.data;
    }

    public async put<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
        const response = await this.instance.put<T>(url, data, config);
        return response.data;
    }

    public async delete<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
        const response = await this.instance.delete<T>(url, config);
        return response.data;
    }
}

export const setUnauthorizedHandler = (callback: () => void) => {
    onUnauthorizedCallback = callback;
};

export const apiClient = new ApiClient();
export default apiClient;