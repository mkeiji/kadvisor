import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import KComposedChartComponent from './charts/k-composed-chart.component';
import BalanceCard from './balance-card/balance-card.component';
import DashboardEntries from './entries/dash-entries.component';
import React, { useState, useEffect } from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import clsx from 'clsx';
import PageSpacer from '../page-spacer/page-spacer.component';
import { KSelect, KSelectItem, ReportsApiService } from '@client/klibs';
import { takeUntil } from 'rxjs/operators';
import { Subject } from 'rxjs';

export default function Dashboard(props: DashboardPropsType) {
    const destroy$ = new Subject<boolean>();
    const service = new ReportsApiService(props.userID);
    const currentYear = new Date().getFullYear();
    const [graphYear, setGraphYear] = useState<number>(currentYear);
    const [showYearDropdown, setShowYearDropdown] = useState<boolean>(false);
    const [yearMenuItems, setYearMenuItems] = useState<KSelectItem[]>([]);
    const fixedHeightPaper = clsx(
        props.classes.paper,
        props.classes.fixedHeight
    );

    useEffect(() => {
        service
            .getAvailableReportYears()
            .pipe(takeUntil(destroy$))
            .subscribe((years: number[]) => {
                const selectMenuItems = [];
                if (years) {
                    years.map((year: number) => {
                        selectMenuItems.push({
                            value: year,
                            displayValue: year.toString()
                        });
                    });
                    setYearMenuItems(selectMenuItems);
                    setShowYearDropdown(true);
                }
                setGraphYear(
                    years.includes(currentYear)
                        ? currentYear
                        : years[years.length - 1]
                );
            });

        return () => {
            destroy$.next(true);
            destroy$.unsubscribe();
        };
    }, []);

    function renderDropdown(): JSX.Element {
        return showYearDropdown ? (
            <KSelect
                label={'Year'}
                items={yearMenuItems}
                onValueChange={setGraphYear}
                value={graphYear}
            />
        ) : null;
    }

    return (
        <PageSpacer classes={props.classes}>
            <Grid container item xs={12} alignContent={'flex-start'}>
                {renderDropdown()}
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
