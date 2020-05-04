import {
    APP_CLASS_ENDPOINT,
    APP_ENTRY_ENDPOINT,
    KEndpointUtil,
    KRxios
} from '@client/klibs';
import { KEntryResponse, KClassResponse, Entry } from './view-model';
import { Observable } from 'rxjs';

class EntryService {
    private krxios: KRxios;
    constructor(userID: number) {
        this.krxios = new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getEntries(): Observable<KEntryResponse> {
        return this.krxios.get(APP_ENTRY_ENDPOINT);
    }

    postEntry(entry: Entry): any {
        return this.krxios.post(APP_ENTRY_ENDPOINT, entry);
    }

    putEntry(entry: Entry): Observable<KEntryResponse> {
        return this.krxios.put(APP_ENTRY_ENDPOINT, entry);
    }

    deleteEntry(entryID: number): Observable<KEntryResponse> {
        return this.krxios.delete(APP_ENTRY_ENDPOINT, {
            id: entryID
        });
    }

    getClasses(): Observable<KClassResponse> {
        return this.krxios.get(`${APP_CLASS_ENDPOINT}?preloaded=true`);
    }
}

export default EntryService;
