import React, { CSSProperties, forwardRef, useEffect, useState } from 'react';
import MaterialTable, { Icons } from 'material-table';
import { Entry, RowData, TableState } from './view-model';
import AddBox from '@material-ui/icons/AddBox';
import ArrowDownward from '@material-ui/icons/ArrowDownward';
import Check from '@material-ui/icons/Check';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import Clear from '@material-ui/icons/Clear';
import DeleteOutline from '@material-ui/icons/DeleteOutline';
import Edit from '@material-ui/icons/Edit';
import FilterList from '@material-ui/icons/FilterList';
import FirstPage from '@material-ui/icons/FirstPage';
import LastPage from '@material-ui/icons/LastPage';
import Remove from '@material-ui/icons/Remove';
import SaveAlt from '@material-ui/icons/SaveAlt';
import Search from '@material-ui/icons/Search';
import ViewColumn from '@material-ui/icons/ViewColumn';
import { FormControl, InputLabel, MenuItem, Select } from '@material-ui/core';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import Grid from '@material-ui/core/Grid';
import PageSpacer from '../page-spacer/page-spacer.component';
import EntryService from './entry.service';
import EntryViewModelService from './view-model.service';
import { combineLatest } from 'rxjs';
import { take } from 'rxjs/operators';

export default function EntryTable(props: EntryComponentPropsType) {
    const service = new EntryService(props.userID);
    const viewModelService = new EntryViewModelService();
    const [table, setTable] = useState<TableState>({} as TableState);
    const [nEntries, setNEntries] = React.useState(10);

    useEffect(() => {
        combineLatest(service.getClasses(), service.getEntries(nEntries))
            .pipe(take(1))
            .subscribe(([resClasses, resEntries]) =>
                setTable(
                    viewModelService.formatTableState(resClasses, resEntries)
                )
            );
    }, [nEntries]);

    function onAdd(newData: RowData) {
        service
            .postEntry(viewModelService.rowDataToEntry(props.userID, newData))
            .pipe(take(1))
            .subscribe((entry: Entry) => {
                setTable((prevState: TableState) => {
                    const data = [...prevState.data];
                    const row = viewModelService.entryToRowData(entry);
                    data.unshift(row);
                    return { ...prevState, data };
                });
            });
        return new Promise(resolve => {
            resolve();
        });
    }

    function onEdit(newData: RowData, oldData: RowData | undefined) {
        return new Promise(resolve => {
            resolve();
            if (oldData) {
                setTable((prevState: TableState) => {
                    const data = [...prevState.data];
                    data[data.indexOf(oldData)] = newData;
                    service.putEntry(
                        viewModelService.rowDataToEntry(
                            props.userID,
                            viewModelService.parseRowDataDate(newData)
                        )
                    );
                    return { ...prevState, data };
                });
            }
        });
    }

    function onDelete(oldData: RowData) {
        return new Promise(resolve => {
            resolve();
            setTable((prevState: TableState) => {
                const data = [...prevState.data];
                data.splice(data.indexOf(oldData), 1);
                service.deleteEntry(oldData.entryID);
                return { ...prevState, data };
            });
        });
    }

    return (
        <PageSpacer classes={props.classes}>
            <Grid item xs={12}>
                <FormControl
                    className={props.classes.formControl}
                    style={selectStyle}
                >
                    <InputLabel># Entries</InputLabel>
                    <Select
                        value={nEntries}
                        onChange={event =>
                            setNEntries(event.target.value as number)
                        }
                    >
                        <MenuItem value={10}>last 10</MenuItem>
                        <MenuItem value={20}>last 20</MenuItem>
                        <MenuItem value={50}>last 50</MenuItem>
                        <MenuItem value={100}>last 100</MenuItem>
                        <MenuItem value={0}>all entries</MenuItem>
                    </Select>
                </FormControl>

                <MaterialTable
                    title="Entries"
                    icons={tableIcons}
                    columns={table.columns}
                    data={table.data}
                    options={{
                        actionsColumnIndex: -1,
                        pageSize: 10,
                        addRowPosition: 'first',
                        exportButton: true
                    }}
                    editable={{
                        onRowAdd: (newData: RowData) => onAdd(newData),
                        onRowUpdate: (
                            newData: RowData,
                            oldData: RowData | undefined
                        ) => onEdit(newData, oldData),
                        onRowDelete: (oldData: RowData) => onDelete(oldData)
                    }}
                />
            </Grid>
        </PageSpacer>
    );
}

interface EntryComponentPropsType {
    userID: number;
    classes: ClassNameMap<any>;
}

const selectStyle = { display: 'flex' } as CSSProperties;

const tableIcons = {
    Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
    Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
    Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
    DetailPanel: forwardRef((props, ref) => (
        <ChevronRight {...props} ref={ref} />
    )),
    Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
    Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
    Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
    FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
    LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
    NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    PreviousPage: forwardRef((props, ref) => (
        <ChevronLeft {...props} ref={ref} />
    )),
    ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
    SortArrow: forwardRef((props, ref) => (
        <ArrowDownward {...props} ref={ref} />
    )),
    ThirdStateCheck: forwardRef((props, ref) => (
        <Remove {...props} ref={ref} />
    )),
    ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />)
} as Icons;
