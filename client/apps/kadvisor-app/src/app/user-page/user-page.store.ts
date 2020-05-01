import { MainNavBarState } from '../main-nav-bar/main-nav-bar.store';

export default class UserPageStore {
    mapStateToProps = (state: MainNavBarState) => ({
        getLoginStore: state.MAIN_NAV_BAR.LOGIN
    });
}
