import { KRxios } from './krxios';
import { Observable } from 'rxjs';
import { Login, AuthError } from '../k-models/login';
import { APP_LOGIN_ENDPOINT } from '../k-utils/router/route.constants';

describe('KRxios', () => {
    let testKRxios: KRxios;

    const tokenKey = 'token';
    const testUrl = 'test/url';
    const testUrl2 = '/testUrl2';
    const testParams = { params: 'test' };
    const testOptions = {};

    beforeEach(() => {
        testKRxios = new KRxios(testUrl, testOptions);
        window.localStorage.__proto__.getItem = jest.fn();
    });

    describe('constructor', () => {
        it('should set url, options and initialize client', () => {
            const result = new KRxios(testUrl, testOptions);

            expect(result.baseUrl).toEqual(testUrl);
            expect(result.options).toEqual(testOptions);
            expect(result._httpClient).not.toBeNull();
        });
    });

    describe('_setRequestInterceptor', () => {
        it('should get token from localStorage', () => {
            jest.spyOn(
                window.localStorage.__proto__,
                'getItem'
            ).mockReturnValue('xxx');
            jest.spyOn(testKRxios._httpClient.interceptors.request, 'use');

            testKRxios._setRequestInterceptor();
            expect(localStorage.getItem).toHaveBeenCalledWith(tokenKey);
            expect(
                testKRxios._httpClient.interceptors.request.use
            ).toHaveBeenCalled();
        });
    });

    describe('_setResponseInterceptor', () => {
        it('should call interceptors.response.use', () => {
            jest.spyOn(testKRxios._httpClient.interceptors.response, 'use');

            testKRxios._setResponseInterceptor();
            expect(
                testKRxios._httpClient.interceptors.response.use
            ).toHaveBeenCalled();
        });
    });

    describe('_makeRequest', () => {
        let result: any;

        const testBody = {};
        const expectedParams = { params: testParams };

        afterEach(() => {
            expect(result instanceof Observable).toBeTruthy();
        });

        it('should set interceptors', () => {
            jest.spyOn(testKRxios._httpClient.interceptors.request, 'use');
            jest.spyOn(testKRxios._httpClient.interceptors.response, 'use');
            jest.spyOn(testKRxios._httpClient, 'get').mockReturnValue(
                new Promise((r, _) => r('ok'))
            );

            result = testKRxios._makeRequest('GET', testUrl, null, null);

            expect(
                testKRxios._httpClient.interceptors.request.use
            ).toHaveBeenCalled();
            expect(
                testKRxios._httpClient.interceptors.response.use
            ).toHaveBeenCalled();
        });

        it('GET - should call client.get and return an observable', () => {
            jest.spyOn(testKRxios._httpClient, 'get').mockReturnValue(
                new Promise((r, _) => r('ok'))
            );

            result = testKRxios._makeRequest('GET', testUrl, testParams, null);
            expect(testKRxios._httpClient.get).toHaveBeenCalledWith(
                testUrl,
                expectedParams
            );
        });

        it('POST - should call client.post and return an observable', () => {
            jest.spyOn(testKRxios._httpClient, 'post').mockReturnValue(
                new Promise((r, _) => r('ok'))
            );

            result = testKRxios._makeRequest(
                'POST',
                testUrl,
                testParams,
                testBody
            );
            expect(testKRxios._httpClient.post).toHaveBeenCalledWith(
                testUrl,
                testBody,
                expectedParams
            );
        });

        it('PUT - should call client.put and return an observable', () => {
            jest.spyOn(testKRxios._httpClient, 'put').mockReturnValue(
                new Promise((r, _) => r('ok'))
            );

            result = testKRxios._makeRequest(
                'PUT',
                testUrl,
                testParams,
                testBody
            );
            expect(testKRxios._httpClient.put).toHaveBeenCalledWith(
                testUrl,
                testBody,
                expectedParams
            );
        });

        it('PATCH - should call client.patch and return an observable', () => {
            jest.spyOn(testKRxios._httpClient, 'patch').mockReturnValue(
                new Promise((r, _) => r('ok'))
            );

            result = testKRxios._makeRequest(
                'PATCH',
                testUrl,
                testParams,
                testBody
            );
            expect(testKRxios._httpClient.patch).toHaveBeenCalledWith(
                testUrl,
                testBody,
                expectedParams
            );
        });

        it('DELETE - should call client.delete and return an observable', () => {
            jest.spyOn(testKRxios._httpClient, 'delete').mockReturnValue(
                new Promise((r, _) => r('ok'))
            );

            result = testKRxios._makeRequest(
                'DELETE',
                testUrl,
                testParams,
                null
            );
            expect(testKRxios._httpClient.delete).toHaveBeenCalledWith(
                testUrl,
                expectedParams
            );
        });
    });

    describe('getToken', () => {
        const testLogin = { email: 'test' } as Partial<Login>;

        it('should return a promise with the token', () => {
            const fakeToken = 'fake-token';
            const expectedUrl = `${testUrl}${APP_LOGIN_ENDPOINT.auth}`;
            jest.spyOn(testKRxios._httpClient, 'post').mockReturnValue(
                new Promise((r, _) => r(fakeToken))
            );

            const result = testKRxios.getToken(testLogin).then((data) => {
                expect(data).toBe(fakeToken);
            });
            expect(testKRxios._httpClient.post).toHaveBeenCalledWith(
                expectedUrl,
                testLogin
            );
            expect(result instanceof Promise).toBeTruthy();
        });

        it('should return a promisse with an auth error', () => {
            const expected = {
                code: 401,
                message: 'incorrect Username or Password'
            } as AuthError;
            jest.spyOn(testKRxios._httpClient, 'post').mockReturnValue(
                new Promise((_, e) => e(expected))
            );

            const result = testKRxios.getToken(testLogin).then((data) => {
                expect(data).toEqual(expected);
            });
            expect(result instanceof Promise).toBeTruthy();
        });
    });

    describe('get, post, put, patch, delete', () => {
        beforeEach(() => {
            jest.spyOn(testKRxios, '_makeRequest');
        });

        describe('get', () => {
            it('should call _makeRequest with baseUrl if available', () => {
                testKRxios.get(testUrl2, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'GET',
                    `${testUrl}${testUrl2}`,
                    testParams
                );
            });

            it('should call _makeRequest without baseUrl if not available', () => {
                testKRxios.baseUrl = '';

                testKRxios.get(testUrl2, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'GET',
                    testUrl2,
                    testParams
                );
            });
        });

        describe('post', () => {
            it('should call _makeRequest with baseUrl if abailable', () => {
                const testBody = {};

                testKRxios.post(testUrl2, testBody, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'POST',
                    `${testUrl}${testUrl2}`,
                    testParams,
                    testBody
                );
            });

            it('should call _makeRequest without baseUrl if not available', () => {
                const testBody = {};
                testKRxios.baseUrl = '';

                testKRxios.post(testUrl2, testBody, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'POST',
                    testUrl2,
                    testParams,
                    testBody
                );
            });
        });

        describe('put', () => {
            it('should call _makeRequest with baseUrl if abailable', () => {
                const testBody = {};

                testKRxios.put(testUrl2, testBody, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'PUT',
                    `${testUrl}${testUrl2}`,
                    testParams,
                    testBody
                );
            });

            it('should call _makeRequest without baseUrl if not available', () => {
                const testBody = {};
                testKRxios.baseUrl = '';

                testKRxios.put(testUrl2, testBody, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'PUT',
                    testUrl2,
                    testParams,
                    testBody
                );
            });
        });

        describe('patch', () => {
            it('should call _makeRequest with baseUrl if abailable', () => {
                const testBody = {};

                testKRxios.patch(testUrl2, testBody, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'PATCH',
                    `${testUrl}${testUrl2}`,
                    testParams,
                    testBody
                );
            });

            it('should call _makeRequest without baseUrl if not available', () => {
                const testBody = {};
                testKRxios.baseUrl = '';

                testKRxios.patch(testUrl2, testBody, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'PATCH',
                    testUrl2,
                    testParams,
                    testBody
                );
            });
        });

        describe('delete', () => {
            it('should call _makeRequest with baseUrl if available', () => {
                testKRxios.delete(testUrl2, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'DELETE',
                    `${testUrl}${testUrl2}`,
                    testParams
                );
            });

            it('should call _makeRequest without baseUrl if not available', () => {
                testKRxios.baseUrl = '';

                testKRxios.delete(testUrl2, testParams).subscribe();
                expect(testKRxios._makeRequest).toHaveBeenCalledWith(
                    'DELETE',
                    testUrl2,
                    testParams
                );
            });
        });
    });
});
