import { Class, Entry, RowData, SubClass, TableState } from './view-model';
import { Column } from 'material-table';

class EntryViewModelService {
    formatTableState(classes: Class[], entries: Entry[]): TableState {
        const [
            classLookups,
            subClassLookups
        ] = this.createClassAndSubClassLookups(classes);

        return {
            columns: [
                { title: 'Description', field: 'description' },
                {
                    title: 'Date',
                    field: 'date',
                    type: 'date',
                    render: rowData => this.formatDate(rowData.date)
                },
                {
                    title: 'Class',
                    field: 'class',
                    lookup: classLookups
                },
                {
                    title: 'Sub-Class',
                    field: 'subClass',
                    lookup: subClassLookups
                },
                {
                    title: 'Amount',
                    field: 'amount',
                    type: 'currency',
                    cellStyle: { textAlign: 'left' }
                }
            ] as Column<RowData>[],

            data: this.entriesToRowDatas(entries)
        };
    }

    rowDataToEntry(userID: number, rowData: RowData): Entry {
        return {
            id: rowData.entryID,
            userID: userID,
            classID: Number(rowData.class),
            subClassID: Number(rowData.subClass),
            date: rowData.date.toISOString(),
            amount: Number(rowData.amount),
            description: rowData.description
        } as Entry;
    }

    parseRowDataDate(rowData: RowData): RowData {
        return {
            ...rowData,
            date: new Date(rowData.date)
        };
    }

    entriesToRowDatas(entries: Entry[]): RowData[] {
        const result = [] as RowData[];
        entries.forEach(e => result.push(this.entryToRowData(e)));
        return result.sort((n1, n2) => +n2.createdAt - +n1.createdAt);
    }

    entryToRowData(entry: Entry): RowData {
        return {
            entryID: entry.id,
            createdAt: new Date(entry.createdAt),
            date: new Date(entry.date),
            description: entry.description,
            class: entry.classID,
            subClass: entry.subClassID,
            amount: entry.amount
        } as RowData;
    }

    private formatDate(date: Date | string): string {
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

        return `${da}-${mo}-${ye}`;
    }

    private createClassAndSubClassLookups(
        classes: Class[]
    ): Record<number, string>[] {
        const classLookups = {} as Record<number, string>;
        const subClassLookups = {} as Record<number, string>;

        classes.forEach((c: Class, i: number) => {
            classLookups[i + 1] = c.name;
            c.subClasses.forEach((sc: SubClass, sci: number) => {
                subClassLookups[sci + 1] = sc.name;
            });
        });
        return [classLookups, subClassLookups];
    }
}

export default EntryViewModelService;
