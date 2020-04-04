import React, {Component} from 'react';
import {Button, Form, FormControl, DropdownButton, Dropdown} from 'react-bootstrap';
import KLoginService from 'klibs/k-login/k-login.service'
import PropTypes from 'prop-types';
import { Formik } from 'formik';
import * as Yup from 'yup';

class KLogin extends Component {
    /* @input */ loginObj = this.props.loginObj;
    /* @output */ onLoginEmitter = (event) => { this.props.onLogin(event); };
    /* @output */ onLogoutEmitter = (event) => { this.props.onLogout(event); };

    state = {};
    service = new KLoginService();
    formInitialValues = { email: "", password: "" };
    formValidationSchema = Yup.object({
            email: Yup.string()
                .email('Invalid email address')
                .required('Required'),
            password: Yup.string()
                .min(3, 'Must be more than 3 char')
                .required('Required'),
        });

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

    login = (loginObj) => {
        this.setState({
            login: {
                ...this.state.login,
                email: loginObj.email,
                password: loginObj.password
            }
        });

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

    setupLoginFormControl = () => {
        let loginControl;
        if (this.state.isLoggedIn) {
            loginControl = (
                <Form inline>
                    <DropdownButton variant="secondary"
                                    id="dropdown-variants-secondary"
                                    title={this.state.login.email}>
                        <Dropdown.Item href="#/profilepage">Profile</Dropdown.Item>
                        <Dropdown.Divider />
                        <Dropdown.Item onClick={this.logout}>logout</Dropdown.Item>
                    </DropdownButton>
                </Form>
            );
        } else {
            loginControl = (
                <Formik
                    initialValues={this.formInitialValues}
                    validationSchema={this.formValidationSchema}
                    onSubmit={(loginObj) => this.login(loginObj)}
                >
                    {formik => (
                        <Form inline onSubmit={formik.handleSubmit}>
                            <FormControl type="text"
                                         placeholder="email"
                                         className="mr-sm-2"
                                         {...formik.getFieldProps('email')} />
                            <FormControl type="password"
                                         className="mr-sm-2"
                                         placeholder="password"
                                         {...formik.getFieldProps('password')} />
                            <Button variant={this.getButtonVariant(formik.errors)} type="submit">login</Button>
                        </Form>
                    )}
                </Formik>
            );
        }
        return loginControl;
    };

    getButtonVariant = (formikErrors) => {
        const hasErrors = formikErrors.email || formikErrors.password;
        if (hasErrors) {
            return "outline-danger";
        } else {
            return "outline-success";
        }
    };

    render() {
        return (
            <div>
                {this.setupLoginFormControl()}
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