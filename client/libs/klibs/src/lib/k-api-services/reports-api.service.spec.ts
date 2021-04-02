import { Observable, of } from 'rxjs';
import { MonthReport } from '../k-models/chart-models';
import { KRxios } from '../k-rxios/krxios';
import { provideMock } from '../k-utils/util-constants';
import { ReportsApiService } from './reports-api.service';
import {
    APP_REPORT_AVAILABLE_ENDPOINT,
    APP_REPORT_ENDPOINT
} from '../k-utils/router/route.constants';

describe('ReportsApiService', () => {
    let testYear: number;
    let testID: number;
    let mockKRxios: KRxios;
    let service: ReportsApiService;

    beforeEach(() => {
        testYear = new Date().getFullYear();
        testID = 1;
        mockKRxios = provideMock({
            get: jest.fn()
        });
        service = new ReportsApiService(testID, mockKRxios);
    });

    describe('getYtdWithForecastReport', () => {
        it('should call krxios get', () => {
            const expectedParam = `${APP_REPORT_ENDPOINT}?type=YFC&year=${testYear}`;
            const expected = of({} as MonthReport);
            jest.spyOn(mockKRxios, 'get').mockReturnValue(expected);

            const result = service.getYtdWithForecastReport(testYear);
            expect(result).toBe(expected);
            expect(mockKRxios.get).toHaveBeenCalledWith(expectedParam);
        });
    });

    describe('getAvailableReportYears', () => {
        let expectedParam: string;
        let expected: Observable<any>;
        let isForecastParam: boolean | undefined;

        afterEach(() => {
            const result = service.getAvailableReportYears(isForecastParam);
            expect(result).toBe(expected);
            expect(mockKRxios.get).toHaveBeenCalledWith(expectedParam);
        });

        it('should call krxios get without param', () => {
            expectedParam = APP_REPORT_AVAILABLE_ENDPOINT;
            expected = of([] as number[]);
            jest.spyOn(mockKRxios, 'get').mockReturnValue(expected);
        });

        it('should call krxios get with forecast param', () => {
            isForecastParam = true;
            expectedParam = `${APP_REPORT_AVAILABLE_ENDPOINT}?forecast=true`;
            expected = of([] as number[]);
            jest.spyOn(mockKRxios, 'get').mockReturnValue(expected);
        });
    });
});
