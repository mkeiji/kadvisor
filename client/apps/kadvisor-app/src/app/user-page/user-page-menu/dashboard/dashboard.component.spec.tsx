import React from 'react';
import { KSelect, ReportsApiService } from '@client/klibs';
import { shallow, ShallowWrapper } from 'enzyme';
import { of } from 'rxjs';
import Dashboard from './dashboard.component';
import { DashboardPropsType } from './dashboard.models';
import KComposedChartComponent from './charts/k-composed-chart.component';
import BalanceCard from './balance-card/balance-card.component';
import DashboardEntriesComponent from './entries/dash-entries.component';

describe('Dashboard', () => {
    let component: Dashboard;
    let props: DashboardPropsType;
    let service: ReportsApiService;

    const testID = 1;
    const testClasses = {};
    const currentYear = new Date().getFullYear();

    beforeEach(() => {
        service = {
            getYtdWithForecastReport: jest.fn(),
            getAvailableReportYears: jest.fn()
        } as ReportsApiService;

        props = {
            userID: testID,
            classes: testClasses,
            reportsApiService: service
        };

        component = new Dashboard(props);
        component.setState = jest.fn();
    });

    describe('constructor', () => {
        it('sets service if not provided in props', () => {
            const newProps = {
                userID: 1,
                classes: {}
            };
            const newComponent = new Dashboard(newProps);
            expect(newComponent.service).not.toBeNull();
        });

        it('sets default state', () => {
            const defaultGraphYear = currentYear;
            const defaultShowYearDropdown = false;
            const defaultYearMenuItems = [];

            expect(component.state.graphYear).toEqual(defaultGraphYear);
            expect(component.state.showYearDropdown).toEqual(
                defaultShowYearDropdown
            );
            expect(component.state.yearMenuItems).toEqual(defaultYearMenuItems);
        });
    });

    describe('componentDidMount', () => {
        it('call service.getAvailableReportYears', () => {
            jest.spyOn(service, 'getAvailableReportYears').mockReturnValue(
                of([])
            );
            component.componentDidMount();
            expect(service.getAvailableReportYears).toHaveBeenCalled();
        });

        it('update component state', () => {
            const expectedMenuItems = [
                {
                    value: currentYear,
                    displayValue: String(currentYear)
                }
            ];
            jest.spyOn(service, 'getAvailableReportYears').mockReturnValue(
                of([currentYear])
            );

            component.componentDidMount();
            expect(component.setState).toHaveBeenCalledTimes(2);
            expect(component.setState).toHaveBeenNthCalledWith(1, {
                yearMenuItems: expectedMenuItems,
                showYearDropdown: true
            });
            expect(component.setState).toHaveBeenNthCalledWith(2, {
                graphYear: currentYear
            });
        });
    });

    describe('componentWillUnmount', () => {
        it('should unsubscribe', () => {
            const nextSpy = jest.spyOn(component.destroy$, 'next');
            const destroySpy = jest.spyOn(component.destroy$, 'unsubscribe');

            component.componentWillUnmount();
            expect(nextSpy).toHaveBeenCalled();
            expect(destroySpy).toHaveBeenCalled();
        });
    });

    describe('renderDropdown', () => {
        let newComponent: ShallowWrapper;

        beforeEach(() => {
            newComponent = shallow(
                <Dashboard userID={testID} classes={testClasses} />
            );
        });

        it('should render a KSelect component', () => {
            newComponent.setState({ showYearDropdown: true });
            expect(newComponent.find(KSelect)).toHaveLength(1);
        });

        it('should not render a KSelect component', () => {
            newComponent.setState({ showYearDropdown: false });
            expect(newComponent.find(KSelect)).toHaveLength(0);
        });
    });

    describe('render', () => {
        let wrapper: ShallowWrapper;

        beforeEach(() => {
            wrapper = shallow(
                <Dashboard userID={testID} classes={testClasses} />
            );
        });

        it('should render KComposedChartComponent', () => {
            expect(wrapper.find(KComposedChartComponent)).toHaveLength(1);
        });

        it('should render BalanceCard', () => {
            expect(wrapper.find(BalanceCard)).toHaveLength(1);
        });

        it('should render DashboardEntries', () => {
            expect(wrapper.find(DashboardEntriesComponent)).toHaveLength(1);
        });
    });
});
