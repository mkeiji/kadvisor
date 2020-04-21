import React from 'react';
import clsx from 'clsx';
import CssBaseline from '@material-ui/core/CssBaseline';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import Badge from '@material-ui/core/Badge';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import NotificationsIcon from '@material-ui/icons/Notifications';
import PropTypes from 'prop-types';
import MenuListItem from './menu-list-item.component';

export default function UserPageMenu(props: any) {
    const [open, setOpen] = React.useState(true);
    const handleDrawerOpen = () => {
        setOpen(true);
    };
    const handleDrawerClose = () => {
        setOpen(false);
    };

    return (
        <div className={props.classes.root}>
            <CssBaseline />
            <AppBar
                position="absolute"
                className={clsx(
                    props.classes.appBar,
                    open && props.classes.appBarShift
                )}
            >
                <Toolbar className={props.classes.toolbar}>
                    <IconButton
                        edge="start"
                        color="inherit"
                        aria-label="open drawer"
                        onClick={handleDrawerOpen}
                        className={clsx(
                            props.classes.menuButton,
                            open && props.classes.menuButtonHidden
                        )}
                    >
                        <MenuIcon />
                    </IconButton>

                    <Typography
                        component="h1"
                        variant="h6"
                        color="inherit"
                        noWrap
                        className={props.classes.title}
                    >
                        {props.title}
                    </Typography>

                    <IconButton color="inherit">
                        <Badge badgeContent={4} color="secondary">
                            <NotificationsIcon />
                        </Badge>
                    </IconButton>
                </Toolbar>
            </AppBar>
            <Drawer
                variant="permanent"
                classes={{
                    paper: clsx(
                        props.classes.drawerPaper,
                        !open && props.classes.drawerPaperClose
                    )
                }}
                open={open}
            >
                <div className={props.classes.toolbarIcon}>
                    <IconButton onClick={handleDrawerClose}>
                        <ChevronLeftIcon />
                    </IconButton>
                </div>

                <MenuListItem userID={props.userID} />
            </Drawer>
            <main className={props.classes.content}>{props.children}</main>
        </div>
    );
}

UserPageMenu.propTypes = {
    userID: PropTypes.number,
    title: PropTypes.string,
    children: PropTypes.node,
    classes: PropTypes.object
};
