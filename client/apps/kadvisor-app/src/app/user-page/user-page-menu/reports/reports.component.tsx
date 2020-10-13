import React from 'react';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import PageSpacer from '../page-spacer/page-spacer.component';
import { Grid } from '@material-ui/core';

export default function Reports(props: ReportsPropsType) {
    return (
        <PageSpacer classes={props.classes}>
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <h1>Under Construction</h1>
                </Grid>
            </Grid>
        </PageSpacer>
    );
}

interface ReportsPropsType {
    userID: number;
    classes: ClassNameMap<any>;
}
