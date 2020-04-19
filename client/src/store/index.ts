import {applyMiddleware, createStore} from "redux";
import thunk from "redux-thunk";
import appReducer from "./reducers/app-reducer";

export default function CreateStore(initialState={}) {
    return applyMiddleware(thunk, middlewareLogger)(createStore)(appReducer, initialState);
}

const middlewareLogger = (store: any) => (next: any) => (action: any) => {
    let result;

    console.groupCollapsed(`dispatched action => ${action.type}`);
    result = next(action);
    console.log(`store: ${JSON.stringify(store.getState())}`);
    console.groupEnd();

    return result;
};