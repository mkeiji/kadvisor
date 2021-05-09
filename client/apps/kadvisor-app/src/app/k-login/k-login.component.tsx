import React, { Component, CSSProperties } from 'react';
import { KLoginFormType, KLoginPropTypes, KLoginState } from './view-models';
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
import {
    GernericErr,
    Login,
    AuthError,
    AuthSuccess,
    Auth
} from '@client/klibs';

class KLogin extends Component<KLoginPropTypes, KLoginState> {
    private readonly TOKEN_KEY = 'token';

    /* @input */
    loginObj = this.props.loginObj;

    /* @output */
    onLoginEmitter = (event: Login) => {
        this.props.onLogin(event);
    };

    /* @output */
    onLogoutEmitter = (event: Login) => {
        this.props.onLogout(event);
    };

    service: KLoginService;
    unsubscribe$: Subject<void>;
    formInitialValues: KLoginFormType;
    formValidationSchema: Yup.ObjectSchema;

    constructor(readonly props: KLoginPropTypes) {
        super(props);
        const isLoggedIn = this.loginObj.isLoggedIn
            ? this.loginObj.isLoggedIn
            : false;
        this.service = props.service ? props.service : new KLoginService();
        this.unsubscribe$ = new Subject<void>();
        this.state = {
            login: {
                email: this.loginObj.email,
                password: this.loginObj.password
            },
            isLoggedIn: isLoggedIn,
            hasWarning: false
        } as KLoginState;
    }

    componentWillMount(): void {
        this.formInitialValues = { email: '', password: '' };
        this.formValidationSchema = Yup.object({
            email: Yup.string()
                .email('Invalid email address')
                .required('Required'),
            password: Yup.string()
                .min(3, 'Must be more than 3 char')
                .required('Required')
        });
    }

    componentWillUnmount(): void {
        this.unsubscribe$.next();
        this.unsubscribe$.complete();
    }

    responseIsError(res: Auth): res is AuthError {
        const obj = res as AuthError;
        return obj.code !== undefined;
    }

    login = async (loginObj: Partial<Login>) => {
        this.setState({
            login: {
                ...this.state.login,
                email: loginObj.email,
                password: loginObj.password
            }
        });

        const tokenResponse = await this.service.getToken(this.state.login);
        if (this.responseIsError(tokenResponse)) {
            this.setState({ hasWarning: true });
        } else {
            const auth = tokenResponse.data as AuthSuccess;
            localStorage.setItem(this.TOKEN_KEY, auth.token);

            this.service
                .login(this.state.login)
                .pipe(takeUntil(this.unsubscribe$))
                .subscribe(
                    (res: Login) => {
                        this.setState({ isLoggedIn: res.isLoggedIn });
                        this.onLoginEmitter(res);
                    },
                    (err: GernericErr) => {
                        if (err.message.includes('status code 400')) {
                            this.setState({ hasWarning: true });
                        }
                    }
                );
        }
    };

    logout = () => {
        this.service
            .logout(this.state.login)
            .pipe(takeUntil(this.unsubscribe$))
            .subscribe(
                (res: Login) => {
                    this.setState({ isLoggedIn: res.isLoggedIn });
                    this.onLogoutEmitter(res);
                    localStorage.removeItem(this.TOKEN_KEY);
                },
                (err: GernericErr) => {
                    console.log(err);
                }
            );
    };

    setupLoginFormControl = () => {
        let loginControl: JSX.Element;
        if (this.state.isLoggedIn) {
            loginControl = (
                <Form inline>
                    <DropdownButton
                        variant="secondary"
                        id="dropdown-variants-secondary"
                        title={this.state.login.email}
                        style={dropdownMenuStyle as CSSProperties}
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
                    onSubmit={(loginObj) => this.login(loginObj)}
                >
                    {(formik) => (
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
