export default class UserPageStore {
    mapStateToProps = (state: any) => ({
        getLoginStore: state.MAIN_NAV_BAR.LOGIN
    });
}
