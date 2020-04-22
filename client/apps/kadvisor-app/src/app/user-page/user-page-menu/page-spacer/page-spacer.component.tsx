import React from 'react';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { KCopyright } from '@client/klibs';
import PropTypes from 'prop-types';

export default function PageSpacer(props: any) {
    return (
        <div>
            <div className={props.classes.appBarSpacer} />
            <Container maxWidth="lg" className={props.classes.container}>
                {props.children}
                <Box pt={4}>
                    <KCopyright />
                </Box>
            </Container>
        </div>
    );
}

PageSpacer.propTypes = {
    children: PropTypes.node,
    classes: PropTypes.object
};
