import React, {Component} from 'react';
import {Button, Form, FormControl, DropdownButton, Dropdown} from 'react-bootstrap';
import KLoginService from 'klibs/k-login/k-login.service'
import PropTypes from 'prop-types';

class KLogin extends Component {
    /* @input */ loginObj = this.props.loginObj;
    /* @output */ onLoginEmitter = (event) => { this.props.onLogin(event); };
    /* @output */ onLogoutEmitter = (event) => { this.props.onLogout(event); };

    service = new KLoginService();
    state = {};

    constructor(props) {
        super(props);
        this.state = {
            login: {
                email: this.loginObj.email,
                password: this.loginObj.password
            },
            isLoggedIn: this.loginObj.isLoggedIn
        };
    }

    componentWillUnmount() {
        this.service.unsubscribe();
    }

    login = () => {
        this.service.login(this.state.login)
            .subscribe(
                res => {
                    this.setState({isLoggedIn: res.login.isLoggedIn});
                    this.onLoginEmitter(res);
                },
                err => {
                    console.log(err);
                }
            );
    };

    logout = () => {
        this.service.logout(this.state.login)
            .subscribe( 
                res => {
                    this.setState({isLoggedIn: res.login.isLoggedIn});
                    this.onLogoutEmitter(res);
                },
                err => {
                    console.log(err);
                }
            );
    };

    getEmail = (event) => {
        const email = event.currentTarget.value;
        this.setState({
            login: {
                ...this.state.login,
                email: email
            }
        });
    };

    getPassword = (event) => {
        const password = event.currentTarget.value;
        this.setState({
            login: {
                ...this.state.login,
                password: password
            }
        });
    };

    setupLoginControl = () => {
        let loginControl;
        if (this.state.isLoggedIn) {
            loginControl = (
                <DropdownButton variant="secondary"
                                id="dropdown-variants-secondary"
                                title={this.state.login.email}>
                    <Dropdown.Item href="#/profilepage">Profile</Dropdown.Item>
                    <Dropdown.Divider />
                    <Dropdown.Item onClick={this.logout}>logout</Dropdown.Item>
                </DropdownButton>
            );
        } else {
            loginControl = (
                <div>
                    <FormControl type="text"
                                 placeholder="email"
                                 className="mr-sm-2"
                                 onChange={this.getEmail} />
                    <FormControl type="password"
                                 className="mr-sm-2"
                                 placeholder="password"
                                 onChange={this.getPassword}/>
                    <Button variant="outline-success" onClick={this.login}>login</Button>
                </div>
            );
        }
        return loginControl;
    };

    render() {
        const loginControl = this.setupLoginControl();
        return (
            <div>
                <Form inline>
                    {loginControl}
                </Form>
            </div>
        );
    }
}

KLogin.propTypes = {
    loginObj: PropTypes.object.isRequired,
    onLogin: PropTypes.func.isRequired,
    onLogout: PropTypes.func.isRequired
};

export default KLogin;