import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import KComposedChartComponent from './charts/k-composed-chart.component';
import BalanceCard from './balance-card/balance-card.component';
import DashboardEntries from './entries/dash-entries.component';
import React from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import clsx from 'clsx';
import PageSpacer from '../page-spacer/page-spacer.component';

// TODO: create a service to fetch from server and feed to children
export default function Dashboard(props: DashboardPropsType) {
    const fixedHeightPaper = clsx(
        props.classes.paper,
        props.classes.fixedHeight
    );
    return (
        <PageSpacer classes={props.classes}>
            <Grid container spacing={3}>
                {/* KComposedChartComponent */}
                <Grid item xs={12} md={8} lg={9}>
                    <Paper className={fixedHeightPaper}>
                        <KComposedChartComponent />
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
