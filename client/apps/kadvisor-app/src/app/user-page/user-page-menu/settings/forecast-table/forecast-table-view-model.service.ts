import {
    Forecast,
    ForecastEntry,
    ForecastTableState,
    KFormatUtil
} from '@client/klibs';
import { orderBy } from 'lodash';
import { ForecastTableComponentState } from './view-model';

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

    handleTableStateUpdate(
        prevState: ForecastTableComponentState,
        oldEntry: ForecastEntry,
        newEntry: ForecastEntry
    ): ForecastTableComponentState {
        const data = prevState.table ? [...prevState.table.data] : [];
        data[data.indexOf(oldEntry)] = newEntry;
        const updatedTable = {
            ...prevState.table,
            data: data
        };
        return { ...prevState, table: updatedTable };
    }

    handleYearMenuItemsStateUpdate(
        prevState: ForecastTableComponentState,
        forecast: Forecast
    ): ForecastTableComponentState {
        const updatedMenuItems = prevState.yearMenuItems;
        updatedMenuItems.push({
            value: forecast.year,
            displayValue: forecast.year.toString()
        });
        return {
            ...prevState,
            yearMenuItems: updatedMenuItems,
            hasForecast: true,
            selectedYear: forecast.year,
            forecastYear: ''
        };
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
