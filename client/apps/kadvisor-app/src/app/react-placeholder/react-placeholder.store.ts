import { setTestMessage } from '../store/actions/test/test-action';
import { TestReducersStoreObj } from '@client/klibs';

export interface ReactPlaceholderState {
    TEST_REDUCERS: TestReducersStoreObj;
}
export default class ReactPlaceholderStore {
    mapStateToProps = (state: ReactPlaceholderState) => ({
        getTestFromStore: state.TEST_REDUCERS.TEST
    });

    mapDispatchToProps = (dispatch: Function) => ({
        testStoreFunc(testMsg: string) {
            dispatch(setTestMessage(testMsg));
        }
    });
}
