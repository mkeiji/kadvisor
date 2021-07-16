import { MonthReport } from '@client/klibs';
import ChartsViewModelService from './charts-view-model.service';

describe('ChartsViewModelService', () => {
    let service: ChartsViewModelService;
    let testReport: MonthReport[];

    beforeEach(() => {
        service = new ChartsViewModelService();
        testReport = service.getEmptyMonthReport();
    });

    describe('getEmptyMonthReport', () => {
        it('should create a MonthReport array with size 12', () => {
            const expectedSize = 12;

            const result = service.getEmptyMonthReport();
            expect(result.length).toEqual(expectedSize);
            for (let i in result) {
                expect(result[i].year).toEqual(0);
                expect(result[i].month).toEqual(Number(i) + 1);
                expect(result[i].income).toEqual(0);
                expect(result[i].expense).toEqual(0);
                expect(result[i].balance).toEqual(0);
            }
        });
    });

    describe('getTicksForNegativeBalance', () => {
        it('should return undefined ticks if balance is greater or equal to 0', () => {
            const [res1Tick1, res1Tick2] = service.getTicksForNegativeBalance(
                testReport
            );
            expect(res1Tick1).toBeUndefined();
            expect(res1Tick2).toBeUndefined();

            testReport[0].balance = 1;
            const [res2Tick1, res2Tick2] = service.getTicksForNegativeBalance(
                testReport
            );
            expect(res2Tick1).toBeUndefined();
            expect(res2Tick2).toBeUndefined();
        });

        it('calls getMaxIncomeOrExpense and getMaxBalance', () => {
            jest.spyOn(service, 'getMaxIncomeOrExpense');
            jest.spyOn(service, 'getMaxBalance');
            testReport[0].balance = -1;

            service.getTicksForNegativeBalance(testReport);
            expect(service.getMaxIncomeOrExpense).toHaveBeenCalledWith(
                testReport
            );
            expect(service.getMaxBalance).toHaveBeenCalledWith(testReport);
        });

        it('calls getTicksForYAxis for maxNegative and maxBalance', () => {
            const expectedMaxNegative = -5;
            const expectedMaxBalance = 10;
            testReport[0].balance = -1;
            jest.spyOn(service, 'getMaxIncomeOrExpense').mockReturnValue(
                expectedMaxNegative * -1
            );
            jest.spyOn(service, 'getMaxBalance').mockReturnValue(
                expectedMaxBalance
            );
            jest.spyOn(service, 'getTicksForYAxis');

            service.getTicksForNegativeBalance(testReport);
            expect(service.getTicksForYAxis).toHaveBeenCalledTimes(2);
            expect(service.getTicksForYAxis).toHaveBeenNthCalledWith(
                1,
                expectedMaxNegative
            );
            expect(service.getTicksForYAxis).toHaveBeenNthCalledWith(
                2,
                expectedMaxBalance
            );
        });
    });

    describe('getMinBalance', () => {
        it('should return the min balance of the report', () => {
            const expectedMin = -2;
            testReport[0].balance = expectedMin;
            testReport[10].balance = 5;

            const result = service.getMinBalance(testReport);
            expect(result).toEqual(expectedMin);
        });
    });

    describe('getMaxBalance', () => {
        it('should return the max balance of the report', () => {
            const expectedMax = 12;
            testReport[0].balance = expectedMax;
            testReport[10].balance = 5;

            const result = service.getMaxBalance(testReport);
            expect(result).toEqual(expectedMax);
        });
    });

    describe('getMaxIncomeOrExpense', () => {
        it('should return highest income if greater then expense', () => {
            const expected = 12;
            testReport[0].income = expected;
            testReport[10].expense = -5;

            const result = service.getMaxIncomeOrExpense(testReport);
            expect(result).toEqual(expected);
        });

        it('should return highest expense if greater then income', () => {
            const expected = 12;
            testReport[0].income = 0;
            testReport[10].expense = expected;

            const result = service.getMaxIncomeOrExpense(testReport);
            expect(result).toEqual(expected);
        });
    });

    describe('getTicksForYAxis', () => {
        it('positive - should return array with 5 rounded numbers', () => {
            const testNumber = 111;
            const expected = [100, 50, 0, -50, -100];

            const result = service.getTicksForYAxis(testNumber);
            expect(result).toEqual(expected);
        });

        it('negative - should return array with 5 rounded numbers', () => {
            const testNumber = -111;
            const expected = [-200, -100, 0, 100, 200];

            const result = service.getTicksForYAxis(testNumber);
            expect(result).toEqual(expected);
        });
    });
});
