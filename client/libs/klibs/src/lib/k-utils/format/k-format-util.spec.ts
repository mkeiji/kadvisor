import { KFormatUtil } from './k-format-util';

describe('KFormatUtil', () => {
    describe('dateDisplayFormat', () => {
        const expected = '22 Mar, 2021';
        const testDate = '2021-03-23T03:53:29.812Z';

        it('should format a Date obj', () => {
            const dateObj = new Date(testDate);

            const result = KFormatUtil.dateDisplayFormat(dateObj);
            expect(result).toEqual(expected);
        });

        it('should format a string Date', () => {
            const result = KFormatUtil.dateDisplayFormat(testDate);
            expect(result).toEqual(expected);
        });
    });

    describe('toCurrency', () => {
        it('should format a number into string currency', () => {
            const expected = '$10.00';
            const testNumber = 10;

            const result = KFormatUtil.toCurrency(testNumber);
            expect(result).toEqual(expected);
        });
    });
});
