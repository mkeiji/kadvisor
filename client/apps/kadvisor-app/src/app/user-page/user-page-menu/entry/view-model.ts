import { Column } from 'material-table';

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
