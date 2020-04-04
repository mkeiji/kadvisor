import {combineReducers} from "redux";

const TEST = (state='test', action) =>
    (action.type === 'SET_TEST') ?
        action.payload :
        state;

const TEST_REDUCERS = combineReducers({
    TEST
});
export default TEST_REDUCERS;