import { Class, LookupEntry } from '../../..';

export class KLookupUtil {
    static createClassAndEntryTypeRecords(
        classes: Class[],
        typeLookups: LookupEntry[]
    ): Record<number, string>[] {
        return [
            this.createClassRecord(classes),
            this.createEntryTypeRecord(typeLookups)
        ];
    }

    static createClassRecord(classes: Class[]): Record<number, string> {
        const classLookups = {} as Record<number, string>;
        classes.map((c: Class, i: number) => (classLookups[i + 1] = c.name));
        return classLookups;
    }

    static createEntryTypeRecord(
        typeLookups: LookupEntry[]
    ): Record<number, string> {
        const entryTypeLookups = {} as Record<number, string>;
        typeLookups.map((l: LookupEntry) => (entryTypeLookups[l.id] = l.text));
        return entryTypeLookups;
    }
}
