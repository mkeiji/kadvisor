import {TEST_ACTION_TYPE} from "../action-types";

export const setTestMessage = (testMessage) =>
    ({
        type: TEST_ACTION_TYPE.SET_TEST,
        payload: testMessage
    });