import { Login } from '@client/klibs';

export interface KLoginState {
    login: Login;
    isLoggedIn: boolean;
    hasWarning: boolean;
    userID?: number;
}

export interface KLoginPropTypes {
    userPageUrl: string;
    loginObj: Partial<Login>;
    onLogin: Function;
    onLogout: Function;
}
