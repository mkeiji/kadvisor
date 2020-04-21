import { BaseModel } from '@client/klibs';

export interface Login {
    email: string;
    password: string;
}

export interface KLoginState {
    login: Login;
    isLoggedIn: boolean;
    hasWarning: boolean;
    userID?: number;
}

export interface KLoginPropTypes {
    userPageUrl: string;
    loginObj: any;
    onLogin: Function;
    onLogout: Function;
}

export interface KLoginResponse {
    login: LoginResponse;
}

interface LoginResponse extends BaseModel {
    userID: number;
    roleID: number;
    email: string;
    userName: string;
    password: string;
    isLoggedIn: boolean;
    LastLogin: string;
}
