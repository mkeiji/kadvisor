import { Column } from 'material-table';
import { Class } from '@client/klibs';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import ClassTableService from './class-table.service';
import ClassTableViewModelService from './class-table-view-model.service';

export interface ClassTableState {
    columns: Array<Column<Class>>;
    data: Class[];
}

export interface ClassesTableComponentState {
    loading: boolean;
    table: ClassTableState;
}

export interface ClassesTablePropsType {
    userID: number;
    classes: ClassNameMap<any>;
    service?: ClassTableService;
    viewModelService?: ClassTableViewModelService;
}
