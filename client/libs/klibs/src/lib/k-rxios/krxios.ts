/*
    WRAPPER to make axios 'promisses' calls become rxjs 'observables'
    fork from: https://github.com/davguij/rxios
*/
import { Observable } from 'rxjs';
import axios, { AxiosRequestConfig, AxiosInstance } from 'axios';
import { Auth, Login, AuthError, AuthSuccess } from '../k-models/login';
import { APP_LOGIN_ENDPOINT } from '../k-utils/router/route.constants';

export class KRxios {
    private readonly TOKEN_KEY = 'token';
    baseUrl: string;
    options: AxiosRequestConfig;
    _httpClient: AxiosInstance;

    constructor(baseUrl = '', options = {}) {
        this.baseUrl = baseUrl;
        this.options = options;
        this._httpClient = axios.create(options);
    }
    _makeRequest(method: any, url: any, queryParams?: any, body?: any) {
        let request: any;
        if (localStorage.getItem(this.TOKEN_KEY) !== null) {
            this._httpClient.interceptors.request.use(
                (c: AxiosRequestConfig) => {
                    c.headers.Authorization = `Bearer ${localStorage.getItem(
                        this.TOKEN_KEY
                    )}`;
                    return c;
                }
            );
        }
        switch (method) {
            case 'GET':
                request = this._httpClient.get(url, { params: queryParams });
                break;
            case 'POST':
                request = this._httpClient.post(url, body, {
                    params: queryParams
                });
                break;
            case 'PUT':
                request = this._httpClient.put(url, body, {
                    params: queryParams
                });
                break;
            case 'PATCH':
                request = this._httpClient.patch(url, body, {
                    params: queryParams
                });
                break;
            case 'DELETE':
                request = this._httpClient.delete(url, { params: queryParams });
                break;
            default:
                throw new Error('Method not supported');
        }
        return new Observable((subscriber: any) => {
            request
                .then((response: any) => {
                    subscriber.next(response.data);
                    subscriber.complete();
                })
                .catch((err: any) => {
                    subscriber.error(err);
                    subscriber.complete();
                });
        });
    }
    async getToken(login: Partial<Login>): Promise<Auth> {
        const request = await this._httpClient
            .post(`${this.baseUrl}${APP_LOGIN_ENDPOINT.auth}`, login)
            .catch(() => {
                return {
                    code: 401,
                    message: 'incorrect Username or Password'
                } as AuthError;
            });
        return request;
    }
    get(url: any, queryParams?: any): Observable<any> {
        return this.baseUrl !== ''
            ? this._makeRequest('GET', `${this.baseUrl}${url}`, queryParams)
            : this._makeRequest('GET', url, queryParams);
    }
    post(url: any, body?: any, queryParams?: any): Observable<any> {
        return this.baseUrl !== ''
            ? this._makeRequest(
                  'POST',
                  `${this.baseUrl}${url}`,
                  queryParams,
                  body
              )
            : this._makeRequest('POST', url, queryParams, body);
    }
    put(url: any, body?: any, queryParams?: any): Observable<any> {
        return this.baseUrl !== ''
            ? this._makeRequest(
                  'PUT',
                  `${this.baseUrl}${url}`,
                  queryParams,
                  body
              )
            : this._makeRequest('PUT', url, queryParams, body);
    }
    patch(url: any, body?: any, queryParams?: any): Observable<any> {
        return this.baseUrl !== ''
            ? this._makeRequest(
                  'PATCH',
                  `${this.baseUrl}${url}`,
                  queryParams,
                  body
              )
            : this._makeRequest('PATCH', url, queryParams, body);
    }
    delete(url: any, queryParams?: any): Observable<any> {
        return this.baseUrl !== ''
            ? this._makeRequest('DELETE', `${this.baseUrl}${url}`, queryParams)
            : this._makeRequest('DELETE', url, queryParams);
    }
}
