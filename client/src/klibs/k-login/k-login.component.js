import React, {Component} from 'react';
import { Button, Form, FormControl } from 'react-bootstrap';
import KLoginService from 'klibs/k-login/k-login.service'
import PropTypes from 'prop-types';

class KLogin extends Component {
    /* @input */ user = this.props.user;
    /* @output */ onLoginEmitter = (event) => { this.props.onLogin(event); }
    /* @output */ onLogoutEmitter = (event) => { this.props.onLogout(event); }

    service = new KLoginService();
    state = {
        isLoggedIn: false
    };

    componentDidMount() {}

    componentWillUnmount() {
        this.service.unsubscribe();
    }

    login = () => {
        this.service.login(this.user)
            .subscribe( 
                res => {
                    this.onLoginEmitter(res);
                },
                err => {
                    console.log(err);
                }
            );

        this.setState({
            isLoggedIn: true
        });
    }

    logout = () => {
        this.service.logout(this.user)
            .subscribe( 
                res => {
                    this.onLogoutEmitter(res);
                },
                err => {
                    console.log(err);
                }
            );

        this.setState({
            isLoggedIn: false
        });
    }

    render() {
        var loginButton;
        if (this.state.isLoggedIn) {
            loginButton = <Button variant="outline-primary" onClick={this.logout}>logout</Button>;
        } else {
            loginButton = <Button variant="outline-success" onClick={this.login}>login</Button>;
        }

        return (
            <div>
                <Form inline>
                    <FormControl type="text" placeholder="email" className="mr-sm-2" />
                    <FormControl type="text" placeholder="password" className="mr-sm-2" />
                    {loginButton}
                </Form>
            </div>
        );
    }
}

KLogin.propTypes = {
    onLogin: PropTypes.func.isRequired,
    onLogout: PropTypes.func.isRequired,
    user: PropTypes.object.isRequired
}

export default KLogin;