export default class UserHomeStore {
    mapStateToProps = (state: any) =>
        ({
            getLoginStore: state.MAIN_NAV_BAR.LOGIN
        });
}