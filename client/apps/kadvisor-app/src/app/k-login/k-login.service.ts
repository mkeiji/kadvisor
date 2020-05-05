import { Observable } from 'rxjs';
import { APP_LOGIN_ENDPOINT, KRxios, Login } from '@client/klibs';

class KLoginService {
    krxios = new KRxios('http://localhost:8081/api');

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
