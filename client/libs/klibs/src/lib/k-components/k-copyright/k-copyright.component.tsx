import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import React from 'react';

export function KCopyright() {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="https://material-ui.com/">
                Kadvisor
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
}
