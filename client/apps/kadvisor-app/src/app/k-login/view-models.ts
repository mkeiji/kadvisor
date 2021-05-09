import { Login } from '@client/klibs';
import KLoginService from './k-login.service';

export interface KLoginState {
    login: Login;
    isLoggedIn: boolean;
    hasWarning: boolean;
    userID?: number;
}

export interface KLoginFormType {
    email: string;
    password: string;
}

export interface KLoginPropTypes {
    userPageUrl: string;
    loginObj: Partial<Login>;
    onLogin: Function;
    onLogout: Function;
    service?: KLoginService;
}
