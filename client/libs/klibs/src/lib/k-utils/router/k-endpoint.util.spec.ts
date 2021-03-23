import { KEndpointUtil } from './k-endpoint.util';

describe('KEndpointUtil', () => {
    const mockEnvDomain = 'undefined';

    it('should return endpoint with correct id', () => {
        const expected = `${mockEnvDomain}/api/kadvisor/1`;
        const testID = 1;

        const result = KEndpointUtil.getUserBaseUrl(testID);
        expect(result).toEqual(expected);
    });
});
