/*
    WRAPPER to make axios 'promisses' calls become rxjs 'observables'
    fork from: https://github.com/davguij/rxios
*/

const axios_1 = require("axios");
const Observable_1 = require("rxjs/Observable");
class KRxios {
    constructor(baseUrl = "", options = {}) {
        this.baseUrl = baseUrl;
        this.options = options;
        this._httpClient = axios_1.default.create(options);
    }
    _makeRequest(method, url, queryParams, body) {
        let request;
        switch (method) {
            case 'GET':
                request = this._httpClient.get(url, { params: queryParams });
                break;
            case 'POST':
                request = this._httpClient.post(url, body, { params: queryParams });
                break;
            case 'PUT':
                request = this._httpClient.put(url, body, { params: queryParams });
                break;
            case 'PATCH':
                request = this._httpClient.patch(url, body, { params: queryParams });
                break;
            case 'DELETE':
                request = this._httpClient.delete(url, { params: queryParams });
                break;
            default:
                throw new Error('Method not supported');
        }
        return new Observable_1.Observable(subscriber => {
            request.then(response => {
                subscriber.next(response.data);
                subscriber.complete();
            }).catch((err) => {
                subscriber.error(err);
                subscriber.complete();
            });
        });
    }
    get(url, queryParams) {
        return this.baseUrl !== "" 
        ? this._makeRequest('GET', `${this.baseUrl}${url}`, queryParams) 
        : this._makeRequest('GET', url, queryParams) ;
    }
    post(url, body, queryParams) {
        return this.baseUrl !== "" 
        ? this._makeRequest('POST', `${this.baseUrl}${url}`, queryParams, body)
        : this._makeRequest('POST', url, queryParams, body);
    }
    put(url, body, queryParams) {
        return this.baseUrl !== "" 
        ? this._makeRequest('PUT', `${this.baseUrl}${url}`, queryParams, body)
        : this._makeRequest('PUT', url, queryParams, body);
    }
    patch(url, body, queryParams) {
        return this.baseUrl !== "" 
        ? this._makeRequest('PATCH', `${this.baseUrl}${url}`, queryParams, body)
        : this._makeRequest('PATCH', url, queryParams, body);
    }
    delete(url, queryParams) {
        return this.baseUrl !== "" 
        ? this._makeRequest('DELETE', `${this.baseUrl}${url}`, queryParams) 
        : this._makeRequest('DELETE', url, queryParams) ;
    }
}
export default KRxios;
