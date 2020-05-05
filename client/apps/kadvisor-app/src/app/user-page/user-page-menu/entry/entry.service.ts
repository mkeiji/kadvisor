import {
    APP_CLASS_ENDPOINT,
    APP_ENTRY_ENDPOINT,
    KEndpointUtil,
    KRxios
} from '@client/klibs';
import { Class, Entry } from './view-model';
import { Observable } from 'rxjs';

class EntryService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getEntries(): Observable<Entry[]> {
        return this.krxios.get(APP_ENTRY_ENDPOINT);
    }

    postEntry(entry: Entry): any {
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
        return this.krxios.get(`${APP_CLASS_ENDPOINT}?preloaded=true`);
    }
}

export default EntryService;
