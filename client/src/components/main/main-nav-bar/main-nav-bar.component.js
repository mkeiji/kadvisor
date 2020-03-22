import React, { Component } from "react";
import { Navbar, Nav, NavDropdown } from 'react-bootstrap';
import KLogin from 'klibs/k-login/k-login.component';

class MainNavBar extends Component {
    user = {
        "email"     : "email@test.com",
        "passoword" : "secret"
    };

    processLogin = (event) => {
        console.log("go to LOGIN");
        console.log(event);
    };

    processLogout = (event) => {
        console.log("go to HOME");
        console.log(event);
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
                            <NavDropdown title="Dropdown" id="basic-nav-dropdown">
                                <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
                                <NavDropdown.Item href="#action/3.2">Another action</NavDropdown.Item>
                                <NavDropdown.Item href="#action/3.3">Something</NavDropdown.Item>
                                <NavDropdown.Divider />
                                <NavDropdown.Item href="#action/3.4">Separated link</NavDropdown.Item>
                            </NavDropdown>
                        </Nav>
                        
                        <KLogin user={this.user} 
                                onLogin={this.processLogin} 
                                onLogout={this.processLogout}/>
                    </Navbar.Collapse>
                </Navbar>
            </div>
        );
    }
}

export default MainNavBar;