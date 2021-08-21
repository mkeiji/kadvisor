import {
    APP_FORECAST_ENDPOINT,
    APP_FORECAST_ENTRY_ENDPOINT,
    Forecast,
    ForecastEntry,
    KRxios
} from '@client/klibs';
import ForecastTableService from './forecast-table.service';

describe('ForecastTableService', () => {
    let mockKrxios: KRxios;
    let service: ForecastTableService;

    const testID = 1;
    const testYear = new Date().getFullYear();

    beforeEach(() => {
        setupMocks();
        service = new ForecastTableService(testID, mockKrxios);
    });

    describe('getForecast', () => {
        it('should call krxios get', () => {
            const expectedUrl = `${APP_FORECAST_ENDPOINT}?year=${testYear}&preloaded=true`;

            service.getForecast(testYear);
            expect(mockKrxios.get).toHaveBeenCalledWith(expectedUrl);
        });
    });

    describe('postForecast', () => {
        it('should call krxios post', () => {
            const expectedForecast = ({
                userID: testID,
                year: testYear
            } as unknown) as Forecast;

            service.postForecast(expectedForecast);
            expect(mockKrxios.post).toHaveBeenCalledWith(
                APP_FORECAST_ENDPOINT,
                expectedForecast
            );
        });
    });

    describe('deleteClass', () => {
        it('should call krxios delete', () => {
            const expectedForecast = ({
                id: testID
            } as unknown) as Forecast;

            service.deleteForecast(testID);
            expect(mockKrxios.delete).toHaveBeenCalledWith(
                APP_FORECAST_ENDPOINT,
                expectedForecast
            );
        });
    });

    describe('putForecastEntry', () => {
        it('should call krxios put', () => {
            const expectedForecastEntry = ({
                forecastID: testID,
                month: 1,
                income: 5.0,
                expense: 0.0
            } as unknown) as ForecastEntry;

            service.putForecastEntry(expectedForecastEntry);
            expect(mockKrxios.put).toHaveBeenCalledWith(
                APP_FORECAST_ENTRY_ENDPOINT,
                expectedForecastEntry
            );
        });
    });

    function setupMocks() {
        mockKrxios = ({
            get: jest.fn(),
            post: jest.fn(),
            put: jest.fn(),
            delete: jest.fn()
        } as unknown) as KRxios;
    }
});
