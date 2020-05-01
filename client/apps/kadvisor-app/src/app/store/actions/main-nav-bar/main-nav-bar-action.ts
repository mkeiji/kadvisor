import { MAIN_NAV_BAR_ACTION_TYPE } from './action-types';
import { Login, StoreAction } from '@client/klibs';

export const SET_LOGIN = (login: Login) =>
    ({
        type: MAIN_NAV_BAR_ACTION_TYPE.SET_LOGIN,
        payload: login
    } as StoreAction);

export const UNSET_LOGIN = (login: Login) =>
    ({
        type: MAIN_NAV_BAR_ACTION_TYPE.UNSET_LOGIN,
        payload: login
    } as StoreAction);
