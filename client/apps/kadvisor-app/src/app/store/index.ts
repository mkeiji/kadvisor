import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import appReducer from './reducers/app-reducer';
import { ApplicationStore, StoreAction } from '@client/klibs';

export default function CreateStore(initialState = {} as ApplicationStore) {
    return applyMiddleware(thunk, middlewareLogger)(createStore)(
        appReducer,
        initialState
    );
}

// Log for debug purposes
const middlewareLogger = (store: ApplicationStore) => (next: Function) => (
    action: StoreAction
) => {
    console.groupCollapsed(`dispatched action => ${action.type}`);
    const log = next(action);
    console.log(`store: ${JSON.stringify(store.getState())}`);
    console.groupEnd();

    return log;
};
