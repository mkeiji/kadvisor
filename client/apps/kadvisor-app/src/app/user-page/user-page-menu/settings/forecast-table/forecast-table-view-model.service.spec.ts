import { Forecast, ForecastEntry } from '@client/klibs';
import ForecastTableViewModelService from './forecast-table-view-model.service';
import { ForecastTableComponentState } from './view-model';

describe('ForecastTableViewModelService', () => {
    let service: ForecastTableViewModelService;

    const testID = 1;
    const testYear = new Date().getFullYear();

    beforeEach(() => {
        service = new ForecastTableViewModelService();
    });

    describe('formatTableState', () => {
        it('should create a ForecastTableState from a Forecast', () => {
            const entries = [
                {
                    forecastID: 1
                }
            ];
            const forecast = ({
                userID: testID,
                year: testYear,
                entries: entries
            } as unknown) as Forecast;

            const result = service.formatTableState(forecast);
            expect(result.columns).toHaveLength(3);
            expect(result.data).toEqual(entries);
        });
    });

    describe('parseAmounts', () => {
        it('should parese income and expense to number', () => {
            const testEntry = ({
                income: 4.01,
                expense: 1.11
            } as unknown) as ForecastEntry;

            const result = service.parseAmounts(testEntry);
            expect(isNaN(result.income)).toBeFalsy();
            expect(isNaN(result.expense)).toBeFalsy();
        });
    });

    describe('createNewForecast', () => {
        it('should create a forecast with 12 entries, one for each month', () => {
            const result = service.createNewForecast(testID, testYear);
            expect(result.entries).toHaveLength(12);
            for (let i = 1; i <= 12; i++) {
                expect(result.entries[i - 1].month).toEqual(i);
            }
        });
    });

    describe('handleTableStateUpdate', () => {
        it('should return state with updated entry', () => {
            const newEntry = {
                forecastID: 1,
                month: 1,
                income: 2
            } as ForecastEntry;
            const oldEntry = {
                forecastID: 1,
                month: 1,
                income: 1
            } as ForecastEntry;
            const prevState = {
                table: {
                    data: [oldEntry]
                }
            } as ForecastTableComponentState;
            const expectedState = {
                table: {
                    data: [newEntry]
                }
            } as ForecastTableComponentState;

            const result = service.handleTableStateUpdate(
                prevState,
                oldEntry,
                newEntry
            );
            expect(result).toEqual(expectedState);
        });
    });

    describe('handleYearMenuItemsStateUpdate', () => {
        it('should return state updated with new yearMenuItems', () => {
            const currentYear = new Date().getFullYear();
            const forecast = {
                year: currentYear
            } as Forecast;
            const prevState = {
                yearMenuItems: []
            } as ForecastTableComponentState;
            const expectedItems = [
                {
                    value: currentYear,
                    displayValue: currentYear.toString()
                }
            ];
            const expectedState = {
                yearMenuItems: expectedItems,
                hasForecast: true,
                selectedYear: currentYear,
                forecastYear: ''
            };

            const result = service.handleYearMenuItemsStateUpdate(
                prevState,
                forecast
            );
            expect(result).toEqual(expectedState);
        });
    });
});
