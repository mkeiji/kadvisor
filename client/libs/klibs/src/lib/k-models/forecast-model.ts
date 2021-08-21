import { BaseModel } from './base-model';
import { Column } from 'material-table';

export interface Forecast extends BaseModel {
    userID: number;
    year: number;
    entries: ForecastEntry[];
}

export interface ForecastEntry extends BaseModel {
    forecastID: number;
    month: number;
    income: number;
    expense: number;
}

export interface ForecastTableState {
    columns: Array<Column<ForecastEntry>>;
    data: ForecastEntry[];
}
