import React, { CSSProperties, useEffect, useState } from 'react';
import { ClassTableState } from './view-model';
import { Class, MATERIAL_TABLE_ICONS } from '@client/klibs';
import MaterialTable from 'material-table';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import ClassTableService from './class-table.service';
import { take } from 'rxjs/operators';
import ClassTableViewModelService from './class-table-view-model.service';

export default function ClassesTable(props: ClassesTablePropsType) {
    const service = new ClassTableService(props.userID);
    const viewModelService = new ClassTableViewModelService();
    const [table, setTable] = useState<ClassTableState>({} as ClassTableState);

    useEffect(() => {
        service
            .getClasses()
            .pipe(take(1))
            .subscribe((classes: Class[]) => {
                setTable(viewModelService.formatTableState(classes));
            });
    }, []);

    function onAdd(newClass: Class) {
        newClass.userID = props.userID;
        return service
            .postClass(newClass)
            .pipe(take(1))
            .toPromise()
            .then((c: Class) => {
                setTable((prevState: ClassTableState) => {
                    const data = [...prevState.data];
                    data.unshift(c);
                    return { ...prevState, data };
                });
            });
    }

    function onEdit(newClass: Class, oldClass: Class | undefined) {
        return service
            .putClass(newClass)
            .pipe(take(1))
            .toPromise()
            .then((c: Class) => {
                setTable((prevState: ClassTableState) => {
                    const data = [...prevState.data];
                    data[data.indexOf(oldClass)] = c;
                    return { ...prevState, data };
                });
            });
    }

    function onDelete(oldClass: Class) {
        return service
            .deleteClass(oldClass.id)
            .pipe(take(1))
            .toPromise()
            .then(() => {
                setTable((prevState: ClassTableState) => {
                    const data = [...prevState.data];
                    data.splice(data.indexOf(oldClass), 1);
                    return { ...prevState, data };
                });
            });
    }

    return (
        <MaterialTable
            title={'Classes'}
            icons={MATERIAL_TABLE_ICONS}
            columns={table.columns}
            data={table.data}
            options={{
                headerStyle: headerStyles,
                actionsColumnIndex: -1,
                pageSize: 10,
                addRowPosition: 'first',
                search: false
            }}
            editable={{
                onRowAdd: (e: Class) => onAdd(e),
                onRowUpdate: (newData: Class, oldData: Class | undefined) =>
                    onEdit(newData, oldData),
                onRowDelete: (oldData: Class) => onDelete(oldData)
            }}
        />
    );
}

interface ClassesTablePropsType {
    userID: number;
    classes: ClassNameMap<any>;
}

const headerStyles = {
    backgroundColor: 'darkgrey',
    color: 'white'
} as CSSProperties;
