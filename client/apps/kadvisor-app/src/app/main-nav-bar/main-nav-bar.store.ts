import {
    SET_LOGIN,
    UNSET_LOGIN
} from '@app/store/actions/main-nav-bar/main-nav-bar-action';
import { Login, MainNavBarStoreObj } from '@client/klibs';

export interface MainNavBarState {
    MAIN_NAV_BAR: MainNavBarStoreObj;
}

export default class MainNavBarStore {
    mapStateToProps = (state: MainNavBarState) => ({
        getLoginStore: state.MAIN_NAV_BAR.LOGIN
    });

    mapDispatchToProps = (dispatch: Function) => ({
        setLoginStore(login: Login) {
            dispatch(SET_LOGIN(login));
        },

        unsetLoginStore(login: Login) {
            dispatch(UNSET_LOGIN(login));
        }
    });
}
