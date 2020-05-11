import { Column } from 'material-table';
import { Class } from '@client/klibs';

export interface ClassTableState {
    columns: Array<Column<Class>>;
    data: Class[];
}
