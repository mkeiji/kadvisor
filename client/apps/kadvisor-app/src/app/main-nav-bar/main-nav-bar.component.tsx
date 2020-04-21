import React, { CSSProperties } from 'react';
import { Navbar, Nav } from 'react-bootstrap';
import MainNavBarStore from './main-nav-bar.store';
import { useHistory } from 'react-router';
import { connect } from 'react-redux';
import KLogin from '../k-login/k-login.component';
import { APP_ROUTES, KRouterPathUtil } from '@client/klibs';

function MainNavBar(props: MainNavBarPropTypes) {
    const history = useHistory();
    let state = {} as any;

    if (props.getLoginStore) {
        Object.entries(props.getLoginStore).map(
            ([_, value]) => (state = value)
        );
    }

    function processLogin(login: any) {
        props.setLoginStore(login);
        history.push(KRouterPathUtil.getUserPage(login.login.userID));
    }

    function processLogout(login: any) {
        history.push(APP_ROUTES.root);
        props.unsetLoginStore(login);
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
                        loginObj={state}
                        onLogin={processLogin}
                        onLogout={processLogout}
                    />
                </Navbar.Collapse>
            </Navbar>
        </div>
    );
}

interface MainNavBarPropTypes {
    getLoginStore: any;
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
