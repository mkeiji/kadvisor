import { APP_CLASS_ENDPOINT, Class, KRxios } from '@client/klibs';
import ClassTableService from './class-table.service';

describe('ClassTableService', () => {
    let mockKrxios: KRxios;
    let service: ClassTableService;

    const testID = 1;

    beforeEach(() => {
        setupMocks();
        service = new ClassTableService(testID, mockKrxios);
    });

    describe('getClasses', () => {
        it('should call krxios get', () => {
            service.getClasses();
            expect(mockKrxios.get).toHaveBeenCalledWith(APP_CLASS_ENDPOINT);
        });
    });

    describe('postClass', () => {
        it('should call krxios post', () => {
            const expectedClass = ({
                userID: testID,
                name: 'testName POST'
            } as unknown) as Class;

            service.postClass(expectedClass);
            expect(mockKrxios.post).toHaveBeenCalledWith(
                APP_CLASS_ENDPOINT,
                expectedClass
            );
        });
    });

    describe('putClass', () => {
        it('should call krxios put', () => {
            const expectedClass = ({
                userID: testID,
                name: 'testName PUT'
            } as unknown) as Class;

            service.putClass(expectedClass);
            expect(mockKrxios.put).toHaveBeenCalledWith(
                APP_CLASS_ENDPOINT,
                expectedClass
            );
        });
    });

    describe('deleteClass', () => {
        it('should call krxios delete', () => {
            const expectedClass = ({
                id: testID
            } as unknown) as Class;

            service.deleteClass(testID);
            expect(mockKrxios.delete).toHaveBeenCalledWith(
                APP_CLASS_ENDPOINT,
                expectedClass
            );
        });
    });

    function setupMocks() {
        mockKrxios = ({
            get: jest.fn(),
            post: jest.fn(),
            put: jest.fn(),
            delete: jest.fn()
        } as unknown) as KRxios;
    }
});
