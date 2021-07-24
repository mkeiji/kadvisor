import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import { of } from 'rxjs';
import EntryService from '../../entry/entry.service';
import { DashboardEntriesPropsType } from '../dashboard.models';
import DashEntriesViewModelService from './dash-entries-view-model.service';
import DashboardEntriesComponent from './dash-entries.component';
import { DashEntryRow } from './view-model';
import { Link, Table } from '@material-ui/core';

describe('DashboardEntries', () => {
    let service: EntryService;
    let viewModelService: DashEntriesViewModelService;
    let props: DashboardEntriesPropsType;
    let component: DashboardEntriesComponent;

    const testID = 1;

    beforeEach(() => {
        service = ({
            getEntryLookups: jest.fn(() => of([])),
            getClasses: jest.fn(() => of([])),
            getEntries: jest.fn(() => of([]))
        } as unknown) as EntryService;
        viewModelService = ({
            formatDashboardRowEntries: jest.fn()
        } as unknown) as DashEntriesViewModelService;
        props = {
            userID: testID,
            service: service,
            viewModelService: viewModelService
        } as DashboardEntriesPropsType;
        component = new DashboardEntriesComponent(props);
    });

    describe('constructor', () => {
        it('should set new services if not provided', () => {
            const newProps = { userID: testID };
            const newComponent = new DashboardEntriesComponent(newProps);

            const expectedService = JSON.stringify(
                new EntryService(newProps.userID)
            );
            const expectedViewModelService = JSON.stringify(
                new DashEntriesViewModelService()
            );
            const actualService = JSON.stringify(newComponent.service);
            const actualViewModelService = JSON.stringify(
                newComponent.viewModelService
            );

            expect(actualService).toEqual(expectedService);
            expect(actualViewModelService).toEqual(expectedViewModelService);
        });

        it('should set provided services', () => {
            expect(component.service).toEqual(service);
            expect(component.viewModelService).toEqual(viewModelService);
        });

        it('should set initial state', () => {
            const expected = {
                rows: []
            };
            expect(component.state).toEqual(expected);
        });
    });

    describe('componentDidMount', () => {
        it('should call service to load data', () => {
            const defaultNEntries = 8;
            component.setState = jest.fn();

            component.componentDidMount();
            expect(service.getEntryLookups).toHaveBeenCalled();
            expect(service.getClasses).toHaveBeenCalled();
            expect(service.getEntries).toHaveBeenCalledWith(defaultNEntries);
        });

        it('should set state', () => {
            const expectedDashEntryRows = ([
                {
                    id: 1
                }
            ] as unknown) as DashEntryRow[];
            jest.spyOn(
                viewModelService,
                'formatDashboardRowEntries'
            ).mockReturnValue(expectedDashEntryRows);
            component.setState = jest.fn();

            component.componentDidMount();
            expect(component.setState).toHaveBeenCalledWith({
                rows: expectedDashEntryRows
            });
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

    describe('render', () => {
        let wrapper: ShallowWrapper;

        beforeEach(() => {
            wrapper = shallow(
                <DashboardEntriesComponent userID={props.userID} />
            );
        });

        it('should render a Table component', () => {
            expect(wrapper.find(Table)).toHaveLength(1);
        });

        it('should render a link for the entries', () => {
            const linkPath = component['getEntriesPath']();
            expect(wrapper.find(Link)).toHaveLength(1);
            expect(wrapper.find(Link).props().href).toBe(linkPath);
        });
    });
});
