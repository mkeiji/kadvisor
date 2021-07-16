import React, { Component } from 'react';
import Link from '@material-ui/core/Link';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Title from '../../Title/title.component';
import EntryService from '../../entry/entry.service';
import DashEntriesViewModelService from './dash-entries-view-model.service';
import { combineLatest, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { APP_PAGES, KRouterPathUtil } from '@client/klibs';
import {
    DashboardEntriesPropsType,
    DashboardEntriesState
} from '../dashboard.models';

class DashboardEntriesComponent extends Component<
    DashboardEntriesPropsType,
    DashboardEntriesState
> {
    service: EntryService;
    viewModelService: DashEntriesViewModelService;
    destroy$ = new Subject<boolean>();
    nEntries = 8;

    constructor(readonly props: DashboardEntriesPropsType) {
        super(props);
        this.service = this.props.service
            ? this.props.service
            : new EntryService(props.userID);
        this.viewModelService = this.props.viewModelService
            ? this.props.viewModelService
            : new DashEntriesViewModelService();
        this.state = { rows: [] };
    }

    componentDidMount() {
        combineLatest([
            this.service.getEntryLookups(),
            this.service.getClasses(),
            this.service.getEntries(this.nEntries)
        ])
            .pipe(takeUntil(this.destroy$))
            .subscribe(([resLookups, resClasses, resEntries]) =>
                this.setState({
                    rows: this.viewModelService.formatDashboardRowEntries(
                        resLookups,
                        resClasses,
                        resEntries
                    )
                })
            );
    }

    componentWillUnmount() {
        this.destroy$.next(true);
        this.destroy$.unsubscribe();
    }

    render() {
        const entriesLinkStyle = { paddingTop: '20px' };
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
                        {this.state.rows.map((row) => (
                            <TableRow key={row.id}>
                                <TableCell>{row.date}</TableCell>
                                <TableCell>{row.description}</TableCell>
                                <TableCell>{row.codeTypeID}</TableCell>
                                <TableCell>{row.strClass}</TableCell>
                                <TableCell align="right">
                                    {row.amount}
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
                <div style={entriesLinkStyle}>
                    <Link color="primary" href={this.getEntriesPath()}>
                        See more entries
                    </Link>
                </div>
            </React.Fragment>
        );
    }

    private getEntriesPath(): string {
        return KRouterPathUtil.getUserPage(
            this.props.userID,
            APP_PAGES.entries
        );
    }
}

export default DashboardEntriesComponent;
