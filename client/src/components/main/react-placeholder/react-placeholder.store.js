import {setTestMessage} from "../../../store/actions/test/test-action";

export default class ReactPlaceholderStore {
    mapStateToProps = (state) =>
        ({
            test: state.test
        });

    mapDispatchToProps = (dispatch) =>
        ({
            testStoreFunc(testMsg) {
                dispatch(
                    setTestMessage(testMsg)
                )
            }
        });
}