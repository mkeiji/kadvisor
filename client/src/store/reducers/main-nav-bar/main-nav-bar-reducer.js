import {MAIN_NAV_BAR_ACTION_TYPE} from "../../actions/action-types";
import {combineReducers} from "redux";

const LOGIN = (state=null, action) => {
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