import { Observable } from 'rxjs';
import { KRxios } from '../k-rxios/krxios';
import { KEndpointUtil } from '../k-utils/router/k-endpoint.util';
import {
    APP_REPORT_ENDPOINT,
    APP_REPORT_AVAILABLE_ENDPOINT
} from '../k-utils/router/route.constants';
import { MonthReport } from '../k-models/chart-models';

export class ReportsApiService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getYtdWithForecastReport(year: number): Observable<MonthReport[]> {
        return this.krxios.get(`${APP_REPORT_ENDPOINT}?type=YFC&year=${year}`);
    }

    getAvailableReportYears(isForecast?: boolean): Observable<number[]> {
        const forecastParam = `?forecast=${isForecast}`;
        return this.krxios.get(
            `${APP_REPORT_AVAILABLE_ENDPOINT}${isForecast ? forecastParam : ''}`
        );
    }
}
