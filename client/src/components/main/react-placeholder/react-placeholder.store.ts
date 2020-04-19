import {setTestMessage} from "../../../store/actions/test/test-action";

export default class ReactPlaceholderStore {
    mapStateToProps = (state: any) =>
        ({
            test: state.test
        });

    mapDispatchToProps = (dispatch: any) =>
        ({
            testStoreFunc(testMsg: string) {
                dispatch(
                    setTestMessage(testMsg)
                )
            }
        });
}