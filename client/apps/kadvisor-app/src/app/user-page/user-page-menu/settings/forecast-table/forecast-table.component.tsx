import React, { CSSProperties, useEffect, useState } from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import MaterialTable from 'material-table';
import { Forecast, ForecastEntry, MATERIAL_TABLE_ICONS } from '@client/klibs';
import ForecastTableService from './forecast-table.service';
import ForecastTableViewModelService, {
    ForecastTableState
} from './forecast-table-view-model.service';
import { take } from 'rxjs/operators';
import { Button } from '@material-ui/core';

export default function ForecastTable(props: ForecastTablePropsType) {
    const service = new ForecastTableService(props.userID);
    const viewModelService = new ForecastTableViewModelService();
    const [hasForecast, setHasForecast] = useState(true);
    const [forecast, setForecast] = useState<Forecast>({} as Forecast);
    const [table, setTable] = useState<ForecastTableState>(
        {} as ForecastTableState
    );

    useEffect(() => {
        service
            .getForecast()
            .pipe(take(1))
            .subscribe(
                (f: Forecast) => {
                    if (f) {
                        setForecast(f);
                        setTable(viewModelService.formatTableState(f));
                    }
                },
                () => setHasForecast(false)
            );
    }, [hasForecast]);

    function onEdit(
        newEntry: ForecastEntry,
        oldEntry: ForecastEntry | undefined
    ) {
        return service
            .putForecastEntry(viewModelService.parseAmounts(newEntry))
            .pipe(take(1))
            .toPromise()
            .then((e: ForecastEntry) => {
                setTable((prevState: ForecastTableState) => {
                    const data = [...prevState.data];
                    data[data.indexOf(oldEntry)] = e;
                    return { ...prevState, data };
                });
            });
    }

    function createForecast() {
        const newForecast = viewModelService.createNewForecast(props.userID);
        service
            .postForecast(newForecast)
            .pipe(take(1))
            .subscribe(() => setHasForecast(true));
    }

    function render(): JSX.Element {
        if (hasForecast) {
            return (
                <MaterialTable
                    title={`Forecast - ${forecast.year}`}
                    icons={MATERIAL_TABLE_ICONS}
                    columns={table.columns}
                    data={table.data}
                    options={{
                        headerStyle: headerStyles,
                        actionsColumnIndex: -1,
                        pageSize: 12,
                        paging: false,
                        search: false
                    }}
                    editable={{
                        onRowUpdate: (
                            newData: ForecastEntry,
                            oldData: ForecastEntry
                        ) => onEdit(newData, oldData)
                    }}
                />
            );
        } else {
            return (
                <Button variant="contained" onClick={createForecast}>
                    Create a Forecast
                </Button>
            );
        }
    }
    return <div>{render()}</div>;
}

interface ForecastTablePropsType {
    userID: number;
    classes: ClassNameMap<any>;
}

const headerStyles = {
    backgroundColor: 'darkgrey',
    color: 'white'
} as CSSProperties;
