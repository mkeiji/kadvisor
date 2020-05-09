import { RowData, TableState } from './view-model';
import { Column } from 'material-table';
import {
    Class,
    Entry,
    KFormatUtil,
    KLookupUtil,
    LookupEntry
} from '@client/klibs';

class EntryViewModelService {
    formatTableState(
        lookups: LookupEntry[],
        classes: Class[],
        entries: Entry[]
    ): TableState {
        const [
            classLookups,
            entryTypeLookups
        ] = KLookupUtil.createClassAndEntryTypeRecords(classes, lookups);

        return {
            columns: [
                { title: 'Description', field: 'description' },
                {
                    title: 'Date',
                    field: 'date',
                    type: 'date',
                    render: rowData =>
                        KFormatUtil.dateDisplayFormat(rowData.date)
                },
                {
                    title: 'Type',
                    field: 'codeTypeID',
                    lookup: entryTypeLookups
                },
                {
                    title: 'Class',
                    field: 'class',
                    lookup: classLookups
                },
                {
                    title: 'Amount',
                    field: 'amount',
                    type: 'currency',
                    cellStyle: { textAlign: 'left' }
                }
            ] as Column<RowData>[],

            data: this.entriesToRowDatas(entries, lookups)
        };
    }

    rowDataToEntry(
        userID: number,
        lookups: LookupEntry[],
        rowData: RowData
    ): Entry {
        return {
            id: rowData.entryID,
            userID: userID,
            entryTypeCodeID: lookups.find(
                l => l.id === Number(rowData.codeTypeID)
            ).code,
            classID: Number(rowData.class),
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

    entriesToRowDatas(entries: Entry[], lookups: LookupEntry[]): RowData[] {
        const result = [] as RowData[];
        entries.forEach(e => result.push(this.entryToRowData(e, lookups)));
        return result.sort((n1, n2) => +n2.createdAt - +n1.createdAt);
    }

    entryToRowData(entry: Entry, lookups: LookupEntry[]): RowData {
        return {
            entryID: entry.id,
            createdAt: new Date(entry.createdAt),
            date: new Date(entry.date),
            description: entry.description,
            class: entry.classID,
            codeTypeID: lookups.find(l => l.code === entry.entryTypeCodeID).id,
            amount: entry.amount
        } as RowData;
    }
}

export default EntryViewModelService;
