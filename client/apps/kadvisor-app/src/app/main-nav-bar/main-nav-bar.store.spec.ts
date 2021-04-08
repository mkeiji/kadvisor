import {
    SET_LOGIN,
    UNSET_LOGIN
} from '@app/store/actions/main-nav-bar/main-nav-bar-action';
import { Login } from '@client/klibs';
import MainNavBarStore, { MainNavBarState } from './main-nav-bar.store';

describe('MainNavBarStore', () => {
    let testStore: MainNavBarStore;

    const testLogin = ({
        userID: 1,
        email: 'test'
    } as unknown) as Login;
    const testState = ({
        MAIN_NAV_BAR: {
            LOGIN: testLogin
        }
    } as unknown) as MainNavBarState;

    beforeEach(() => {
        testStore = new MainNavBarStore();
    });

    it('mapStateToProps', () => {
        const expected = { getLoginStore: { email: 'test', userID: 1 } };
        const ktest = testStore.mapStateToProps(testState);
        expect(ktest).toEqual(expected);
    });

    it('mapDispatchToProps - setLoginStore', () => {
        const testDispatch = jest.fn();
        testStore.mapDispatchToProps(testDispatch).setLoginStore(testLogin);
        expect(testDispatch).toHaveBeenCalledWith(SET_LOGIN(testLogin));
    });

    it('mapDispatchToProps - unsetLoginStore', () => {
        const testDispatch = jest.fn();
        testStore.mapDispatchToProps(testDispatch).unsetLoginStore(testLogin);
        expect(testDispatch).toHaveBeenCalledWith(UNSET_LOGIN(testLogin));
    });
});
