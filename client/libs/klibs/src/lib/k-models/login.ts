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

export interface Auth {
    data: any;
}

export interface AuthSuccess extends Auth {
    code: number;
    expire: string;
    token: string;
}

export interface AuthError extends Auth {
    code: number;
    message: string;
}
