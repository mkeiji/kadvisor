import EntryViewModelService from './entry-view-model.service';
import { Class, Entry, KLookupUtil, LookupEntry } from '@client/klibs';
import { RowData } from './view-model';

describe('EntryViewModelService', () => {
    let service: EntryViewModelService;

    const today = new Date();
    const todayStr = today.toString();
    const todayUnix = Math.round(new Date().getTime() / 1000);
    const testUserId = 1;
    const testLookupId = 1;
    const testCode = 'testTypeCode';
    const testLookupEntry = {
        id: testLookupId,
        text: 'test',
        code: testCode
    };
    const testEntry = {
        id: 1,
        createdAt: todayUnix,
        userID: testUserId,
        entryTypeCodeID: testCode,
        classID: 1,
        amount: 5.0,
        date: todayStr,
        description: 'testDescription',
        obs: 'testObs'
    } as Entry;

    beforeEach(() => {
        service = new EntryViewModelService();
    });

    describe('formatTableState', () => {
        it('should create/format correct columns', () => {
            const entryTypeRecord = { 1: 'test' };
            const classRecord = { 1: 'test' };
            const testClass = {
                userID: testUserId,
                name: 'testName',
                description: 'testDescription'
            } as Class;

            jest.spyOn(
                KLookupUtil,
                'createClassAndEntryTypeRecords'
            ).mockReturnValue([classRecord, entryTypeRecord]);
            jest.spyOn(service, 'entriesToRowDatas');

            const result = service.formatTableState(
                [testLookupEntry],
                [testClass],
                [testEntry]
            );
            const descriptionColumn = result.columns[0];
            const dateColumn = result.columns[1];
            const typeColumn = result.columns[2];
            const classColumn = result.columns[3];
            const amountColumn = result.columns[4];

            expect(descriptionColumn.title).toEqual('Description');
            expect(dateColumn.title).toEqual('Date');
            expect(typeColumn.title).toEqual('Type');
            expect(classColumn.title).toEqual('Class');
            expect(amountColumn.title).toEqual('Amount');

            expect(dateColumn.render).not.toBeUndefined();
            expect(typeColumn.lookup).toEqual(entryTypeRecord);
            expect(classColumn.lookup).toEqual(entryTypeRecord);
            expect(service.entriesToRowDatas).toHaveBeenCalledWith(
                [testEntry],
                [testLookupEntry]
            );
        });
    });

    describe('rowDataToEntry', () => {
        it('should map RowData to Entry', () => {
            const testRowData = {
                entryID: 1,
                description: 'test',
                createdAt: today,
                date: today,
                codeTypeID: testLookupId,
                class: 2,
                amount: 5.0
            };

            const result = service.rowDataToEntry(
                testUserId,
                [testLookupEntry],
                testRowData
            );

            expect(result.id).toEqual(testRowData.entryID);
            expect(result.userID).toEqual(testUserId);
            expect(result.entryTypeCodeID).toEqual(testCode);
            expect(result.classID).toEqual(Number(testRowData.class));
            expect(result.date).toEqual(testRowData.date.toISOString());
            expect(result.amount).toEqual(Number(testRowData.amount));
            expect(result.description).toEqual(testRowData.description);
        });
    });

    describe('parseRowDataDate', () => {
        it('should return rowData date as Date format', () => {
            const testRowData = { date: today } as RowData;

            const result = service.parseRowDataDate(testRowData);
            expect(result.date instanceof Date).toBeTruthy();
            expect(result.date).toEqual(today);
        });
    });

    describe('entriesToRowDatas', () => {
        it('should call entryToRowData as many times as there are entries', () => {
            jest.spyOn(service, 'entryToRowData').mockReturnValue(
                {} as RowData
            );
            const expectedSize = 2;
            const testEntries = [{}, {}] as Entry[];
            const testLookups = [{}, {}] as LookupEntry[];

            const result = service.entriesToRowDatas(testEntries, testLookups);

            expect(result.length).toEqual(expectedSize);
            expect(service.entryToRowData).toHaveBeenCalledTimes(expectedSize);
        });
    });

    describe('entryToRowData', () => {
        it('should map an entry to a rowData', () => {
            const result = service.entryToRowData(testEntry, [testLookupEntry]);

            expect(result.entryID).toEqual(testEntry.id);
            expect(result.createdAt).toEqual(new Date(testEntry.createdAt));
            expect(result.date).toEqual(new Date(testEntry.date));
            expect(result.description).toEqual(testEntry.description);
            expect(result.class).toEqual(testEntry.classID);
            expect(result.codeTypeID).toEqual(testLookupEntry.id);
            expect(result.amount).toEqual(testEntry.amount);
        });
    });
});
