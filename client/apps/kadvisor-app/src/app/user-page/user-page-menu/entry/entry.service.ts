import {
    APP_CLASS_ENDPOINT,
    APP_ENTRY_ENDPOINT,
    APP_LOOKUP_ENDPOINT,
    Class,
    Entry,
    KEndpointUtil,
    KRxios,
    LookupEntry
} from '@client/klibs';
import { Observable } from 'rxjs';

class EntryService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getEntries(nEntries?: number): Observable<Entry[]> {
        return this.krxios.get(`${APP_ENTRY_ENDPOINT}?limit=${nEntries}`);
    }

    postEntry(entry: Entry): Observable<Entry> {
        return this.krxios.post(APP_ENTRY_ENDPOINT, entry);
    }

    putEntry(entry: Entry): Observable<Entry> {
        return this.krxios.put(APP_ENTRY_ENDPOINT, entry);
    }

    deleteEntry(entryID: number): Observable<Entry> {
        return this.krxios.delete(APP_ENTRY_ENDPOINT, {
            id: entryID
        });
    }

    getClasses(): Observable<Class[]> {
        return this.krxios.get(APP_CLASS_ENDPOINT);
    }

    getEntryLookups(): Observable<LookupEntry[]> {
        return this.krxios.get(
            `${APP_LOOKUP_ENDPOINT}?codeGroup=EntryTypeCodeID`
        );
    }
}

export default EntryService;
