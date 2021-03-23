import { Class } from '../../k-models/class-model';
import { LookupEntry } from '../../k-models/gerneric-models';
import { KLookupUtil } from './k-lookup.util';

describe('KLookupUtil', () => {
    const testID = 1;
    const testText = 'test';
    const testClasses = [
        ({
            name: testText
        } as unknown) as Class
    ];
    const testLookups = [
        {
            id: testID,
            text: testText
        } as LookupEntry
    ];

    describe('createClassAndEntryTypeRecords', () => {
        it('should return a list of class and lookup records', () => {
            const expected = [{ 1: testText }, { 1: testText }] as Record<
                number,
                string
            >[];

            const result = KLookupUtil.createClassAndEntryTypeRecords(
                testClasses,
                testLookups
            );
            expect(result).toEqual(expected);
        });
    });

    describe('createClassRecord', () => {
        it('should map classes into a record', () => {
            const expectedIndex = 1;
            const result = KLookupUtil.createClassRecord(testClasses);
            expect(result[expectedIndex]).toEqual(testText);
        });
    });

    describe('createEntryTypeRecord', () => {
        it('should map lookups into a record', () => {
            const result = KLookupUtil.createEntryTypeRecord(testLookups);
            expect(result[testID]).toEqual(testText);
        });
    });
});
