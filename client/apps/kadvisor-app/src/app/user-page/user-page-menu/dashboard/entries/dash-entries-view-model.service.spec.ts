import { Class, Entry, KFormatUtil, KLookupUtil } from '@client/klibs';
import DashEntriesViewModelService from './dash-entries-view-model.service';

describe('DashEntriesViewModelService', () => {
    let service: DashEntriesViewModelService;

    const testID = 1;
    const testAmount = 10.0;
    const today = new Date();
    const testObs = 'obs';
    const testName = 'test';
    const testDescription = 'testDescription';
    const testTypeCodeID = 'testTypeCodeID';
    const testLookupText = 'testLookupText';
    const testClass = {
        userID: testID,
        name: testName,
        description: testDescription
    } as Class;
    const testEntry = ({
        id: testID,
        userID: testID,
        entryTypeCodeID: testTypeCodeID,
        classID: testID,
        amount: testAmount,
        date: today,
        description: testDescription,
        obs: testObs
    } as unknown) as Entry;
    const testTypeLookup = {
        id: testID,
        text: testLookupText,
        code: testTypeCodeID
    };

    beforeEach(() => {
        service = new DashEntriesViewModelService();
    });

    describe('formatDashboardRowEntries', () => {
        it('should call createClassRecord and entryToDashEntryRow', () => {
            const utilSpy = jest.spyOn(KLookupUtil, 'createClassRecord');
            const serviceSpy = jest.spyOn(service, 'entryToDashEntryRow');

            service.formatDashboardRowEntries(
                [testTypeLookup],
                [testClass],
                [testEntry]
            );

            expect(utilSpy).toHaveBeenCalledWith([testClass]);
            expect(serviceSpy).toHaveBeenCalledTimes(1);
            expect(serviceSpy).toHaveBeenCalledWith(
                testEntry,
                KLookupUtil.createClassRecord([testClass]),
                [testTypeLookup]
            );
        });
    });

    describe('entryToDashEntryRow', () => {
        it('should map entry calling dateDisplayFormat and toCurrency from util', () => {
            const expected = {
                date: KFormatUtil.dateDisplayFormat(today),
                description: testDescription,
                codeTypeID: testLookupText,
                strClass: testName,
                amount: '$10.00'
            };
            const classLookup = KLookupUtil.createClassRecord([testClass]);
            const dateFormatSpy = jest.spyOn(KFormatUtil, 'dateDisplayFormat');
            const toCurrencySpy = jest.spyOn(KFormatUtil, 'toCurrency');

            const result = service.entryToDashEntryRow(testEntry, classLookup, [
                testTypeLookup
            ]);

            expect(dateFormatSpy).toHaveBeenCalledWith(today);
            expect(toCurrencySpy).toHaveBeenCalledWith(testAmount);
            expect(result).toEqual(expected);
        });
    });
});
