import { MAIN_NAV_BAR_ACTION_TYPE } from '../action-types';

export const SET_LOGIN = (login: any) => ({
    type: MAIN_NAV_BAR_ACTION_TYPE.SET_LOGIN,
    payload: login
});

export const UNSET_LOGIN = (login: any) => ({
    type: MAIN_NAV_BAR_ACTION_TYPE.UNSET_LOGIN,
    payload: login
});
