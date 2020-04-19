import {combineReducers} from "redux";
import MAIN_NAV_BAR from "./main-nav-bar/main-nav-bar-reducer";
import TEST_REDUCERS from "./test/test-reducer";

const appReducer = combineReducers({
    TEST_REDUCERS,
    MAIN_NAV_BAR
});
export default appReducer;
