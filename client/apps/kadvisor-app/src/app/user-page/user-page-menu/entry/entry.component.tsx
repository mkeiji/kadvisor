import React, { CSSProperties, useState } from 'react';
import MaterialTable from 'material-table';
import { RowData, TableState } from './view-model';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import Grid from '@material-ui/core/Grid';
import PageSpacer from '../page-spacer/page-spacer.component';
import EntryService from './entry.service';
import EntryViewModelService from './entry-view-model.service';
import { Subject, combineLatest } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import {
    Entry,
    KSpinner,
    LookupEntry,
    MATERIAL_TABLE_ICONS,
    KSelect
} from '@client/klibs';

export default function EntryComponent(props: EntryComponentPropsType) {
    const destroy$ = new Subject<boolean>();
    const service = new EntryService(props.userID);
    const viewModelService = new EntryViewModelService();
    const selectMenuItems = [
        { value: 10, displayValue: 'last 10' },
        { value: 20, displayValue: 'last 20' },
        { value: 50, displayValue: 'last 50' },
        { value: 100, displayValue: 'last 100' },
        { value: 0, displayValue: 'all entries' }
    ];
    const [loading, setLoading] = useState(true);
    const [table, setTable] = useState<TableState>({} as TableState);
    const [nEntries, setNEntries] = useState<number>(10);
    const [lookups, setLookups] = useState<LookupEntry[]>([]);

    React.useEffect(() => {
        combineLatest([
            service.getEntryLookups(),
            service.getClasses(),
            service.getEntries(nEntries)
        ])
            .pipe(takeUntil(destroy$))
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

        return () => {
            destroy$.next(true);
            destroy$.unsubscribe();
        };
    }, [nEntries]);

    async function onAdd(newData: RowData) {
        return service
            .postEntry(
                viewModelService.rowDataToEntry(props.userID, lookups, newData)
            )
            .pipe(takeUntil(destroy$))
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
            .pipe(takeUntil(destroy$))
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
            .pipe(takeUntil(destroy$))
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
                <KSelect
                    label={'# Entries'}
                    items={selectMenuItems}
                    onValueChange={setNEntries}
                    value={nEntries}
                    class={props.classes.formControl}
                    style={selectStyle}
                />
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
