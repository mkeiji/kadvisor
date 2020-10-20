import React, { useEffect, useState } from 'react';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Title from '../../Title/title.component';
import { KFormatUtil, UserBalance } from '@client/klibs';
import BalanceCardService from './balance-card.service';

function preventDefault(event: any) {
    event.preventDefault();
}

export default function BalanceCard(props: BalanceCardPropType) {
    const service = new BalanceCardService(props.userID);
    const classes = useStyles();
    const [userBalance, setBalance] = useState<UserBalance>({} as UserBalance);

    useEffect(() => {
        service.getUserBalance().subscribe(
            (u: UserBalance) => setBalance(u),
            () => setBalance({ userID: props.userID, balance: 0 })
        );
    }, []);

    function getBalanceAsNumber(): number {
        return !isNaN(userBalance.balance) ? userBalance.balance : 0;
    }

    return (
        <React.Fragment>
            <Title>Balance Card</Title>
            <Typography
                component="p"
                variant="h4"
                style={{ paddingTop: '30px' }}
            >
                {KFormatUtil.toCurrency(getBalanceAsNumber())}
            </Typography>
            <Typography
                color="textSecondary"
                className={classes.depositContext}
            >
                on {KFormatUtil.dateDisplayFormat(new Date())}
            </Typography>
            <div>
                {/*TODO: add reports page link when available*/}
                <Link color="primary" href="#" onClick={preventDefault}>
                    View reports
                </Link>
            </div>
        </React.Fragment>
    );
}

const useStyles = makeStyles({
    depositContext: {
        flex: 1
    }
});

interface BalanceCardPropType {
    userID: number;
}
