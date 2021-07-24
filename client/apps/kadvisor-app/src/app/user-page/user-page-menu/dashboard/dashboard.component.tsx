import React, { Component } from 'react';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import KComposedChartComponent from './charts/k-composed-chart.component';
import BalanceCard from './balance-card/balance-card.component';
import DashboardEntriesComponent from './entries/dash-entries.component';
import clsx from 'clsx';
import PageSpacer from '../page-spacer/page-spacer.component';
import { KSelect, ReportsApiService } from '@client/klibs';
import { takeUntil } from 'rxjs/operators';
import { Subject } from 'rxjs';
import { DashboardPropsType, DashboardState } from './dashboard.models';

export default class Dashboard extends Component<
    DashboardPropsType,
    DashboardState
> {
    service: ReportsApiService;
    destroy$ = new Subject<boolean>();
    currentYear = new Date().getFullYear();
    fixedHeightPaper = clsx(
        this.props.classes.paper,
        this.props.classes.fixedHeight
    );

    constructor(readonly props: DashboardPropsType) {
        super(props);
        this.service = this.props.reportsApiService
            ? this.props.reportsApiService
            : new ReportsApiService(this.props.userID);

        this.state = {
            graphYear: this.currentYear,
            showYearDropdown: false,
            yearMenuItems: []
        };
    }

    componentDidMount() {
        this.service
            .getAvailableReportYears()
            .pipe(takeUntil(this.destroy$))
            .subscribe((years: number[]) => {
                const selectMenuItems = [];
                if (years) {
                    years.map((year: number) => {
                        selectMenuItems.push({
                            value: year,
                            displayValue: year.toString()
                        });
                    });

                    this.setState({
                        yearMenuItems: selectMenuItems,
                        showYearDropdown: true
                    });
                }

                this.setState({
                    graphYear: years.includes(this.currentYear)
                        ? this.currentYear
                        : years[years.length - 1]
                });
            });
    }

    componentWillUnmount() {
        this.destroy$.next(true);
        this.destroy$.unsubscribe();
    }

    renderDropdown(): JSX.Element {
        return this.state.showYearDropdown ? (
            <KSelect
                label={'Year'}
                items={this.state.yearMenuItems}
                onValueChange={(v: number) => this.setState({ graphYear: v })}
                value={this.state.graphYear}
            />
        ) : null;
    }

    render() {
        return (
            <PageSpacer classes={this.props.classes}>
                <Grid container item xs={12} alignContent={'flex-start'}>
                    {this.renderDropdown()}
                </Grid>
                <Grid container spacing={3}>
                    {/* KComposedChartComponent */}
                    <Grid item xs={12} md={8} lg={9}>
                        <Paper className={this.fixedHeightPaper}>
                            <KComposedChartComponent
                                userID={this.props.userID}
                                year={this.state.graphYear}
                            />
                        </Paper>
                    </Grid>
                    {/* Recent BalanceCardComponent */}
                    <Grid item xs={12} md={4} lg={3}>
                        <Paper className={this.fixedHeightPaper}>
                            <BalanceCard userID={this.props.userID} />
                        </Paper>
                    </Grid>
                    {/* Recent EntriesComponent */}
                    <Grid item xs={12}>
                        <Paper className={this.props.classes.paper}>
                            <DashboardEntriesComponent
                                userID={this.props.userID}
                            />
                        </Paper>
                    </Grid>
                </Grid>
            </PageSpacer>
        );
    }
}
