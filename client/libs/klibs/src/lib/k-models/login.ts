import { BaseModel } from './base-model';

export interface Login extends BaseModel {
    userID: number;
    roleID: number;
    email: string;
    userName: string;
    password: string;
    isLoggedIn: boolean;
    lastLogin: string;
}
