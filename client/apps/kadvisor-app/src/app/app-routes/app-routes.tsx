import React, { Component } from 'react';
import ReactPlaceholder from '../react-placeholder/react-placeholder';
import './app-routes.css';
import MainNavBar from '../main-nav-bar/main-nav-bar.component';
import UserPage from '../user-page/user-page';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { APP_ROUTES } from '@client/klibs';

class AppRoutes extends Component {
    render() {
        return (
            <Router>
                <div className="Main">
                    <MainNavBar />
                    <Route
                        exact
                        path={APP_ROUTES.root}
                        component={ReactPlaceholder}
                    />
                    <Route
                        path={APP_ROUTES.about}
                        render={() => (
                            <React.Fragment>
                                <h1 style={{ paddingTop: '100px' }}>
                                    ABOUT PAGE
                                </h1>
                            </React.Fragment>
                        )}
                    />
                    <Route
                        exact
                        path={APP_ROUTES.userPage}
                        component={UserPage}
                    />
                </div>
            </Router>
        );
    }
}

export default AppRoutes;
