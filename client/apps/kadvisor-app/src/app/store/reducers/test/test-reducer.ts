import { combineReducers } from 'redux';
import { StoreAction } from '@client/klibs';

const TEST = (state = 'test', action: StoreAction) =>
    action.type === 'SET_TEST' ? action.payload : state;

const TEST_REDUCERS = combineReducers({
    TEST
});
export default TEST_REDUCERS;
