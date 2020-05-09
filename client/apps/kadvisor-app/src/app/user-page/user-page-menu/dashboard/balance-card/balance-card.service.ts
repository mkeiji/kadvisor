import {
    APP_REPORT_ENDPOINT,
    KEndpointUtil,
    KRxios,
    UserBalance
} from '@client/klibs';
import { Observable } from 'rxjs';

class BalanceCardService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getUserBalance(): Observable<UserBalance> {
        return this.krxios.get(`${APP_REPORT_ENDPOINT}?type=balance`);
    }
}

export default BalanceCardService;
