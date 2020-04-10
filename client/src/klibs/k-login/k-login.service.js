import KRxios from "klibs/k-rxios/krxios";


class KLoginService {
    baseUrl = 'http://localhost:8081/api';
    loginEndpoint = '/login';
    logoutEndpoint = '/logout';

    constructor() {
        this.krxios = new KRxios(this.baseUrl);
    }

    login(user) {
        return this.krxios.post(this.loginEndpoint, JSON.stringify(user));
    }

    loginUnsub() {
        this.krxios.post(this.loginEndpoint).unsubscribe();
    }

    logout(user) {
        return this.krxios.post(this.logoutEndpoint, JSON.stringify(user));
    }

    logoutUnsub() {
        this.krxios.post(this.logoutEndpoint).unsubscribe();
    }

    unsubscribe() {
        this.loginUnsub();
        this.logoutUnsub();
    }
}

export default KLoginService;