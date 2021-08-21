import React, { Component } from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import {
    Forecast,
    ReportsApiService,
    ForecastTableState,
    ForecastEntry,
    KSelect,
    KSpinner
} from '@client/klibs';
import { of, throwError } from 'rxjs';
import ForecastTableViewModelService from './forecast-table-view-model.service';
import ForecastTable from './forecast-table.component';
import ForecastTableService from './forecast-table.service';
import {
    ForecastTableComponentState,
    ForecastTablePropsType
} from './view-model';
import { Button, TextField } from '@material-ui/core';
import MaterialTable from 'material-table';

describe('ForecastTable', () => {
    let props: ForecastTablePropsType;
    let service: ForecastTableService;
    let reportsService: ReportsApiService;
    let viewModelService: ForecastTableViewModelService;
    let component: ForecastTable;

    const testID = 1;
    const currentYear = new Date().getFullYear();

    beforeEach(() => {
        service = ({
            getForecast: jest.fn().mockReturnValue(of({})),
            putForecastEntry: jest.fn().mockReturnValue(of({})),
            postForecast: jest.fn().mockReturnValue(of({}))
        } as unknown) as ForecastTableService;
        reportsService = ({
            getAvailableReportYears: jest.fn().mockReturnValue(of({}))
        } as unknown) as ReportsApiService;
        viewModelService = ({
            formatTableState: jest.fn(),
            parseAmounts: jest.fn(),
            createNewForecast: jest.fn(),
            handleTableStateUpdate: jest.fn(),
            handleYearMenuItemsStateUpdate: jest.fn()
        } as unknown) as ForecastTableViewModelService;
        props = {
            userID: testID,
            classes: {},
            service: service,
            reportsService: reportsService,
            viewModelService: viewModelService
        } as ForecastTablePropsType;

        component = new ForecastTable(props);
        // component.setState = jest.fn();
        jest.spyOn(component, 'setState').mockImplementation();
    });

    describe('constructor', () => {
        it('sets service if not provided in props', () => {
            const newProps = {
                userID: 1,
                classes: {}
            };
            const newComponent = new ForecastTable(newProps);
            expect(newComponent.service).not.toBeNull();
            expect(newComponent.reportsService).not.toBeNull();
            expect(newComponent.viewModelService).not.toBeNull();
        });

        it('sets default state', () => {
            const defaultYearMenuItems = [];
            const defaultLoading = true;
            const defaultHasForecast = true;
            const defaultSelectedYear = currentYear;
            const defaultForecastYear = '';
            const defaultTable = {};

            expect(component.state.yearMenuItems).toEqual(defaultYearMenuItems);
            expect(component.state.loading).toEqual(defaultLoading);
            expect(component.state.hasForecast).toEqual(defaultHasForecast);
            expect(component.state.selectedYear).toEqual(defaultSelectedYear);
            expect(component.state.forecastYear).toEqual(defaultForecastYear);
            expect(component.state.table).toEqual(defaultTable);
        });
    });

    describe('componentDidMount', () => {
        it('should load data on init', () => {
            const forecastSpy = jest.spyOn(component as any, 'getForecast');
            const reportSpy = jest.spyOn(
                component as any,
                'getAvailableReportYears'
            );

            component.componentDidMount();
            expect(forecastSpy).toHaveBeenCalled();
            expect(reportSpy).toHaveBeenCalled();
        });
    });

    describe('componentDidUpdate', () => {
        beforeEach(() => {
            jest.spyOn(component as any, 'getForecast');
            jest.spyOn(component as any, 'getAvailableReportYears');
        });

        it('should reload data if hasForecast has changed', () => {
            const prevState = {} as ForecastTableComponentState;
            component.state = {
                hasForecast: true
            } as ForecastTableComponentState;

            component.componentDidUpdate(
                {} as ForecastTablePropsType,
                prevState
            );
            expect(component['getForecast']).toHaveBeenCalled();
            expect(component['getAvailableReportYears']).toHaveBeenCalled();
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

    describe('getForecast', () => {
        it('success - should call service and update loading, year and table state', () => {
            const table = {} as ForecastTableState;
            const forecast = {
                userID: testID,
                year: currentYear,
                entries: []
            } as Forecast;
            jest.spyOn(service, 'getForecast').mockReturnValue(of(forecast));
            jest.spyOn(viewModelService, 'formatTableState').mockReturnValue(
                table
            );

            component['getForecast']();
            expect(service.getForecast).toHaveBeenCalled();
            expect(component.setState).toHaveBeenCalledWith({
                selectedYear: currentYear,
                table: table,
                loading: false
            });
        });

        it('fail - should call service and update hasForecast and loading state', () => {
            jest.spyOn(service, 'getForecast').mockReturnValue(throwError(''));

            component['getForecast']();
            expect(service.getForecast).toHaveBeenCalled();
            expect(component.setState).toHaveBeenCalledWith({
                hasForecast: false,
                loading: false
            });
        });
    });

    describe('getAvailableReportYears', () => {
        it('should call reportsService and mapYearMenuItems', () => {
            jest.spyOn(
                reportsService,
                'getAvailableReportYears'
            ).mockReturnValue(of([currentYear]));
            component['mapYearMenuItems'] = jest.fn();
            component['getAvailableReportYears']();
            expect(reportsService.getAvailableReportYears).toHaveBeenCalledWith(
                true
            );
            expect(component['mapYearMenuItems']).toHaveBeenCalledWith([
                currentYear
            ]);
        });
    });

    describe('mapYearMenuItems', () => {
        it('should map year to selectMenuItems and update state', () => {
            const expectedMenuItems = [
                {
                    value: currentYear,
                    displayValue: currentYear.toString()
                }
            ];
            const expectedStateUpdate = {
                yearMenuItems: expectedMenuItems
            };
            component['mapYearMenuItems']([currentYear]);
            expect(component.setState).toHaveBeenCalledWith(
                expectedStateUpdate
            );
        });
    });

    describe('Render & State Management', () => {
        let wrapper: ShallowWrapper;
        let mountedComponent: Component;

        beforeEach(() => {
            wrapper = shallow(
                <ForecastTable
                    userID={testID}
                    classes={{}}
                    service={service}
                    reportsService={reportsService}
                    viewModelService={viewModelService}
                />
            );
            mountedComponent = wrapper.instance();
            resetMountedComponentState();
            jest.spyOn(mountedComponent, 'setState');
        });

        describe('onEdit', () => {
            it('should call service.putForecastEntry and update state', async () => {
                const newEntry = { income: 10 } as ForecastEntry;
                const oldEntry = { income: 5 } as ForecastEntry;
                jest.spyOn(service, 'putForecastEntry').mockReturnValue(
                    of(newEntry)
                );
                jest.spyOn(viewModelService, 'parseAmounts').mockReturnValue(
                    newEntry
                );
                mountedComponent.state = {} as ForecastTableComponentState;

                await mountedComponent['onEdit'](newEntry, oldEntry);
                expect(viewModelService.parseAmounts).toHaveBeenCalledWith(
                    newEntry
                );
                expect(mountedComponent.setState).toHaveBeenCalled();
                expect(
                    viewModelService.handleTableStateUpdate
                ).toHaveBeenCalledWith({}, oldEntry, newEntry);
            });
        });

        describe('createForecast', () => {
            const newForecast = {} as Forecast;

            beforeEach(() => {
                mountedComponent['isValidYear'] = jest
                    .fn()
                    .mockReturnValue(true);
                mountedComponent.state = { forecastYear: currentYear };
                jest.spyOn(
                    viewModelService,
                    'createNewForecast'
                ).mockReturnValue(newForecast);
            });

            it('it should call viewModelService.createForecast', () => {
                mountedComponent['createForecast']();
                expect(viewModelService.createNewForecast).toHaveBeenCalledWith(
                    testID,
                    currentYear
                );
            });

            it('it should call service.postForecast', () => {
                mountedComponent['createForecast']();
                expect(service.postForecast).toHaveBeenCalledWith(newForecast);
            });

            it('it should setState and call handleYearMenuItemsStateUpdate', () => {
                const prevState = {} as ForecastTableComponentState;
                mountedComponent.state = prevState;

                mountedComponent['createForecast']();
                expect(mountedComponent.setState).toHaveBeenCalled();
                expect(
                    viewModelService.handleYearMenuItemsStateUpdate
                ).toHaveBeenCalledWith(prevState, newForecast);
            });
        });

        describe('isValidYear', () => {
            it('should return true if year is a number and has size four', () => {
                const validYear = currentYear;
                mountedComponent.state = {
                    forecastYear: validYear
                };

                const result = mountedComponent['isValidYear']();
                expect(result).toBeTruthy();
            });

            it('should return false if year is not a number', () => {
                const invalidYear = '';
                mountedComponent.state = {
                    forecastYear: invalidYear
                };

                const result = mountedComponent['isValidYear']();
                expect(result).toBeFalsy();
            });

            it('should return false if year has size higher than four', () => {
                const invalidYear = 12345;
                mountedComponent.state = {
                    forecastYear: invalidYear
                };

                const result = mountedComponent['isValidYear']();
                expect(result).toBeFalsy();
            });
        });

        describe('setTextFieldDisplayValue', () => {
            it('should update forecastYear state if value is valid', () => {
                const spy = jest.spyOn(mountedComponent, 'setState');
                const value = 1234;
                mountedComponent['setTextFieldDisplayValue'](value);
                expect(spy).toHaveBeenCalledWith({ forecastYear: value });
            });

            it('should clear forecastYear state if value is not a number', () => {
                const value = 'abc';
                const spy = jest.spyOn(mountedComponent, 'setState');
                mountedComponent['setTextFieldDisplayValue'](value);
                expect(spy).toHaveBeenCalledWith({ forecastYear: '' });
            });

            it('should clear forecastYear state if value is equal to zero', () => {
                const value = 0;
                const spy = jest.spyOn(mountedComponent, 'setState');
                mountedComponent['setTextFieldDisplayValue'](value);
                expect(spy).toHaveBeenCalledWith({ forecastYear: '' });
            });
        });

        describe('renderYearSelectionDropDown', () => {
            it('should render a KSelect component', () => {
                const element = mountedComponent[
                    'renderYearSelectionDropDown'
                ]();
                const nWrapper = shallow(element);
                expect(nWrapper.find(KSelect)).toHaveLength(1);
            });
        });

        describe('renderCreateForecast', () => {
            it('should render TextField', () => {
                const element = mountedComponent['renderCreateForecast']();
                const nWrapper = shallow(element);

                expect(nWrapper.find(TextField)).toHaveLength(1);
            });

            it('should render Button for forecast', () => {
                const expectedBtnText = '+ forecast';

                const element = mountedComponent['renderCreateForecast']();
                const nWrapper = shallow(element);

                expect(nWrapper.find(Button)).toHaveLength(1);
                expect(nWrapper.find(Button).text()).toEqual(expectedBtnText);
            });
        });

        describe('renderForecast', () => {
            it('should render MaterialTable', () => {
                const element = mountedComponent['renderForecast']();
                const nWrapper = shallow(element);

                expect(nWrapper.find(MaterialTable)).toHaveLength(1);
            });
        });

        describe('render', () => {
            let renderWrapper: ShallowWrapper;

            beforeEach(() => {
                renderWrapper = shallow(
                    <ForecastTable userID={testID} classes={{}} />
                );
            });

            it('should render KSpinner when initialized', () => {
                expect(renderWrapper.find(KSpinner)).toHaveLength(1);
            });

            it('should render MaterialTable when loading is false and hasForecast is true', () => {
                renderWrapper.setState({ loading: false, hasForecast: true });
                expect(renderWrapper.find(MaterialTable)).toHaveLength(1);
            });
        });

        function resetMountedComponentState() {
            mountedComponent.state = {
                yearMenuItems: [],
                loading: true,
                hasForecast: true,
                selectedYear: currentYear,
                forecastYear: '',
                table: {} as ForecastTableState
            };
        }
    });
});
