import React from 'react';
import { Provider } from 'react-redux';
import AppRoutes from './app-routes/app-routes';
import CreateStore from './store';
import { ApplicationStore } from '@client/klibs';

/* Store Setup
 * ------------*/
const initialStoreState = localStorage['kadvisor-store']
    ? (JSON.parse(localStorage['kadvisor-store']) as ApplicationStore)
    : ({} as ApplicationStore);
const store = CreateStore(initialStoreState);
const saveState = () =>
    (localStorage['kadvisor-store'] = JSON.stringify(store.getState()));
store.subscribe(saveState);

//window.store = store; //(debug only)
/* Store Setup -- END */

export const App = () => {
    return (
        <Provider store={store}>
            <AppRoutes />
        </Provider>
    );
};

export default App;
