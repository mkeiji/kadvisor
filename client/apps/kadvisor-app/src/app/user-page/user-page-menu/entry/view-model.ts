import { Column } from 'material-table';
import { BaseModel } from '@client/klibs';

export interface RowData {
    entryID: number;
    description: string;
    createdAt: Date;
    date: Date;
    codeTypeID: number;
    class: number;
    amount: number;
}

export interface TableState {
    columns: Array<Column<RowData>>;
    data: RowData[];
}

export interface Class extends BaseModel {
    userID: number;
    name: string;
    description: string;
}

export interface Entry extends BaseModel {
    id: number;
    userID: number;
    entryTypeCodeID: string;
    classID: number;
    amount: number;
    date: string;
    description: string;
    obs: string;
}
