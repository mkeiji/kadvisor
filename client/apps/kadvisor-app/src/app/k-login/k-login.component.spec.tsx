import { Auth, AuthError, Login } from '@client/klibs';
import KLogin from './k-login.component';
import { KLoginPropTypes } from './view-models';
import { of, Subject, throwError } from 'rxjs';
import KLoginService from './k-login.service';
import { shallow, ShallowWrapper } from 'enzyme';
import {
    Dropdown,
    DropdownButton,
    Form,
    FormControl,
    Button,
    Toast
} from 'react-bootstrap';
import { FormikErrors } from 'formik';
import React from 'react';

describe('KLogin', () => {
    let props: KLoginPropTypes;
    let component: KLogin;
    let mockService: KLoginService;

    const testPageUrl = 'testPageUrl';
    const testEmail = 'test@Email';
    const testPwd = 'testPwd';
    const testLogin = ({
        email: testEmail,
        password: testPwd
    } as unknown) as Login;

    beforeEach(() => {
        mockService = ({
            getToken: jest.fn().mockReturnValue({} as Auth),
            login: jest.fn(),
            logout: jest.fn()
        } as unknown) as KLoginService;

        props = {
            userPageUrl: testPageUrl,
            loginObj: testLogin,
            onLogin: jest.fn(),
            onLogout: jest.fn(),
            service: mockService
        };

        component = new KLogin(props);
        component.setState = jest.fn();
        component.onLoginEmitter = jest.fn();
        component.onLogoutEmitter = jest.fn();
    });

    describe('constructor', () => {
        it('should set: service, unsubscribe$ and initial state', () => {
            props.loginObj.isLoggedIn = true;

            const newComponent = new KLogin(props);

            expect(newComponent.service).not.toBeNull();
            expect(newComponent.unsubscribe$).not.toBe(undefined);
            expect(newComponent.state.login.email).toEqual(testEmail);
            expect(newComponent.state.login.password).toEqual(testPwd);
            expect(newComponent.state.isLoggedIn).toBeTruthy();
            expect(newComponent.state.hasWarning).toBeFalsy();
        });
    });

    describe('componentWillMount', () => {
        it('should set form initial values', () => {
            const expected = { email: '', password: '' };
            component.componentWillMount();
            expect(component.formInitialValues).toEqual(expected);
        });

        it('should set form validation schema', () => {
            component.componentWillMount();
            expect(component.formValidationSchema).not.toBeNull();
        });
    });

    describe('componentWillUnmount', () => {
        it('should call subject next and complete', () => {
            component.unsubscribe$ = ({
                next: jest.fn(),
                complete: jest.fn()
            } as unknown) as Subject<void>;

            component.componentWillUnmount();

            expect(component.unsubscribe$.next).toHaveBeenCalled();
            expect(component.unsubscribe$.complete).toHaveBeenCalled();
        });
    });

    describe('responseIsError', () => {
        it('should return true if response is of type AuthError', () => {
            const testResponse = ({
                code: 500
            } as unknown) as AuthError;
            const result = component.responseIsError(testResponse);
            expect(result).toBeTruthy();
        });

        it('should return false if response is not of type AuthError', () => {
            const testResponse = ({} as unknown) as Auth;
            const result = component.responseIsError(testResponse);
            expect(result).toBeFalsy();
        });
    });

    describe('login', () => {
        const secondCallIndex = 2;
        const testToken = 'test';
        const tokenResponse = Promise.resolve({
            data: {
                token: testToken
            }
        });

        it('should update component state', async () => {
            jest.spyOn(component, 'responseIsError').mockReturnValue(true);
            const newEmail = 'newEmail';
            const newPwd = 'newPwd';
            testLogin.email = newEmail;
            testLogin.password = newPwd;

            await component.login(testLogin);

            expect(component.setState).toHaveBeenCalledWith(
                expect.objectContaining({
                    login: {
                        email: newEmail,
                        password: newPwd
                    }
                })
            );
        });

        it('should call service getToken', async () => {
            jest.spyOn(component, 'responseIsError').mockReturnValue(true);
            await component.login(testLogin);
            expect(component.service.getToken).toHaveBeenCalledWith(
                component.state.login
            );
        });

        it('should update component state with hasWarning equals true if response is an error', async () => {
            const expectedStateUpdate = { hasWarning: true };
            jest.spyOn(component, 'responseIsError').mockReturnValue(true);

            await component.login(testLogin);

            expect(component.setState).toHaveBeenCalledTimes(2);
            expect(component.setState).toHaveBeenNthCalledWith(
                secondCallIndex,
                expectedStateUpdate
            );
        });

        it('should call localStorage.setItem if response is not error', async () => {
            const testTokenKey = 'token';
            jest.spyOn(mockService, 'login').mockReturnValue(of());
            jest.spyOn(mockService, 'getToken').mockReturnValue(tokenResponse);
            Storage.prototype.setItem = jest.fn();

            await component.login(testLogin);

            expect(localStorage.setItem).toHaveBeenCalledWith(
                testTokenKey,
                testToken
            );
        });

        it(
            'should call service.login, update state isLoggedIn equals ' +
                'true and call onLoginEmitter if response is ok',
            async () => {
                const expectedStateUpdate = { isLoggedIn: true };
                const loginResponse = ({
                    isLoggedIn: true
                } as unknown) as Login;
                jest.spyOn(mockService, 'login').mockReturnValue(
                    of(loginResponse)
                );
                jest.spyOn(mockService, 'getToken').mockReturnValue(
                    tokenResponse
                );
                Storage.prototype.setItem = jest.fn();

                await component.login(testLogin);

                expect(component.setState).toHaveBeenNthCalledWith(
                    secondCallIndex,
                    expectedStateUpdate
                );
                expect(component.onLoginEmitter).toHaveBeenCalledWith(
                    loginResponse
                );
            }
        );

        it(
            'should call service.login and update state with hasWarning equals ' +
                'true if response is an error',
            async () => {
                const expectedStateUpdate = { hasWarning: true };
                jest.spyOn(mockService, 'login').mockReturnValue(
                    throwError({
                        message: 'status code 400'
                    })
                );
                jest.spyOn(mockService, 'getToken').mockReturnValue(
                    tokenResponse
                );
                Storage.prototype.setItem = jest.fn();

                await component.login(testLogin);

                expect(component.setState).toHaveBeenNthCalledWith(
                    secondCallIndex,
                    expectedStateUpdate
                );
            }
        );
    });

    describe('logout', () => {
        it(
            'should update isLoggedIn state to false, call onLogoutEmitter ' +
                'localStorage',
            () => {
                const testTokenKey = 'token';
                const expectedStateUpdate = { isLoggedIn: false };
                const logoutResponse = ({
                    isLoggedIn: false
                } as unknown) as Login;
                jest.spyOn(mockService, 'logout').mockReturnValue(
                    of(logoutResponse)
                );
                Storage.prototype.removeItem = jest.fn();

                component.logout();

                expect(component.setState).toHaveBeenCalledWith(
                    expectedStateUpdate
                );
                expect(component.onLogoutEmitter).toHaveBeenCalledWith(
                    logoutResponse
                );
                expect(localStorage.removeItem).toHaveBeenCalledWith(
                    testTokenKey
                );
            }
        );
    });

    describe('setupLoginFormControl', () => {
        describe('isLoggedIn equals true', () => {
            let newComponent: KLogin;
            let loginControl: JSX.Element;
            let result: ShallowWrapper;

            beforeEach(() => {
                props.loginObj.isLoggedIn = true;
                props.loginObj.email = testEmail;

                newComponent = new KLogin(props);
                newComponent.setState = jest.fn();

                loginControl = newComponent.setupLoginFormControl();
                result = shallow(loginControl);
            });

            it('should display dropdown button', () => {
                expect(result.find(DropdownButton)).toHaveLength(1);
            });

            it('should display dropdown button variant as secondary', () => {
                expect(result.find(DropdownButton).prop('variant')).toEqual(
                    'secondary'
                );
            });

            it('should display dropdown button title with login email', () => {
                expect(result.find(DropdownButton).prop('title')).toEqual(
                    testEmail
                );
            });

            it('should display dropdown button with an item called Profile', () => {
                expect(result.find(Dropdown.Item).first().text()).toEqual(
                    'Profile'
                );
            });

            it('should display dropdown button containing a divider', () => {
                expect(result.find(Dropdown.Divider)).toHaveLength(1);
            });

            it('should display dropdown button with an item called logout', () => {
                expect(result.find(Dropdown.Item).last().text()).toEqual(
                    'logout'
                );
            });

            it('should call service.logout if logout item is clicked', () => {
                const logoutResponse = ({
                    isLoggedIn: false
                } as unknown) as Login;
                jest.spyOn(mockService, 'logout').mockReturnValue(
                    of(logoutResponse)
                );

                result.find(Dropdown.Item).last().simulate('click');
                expect(mockService.logout).toHaveBeenCalled();
            });
        });

        describe('isLoggedIn equals false', () => {
            let newComponent: KLogin;
            let loginControl: JSX.Element;
            let result: ShallowWrapper;

            beforeEach(() => {
                props.loginObj.isLoggedIn = false;
                newComponent = new KLogin(props);

                loginControl = newComponent.setupLoginFormControl();
                result = shallow(loginControl);
            });

            it('should display login form', () => {
                expect(result.find(Form)).toHaveLength(1);
                expect(result.find(FormControl)).toHaveLength(2);
                expect(result.find(FormControl).first().prop('type')).toEqual(
                    'text'
                );
                expect(result.find(FormControl).last().prop('type')).toEqual(
                    'password'
                );
                expect(result.find(Button)).toHaveLength(1);
                expect(result.find(Button).prop('type')).toEqual('submit');
            });
        });
    });

    describe('getButtonVariant', () => {
        it('should return outline-danger if form has errors', () => {
            const err = ({ email: 'err' } as unknown) as FormikErrors<Login>;
            const result = component.getButtonVariant(err);
            expect(result).toEqual('outline-danger');
        });

        it('should return outline-success if form has no errors', () => {
            const err = ({} as unknown) as FormikErrors<Login>;
            const result = component.getButtonVariant(err);
            expect(result).toEqual('outline-success');
        });
    });

    describe('setupWarningToast', () => {
        let newComponent: ShallowWrapper;

        beforeEach(() => {
            newComponent = shallow(
                <KLogin
                    userPageUrl={''}
                    loginObj={props.loginObj}
                    onLogin={props.onLogin}
                    onLogout={props.onLogout}
                />
            );
        });

        it('should NOT show warning toast if component has no warnings', () => {
            expect(newComponent.find(Toast)).toHaveLength(1);
            expect(newComponent.find(Toast).prop('show')).toBeFalsy();
        });

        it('should show warning toast if component has warnings', () => {
            newComponent.setState({ hasWarning: true });
            expect(newComponent.find(Toast)).toHaveLength(1);
            expect(newComponent.find(Toast).prop('show')).toBeTruthy();
        });
    });

    describe('render', () => {
        it('should call setupLoginFormControl and setupWarningToast', () => {
            jest.spyOn(component, 'setupLoginFormControl');
            jest.spyOn(component, 'setupWarningToast');

            component.render();
            expect(component.setupLoginFormControl).toHaveBeenCalled();
            expect(component.setupWarningToast).toHaveBeenCalled();
        });
    });
});
