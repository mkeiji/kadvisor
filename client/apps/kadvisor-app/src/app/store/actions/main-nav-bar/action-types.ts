export const MAIN_NAV_BAR_ACTION_TYPE = {
    SET_LOGIN: 'SET_LOGIN',
    UNSET_LOGIN: 'UNSET_LOGIN'
} as MainNavBarActionType;

export interface MainNavBarActionType {
    SET_LOGIN: string;
    UNSET_LOGIN: string;
}
