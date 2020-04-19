import React from "react";
import {Navbar, Nav} from 'react-bootstrap';
import MainNavBarStore from "./main-nav-bar.store";
import {useHistory} from "react-router";
import {connect} from "react-redux";
import APP_ROUTES from "AppRoutes";
import KLogin from "klibs/k-login/k-login.component";

function MainNavBar(props: MainNavBarPropTypes) {
    let state = {} as any;
    let history = useHistory();

    if (props.getLoginStore) {
        Object.entries(props.getLoginStore).map(([_,value]) =>
            state = value);
    }

    function processLogin(login: any) {
        props.setLoginStore(login);
        history.push(APP_ROUTES.userPage.replace(":id", login.login.userID));
    }

    function processLogout(login: any) {
        props.unsetLoginStore(login);
        history.push(APP_ROUTES.root);
    }

    return (
        <div>
            <Navbar fixed="top" bg="dark" variant="dark">
                <Navbar.Brand href="#home">Kadvisor</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="mr-auto">
                        <Nav.Link href="/">Home</Nav.Link>
                        <Nav.Link href="/about">About</Nav.Link>
                    </Nav>

                    <KLogin userPageUrl={APP_ROUTES.userPage}
                            loginObj={state}
                            onLogin={processLogin}
                            onLogout={processLogout}/>
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

const store = new MainNavBarStore();
export default connect(store.mapStateToProps, store.mapDispatchToProps)(MainNavBar);