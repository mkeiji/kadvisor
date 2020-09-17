import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import KComposedChartComponent from './charts/k-composed-chart.component';
import BalanceCard from './balance-card/balance-card.component';
import DashboardEntries from './entries/dash-entries.component';
import React, { useState, useEffect } from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import clsx from 'clsx';
import PageSpacer from '../page-spacer/page-spacer.component';
import { KSelect, KSelectItem } from '@client/klibs';
import ChartService from './charts/chart.service';
import { take } from 'rxjs/operators';

export default function Dashboard(props: DashboardPropsType) {
    const chartService = new ChartService(props.userID);
    const currentYear = new Date().getFullYear();
    const [graphYear, setGraphYear] = useState<number>(currentYear);
    const [yearMenuItems, setYearMenuItems] = useState<KSelectItem[]>([]);
    const fixedHeightPaper = clsx(
        props.classes.paper,
        props.classes.fixedHeight
    );

    useEffect(() => {
        chartService
            .getAvailableReportYears()
            .pipe(take(1))
            .subscribe((years: number[]) => {
                const selectMenuItems = [];
                years.map((year: number) => {
                    selectMenuItems.push({
                        value: year,
                        displayValue: year.toString()
                    });
                });
                setYearMenuItems(selectMenuItems);
                setGraphYear(currentYear);
            });
    }, []);

    return (
        <PageSpacer classes={props.classes}>
            <Grid container item xs={12} alignContent={'flex-start'}>
                <KSelect
                    label={'Year'}
                    items={yearMenuItems}
                    onValueChange={setGraphYear}
                    initialValue={graphYear}
                />
            </Grid>
            <Grid container spacing={3}>
                {/* KComposedChartComponent */}
                <Grid item xs={12} md={8} lg={9}>
                    <Paper className={fixedHeightPaper}>
                        <KComposedChartComponent
                            userID={props.userID}
                            year={graphYear}
                        />
                    </Paper>
                </Grid>
                {/* Recent BalanceCardComponent */}
                <Grid item xs={12} md={4} lg={3}>
                    <Paper className={fixedHeightPaper}>
                        <BalanceCard userID={props.userID} />
                    </Paper>
                </Grid>
                {/* Recent EntriesComponent */}
                <Grid item xs={12}>
                    <Paper className={props.classes.paper}>
                        <DashboardEntries userID={props.userID} />
                    </Paper>
                </Grid>
            </Grid>
        </PageSpacer>
    );
}

interface DashboardPropsType {
    userID: number;
    classes: ClassNameMap<any>;
}
