import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import appReducer from './reducers/app-reducer';

export default function CreateStore(initialState = {}) {
    return applyMiddleware(thunk, middlewareLogger)(createStore)(
        appReducer,
        initialState
    );
}

// Log for debug purposes
const middlewareLogger = (store: any) => (next: any) => (action: any) => {
    console.groupCollapsed(`dispatched action => ${action.type}`);
    const log = next(action);
    console.log(`store: ${JSON.stringify(store.getState())}`);
    console.groupEnd();

    return log;
};
