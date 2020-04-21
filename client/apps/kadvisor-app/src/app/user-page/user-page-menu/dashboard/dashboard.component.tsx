import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import KComposedChartComponent from '../charts/k-composed-chart.component';
import BalanceCard from '../balance-card/balance-card.component';
import Entries from '../entries/entries.component';
import Box from '@material-ui/core/Box';
import React from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import clsx from 'clsx';
import { KCopyright } from '@client/klibs';

export default function Dashboard(props: DashboardPropsType) {
    const fixedHeightPaper = clsx(
        props.classes.paper,
        props.classes.fixedHeight
    );
    return (
        <div>
            <div className={props.classes.appBarSpacer} />
            <Container maxWidth="lg" className={props.classes.container}>
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
                            <BalanceCard />
                        </Paper>
                    </Grid>
                    {/* Recent EntriesComponent */}
                    <Grid item xs={12}>
                        <Paper className={props.classes.paper}>
                            <Entries />
                        </Paper>
                    </Grid>
                </Grid>
                <Box pt={4}>
                    <KCopyright />
                </Box>
            </Container>
        </div>
    );
}

interface DashboardPropsType {
    classes: ClassNameMap<any>;
}
