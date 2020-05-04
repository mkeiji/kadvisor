import { Column } from 'material-table';
import { BaseModel } from '@client/klibs';

export interface RowData {
    entryID: number;
    description: string;
    createdAt: Date;
    date: Date;
    class: number;
    subClass: number;
    amount: number;
}

export interface TableState {
    columns: Array<Column<RowData>>;
    data: RowData[];
}

export interface KClassResponse {
    classes: Class[];
}

export interface Class extends BaseModel {
    userID: number;
    name: string;
    description: string;
    subClasses: SubClass[];
}

export interface SubClass extends BaseModel {
    classID: number;
    userID: number;
    name: string;
    description: string;
}

export interface KEntryResponse {
    entries: Entry[];
}

export interface Entry extends BaseModel {
    id: number;
    userID: number;
    classID: number;
    subClassID: number;
    amount: number;
    date: string;
    description: string;
    obs: string;
}
