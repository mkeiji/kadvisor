import { BaseModel } from './base-model';

export interface Entry extends BaseModel {
    id: number;
    userID: number;
    entryTypeCodeID: string;
    classID: number;
    amount: number;
    date: string;
    description: string;
    obs: string;
}
