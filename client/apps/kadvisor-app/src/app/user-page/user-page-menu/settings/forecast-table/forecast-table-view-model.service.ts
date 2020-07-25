import { Forecast, ForecastEntry, KFormatUtil } from '@client/klibs';
import { Column } from 'material-table';
import { orderBy } from 'lodash';

export interface ForecastTableState {
    columns: Array<Column<ForecastEntry>>;
    data: ForecastEntry[];
}

class ForecastTableViewModelService {
    formatTableState(forecast: Forecast): ForecastTableState {
        return {
            columns: [
                {
                    title: 'Month',
                    field: 'month',
                    editable: 'never',
                    render: (data) => months[data.month - 1]
                },
                {
                    title: 'Income',
                    field: 'income',
                    editable: 'onUpdate',
                    type: 'numeric',
                    render: (data) =>
                        KFormatUtil.toCurrency(data.income ? data.income : 0)
                },
                {
                    title: 'Expense',
                    field: 'expense',
                    editable: 'onUpdate',
                    type: 'numeric',
                    render: (data) =>
                        KFormatUtil.toCurrency(data.expense ? data.expense : 0)
                }
            ],
            data: orderBy(forecast.entries, ['month'], ['asc'])
        };
    }

    parseAmounts(entry: ForecastEntry): ForecastEntry {
        entry.income = Number(entry.income);
        entry.expense = Number(entry.expense);
        return entry;
    }

    createNewForecast(userID: number, year: number): Forecast {
        const entries = [] as ForecastEntry[];
        for (let i = 1; i <= 12; i++) {
            entries.push({
                month: i
            } as ForecastEntry);
        }

        return {
            userID: userID,
            year: year,
            entries: entries
        } as Forecast;
    }
}

const months: string[] = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December'
];

export default ForecastTableViewModelService;
