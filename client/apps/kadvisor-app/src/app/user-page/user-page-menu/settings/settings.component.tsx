import React from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import PageSpacer from '../page-spacer/page-spacer.component';
import { Grid } from '@material-ui/core';
import ClassesTable from './classes-table/classes-table.component';
import ForecastTable from './forecast-table/forecast-table.component';

export default function Settings(props: SettingsPropsType) {
    return (
        <PageSpacer classes={props.classes}>
            <Grid container spacing={3}>
                <Grid item xs={6}>
                    <ForecastTable
                        userID={props.userID}
                        classes={props.classes}
                    />
                </Grid>
                <Grid item xs={6}>
                    <ClassesTable
                        userID={props.userID}
                        classes={props.classes}
                    />
                </Grid>
            </Grid>
        </PageSpacer>
    );
}

interface SettingsPropsType {
    userID: number;
    classes: ClassNameMap<any>;
}
