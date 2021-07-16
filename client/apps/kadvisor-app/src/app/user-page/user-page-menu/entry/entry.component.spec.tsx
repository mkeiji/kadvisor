import React from 'react';
import { mount, ReactWrapper, shallow } from 'enzyme';
import EntryComponent from './entry.component';
import EntryService from './entry.service';
import EntryViewModelService from './entry-view-model.service';
import { mocked } from 'ts-jest/utils';
import { of } from 'rxjs';
import PageSpacer from '../page-spacer/page-spacer.component';
import MaterialTable from 'material-table';
jest.mock('./entry.service');
jest.mock('./entry-view-model.service');

describe('EntryComponent', () => {
    let useEffect: any;
    let mockService: EntryService;
    let mockViewModelService: EntryViewModelService;

    const testID = 1;
    const testStyles = {};
    const mockServiceMgr = mocked(EntryService, true);
    const mockViewModelServiceMgr = mocked(EntryViewModelService, true);
    const mockUseEffect = () => {
        useEffect.mockImplementationOnce((f: any) => f());
    };

    beforeEach(() => {
        mockServiceMgr.mockClear();
        mockViewModelServiceMgr.mockClear();
    });

    afterAll(() => {
        jest.unmock('./entry.service');
        jest.unmock('./entry-view-model.service');
    });

    describe('creates services', () => {
        beforeEach(() => {
            mockService = ({
                getEntries: jest.fn(() => of([])),
                getClasses: jest.fn(() => of([])),
                getEntryLookups: jest.fn(() => of([]))
            } as unknown) as EntryService;
            mockViewModelService = ({
                formatTableState: jest.fn(() => ({}))
            } as unknown) as EntryViewModelService;

            mockServiceMgr.mockImplementationOnce(() => mockService);
            mockViewModelServiceMgr.mockImplementationOnce(
                () => mockViewModelService
            );

            useEffect = jest.spyOn(React, 'useEffect');
            mockUseEffect();

            shallow(
                <EntryComponent
                    userID={testID}
                    classes={testStyles}
                ></EntryComponent>
            );
        });

        it('EntryService', () => {
            expect(EntryService).toHaveBeenCalled();
        });

        it('EntryViewModelService', () => {
            expect(EntryViewModelService).toHaveBeenCalled();
        });
    });

    describe('call services', () => {
        const testLookups = [{ id: testID }];
        const testClasses = [{ userId: testID }];
        const testEntries = [{ id: testID }];

        beforeEach(() => {
            mockService = ({
                getEntries: jest.fn(() => of(testEntries)),
                getClasses: jest.fn(() => of(testClasses)),
                getEntryLookups: jest.fn(() => of(testLookups))
            } as unknown) as EntryService;
            mockViewModelService = ({
                formatTableState: jest.fn(() => ({}))
            } as unknown) as EntryViewModelService;

            mockServiceMgr.mockImplementationOnce(() => mockService);
            mockViewModelServiceMgr.mockImplementationOnce(
                () => mockViewModelService
            );

            useEffect = jest.spyOn(React, 'useEffect');
            mockUseEffect();

            shallow(
                <EntryComponent
                    userID={testID}
                    classes={testStyles}
                ></EntryComponent>
            );
        });

        it('entryService.getEntries', () => {
            expect(mockService.getEntries).toHaveBeenCalled();
        });

        it('entryService.getClasses', () => {
            expect(mockService.getClasses).toHaveBeenCalled();
        });

        it('entryService.getEntryLookups', () => {
            expect(mockService.getEntryLookups).toHaveBeenCalled();
        });

        it('viewModelService.formatTableState', () => {
            expect(mockViewModelService.formatTableState).toHaveBeenCalledWith(
                testLookups,
                testClasses,
                testEntries
            );
        });
    });

    describe('renders', () => {
        let wrapper: ReactWrapper;

        beforeEach(() => {
            mockService = ({
                getEntries: jest.fn(() => of([])),
                getClasses: jest.fn(() => of([])),
                getEntryLookups: jest.fn(() => of([])),
                postEntry: jest.fn()
            } as unknown) as EntryService;
            mockViewModelService = ({
                formatTableState: jest.fn(() => ({}))
            } as unknown) as EntryViewModelService;

            mockServiceMgr.mockImplementationOnce(() => mockService);
            mockViewModelServiceMgr.mockImplementationOnce(
                () => mockViewModelService
            );

            wrapper = mount(
                <EntryComponent
                    userID={testID}
                    classes={testStyles}
                ></EntryComponent>
            );
        });

        it('pageSpacer component', () => {
            expect(wrapper.find(PageSpacer)).toHaveLength(1);
        });

        it('materialTable component', () => {
            expect(wrapper.find(MaterialTable)).toHaveLength(1);
        });
    });
});
