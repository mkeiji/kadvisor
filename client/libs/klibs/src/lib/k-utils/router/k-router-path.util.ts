import { APP_ROUTES } from './route.constants';

export class KRouterPathUtil {
    static getUserPage(userID: number, page?: string): string {
        const path = APP_ROUTES.userPage.replace(':id', userID.toString());
        return page ? path.replace(':page?', page) : path.replace(':page?', '');
    }
}
