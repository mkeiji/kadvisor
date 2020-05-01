import { Login } from '@client/klibs';
import { Store } from 'redux';

export interface ApplicationStore extends Store {
    TEST_REDUCERS: TestReducersStoreObj;
    MAIN_NAV_BAR: MainNavBarStoreObj;
}

export interface TestReducersStoreObj {
    TEST: string;
}

export interface MainNavBarStoreObj {
    LOGIN: Login;
}
