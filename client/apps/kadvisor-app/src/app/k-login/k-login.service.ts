import { Observable } from 'rxjs';
import {
    APP_LOGIN_ENDPOINT,
    KRxios,
    Login,
    Auth,
    APP_BACKEND
} from '@client/klibs';

class KLoginService {
    constructor(private readonly krxios?: KRxios) {
        this.krxios = krxios ? krxios : new KRxios(APP_BACKEND);
    }

    login(user: Partial<Login>): Observable<Login> {
        return this.krxios.post(APP_LOGIN_ENDPOINT.login, JSON.stringify(user));
    }

    logout(user: Partial<Login>): Observable<Login> {
        return this.krxios.post(
            APP_LOGIN_ENDPOINT.logout,
            JSON.stringify(user)
        );
    }

    getToken(user: Partial<Login>): Promise<Auth> {
        return this.krxios.getToken(user);
    }
}

export default KLoginService;
