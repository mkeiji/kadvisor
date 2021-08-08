import ClassTableViewModelService from './class-table-view-model.service';

describe('ClassTableViewModelService', () => {
    let today: number;
    let yesterday: number;
    let service: ClassTableViewModelService;

    const testID = 1;
    const testUserID = 2;
    const testname = 'test';
    const testdescription = 'testdescription';
    const testClass = {
        id: testID,
        createdAt: yesterday,
        updatedAt: yesterday,
        userID: testUserID,
        name: testname,
        description: testdescription
    };
    const testClass2 = {
        id: testID,
        createdAt: today,
        updatedAt: today,
        userID: testUserID,
        name: testname,
        description: testdescription
    };

    beforeEach(() => {
        const yesterdayDate = new Date();
        yesterdayDate.setDate(yesterdayDate.getDate() - 1);

        today = Math.round(new Date().getTime() / 1000);
        yesterday = Math.round(yesterdayDate.getTime() / 1000);

        service = new ClassTableViewModelService();
    });

    describe('mapClassesToClassTableState', () => {
        it('should map classes to classTableState', () => {
            const expected = {
                columns: [
                    { title: 'Classname', field: 'name' },
                    { title: 'Description', field: 'description' }
                ],
                data: [testClass]
            };

            const result = service.mapClassesToClassTableState([testClass]);
            expect(result).toEqual(expected);
        });
    });

    describe('sortedClasses', () => {
        it('it should return classes sorted descendently by creation date', () => {
            const unordered = [testClass2, testClass];
            const expected = [testClass, testClass2];
            const result = service['sortedClasses'](unordered);

            expect(result).toEqual(expected);
        });
    });
});
