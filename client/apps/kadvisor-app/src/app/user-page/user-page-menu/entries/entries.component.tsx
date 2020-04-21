import React from 'react';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Title from '../Title/title.component';

function createData(
    id: any,
    date: any,
    description: any,
    inOut: any,
    classification: any,
    amount: any
) {
    return { id, date, description, inOut, classification, amount };
}

const rows = [
    createData(0, '16 Mar, 2019', 'rbc lunch', 'expense', 'food', 312.44),
    createData(1, '16 Mar, 2019', 'rbc lunch', 'expense', 'food', 866.99),
    createData(2, '16 Mar, 2019', 'computer', 'expense', 'objects', 100.81),
    createData(3, '16 Mar, 2019', 'sold gpu', 'income', 'objects', 654.39),
    createData(4, '15 Mar, 2019', 'monitor', 'expense', 'objects', 212.79)
];

function preventDefault(event: any) {
    event.preventDefault();
}

const useStyles = makeStyles(theme => ({
    seeMore: {
        marginTop: theme.spacing(3)
    }
}));

export default function Entries() {
    const classes = useStyles();
    return (
        <React.Fragment>
            <Title>Recent Entries</Title>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>Date</TableCell>
                        <TableCell>Description</TableCell>
                        <TableCell>In / Out</TableCell>
                        <TableCell>Class</TableCell>
                        <TableCell align="right">Amount</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.map(row => (
                        <TableRow key={row.id}>
                            <TableCell>{row.date}</TableCell>
                            <TableCell>{row.description}</TableCell>
                            <TableCell>{row.inOut}</TableCell>
                            <TableCell>{row.classification}</TableCell>
                            <TableCell align="right">{row.amount}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            <div className={classes.seeMore}>
                <Link color="primary" href="#" onClick={preventDefault}>
                    See more entries
                </Link>
            </div>
        </React.Fragment>
    );
}
