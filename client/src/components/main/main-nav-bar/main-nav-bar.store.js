import {SET_LOGIN, UNSET_LOGIN} from "../../../store/actions/main-nav-bar/main-nav-bar-action";

export default class MainNavBarStore {
    mapStateToProps = (state) =>
        ({
            getLoginStore: state.MAIN_NAV_BAR.LOGIN
        });

    mapDispatchToProps = (dispatch) =>
        ({
            setLoginStore(login) {
                dispatch(
                    SET_LOGIN(login)
                )
            },

            unsetLoginStore(login) {
                dispatch(
                    UNSET_LOGIN(login)
                )
            }
        });
}