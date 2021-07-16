import { KRxios, APP_REPORT_ENDPOINT } from '@client/klibs';
import BalanceCardService from './balance-card.service';

describe('BalanceCardService', () => {
    let service: BalanceCardService;
    let mockKrxios: KRxios;

    const testID = 1;

    beforeEach(() => {
        mockKrxios = ({
            get: jest.fn()
        } as unknown) as KRxios;
        service = new BalanceCardService(testID, mockKrxios);
    });

    describe('constructor', () => {
        it('set krxios instance', () => {
            expect(service['krxios']).not.toBeUndefined();
            expect(service['krxios']).not.toBeNull();
        });
    });

    describe('getUserBalance', () => {
        it('calls krxios.get', () => {
            const expectedEndpoint = `${APP_REPORT_ENDPOINT}?type=BALANCE`;
            service.getUserBalance();
            expect(mockKrxios.get).toHaveBeenCalledWith(expectedEndpoint);
        });
    });
});
