import React, { Component, CSSProperties } from 'react';
import {
    ClassesTableComponentState,
    ClassesTablePropsType,
    ClassTableState
} from './view-model';
import { Class, KSpinner, MATERIAL_TABLE_ICONS } from '@client/klibs';
import MaterialTable from 'material-table';
import ClassTableService from './class-table.service';
import { takeUntil } from 'rxjs/operators';
import ClassTableViewModelService from './class-table-view-model.service';
import { Subject } from 'rxjs';

export default class ClassesTable extends Component<
    ClassesTablePropsType,
    ClassesTableComponentState
> {
    service: ClassTableService;
    viewModelService: ClassTableViewModelService;

    unsubscribe$ = new Subject<boolean>();
    headerStyles = {
        backgroundColor: 'darkgrey',
        color: 'white'
    } as CSSProperties;

    constructor(props: ClassesTablePropsType) {
        super(props);
        this.service = this.props.service
            ? this.props.service
            : new ClassTableService(this.props.userID);
        this.viewModelService = this.props.viewModelService
            ? this.props.viewModelService
            : new ClassTableViewModelService();

        this.state = {
            loading: true,
            table: {} as ClassTableState
        };
    }

    componentDidMount() {
        this.service
            .getClasses()
            .pipe(takeUntil(this.unsubscribe$))
            .subscribe((classes: Class[]) => {
                this.setState({
                    table: this.viewModelService.mapClassesToClassTableState(
                        classes
                    ),
                    loading: false
                });
            });
    }

    componentWillUnmount() {
        this.unsubscribe$.next(true);
        this.unsubscribe$.unsubscribe();
    }

    async onAdd(newClass: Class) {
        newClass.userID = this.props.userID;
        await this.service
            .postClass(newClass)
            .pipe(takeUntil(this.unsubscribe$))
            .toPromise()
            .then((c: Class) => {
                this.setState((prevState: ClassesTableComponentState) => {
                    const data = prevState.table
                        ? [...prevState.table.data]
                        : [];
                    data.unshift(c);
                    return this.handleTableStateClassUpdate(prevState, data);
                });
            });
    }

    async onEdit(newClass: Class, oldClass: Class | undefined) {
        await this.service
            .putClass(newClass)
            .pipe(takeUntil(this.unsubscribe$))
            .toPromise()
            .then((c: Class) => {
                this.setState((prevState: ClassesTableComponentState) => {
                    const data = prevState.table
                        ? [...prevState.table.data]
                        : [];
                    data[data.indexOf(oldClass)] = c;
                    return this.handleTableStateClassUpdate(prevState, data);
                });
            });
    }

    async onDelete(oldClass: Class) {
        await this.service
            .deleteClass(oldClass.id)
            .pipe(takeUntil(this.unsubscribe$))
            .toPromise()
            .then(() => {
                this.setState((prevState: ClassesTableComponentState) => {
                    const data = prevState.table
                        ? [...prevState.table.data]
                        : [];
                    data.splice(data.indexOf(oldClass), 1);
                    return this.handleTableStateClassUpdate(prevState, data);
                });
            });
    }

    render(): JSX.Element {
        if (this.state.loading) {
            return <KSpinner />;
        } else {
            return (
                <MaterialTable
                    title={'Classes'}
                    icons={MATERIAL_TABLE_ICONS}
                    columns={this.state.table ? this.state.table.columns : []}
                    data={this.state.table ? this.state.table.data : []}
                    options={{
                        headerStyle: this.headerStyles,
                        actionsColumnIndex: -1,
                        pageSize: 10,
                        addRowPosition: 'first',
                        search: false
                    }}
                    editable={{
                        onRowAdd: (e: Class) => this.onAdd(e),
                        onRowUpdate: (
                            newData: Class,
                            oldData: Class | undefined
                        ) => this.onEdit(newData, oldData),
                        onRowDelete: (oldData: Class) => this.onDelete(oldData)
                    }}
                />
            );
        }
    }

    private handleTableStateClassUpdate(
        prevState: ClassesTableComponentState,
        data: Class[]
    ): ClassesTableComponentState {
        const updatedTable = {
            ...prevState.table,
            data: data
        };
        return { ...prevState, table: updatedTable };
    }
}
