import {
    APP_CLASS_ENDPOINT,
    Class,
    KEndpointUtil,
    KRxios
} from '@client/klibs';
import { Observable } from 'rxjs';

class ClassTableService {
    private krxios: KRxios;
    constructor(userID: number, krxios?: KRxios) {
        this.krxios = krxios
            ? krxios
            : new KRxios(KEndpointUtil.getUserBaseUrl(userID));
    }

    getClasses(): Observable<Class[]> {
        return this.krxios.get(APP_CLASS_ENDPOINT);
    }

    postClass(sClass: Class): Observable<Class> {
        return this.krxios.post(APP_CLASS_ENDPOINT, sClass);
    }

    putClass(uClass: Class): Observable<Class> {
        return this.krxios.put(APP_CLASS_ENDPOINT, uClass);
    }

    deleteClass(classID: number): Observable<Class> {
        return this.krxios.delete(APP_CLASS_ENDPOINT, {
            id: classID
        });
    }
}

export default ClassTableService;
