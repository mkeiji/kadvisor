import React, { useEffect, useState } from 'react';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Title from '../../Title/title.component';
import EntryService from '../../entry/entry.service';
import DashEntriesViewModelService from './dash-entries-view-model.service';
import { DashEntryRow } from './view-model';
import { combineLatest, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { APP_PAGES, KRouterPathUtil } from '@client/klibs';

export default function DashboardEntries(props: DashboardEntriesPropsType) {
    const destroy$ = new Subject<boolean>();
    const service = new EntryService(props.userID);
    const viewModelService = new DashEntriesViewModelService();
    const classes = useStyles();
    const nEntries = 8;
    const [rows, setRows] = useState<DashEntryRow[]>([]);

    useEffect(() => {
        combineLatest(
            service.getEntryLookups(),
            service.getClasses(),
            service.getEntries(nEntries)
        )
            .pipe(takeUntil(destroy$))
            .subscribe(([resLookups, resClasses, resEntries]) =>
                setRows(
                    viewModelService.formatDashboardRowEntries(
                        resLookups,
                        resClasses,
                        resEntries
                    )
                )
            );

        return () => {
            destroy$.next(true);
            destroy$.unsubscribe();
        };
    }, []);

    function getEntriesPath(): string {
        return KRouterPathUtil.getUserPage(props.userID, APP_PAGES.entries);
    }

    return (
        <React.Fragment>
            <Title>Recent Entries</Title>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>Date</TableCell>
                        <TableCell>Description</TableCell>
                        <TableCell>Type</TableCell>
                        <TableCell>Class</TableCell>
                        <TableCell align="right">Amount</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.map((row) => (
                        <TableRow key={row.id}>
                            <TableCell>{row.date}</TableCell>
                            <TableCell>{row.description}</TableCell>
                            <TableCell>{row.codeTypeID}</TableCell>
                            <TableCell>{row.strClass}</TableCell>
                            <TableCell align="right">{row.amount}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            <div className={classes.seeMore}>
                <Link color="primary" href={getEntriesPath()}>
                    See more entries
                </Link>
            </div>
        </React.Fragment>
    );
}

interface DashboardEntriesPropsType {
    userID: number;
}

const useStyles = makeStyles((theme) => ({
    seeMore: {
        marginTop: theme.spacing(3)
    }
}));
