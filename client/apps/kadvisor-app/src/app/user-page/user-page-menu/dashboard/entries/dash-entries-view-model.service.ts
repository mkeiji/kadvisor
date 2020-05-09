import { DashEntryRow } from './view-model';
import {
    Class,
    Entry,
    KFormatUtil,
    KLookupUtil,
    LookupEntry
} from '@client/klibs';

class DashEntriesViewModelService {
    formatDashboardRowEntries(
        typeLookups: LookupEntry[],
        classes: Class[],
        entries: Entry[]
    ): DashEntryRow[] {
        const result = [] as DashEntryRow[];
        const classLookups = KLookupUtil.createClassRecord(classes);

        entries.map((e, i) =>
            result.push({
                ...this.entryToDashEntryRow(e, classLookups, typeLookups),
                id: i
            })
        );

        return result;
    }

    entryToDashEntryRow(
        entry: Entry,
        classLookup: Record<number, string>,
        typeLookup: LookupEntry[]
    ): DashEntryRow {
        return {
            date: KFormatUtil.dateDisplayFormat(entry.date),
            description: entry.description,
            codeTypeID: typeLookup.find(
                (l: LookupEntry) => l.code === entry.entryTypeCodeID
            ).text,
            strClass: classLookup[entry.classID],
            amount: KFormatUtil.toCurrency(entry.amount)
        } as DashEntryRow;
    }
}

export default DashEntriesViewModelService;
