import {
    APP_FORECAST_ENDPOINT,
    APP_FORECAST_ENTRY_ENDPOINT,
    Forecast,
    ForecastEntry,
    KEndpointUtil,
    KRxios
} from '@client/klibs';
import { Observable } from 'rxjs';

class ForecastTableService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getForecast(year: number): Observable<Forecast> {
        return this.krxios.get(
            `${APP_FORECAST_ENDPOINT}?year=${year}&preloaded=true`
        );
    }

    postForecast(forecast: Forecast): Observable<Forecast> {
        return this.krxios.post(APP_FORECAST_ENDPOINT, forecast);
    }

    deleteForecast(forecastID: number): Observable<Forecast> {
        return this.krxios.delete(APP_FORECAST_ENDPOINT, {
            id: forecastID
        });
    }

    putForecastEntry(entry: ForecastEntry): Observable<ForecastEntry> {
        return this.krxios.put(APP_FORECAST_ENTRY_ENDPOINT, entry);
    }
}

export default ForecastTableService;
