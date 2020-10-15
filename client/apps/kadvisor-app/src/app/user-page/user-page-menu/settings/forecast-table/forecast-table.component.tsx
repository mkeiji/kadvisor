import React, { CSSProperties, useEffect, useState } from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import MaterialTable from 'material-table';
import {
    Forecast,
    ForecastEntry,
    KSpinner,
    MATERIAL_TABLE_ICONS,
    ReportsApiService,
    KSelect,
    KSelectItem
} from '@client/klibs';
import ForecastTableService from './forecast-table.service';
import ForecastTableViewModelService, {
    ForecastTableState
} from './forecast-table-view-model.service';
import { takeUntil } from 'rxjs/operators';
import { Button, TextField, Container } from '@material-ui/core';
import { Subject } from 'rxjs';
import { Row, Col } from 'react-bootstrap';

export default function ForecastTable(props: ForecastTablePropsType) {
    const destroy$ = new Subject<boolean>();
    const currentYear = new Date().getFullYear();
    const forecastService = new ForecastTableService(props.userID);
    const reportsService = new ReportsApiService(props.userID);
    const viewModelService = new ForecastTableViewModelService();
    const [yearMenuItems, setYearMenuItems] = useState<KSelectItem[]>([]);
    const [loading, setLoading] = useState(true);
    const [hasForecast, setHasForecast] = useState(true);
    const [selectedYear, setSelectedYear] = useState<number>(currentYear);
    const [forecastYear, setForecastYear] = useState<number | string>('');
    const [table, setTable] = useState<ForecastTableState>(
        {} as ForecastTableState
    );

    useEffect(() => {
        getForecast();
        getAvailableReportYears();

        return () => {
            destroy$.next(true);
            destroy$.unsubscribe();
        };
    }, [hasForecast, selectedYear]);

    async function getForecast() {
        forecastService
            .getForecast(selectedYear)
            .pipe(takeUntil(destroy$))
            .subscribe(
                (f: Forecast) => {
                    if (f) {
                        setSelectedYear(f.year);
                        setTable(viewModelService.formatTableState(f));
                        setLoading(false);
                    }
                },
                () => {
                    setHasForecast(false);
                    setLoading(false);
                }
            );
    }

    async function getAvailableReportYears() {
        reportsService
            .getAvailableReportYears()
            .pipe(takeUntil(destroy$))
            .subscribe((years: number[]) => {
                mapYearMenuItems(years);
            });
    }

    function mapYearMenuItems(years: number[]) {
        const selectMenuItems = [];
        if (years) {
            years.map((year: number) => {
                selectMenuItems.push({
                    value: year,
                    displayValue: year.toString()
                });
            });
            setYearMenuItems(selectMenuItems);
        }
    }

    async function onEdit(
        newEntry: ForecastEntry,
        oldEntry: ForecastEntry | undefined
    ) {
        return forecastService
            .putForecastEntry(viewModelService.parseAmounts(newEntry))
            .pipe(takeUntil(destroy$))
            .toPromise()
            .then((e: ForecastEntry) => {
                setTable((prevState: ForecastTableState) => {
                    const data = [...prevState.data];
                    data[data.indexOf(oldEntry)] = e;
                    return { ...prevState, data };
                });
            });
    }

    async function createForecast() {
        if (isValidYear()) {
            const newForecast = viewModelService.createNewForecast(
                props.userID,
                forecastYear as number
            );
            forecastService
                .postForecast(newForecast)
                .pipe(takeUntil(destroy$))
                .subscribe(
                    (f: Forecast) => {
                        yearMenuItems.push({
                            value: f.year,
                            displayValue: f.year.toString()
                        });
                        setTextFieldDisplayValue();
                        setHasForecast(true);
                        setSelectedYear(f.year);
                    },
                    () => setTextFieldDisplayValue()
                );
        }
    }

    function isValidYear(): boolean {
        const len = forecastYear.toString().length;
        return !isNaN(forecastYear as number) && len === 4;
    }

    function setTextFieldDisplayValue(value?: number | string) {
        if (!isNaN(value as number) && value !== 0) {
            setForecastYear(value);
        } else {
            setForecastYear('');
        }
    }

    function renderYearSelectionDropDown(): JSX.Element {
        return (
            <Col>
                <KSelect
                    items={yearMenuItems}
                    onValueChange={setSelectedYear}
                    value={selectedYear}
                    formVariant="outlined"
                />
            </Col>
        );
    }

    function renderCreateForecast(): JSX.Element {
        return (
            <form>
                <Container>
                    <Row style={{ display: 'flex', alignItems: 'center' }}>
                        {yearMenuItems.length > 0
                            ? renderYearSelectionDropDown()
                            : null}
                        <Col className="align-self-center">
                            <TextField
                                placeholder="year"
                                style={yearFieldStyles}
                                value={forecastYear}
                                onChange={(val) =>
                                    setTextFieldDisplayValue(
                                        Number(val.target.value.trim())
                                    )
                                }
                            />
                            <Button
                                color={
                                    isValidYear() || forecastYear === ''
                                        ? 'primary'
                                        : 'secondary'
                                }
                                variant="contained"
                                onClick={createForecast}
                            >
                                + forecast
                            </Button>
                        </Col>
                    </Row>
                </Container>
            </form>
        );
    }

    function renderForecast(): JSX.Element {
        if (hasForecast) {
            return (
                <div>
                    {renderCreateForecast()}
                    <MaterialTable
                        title={'Forecast'}
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
                </div>
            );
        } else {
            return renderCreateForecast();
        }
    }

    return loading ? <KSpinner /> : renderForecast();
}

interface ForecastTablePropsType {
    userID: number;
    classes: ClassNameMap<any>;
}

const headerStyles = {
    backgroundColor: 'darkgrey',
    color: 'white'
} as CSSProperties;

const yearFieldStyles = {
    width: '55px',
    paddingRight: '10px'
} as CSSProperties;
