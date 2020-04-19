import React from 'react';
import {Component} from 'react';
import './app-routes/AppRoutes.css';
import AppRoutes from "./app-routes/AppRoutes";
import {Provider} from "react-redux";
import CreateStore from "store/index";

/* Store Setup
* ------------*/
const initialStoreState =
    (localStorage["kadvisor-store"]) ?
    JSON.parse(localStorage["kadvisor-store"]) :
    {};
const store = CreateStore(initialStoreState);
const saveState = () => localStorage["kadvisor-store"] = JSON.stringify(store.getState());
store.subscribe(saveState);

//window.store = store; // TODO: delete/comment (debug only)
/* Store Setup -- END */

class Main extends Component {
    render() {
        return (
            <Provider store={store}>
                <AppRoutes/>
            </Provider>
        );
    }
}

export default Main;
