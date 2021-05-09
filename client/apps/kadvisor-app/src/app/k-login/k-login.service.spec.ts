import { APP_LOGIN_ENDPOINT, KRxios, Login } from '@client/klibs';
import KLoginService from './k-login.service';

describe('KLoginService', () => {
    let mockKrxios: KRxios;
    let service: KLoginService;

    const testEmail = 'test@email.com';
    const testPwd = 'test';
    const testUser = ({
        email: testEmail,
        password: testPwd
    } as unknown) as Login;

    beforeEach(() => {
        mockKrxios = ({
            post: jest.fn(),
            getToken: jest.fn()
        } as unknown) as KRxios;

        service = new KLoginService(mockKrxios);
    });

    describe('login', () => {
        it('should call krxios.post with login endpoint and user', () => {
            service.login(testUser);
            expect(mockKrxios.post).toHaveBeenCalledWith(
                APP_LOGIN_ENDPOINT.login,
                JSON.stringify(testUser)
            );
        });
    });

    describe('logout', () => {
        it('should call krxios.post with logout endpoint and user', () => {
            service.logout(testUser);
            expect(mockKrxios.post).toHaveBeenCalledWith(
                APP_LOGIN_ENDPOINT.logout,
                JSON.stringify(testUser)
            );
        });
    });

    describe('getToken', () => {
        it('should call krxios.getToken with a user', () => {
            service.getToken(testUser);
            expect(mockKrxios.getToken).toHaveBeenCalledWith(testUser);
        });
    });
});
