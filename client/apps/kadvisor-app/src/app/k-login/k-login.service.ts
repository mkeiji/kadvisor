import { Login } from './view-models';
import { Observable } from 'rxjs';
import { KRxios } from '@client/klibs';

class KLoginService {
    krxios = new KRxios('http://localhost:8081/api');
    loginEndpoint = '/login';
    logoutEndpoint = '/logout';

    login(user: Login): Observable<any> {
        return this.krxios.post(this.loginEndpoint, JSON.stringify(user));
    }

    logout(user: any): Observable<any> {
        return this.krxios.post(this.logoutEndpoint, JSON.stringify(user));
    }
}

export default KLoginService;
