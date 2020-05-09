export class KFormatUtil {
    static dateDisplayFormat(date: Date | string): string {
        const asDate = new Date(date);
        const ye = new Intl.DateTimeFormat('en', { year: 'numeric' }).format(
            asDate
        );
        const mo = new Intl.DateTimeFormat('en', { month: 'short' }).format(
            asDate
        );
        const da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(
            asDate
        );

        return `${da} ${mo}, ${ye}`;
    }

    static toCurrency(value: number): string {
        const formatter = new Intl.NumberFormat('en-US', {
            style: 'currency',
            currency: 'USD',
            minimumFractionDigits: 2
        });

        return formatter.format(value);
    }
}
