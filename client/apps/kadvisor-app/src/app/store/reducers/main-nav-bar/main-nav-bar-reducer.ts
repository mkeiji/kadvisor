import { combineReducers } from 'redux';
import { MAIN_NAV_BAR_ACTION_TYPE } from '../../actions/main-nav-bar/action-types';
import { StoreAction } from '@client/klibs';

const LOGIN = (state = null, action: StoreAction) => {
    switch (action.type) {
        case MAIN_NAV_BAR_ACTION_TYPE.SET_LOGIN:
            return action.payload;

        case MAIN_NAV_BAR_ACTION_TYPE.UNSET_LOGIN:
            return null;

        default:
            return state;
    }
};

const MAIN_NAV_BAR = combineReducers({
    LOGIN
});
export default MAIN_NAV_BAR;
