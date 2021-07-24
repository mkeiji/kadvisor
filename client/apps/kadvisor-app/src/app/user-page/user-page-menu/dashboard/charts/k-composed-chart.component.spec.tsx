import React from 'react';
import { MonthReport, ReportsApiService } from '@client/klibs';
import { shallow, ShallowWrapper } from 'enzyme';
import { clone } from 'lodash';
import { of, throwError } from 'rxjs';
import { KComposedChartPropsType } from '../dashboard.models';
import ChartsViewModelService from './charts-view-model.service';
import KComposedChartComponent from './k-composed-chart.component';
import { ComposedChart } from 'recharts';

describe('KComposedChartComponent', () => {
    let service: ReportsApiService;
    let viewModelService: ChartsViewModelService;
    let props: KComposedChartPropsType;
    let component: KComposedChartComponent;

    const testID = 1;
    const testYear = new Date().getUTCFullYear();

    beforeEach(() => {
        service = ({
            getYtdWithForecastReport: jest.fn(() => of())
        } as unknown) as ReportsApiService;
        viewModelService = ({
            getMinBalance: jest.fn(),
            getTicksForNegativeBalance: jest.fn(),
            getEmptyMonthReport: jest.fn()
        } as unknown) as ChartsViewModelService;
        props = {
            userID: testID,
            year: testYear,
            service: service,
            viewModelService: viewModelService
        };

        component = new KComposedChartComponent(props);
    });

    describe('constructor', () => {
        it('should set provided services', () => {
            expect(component.service).toEqual(service);
            expect(component.viewModelService).toEqual(viewModelService);
        });

        it('should set new services if not provided', () => {
            const newProps = { userID: testID, year: testYear };
            const newComponent = new KComposedChartComponent(newProps);

            const expectedService = JSON.stringify(
                new ReportsApiService(newProps.userID)
            );
            const expectedViewModelService = JSON.stringify(
                new ChartsViewModelService()
            );
            const actualService = JSON.stringify(newComponent.service);
            const actualViewModelService = JSON.stringify(
                newComponent.viewModelService
            );

            expect(actualService).toEqual(expectedService);
            expect(actualViewModelService).toEqual(expectedViewModelService);
        });

        it('should set initial state', () => {
            const expected = {
                data: [],
                minDomain: 0,
                leftTicks: undefined,
                rightTicks: undefined
            };
            expect(component.state).toEqual(expected);
        });
    });

    describe('componentDidMount', () => {
        it('should call loadData', () => {
            jest.spyOn(component as any, 'loadData');
            component.componentDidMount();
            expect(component['loadData']).toHaveBeenCalled();
        });
    });

    describe('componentWillUnmount', () => {
        it('should call destroy$.next and destroy$.unsubscribe', () => {
            const nextSpy = jest.spyOn(component.destroy$, 'next');
            const unsubSpy = jest.spyOn(component.destroy$, 'unsubscribe');

            component.componentWillUnmount();
            expect(nextSpy).toHaveBeenCalled();
            expect(unsubSpy).toHaveBeenCalled();
        });
    });

    describe('componentDidUpdate', () => {
        it('should call loadData if new year input is different from previous', () => {
            jest.spyOn(component as any, 'loadData');
            const prevProps = clone(props);
            props.year = testYear + 1;

            component.componentDidUpdate(prevProps);
            expect(component['loadData']).toHaveBeenCalled();
        });
    });

    describe('render', () => {
        let wrapper: ShallowWrapper;

        beforeEach(() => {
            wrapper = shallow(
                <KComposedChartComponent
                    userID={props.userID}
                    year={props.year}
                />
            );
        });

        it('should render a ComposedChart component', () => {
            expect(wrapper.find(ComposedChart)).toHaveLength(1);
        });
    });

    describe('loadData', () => {
        const testMonthReport: MonthReport[] = [];
        const minDomainDataIndex = 1;
        const leftRightTicksIndex = 2;
        const errorIndex = 1;

        beforeEach(() => {
            for (let i = 1; i <= 12; i++) {
                testMonthReport.push({
                    year: 0,
                    month: i,
                    income: 0,
                    expense: 0,
                    balance: 0
                } as MonthReport);
            }
            jest.spyOn(
                component.service,
                'getYtdWithForecastReport'
            ).mockReturnValue(of(testMonthReport));
        });

        it('should call getYtdWithForecastReport, getTicksForNegativeBalance', () => {
            const minBalanceSpy = jest.spyOn(
                component.viewModelService,
                'getMinBalance'
            );
            const ticksForNegBalSpy = jest.spyOn(
                component.viewModelService,
                'getTicksForNegativeBalance'
            );
            component.setState = jest.fn();

            component['loadData']();
            expect(minBalanceSpy).toHaveBeenCalledWith(testMonthReport);
            expect(ticksForNegBalSpy).toHaveBeenCalledWith(testMonthReport);
        });

        it('should set minDomain state to zero if minBalance is greater than zero', () => {
            component.setState = jest.fn();
            jest.spyOn(viewModelService, 'getMinBalance').mockReturnValue(10);

            component['loadData']();
            expect(component.setState).toHaveBeenNthCalledWith(
                minDomainDataIndex,
                {
                    minDomain: 0,
                    data: testMonthReport
                }
            );
        });

        it('should set minDomain state with value if is less or equal zero', () => {
            const expected = -1;
            component.setState = jest.fn();
            jest.spyOn(viewModelService, 'getMinBalance').mockReturnValue(
                expected
            );

            component['loadData']();
            expect(component.setState).toHaveBeenNthCalledWith(
                minDomainDataIndex,
                {
                    minDomain: expected,
                    data: testMonthReport
                }
            );
        });

        it('should call getTicksForNegativeBalance and set left and right ticks', () => {
            const leftTick = [1, 2, 3, 4, 5];
            const rightTick = [6, 7, 8, 9, 10];
            component.setState = jest.fn();
            jest.spyOn(viewModelService, 'getMinBalance').mockReturnValue(10);
            jest.spyOn(
                viewModelService,
                'getTicksForNegativeBalance'
            ).mockReturnValue([leftTick, rightTick]);

            component['loadData']();
            expect(component.setState).toHaveBeenNthCalledWith(
                leftRightTicksIndex,
                {
                    leftTicks: leftTick,
                    rightTicks: rightTick
                }
            );
        });

        it('on error, should call getEmptyMonthReport and set it as the data', () => {
            component.setState = jest.fn();
            jest.spyOn(
                component.service,
                'getYtdWithForecastReport'
            ).mockReturnValue(throwError(new Error('test')));
            jest.spyOn(viewModelService, 'getEmptyMonthReport').mockReturnValue(
                testMonthReport
            );

            component['loadData']();
            expect(component.setState).toHaveBeenNthCalledWith(errorIndex, {
                data: testMonthReport
            });
        });
    });
});
