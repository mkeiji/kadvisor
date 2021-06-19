import { Entry, KEndpointUtil, KRxios } from '@client/klibs';
import EntryService from './entry.service';

describe('EntryService', () => {
    let mockKrxios: KRxios;
    let service: EntryService;

    const testUserID = 1;

    beforeEach(() => {
        mockKrxios = ({
            get: jest.fn(),
            put: jest.fn(),
            post: jest.fn(),
            delete: jest.fn(),
            getToken: jest.fn()
        } as unknown) as KRxios;

        service = new EntryService(testUserID, mockKrxios);
    });

    describe('getEntries', () => {
        it('should call krxios get with correct Uri', () => {
            const testNEntries = 2;
            const expectedUri = `/entry?limit=${testNEntries}`;
            service.getEntries(testNEntries);
            expect(mockKrxios.get).toHaveBeenCalledWith(expectedUri);
        });
    });

    describe('postEntry', () => {
        it('should call krxios post with correct Uri and obj', () => {
            const testEntriy = {} as Entry;
            const expectedUri = `/entry`;
            service.postEntry(testEntriy);
            expect(mockKrxios.post).toHaveBeenCalledWith(
                expectedUri,
                testEntriy
            );
        });
    });

    describe('putEntry', () => {
        it('should call krxios put with correct Uri and obj', () => {
            const testEntriy = {} as Entry;
            const expectedUri = `/entry`;
            service.putEntry(testEntriy);
            expect(mockKrxios.put).toHaveBeenCalledWith(
                expectedUri,
                testEntriy
            );
        });
    });

    describe('deleteEntry', () => {
        it('should call krxios delete with correct Uri and ID', () => {
            const testEntriyID = 1;
            const expectedUri = `/entry`;
            service.deleteEntry(testEntriyID);
            expect(mockKrxios.delete).toHaveBeenCalledWith(expectedUri, {
                id: testEntriyID
            });
        });
    });

    describe('getClasses', () => {
        it('should call krxios get with correct classes endpoint', () => {
            const expectedEndpoint = `/class`;
            service.getClasses();
            expect(mockKrxios.get).toHaveBeenCalledWith(expectedEndpoint);
        });
    });

    describe('getEntryLookups', () => {
        it('should call krxios get with correct endpoint uri', () => {
            const expectedUri = `/lookup?codeGroup=EntryTypeCodeID`;
            service.getEntryLookups();
            expect(mockKrxios.get).toHaveBeenCalledWith(expectedUri);
        });
    });
});
