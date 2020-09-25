import { Observable } from 'rxjs';
import { APP_LOGIN_ENDPOINT, KRxios, Login, APP_BACKEND } from '@client/klibs';

class KLoginService {
    krxios = new KRxios(APP_BACKEND);

    login(user: Partial<Login>): Observable<Login> {
        return this.krxios.post(APP_LOGIN_ENDPOINT.login, JSON.stringify(user));
    }

    logout(user: Partial<Login>): Observable<Login> {
        return this.krxios.post(
            APP_LOGIN_ENDPOINT.logout,
            JSON.stringify(user)
        );
    }
}

export default KLoginService;
