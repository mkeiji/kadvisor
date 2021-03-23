import { KRouterPathUtil } from './k-router-path.util';

describe('KRouterPathUtil', () => {
    const testID = 1;
    const testPage = 'testPage';

    it('should return formated path id only', () => {
        const expected = '/user/1/home/';
        const result = KRouterPathUtil.getUserPage(testID);
        expect(result).toEqual(expected);
    });

    it('should return formated path with page', () => {
        const expected = '/user/1/home/testPage';
        const result = KRouterPathUtil.getUserPage(testID, testPage);
        expect(result).toEqual(expected);
    });
});
