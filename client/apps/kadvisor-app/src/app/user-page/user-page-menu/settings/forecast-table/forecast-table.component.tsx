import React, { Component, CSSProperties } from 'react';
import MaterialTable from 'material-table';
import {
    Forecast,
    ForecastEntry,
    MATERIAL_TABLE_ICONS,
    ReportsApiService,
    KSelect,
    ForecastTableState,
    KSpinner
} from '@client/klibs';
import ForecastTableService from './forecast-table.service';
import ForecastTableViewModelService from './forecast-table-view-model.service';
import { takeUntil } from 'rxjs/operators';
import { Button, TextField, Container } from '@material-ui/core';
import { Subject } from 'rxjs';
import { Row, Col } from 'react-bootstrap';
import {
    ForecastTableComponentState,
    ForecastTablePropsType
} from './view-model';

export default class ForecastTable extends Component<
    ForecastTablePropsType,
    ForecastTableComponentState
> {
    service: ForecastTableService;
    reportsService: ReportsApiService;
    viewModelService: ForecastTableViewModelService;

    destroy$ = new Subject<boolean>();
    currentYear = new Date().getFullYear();
    headerStyles = {
        backgroundColor: 'darkgrey',
        color: 'white'
    } as CSSProperties;
    yearFieldStyles = {
        width: '55px',
        paddingRight: '10px'
    } as CSSProperties;

    constructor(props: ForecastTablePropsType) {
        super(props);
        this.service = props.service
            ? props.service
            : new ForecastTableService(props.userID);
        this.reportsService = props.reportsService
            ? props.reportsService
            : new ReportsApiService(props.userID);
        this.viewModelService = props.viewModelService
            ? props.viewModelService
            : new ForecastTableViewModelService();

        this.state = {
            yearMenuItems: [],
            loading: true,
            hasForecast: true,
            selectedYear: this.currentYear,
            forecastYear: '',
            table: {} as ForecastTableState
        };

        this.createForecast = this.createForecast.bind(this);
    }

    componentDidMount() {
        this.getForecast();
        this.getAvailableReportYears();
    }

    componentDidUpdate(
        _: ForecastTablePropsType,
        prevState: ForecastTableComponentState
    ) {
        if (
            this.state.hasForecast !== prevState.hasForecast ||
            this.state.selectedYear !== prevState.selectedYear
        ) {
            this.getForecast();
            this.getAvailableReportYears();
        }
    }

    componentWillUnmount() {
        this.destroy$.next(true);
        this.destroy$.unsubscribe();
    }

    private getForecast() {
        this.service
            .getForecast(this.state.selectedYear)
            .pipe(takeUntil(this.destroy$))
            .subscribe(
                (f: Forecast) => {
                    if (f) {
                        this.setState({
                            selectedYear: f.year,
                            table: this.viewModelService.formatTableState(f),
                            loading: false
                        });
                    }
                },
                () => {
                    this.setState({
                        hasForecast: false,
                        loading: false
                    });
                }
            );
    }

    private getAvailableReportYears() {
        this.reportsService
            .getAvailableReportYears(true)
            .pipe(takeUntil(this.destroy$))
            .subscribe((years: number[]) => {
                this.mapYearMenuItems(years);
            });
    }

    private mapYearMenuItems(years: number[]) {
        const selectMenuItems = [];
        if (years) {
            years.map((year: number) => {
                selectMenuItems.push({
                    value: year,
                    displayValue: year.toString()
                });
            });
            this.setState({ yearMenuItems: selectMenuItems });
        }
    }

    private async onEdit(
        newEntry: ForecastEntry,
        oldEntry: ForecastEntry | undefined
    ) {
        return this.service
            .putForecastEntry(this.viewModelService.parseAmounts(newEntry))
            .pipe(takeUntil(this.destroy$))
            .toPromise()
            .then((e: ForecastEntry) => {
                this.setState((prevState: ForecastTableComponentState) =>
                    this.viewModelService.handleTableStateUpdate(
                        prevState,
                        oldEntry,
                        e
                    )
                );
            });
    }

    private createForecast() {
        if (this.isValidYear()) {
            const newForecast = this.viewModelService.createNewForecast(
                this.props.userID,
                this.state.forecastYear as number
            );
            this.service
                .postForecast(newForecast)
                .pipe(takeUntil(this.destroy$))
                .subscribe(
                    (f: Forecast) => {
                        this.setState(
                            (prevState: ForecastTableComponentState) =>
                                this.viewModelService.handleYearMenuItemsStateUpdate(
                                    prevState,
                                    f
                                )
                        );
                    },
                    () => this.setTextFieldDisplayValue()
                );
        }
    }

    private isValidYear(): boolean {
        const len = this.state.forecastYear.toString().length;
        return !isNaN(this.state.forecastYear as number) && len === 4;
    }

    private setTextFieldDisplayValue(value?: number | string) {
        if (!isNaN(value as number) && value !== 0) {
            this.setState({ forecastYear: value });
        } else {
            this.setState({ forecastYear: '' });
        }
    }

    private renderYearSelectionDropDown(): JSX.Element {
        return (
            <Col>
                <KSelect
                    items={this.state.yearMenuItems}
                    onValueChange={(value: number) =>
                        this.setState({ selectedYear: value })
                    }
                    value={this.state.selectedYear}
                    formVariant="outlined"
                />
            </Col>
        );
    }

    private renderCreateForecast(): JSX.Element {
        return (
            <form>
                <Container>
                    <Row style={{ display: 'flex', alignItems: 'center' }}>
                        {this.state.yearMenuItems.length > 0
                            ? this.renderYearSelectionDropDown()
                            : null}
                        <Col className="align-self-center">
                            <TextField
                                placeholder="year"
                                style={this.yearFieldStyles}
                                value={this.state.forecastYear}
                                onChange={(val) =>
                                    this.setTextFieldDisplayValue(
                                        Number(val.target.value.trim())
                                    )
                                }
                            />
                            <Button
                                color={
                                    this.isValidYear() ||
                                    this.state.forecastYear === ''
                                        ? 'primary'
                                        : 'secondary'
                                }
                                variant="contained"
                                onClick={this.createForecast}
                            >
                                + forecast
                            </Button>
                        </Col>
                    </Row>
                </Container>
            </form>
        );
    }

    private renderForecast(): JSX.Element {
        if (this.state.hasForecast) {
            return (
                <div>
                    {this.renderCreateForecast()}
                    <MaterialTable
                        title={'Forecast'}
                        icons={MATERIAL_TABLE_ICONS}
                        columns={this.state.table.columns}
                        data={this.state.table.data}
                        options={{
                            headerStyle: this.headerStyles,
                            actionsColumnIndex: -1,
                            pageSize: 12,
                            paging: false,
                            search: false
                        }}
                        editable={{
                            onRowUpdate: (
                                newData: ForecastEntry,
                                oldData: ForecastEntry
                            ) => this.onEdit(newData, oldData)
                        }}
                    />
                </div>
            );
        } else {
            return this.renderCreateForecast();
        }
    }

    render(): JSX.Element {
        return this.state.loading ? <KSpinner /> : this.renderForecast();
    }
}
