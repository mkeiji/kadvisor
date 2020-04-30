import { Observable } from 'rxjs';
import { KLoginResponse, KRxios, Login } from '@client/klibs';

class KLoginService {
    krxios = new KRxios('http://localhost:8081/api');
    loginEndpoint = '/login';
    logoutEndpoint = '/logout';

    login(user: Partial<Login>): Observable<KLoginResponse> {
        return this.krxios.post(this.loginEndpoint, JSON.stringify(user));
    }

    logout(user: Partial<Login>): Observable<KLoginResponse> {
        return this.krxios.post(this.logoutEndpoint, JSON.stringify(user));
    }
}

export default KLoginService;
