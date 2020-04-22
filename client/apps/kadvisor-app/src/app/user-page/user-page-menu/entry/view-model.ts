import { Column } from 'material-table';

export interface RowData {
    description: string;
    date: string;
    class: string;
    subClass: string;
    amount: number;
}

export interface TableState {
    columns: Array<Column<RowData>>;
    data: RowData[];
}
