import {
    SET_LOGIN,
    UNSET_LOGIN
} from '@app/store/actions/main-nav-bar/main-nav-bar-action';

export default class MainNavBarStore {
    mapStateToProps = (state: any) => ({
        getLoginStore: state.MAIN_NAV_BAR.LOGIN
    });

    mapDispatchToProps = (dispatch: any) => ({
        setLoginStore(login: any) {
            dispatch(SET_LOGIN(login));
        },

        unsetLoginStore(login: any) {
            dispatch(UNSET_LOGIN(login));
        }
    });
}
