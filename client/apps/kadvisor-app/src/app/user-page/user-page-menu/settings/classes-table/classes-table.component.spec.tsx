import React, { Component } from 'react';
import { Class, KSpinner } from '@client/klibs';
import { shallow, ShallowWrapper } from 'enzyme';
import { of } from 'rxjs';
import ClassTableViewModelService from './class-table-view-model.service';
import ClassTableService from './class-table.service';
import ClassesTable from './classes-table.component';
import { ClassesTablePropsType, ClassTableState } from './view-model';
import MaterialTable from 'material-table';

describe('ClassesTable', () => {
    let props: ClassesTablePropsType;
    let service: ClassTableService;
    let viewModelService: ClassTableViewModelService;
    let component: ClassesTable;

    const testID = 1;
    const testUserID = 2;
    const testname = 'testname';
    const testClass = ({
        id: testID,
        userID: testUserID,
        name: testname
    } as unknown) as Class;

    beforeEach(() => {
        service = ({
            getClasses: jest.fn(),
            postClass: jest.fn(),
            putClass: jest.fn(),
            deleteClass: jest.fn()
        } as unknown) as ClassTableService;
        viewModelService = ({
            mapClassesToClassTableState: jest.fn()
        } as unknown) as ClassTableViewModelService;
        props = ({
            userID: testID,
            classes: {},
            service: service,
            viewModelService: viewModelService
        } as unknown) as ClassesTablePropsType;

        component = new ClassesTable(props);
        component.setState = jest.fn();
        jest.spyOn(service, 'getClasses').mockReturnValue(of([testClass]));
    });

    describe('constructor', () => {
        it('sets service if not provided in props', () => {
            const newProps = {
                userID: 1,
                classes: {}
            };
            const newComponent = new ClassesTable(newProps);
            expect(newComponent.service).not.toBeNull();
        });

        it('sets default state', () => {
            const defaultLoading = true;
            const defaultTable = {};

            expect(component.state.loading).toEqual(defaultLoading);
            expect(component.state.table).toEqual(defaultTable);
        });
    });

    describe('componentDidMount', () => {
        it('call service.getClasses', () => {
            component.componentDidMount();
            expect(service.getClasses).toHaveBeenCalled();
        });

        it('call viewModelService.mapClassesToClassTableState and update state', () => {
            const expectedTable = {} as ClassTableState;
            jest.spyOn(
                viewModelService,
                'mapClassesToClassTableState'
            ).mockReturnValue(expectedTable);
            component.componentDidMount();
            expect(
                viewModelService.mapClassesToClassTableState
            ).toHaveBeenCalled();
            expect(component.setState).toHaveBeenCalledWith({
                table: expectedTable,
                loading: false
            });
        });
    });

    describe('componentWillUnmount', () => {
        it('should unsubscribe', () => {
            const nextSpy = jest.spyOn(component.unsubscribe$, 'next');
            const destroySpy = jest.spyOn(
                component.unsubscribe$,
                'unsubscribe'
            );

            component.componentWillUnmount();
            expect(nextSpy).toHaveBeenCalled();
            expect(destroySpy).toHaveBeenCalled();
        });
    });

    describe('component state', () => {
        let mountedComponent: Component;

        beforeEach(() => {
            mountedComponent = getMountedComponent();
            jest.spyOn(mountedComponent as any, 'handleTableStateClassUpdate');
        });

        describe('onAdd', () => {
            beforeEach(() => {
                jest.spyOn(service, 'postClass').mockReturnValue(of(testClass));
            });

            it('should call service.postClass', async () => {
                await mountedComponent['onAdd'](testClass);
                expect(service.postClass).toHaveBeenCalledWith(testClass);
            });

            it('should update state', async () => {
                const expectedNewData = [testClass];
                const expectedPrevState = {
                    loading: false,
                    table: undefined
                };

                await mountedComponent['onAdd'](testClass);
                expect(
                    mountedComponent['handleTableStateClassUpdate']
                ).toHaveBeenCalledWith(expectedPrevState, expectedNewData);
            });
        });

        describe('onEdit', () => {
            const newClass = ({
                userID: testUserID,
                name: 'new name'
            } as unknown) as Class;

            beforeEach(() => {
                jest.spyOn(service, 'putClass').mockReturnValue(of(newClass));
            });

            it('should call service.putClass', async () => {
                await mountedComponent['onEdit'](newClass, testClass);
                expect(service.putClass).toHaveBeenCalledWith(newClass);
            });

            it('should update state', async () => {
                const expectedNewData = [newClass];
                const expectedPrevState = {
                    loading: false,
                    table: {
                        data: [testClass]
                    }
                };
                mountedComponent.state = expectedPrevState;

                await mountedComponent['onEdit'](newClass, testClass);
                expect(
                    mountedComponent['handleTableStateClassUpdate']
                ).toHaveBeenCalledWith(expectedPrevState, expectedNewData);
            });
        });

        describe('onDelete', () => {
            beforeEach(() => {
                jest.spyOn(service, 'deleteClass').mockReturnValue(
                    of(testClass)
                );
            });

            it('should call service.putClass', async () => {
                await mountedComponent['onDelete'](testClass);
                expect(service.deleteClass).toHaveBeenCalledWith(testClass.id);
            });

            it('should update state', async () => {
                const expectedNewData = [];
                const expectedPrevState = {
                    loading: false,
                    table: {
                        data: [testClass]
                    }
                };
                mountedComponent.state = expectedPrevState;

                await mountedComponent['onDelete'](testClass);
                expect(
                    mountedComponent['handleTableStateClassUpdate']
                ).toHaveBeenCalledWith(expectedPrevState, expectedNewData);
            });
        });
    });

    describe('render', () => {
        let wrapper: ShallowWrapper;

        beforeEach(() => {
            wrapper = getWrapperComponent();
        });

        it('should render KSpinner', () => {
            wrapper.setState({
                loading: true
            });
            expect(wrapper.find(KSpinner)).toHaveLength(1);
        });

        it('should render MaterialTable', () => {
            wrapper.setState({
                loading: false,
                table: {
                    columns: [
                        { title: 'Classname', field: 'name' },
                        { title: 'Description', field: 'description' }
                    ],
                    data: [testClass]
                }
            });
            expect(wrapper.find(MaterialTable)).toHaveLength(1);
        });
    });

    function getMountedComponent(): Component {
        const wrapper = getWrapperComponent();
        return wrapper.instance();
    }

    function getWrapperComponent(): ShallowWrapper {
        return shallow(
            <ClassesTable
                userID={testUserID}
                classes={{}}
                service={service}
                viewModelService={viewModelService}
            />
        );
    }
});
