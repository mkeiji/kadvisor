import { BaseModel } from './base-model';

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
