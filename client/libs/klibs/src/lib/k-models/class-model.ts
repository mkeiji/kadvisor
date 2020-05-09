import { BaseModel } from './base-model';

export interface Class extends BaseModel {
    userID: number;
    name: string;
    description: string;
}
