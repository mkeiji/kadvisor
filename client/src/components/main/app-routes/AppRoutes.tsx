import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Component } from 'react';
import './AppRoutes.css';
import MainNavBar from '../main-nav-bar/main-nav-bar.component';
import ReactPlaceholder from '../react-placeholder/react-placeholder.component';
import UserHome from "../user-home/user-home";
import APP_ROUTES from "AppRoutes";

class AppRoutes extends Component {
    render() {
        return (
            <Router>
                <div className="Main">
                    <MainNavBar />
                    <Route exact path={APP_ROUTES.root} component={ReactPlaceholder} />
                    <Route path={APP_ROUTES.about} render={ () => (
                        <React.Fragment>
                            <h1 style={{paddingTop: '100px'}}>ABOUT PAGE</h1>
                        </React.Fragment>
                    )} />
                    <Route exact path={APP_ROUTES.userPage} component={UserHome} />
                </div>
            </Router>
        );
    }
}

export default AppRoutes;
