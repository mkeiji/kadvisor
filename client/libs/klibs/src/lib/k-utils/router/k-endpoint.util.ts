import { APP_BACKEND_BASE } from './route.constants';

export class KEndpointUtil {
    static getUserBaseUrl(userID: number): string {
        return APP_BACKEND_BASE.replace(':uid', userID.toString());
    }
}
