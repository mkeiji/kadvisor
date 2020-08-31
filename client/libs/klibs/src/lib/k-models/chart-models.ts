import { BaseModel } from './base-model';

export interface MonthReport extends BaseModel {
    year: number;
    month: number;
    income: number;
    expense: number;
    balance: number;
}
