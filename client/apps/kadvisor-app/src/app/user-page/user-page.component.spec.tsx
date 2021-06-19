import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import UserPage from './user-page.component';
import UserPageMenu from './user-page-menu/user-page-menu.component';
import { APP_PAGES, Login, Match } from '@client/klibs';
import Dashboard from './user-page-menu/dashboard/dashboard.component';
import EntryComponent from './user-page-menu/entry/entry.component';
import Settings from './user-page-menu/settings/settings.component';
import Reports from './user-page-menu/reports/reports.component';

describe('UserPage', () => {
    let wrapper: ShallowWrapper;
    let setWrapper: Function;
    let useEffect: any;
    let testMatchProp: Match<any>;

    const NUMBER_OF_WRAPPERS = 4;
    const TEST_USER_ID = 1;
    const testLogin = ({
        userID: TEST_USER_ID
    } as unknown) as Login;
    const mockUseEffect = () => {
        useEffect.mockImplementationOnce((f: any) => f());
    };

    beforeEach(() => {
        // mocking useEffect (needs 4 mocks since its using 'wrappedComponent')
        useEffect = jest.spyOn(React, 'useEffect');
        [...Array(NUMBER_OF_WRAPPERS)].forEach(() => {
            mockUseEffect();
        });

        testMatchProp = {
            params: {
                id: TEST_USER_ID,
                page: APP_PAGES.home
            },
            isExact: true,
            path: 'test',
            url: 'test'
        };

        setWrapper = () => {
            wrapper = shallow(
                <UserPage.WrappedComponent
                    match={testMatchProp}
                    getLoginStore={testLogin}
                />
            );
        };
    });

    describe('Renders', () => {
        it('should renders UserPageMenu component', () => {
            setWrapper();
            expect(wrapper.find(UserPageMenu)).toHaveLength(1);
        });

        describe('getPage', () => {
            it('should display please login msg', () => {
                testMatchProp.params.id = 99;
                setWrapper();
                expect(wrapper.find('h1').text()).toBe('PLEASE LOGIN');
            });

            it('should render Dashboard component', () => {
                testMatchProp.params.page = APP_PAGES.dashboard;
                setWrapper();
                expect(wrapper.find(Dashboard)).toHaveLength(1);
            });

            it('should render EntryTable component', () => {
                testMatchProp.params.page = APP_PAGES.entries;
                setWrapper();
                expect(wrapper.find(EntryComponent)).toHaveLength(1);
            });

            it('should render Settings component', () => {
                testMatchProp.params.page = APP_PAGES.settings;
                setWrapper();
                expect(wrapper.find(Settings)).toHaveLength(1);
            });

            it('should render Reports component', () => {
                testMatchProp.params.page = APP_PAGES.reports;
                setWrapper();
                expect(wrapper.find(Reports)).toHaveLength(1);
            });

            it('should display page not found msg', () => {
                testMatchProp.params.page = 'invalidPage';
                setWrapper();
                expect(wrapper.find('h1').text()).toBe('PAGE NOT FOUND');
            });
        });
    });
});
