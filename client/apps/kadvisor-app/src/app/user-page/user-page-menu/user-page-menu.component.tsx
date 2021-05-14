import React, { Component } from 'react';
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
import MenuListItem from './menu-list-item/menu-list-item.component';

export default class UserPageMenu extends Component<
    UserPageMenuPropTypes,
    UserPageMenuState
> {
    constructor(readonly props: any) {
        super(props);
        this.state = { open: true };
        this.handleDrawerOpen = this.handleDrawerOpen.bind(this);
        this.handleDrawerClose = this.handleDrawerClose.bind(this);
    }

    handleDrawerOpen = () => {
        this.setState({ open: true });
    };

    handleDrawerClose = () => {
        this.setState({ open: false });
    };

    render() {
        return (
            <div className={this.props.classes.root}>
                <CssBaseline />
                <AppBar
                    position="absolute"
                    className={clsx(
                        this.props.classes.appBar,
                        this.state.open && this.props.classes.appBarShift
                    )}
                >
                    <Toolbar className={this.props.classes.toolbar}>
                        <IconButton
                            id="drawer-open-button"
                            edge="start"
                            color="inherit"
                            aria-label="open drawer"
                            onClick={this.handleDrawerOpen}
                            className={clsx(
                                this.props.classes.menuButton,
                                this.state.open &&
                                    this.props.classes.menuButtonHidden
                            )}
                        >
                            <MenuIcon />
                        </IconButton>

                        <Typography
                            component="h1"
                            variant="h6"
                            color="inherit"
                            noWrap
                            className={this.props.classes.title}
                        >
                            {this.props.title}
                        </Typography>

                        <IconButton id="notification-button" color="inherit">
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
                            this.props.classes.drawerPaper,
                            !this.state.open &&
                                this.props.classes.drawerPaperClose
                        )
                    }}
                    open={this.state.open}
                >
                    <div
                        id="drawer-close-button-container"
                        className={this.props.classes.toolbarIcon}
                    >
                        <IconButton
                            id="drawer-close-button"
                            onClick={this.handleDrawerClose}
                        >
                            <ChevronLeftIcon />
                        </IconButton>
                    </div>

                    <MenuListItem userID={this.props.userID} />
                </Drawer>
                <main
                    id="children-content"
                    className={this.props.classes.content}
                >
                    {this.props.children}
                </main>
            </div>
        );
    }
}

export interface UserPageMenuState {
    open: boolean;
}

export interface UserPageMenuPropTypes {
    userID: number;
    title: string;
    children: Node;
    classes: object;
}
