import React, { CSSProperties } from 'react';
import { Navbar, Nav } from 'react-bootstrap';
import MainNavBarStore from './main-nav-bar.store';
import { useHistory } from 'react-router';
import { connect } from 'react-redux';
import KLogin from '../k-login/k-login.component';
import {
    APP_ROUTES,
    KLoginResponse,
    KRouterPathUtil,
    Login
} from '@client/klibs';

function MainNavBar(props: MainNavBarPropTypes) {
    const history = useHistory();
    let login = {} as Login;

    if (props.getLoginStore) {
        login = props.getLoginStore;
    }

    function processLogin(login: KLoginResponse) {
        props.setLoginStore(login.login);
        history.push(KRouterPathUtil.getUserPage(login.login.userID));
    }

    function processLogout(login: KLoginResponse) {
        history.push(APP_ROUTES.root);
        props.unsetLoginStore(login.login);
    }

    return (
        <div>
            <Navbar fixed="top" bg="dark" variant="dark" style={navBarStyle}>
                <Navbar.Brand href="#home">Kadvisor</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="mr-auto">
                        <Nav.Link href="/">Home</Nav.Link>
                        <Nav.Link href="/about">About</Nav.Link>
                    </Nav>

                    <KLogin
                        userPageUrl={APP_ROUTES.userPage}
                        loginObj={login}
                        onLogin={processLogin}
                        onLogout={processLogout}
                    />
                </Navbar.Collapse>
            </Navbar>
        </div>
    );
}

interface MainNavBarPropTypes {
    getLoginStore: Login;
    setLoginStore: Function;
    unsetLoginStore: Function;
}

const navBarStyle = {
    position: 'static'
} as CSSProperties;

const store = new MainNavBarStore();
export default connect(
    store.mapStateToProps,
    store.mapDispatchToProps
)(MainNavBar);
