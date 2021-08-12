import { Class } from '@client/klibs';
import { ClassTableState } from './view-model';
import { orderBy } from 'lodash';

class ClassTableViewModelService {
    mapClassesToClassTableState(classes: Class[]): ClassTableState {
        return {
            columns: [
                { title: 'Classname', field: 'name' },
                { title: 'Description', field: 'description' }
            ],
            data: this.sortedClasses(classes)
        } as ClassTableState;
    }

    private sortedClasses(classes: Class[]): Class[] {
        return orderBy(classes, ['createdAt'], ['desc']);
    }
}

export default ClassTableViewModelService;
