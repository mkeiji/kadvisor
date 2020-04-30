import React, { Component, CSSProperties } from 'react';
import { KLoginPropTypes, KLoginState } from './view-models';
import * as Yup from 'yup';
import {
    Toast,
    Form,
    DropdownButton,
    Dropdown,
    FormControl,
    Button
} from 'react-bootstrap';
import { Formik, FormikErrors } from 'formik';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import KLoginService from './k-login.service';
import { GernericErr, KLoginResponse, Login } from '@client/klibs';

class KLogin extends Component<KLoginPropTypes, KLoginState> {
    /* @input */ loginObj = this.props.loginObj;
    /* @output */ onLoginEmitter = (event: any) => {
        this.props.onLogin(event);
    };
    /* @output */ onLogoutEmitter = (event: any) => {
        this.props.onLogout(event);
    };

    service = new KLoginService();
    unsubscribe$ = new Subject<void>();

    formInitialValues = { email: '', password: '' };
    formValidationSchema = Yup.object({
        email: Yup.string()
            .email('Invalid email address')
            .required('Required'),
        password: Yup.string()
            .min(3, 'Must be more than 3 char')
            .required('Required')
    });

    constructor(readonly props: KLoginPropTypes) {
        super(props);
        this.state = {
            login: {
                email: this.loginObj.email,
                password: this.loginObj.password
            },
            isLoggedIn: this.loginObj.isLoggedIn,
            hasWarning: false
        } as KLoginState;
    }

    componentWillUnmount(): void {
        this.unsubscribe$.next();
        this.unsubscribe$.complete();
    }

    login = (loginObj: Partial<Login>) => {
        this.setState({
            login: {
                ...this.state.login,
                email: loginObj.email,
                password: loginObj.password
            }
        });

        this.service
            .login(this.state.login)
            .pipe(takeUntil(this.unsubscribe$))
            .subscribe(
                (res: KLoginResponse) => {
                    this.setState({ isLoggedIn: res.login.isLoggedIn });
                    this.onLoginEmitter(res);
                },
                (err: GernericErr) => {
                    if (err.message.includes('status code 400')) {
                        this.setState({ hasWarning: true });
                    }
                }
            );
    };

    logout = () => {
        this.service
            .logout(this.state.login)
            .pipe(takeUntil(this.unsubscribe$))
            .subscribe(
                (res: KLoginResponse) => {
                    this.setState({ isLoggedIn: res.login.isLoggedIn });
                    this.onLogoutEmitter(res);
                },
                (err: GernericErr) => {
                    console.log(err);
                }
            );
    };

    setupLoginFormControl = () => {
        let loginControl;
        if (this.state.isLoggedIn) {
            loginControl = (
                <Form inline>
                    <DropdownButton
                        variant="secondary"
                        id="dropdown-variants-secondary"
                        title={this.state.login.email}
                        style={dropdownMenuStyle}
                    >
                        <Dropdown.Item href="#/profilepage">
                            Profile
                        </Dropdown.Item>
                        <Dropdown.Divider />
                        <Dropdown.Item onClick={this.logout}>
                            logout
                        </Dropdown.Item>
                    </DropdownButton>
                </Form>
            );
        } else {
            loginControl = (
                <Formik
                    initialValues={this.formInitialValues}
                    validationSchema={this.formValidationSchema}
                    onSubmit={loginObj => this.login(loginObj)}
                >
                    {formik => (
                        <Form inline onSubmit={formik.handleSubmit}>
                            <FormControl
                                type="text"
                                placeholder="email"
                                className="mr-sm-2"
                                {...formik.getFieldProps('email')}
                            />
                            <FormControl
                                type="password"
                                className="mr-sm-2"
                                placeholder="password"
                                {...formik.getFieldProps('password')}
                            />
                            <Button
                                variant={this.getButtonVariant(formik.errors)}
                                type="submit"
                            >
                                login
                            </Button>
                        </Form>
                    )}
                </Formik>
            );
        }
        return loginControl;
    };

    getButtonVariant = (formikErrors: FormikErrors<Login>) => {
        const hasErrors = formikErrors.email || formikErrors.password;
        if (hasErrors) {
            return 'outline-danger';
        } else {
            return 'outline-success';
        }
    };

    setupWarningToast = () => {
        const toastStatus = this.state.hasWarning
            ? visibleWarningStyle
            : hiddenWarningStyle;
        return (
            <Toast
                autohide
                onClose={() => this.setState({ hasWarning: false })}
                show={this.state.hasWarning}
                delay={2000}
                style={toastStatus}
            >
                <Toast.Header>
                    <img className="rounded mr-2" alt="" />
                    <strong className="mr-auto">
                        Invalid username or password
                    </strong>
                </Toast.Header>
            </Toast>
        );
    };

    render() {
        return (
            <div>
                {this.setupLoginFormControl()}
                {this.setupWarningToast()}
            </div>
        );
    }
}

const hiddenWarningStyle = {
    visibility: 'hidden',
    position: 'fixed'
} as CSSProperties;
const visibleWarningStyle = {
    visibility: 'visible',
    position: 'fixed'
} as CSSProperties;
const dropdownMenuStyle = { zIndex: '9999' };

export default KLogin;
