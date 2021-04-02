import { KFormatUtil } from './k-format-util';

describe('KFormatUtil', () => {
    describe('dateDisplayFormat', () => {
        const utcDate = new Date(Date.UTC(0, 0, 0, 0, 0, 0));
        const expected = getExpectedDate(utcDate);

        it('should format a Date obj', () => {
            const result = KFormatUtil.dateDisplayFormat(utcDate);
            expect(result).toEqual(expected);
        });

        it('should format a string Date', () => {
            const result = KFormatUtil.dateDisplayFormat(utcDate.toUTCString());
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

function getExpectedDate(date: Date): string {
    const asDate = new Date(date);
    const ye = new Intl.DateTimeFormat('en', { year: 'numeric' }).format(
        asDate
    );
    const mo = new Intl.DateTimeFormat('en', { month: 'short' }).format(asDate);
    const da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(asDate);

    return `${da} ${mo}, ${ye}`;
}
