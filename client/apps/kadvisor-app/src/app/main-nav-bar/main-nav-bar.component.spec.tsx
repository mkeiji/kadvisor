import React from 'react';
import { mount, ReactWrapper } from 'enzyme';
import MainNavBar from './main-nav-bar.component';
import configureStore from 'redux-mock-store';
import { Provider } from 'react-redux';
import { Login } from '@client/klibs';
import { MainNavBarState } from './main-nav-bar.store';
import { Navbar, Nav } from 'react-bootstrap';
import KLogin from '../k-login/k-login.component';
import { Router } from 'react-router-dom';
import { createMemoryHistory, MemoryHistory } from 'history';

describe('MainNavBar', () => {
    let wrapper: ReactWrapper;

    const testLogin = ({
        userID: 1,
        email: 'test'
    } as unknown) as Login;
    const testState = ({
        MAIN_NAV_BAR: {
            LOGIN: testLogin
        }
    } as unknown) as MainNavBarState;
    const mockStore = configureStore();
    const testStore = mockStore(testState);

    beforeEach(() => {
        wrapper = mount(
            <Provider store={testStore}>
                <MainNavBar />
            </Provider>
        );
    });

    describe('Init', () => {
        it('should pass a Login from props if exist', () => {
            expect(wrapper.find(KLogin).props().loginObj).toEqual(testLogin);
        });
    });

    describe('Renders', () => {
        it('should render a img tag', () => {
            expect(wrapper.find('#navBarKLogo')).toHaveLength(1);
        });

        it('should render a Navbar.Brand', () => {
            expect(wrapper.find(Navbar.Brand)).toHaveLength(1);
        });

        it('should render a Navbar.Toggle', () => {
            expect(wrapper.find(Navbar.Toggle)).toHaveLength(1);
        });

        it('should render a Navbar.Collapse', () => {
            expect(wrapper.find(Navbar.Collapse)).toHaveLength(1);
        });

        it('should render a Home link', () => {
            expect(wrapper.find(Nav.Link).first().text()).toBe('Home');
        });

        it('should render a About link', () => {
            const aboutIndex = 1;
            expect(wrapper.find(Nav.Link).at(aboutIndex).text()).toBe('About');
        });
    });

    describe('eventHandlers', () => {
        let newWrapper: ReactWrapper;
        let mockHistory: MemoryHistory;

        beforeEach(() => {
            mockHistory = createMemoryHistory();
            newWrapper = mount(
                <Provider store={testStore}>
                    <Router history={mockHistory}>
                        <MainNavBar />
                    </Router>
                </Provider>
            );
        });

        describe('processLogin', () => {
            it('should call setLogin action', () => {
                const expectedStoreAction = {
                    type: 'SET_LOGIN',
                    payload: { userID: 1, email: 'test' }
                };
                // 'simulate onLogin' and return testLogin
                newWrapper.find(KLogin).prop('onLogin')(testLogin);
                // get last performed store action
                expect(testStore.getActions().pop()).toEqual(
                    expectedStoreAction
                );
            });

            it('should call history push', () => {
                const pushSpy = jest.spyOn(mockHistory, 'push');
                const expectedRoute = '/user/1/home/dashboard';

                newWrapper.find(KLogin).prop('onLogin')(testLogin);
                expect(pushSpy).toHaveBeenCalledWith(expectedRoute);
            });
        });

        describe('processLogout', () => {
            it('should call unsetLoginStore', () => {
                const expectedStoreAction = {
                    type: 'UNSET_LOGIN',
                    payload: { userID: 1, email: 'test' }
                };
                newWrapper.find(KLogin).prop('onLogout')(testLogin);
                expect(testStore.getActions().pop()).toEqual(
                    expectedStoreAction
                );
            });

            it('should call history push', () => {
                const pushSpy = jest.spyOn(mockHistory, 'push');
                const expectedRoute = '/';

                newWrapper.find(KLogin).prop('onLogout')(testLogin);
                expect(pushSpy).toHaveBeenCalledWith(expectedRoute);
            });
        });
    });
});
