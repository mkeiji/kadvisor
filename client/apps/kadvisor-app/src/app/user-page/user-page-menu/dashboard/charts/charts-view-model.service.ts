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

    getTicksForNegativeBalance(
        m: MonthReport[]
    ): [number[] | undefined, number[] | undefined] {
        const minBalance = this.getMinBalance(m);
        if (minBalance >= 0) {
            return [undefined, undefined];
        }

        const maxNegativeIncExp = this.getMaxIncomeOrExpense(m) * -1;
        return [
            this.getTicksForYAxis(maxNegativeIncExp),
            this.getTicksForYAxis(minBalance)
        ];
    }

    getMinBalance(m: MonthReport[]): number {
        return Math.min.apply(
            Math,
            m.map((obj) => obj.balance)
        );
    }

    getMaxIncomeOrExpense(m: MonthReport[]): number {
        const maxExpense = Math.max.apply(
            Math,
            m.map((obj) => obj.expense)
        );
        const maxIncome = Math.max.apply(
            Math,
            m.map((obj) => obj.income)
        );

        return Math.max.apply(Math, [maxExpense, maxIncome]);
    }

    getTicksForYAxis(n: number): number[] {
        const positiveN = n * -1;
        const nearestHundred = Math.ceil(positiveN / 100) * 100;

        return [0, 0, 0, 0, 0].map((_, i) => {
            switch (i) {
                case 0:
                    return nearestHundred * -1;
                case 1:
                    return (nearestHundred / 2) * -1;
                case 3:
                    return nearestHundred / 2;
                case 4:
                    return nearestHundred;
                default:
                    return 0;
            }
        });
    }
}

export default ChartsViewModelService;
