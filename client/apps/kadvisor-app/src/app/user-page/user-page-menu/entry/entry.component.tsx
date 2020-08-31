import React, { CSSProperties, useEffect, useState } from 'react';
import MaterialTable from 'material-table';
import { RowData, TableState } from './view-model';
import { FormControl, InputLabel, MenuItem, Select } from '@material-ui/core';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import Grid from '@material-ui/core/Grid';
import PageSpacer from '../page-spacer/page-spacer.component';
import EntryService from './entry.service';
import EntryViewModelService from './view-model.service';
import { combineLatest } from 'rxjs';
import { take } from 'rxjs/operators';
import {
    Entry,
    KSpinner,
    LookupEntry,
    MATERIAL_TABLE_ICONS
} from '@client/klibs';

export default function EntryTable(props: EntryComponentPropsType) {
    const service = new EntryService(props.userID);
    const viewModelService = new EntryViewModelService();
    const [loading, setLoading] = useState(true);
    const [table, setTable] = useState<TableState>({} as TableState);
    const [nEntries, setNEntries] = useState<number>(10);
    const [lookups, setLookups] = useState<LookupEntry[]>([]);

    useEffect(() => {
        combineLatest([
            service.getEntryLookups(),
            service.getClasses(),
            service.getEntries(nEntries)
        ])
            .pipe(take(1))
            .subscribe(([resLookups, resClasses, resEntries]) => {
                setLookups(resLookups);
                setTable(
                    viewModelService.formatTableState(
                        resLookups,
                        resClasses,
                        resEntries
                    )
                );
                setLoading(false);
            });
    }, [nEntries]);

    async function onAdd(newData: RowData) {
        return service
            .postEntry(
                viewModelService.rowDataToEntry(props.userID, lookups, newData)
            )
            .pipe(take(1))
            .toPromise()
            .then((entry: Entry) => {
                setTable((prevState: TableState) => {
                    const data = [...prevState.data];
                    const row = viewModelService.entryToRowData(entry, lookups);
                    data.unshift(row);
                    return { ...prevState, data };
                });
            });
    }

    async function onEdit(newData: RowData, oldData: RowData | undefined) {
        return service
            .putEntry(
                viewModelService.rowDataToEntry(
                    props.userID,
                    lookups,
                    viewModelService.parseRowDataDate(newData)
                )
            )
            .pipe(take(1))
            .toPromise()
            .then((entry: Entry) => {
                setTable((prevState: TableState) => {
                    const data = [...prevState.data];
                    data[
                        data.indexOf(oldData)
                    ] = viewModelService.entryToRowData(entry, lookups);
                    return { ...prevState, data };
                });
            });
    }

    async function onDelete(oldData: RowData) {
        return service
            .deleteEntry(oldData.entryID)
            .pipe(take(1))
            .toPromise()
            .then(() => {
                setTable((prevState: TableState) => {
                    const data = [...prevState.data];
                    data.splice(data.indexOf(oldData), 1);
                    return { ...prevState, data };
                });
            });
    }

    function renderTable(): JSX.Element {
        if (loading) {
            return <KSpinner />;
        } else {
            return (
                <MaterialTable
                    title="Entries"
                    icons={MATERIAL_TABLE_ICONS}
                    columns={table.columns}
                    data={table.data}
                    options={{
                        headerStyle: headerStyles,
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
            );
        }
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
                        onChange={(event) =>
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

                {renderTable()}
            </Grid>
        </PageSpacer>
    );
}

interface EntryComponentPropsType {
    userID: number;
    classes: ClassNameMap<any>;
}

const headerStyles = {
    backgroundColor: 'gray',
    color: 'white'
} as CSSProperties;

const selectStyle = { display: 'flex' } as CSSProperties;
