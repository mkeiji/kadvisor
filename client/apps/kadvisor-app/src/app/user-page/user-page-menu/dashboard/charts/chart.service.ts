import {
    KRxios,
    KEndpointUtil,
    MonthReport,
    APP_REPORT_ENDPOINT
} from '@client/klibs';
import { Observable } from 'rxjs';

export default class ChartService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getYtdWithForecastReport(year: number): Observable<MonthReport[]> {
        return this.krxios.get(`${APP_REPORT_ENDPOINT}?type=YFC&year=${year}`);
    }
}
