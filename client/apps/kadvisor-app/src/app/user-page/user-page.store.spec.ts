import { Login } from '@client/klibs';
import { MainNavBarState } from '../main-nav-bar/main-nav-bar.store';
import UserPageStore from './user-page.store';

describe('UserPageStore', () => {
    let testStore: UserPageStore;

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
        testStore = new UserPageStore();
    });

    it('mapStateToProps', () => {
        const expected = { getLoginStore: { email: 'test', userID: 1 } };
        const ktest = testStore.mapStateToProps(testState);
        expect(ktest).toEqual(expected);
    });
});
