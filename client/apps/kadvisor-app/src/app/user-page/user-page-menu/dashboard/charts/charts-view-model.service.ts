import { MonthReport } from '@client/klibs';

class ChartsViewModelService {
    getEmptyMonthReport(): MonthReport[] {
        let result = [];

        for (let i = 1; i <= 12; i++) {
            result.push({
                year: 0,
                month: i,
                income: 0,
                expense: 0,
                balance: 0
            } as MonthReport);
        }

        return result;
    }
}

export default ChartsViewModelService;
