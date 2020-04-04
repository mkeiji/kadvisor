import React, { Component } from "react";
import {Navbar, Nav} from 'react-bootstrap';
import KLogin from 'klibs/k-login/k-login.component';
import {connect} from "react-redux";
import PropTypes from 'prop-types';
import MainNavBarStore from "./main-nav-bar.store";

class MainNavBar extends Component {
    state = {};
    constructor(props) {
        super(props);
        if (this.props.getLoginStore) {
            Object.entries(this.props.getLoginStore).map(([_,value]) =>
                this.state = value);
        }
    }

    processLogin = (login) => {
        console.log("go to LOGIN");
        this.props.setLoginStore(login);
    };

    processLogout = (login) => {
        console.log("go to HOME");
        this.props.unsetLoginStore(login);
    };

    render() {
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
                        
                        <KLogin loginObj={this.state}
                                onLogin={this.processLogin}
                                onLogout={this.processLogout}/>
                    </Navbar.Collapse>
                </Navbar>
            </div>
        );
    }
}

MainNavBar.propTypes = {
    getLoginStore: PropTypes.object,
    setLoginStore: PropTypes.func,
    unsetLoginStore: PropTypes.func
};

const store = new MainNavBarStore();
export default connect(store.mapStateToProps, store.mapDispatchToProps)(MainNavBar);