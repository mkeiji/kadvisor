import {
    APP_REPORT_ENDPOINT,
    KEndpointUtil,
    KRxios,
    UserBalance
} from '@client/klibs';
import { Observable } from 'rxjs';

class BalanceCardService {
    private krxios: KRxios;

    constructor(userID: number, krxios?: KRxios) {
        this.krxios = krxios
            ? krxios
            : new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getUserBalance(): Observable<UserBalance> {
        return this.krxios.get(`${APP_REPORT_ENDPOINT}?type=BALANCE`);
    }
}

export default BalanceCardService;
