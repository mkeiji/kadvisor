import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import UserPageMenu from './user-page-menu.component';
import {
    AppBar,
    Badge,
    CssBaseline,
    Drawer,
    Toolbar,
    Typography
} from '@material-ui/core';
import NotificationsIcon from '@material-ui/icons/Notifications';
import MenuListItem from './menu-list-item/menu-list-item.component';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';

describe('UserPageMenu', () => {
    let component: UserPageMenu;

    const testUserID = 1;
    const testClass = {
        root: 'testRoot',
        appBar: 'testAppBar',
        appBarShift: 'testAppBarShift',
        toolbar: 'testToolbar',
        menuButton: 'menuButton',
        menuButtonHidden: 'menuButtonHidden',
        title: 'testTitle',
        toolbarIcon: 'testToolbarIcon',
        drawerPaper: 'testDrawerPaper',
        content: 'testContent'
    };
    const props = {
        userID: testUserID,
        title: 'Dashboard',
        classes: testClass
    };

    beforeEach(() => {
        component = new UserPageMenu(props);
    });

    describe('constructor', () => {
        it('should set initial state', () => {
            const expectedState = { open: true };
            expect(component.state).toEqual(expectedState);
        });
    });

    describe('handleDrawerOpen', () => {
        it('should set state to true', () => {
            jest.spyOn(component, 'setState').mockImplementation();
            const expectedState = { open: true };

            component.handleDrawerOpen();
            expect(component.setState).toHaveBeenCalledWith(expectedState);
        });
    });

    describe('handleDrawerClose', () => {
        it('should set state to false', () => {
            jest.spyOn(component, 'setState').mockImplementation();
            const expectedState = { open: false };

            component.handleDrawerClose();
            expect(component.setState).toHaveBeenCalledWith(expectedState);
        });
    });

    describe('renders', () => {
        let renderWrapper: ShallowWrapper;

        beforeEach(() => {
            renderWrapper = shallow(
                <UserPageMenu
                    userID={testUserID}
                    title={'Dashboard'}
                    classes={testClass}
                />
            );
        });

        it('should render a CssBaseline component', () => {
            expect(renderWrapper.find(CssBaseline)).toHaveLength(1);
        });

        it('should render AppBar component', () => {
            const expectedPosition = 'absolute';
            const expectedClasses = `${testClass.appBar} ${testClass.appBarShift}`;

            expect(renderWrapper.find(AppBar)).toHaveLength(1);
            expect(renderWrapper.find(AppBar).prop('position')).toEqual(
                expectedPosition
            );
            expect(renderWrapper.find(AppBar).prop('className')).toEqual(
                expectedClasses
            );
        });

        it('should render Toolbar', () => {
            expect(renderWrapper.find(Toolbar)).toHaveLength(1);
            expect(renderWrapper.find(Toolbar).prop('className')).toEqual(
                testClass.toolbar
            );
        });

        it('should render drawer-open-button', () => {
            const expectedClasses = `${testClass.menuButton} ${testClass.menuButtonHidden}`;
            const drawerOpenButton = renderWrapper.findWhere(
                (n) => n.prop('id') === 'drawer-open-button'
            );

            expect(drawerOpenButton).toHaveLength(1);
            expect(drawerOpenButton.prop('edge')).toEqual('start');
            expect(drawerOpenButton.prop('color')).toEqual('inherit');
            expect(drawerOpenButton.prop('className')).toEqual(expectedClasses);
        });

        it('should render typography', () => {
            const expectedClasses = testClass.title;
            const typography = renderWrapper.find(Typography);

            expect(typography).toHaveLength(1);
            expect(typography.prop('component')).toEqual('h1');
            expect(typography.prop('variant')).toEqual('h6');
            expect(typography.prop('color')).toEqual('inherit');
            expect(typography.prop('className')).toEqual(expectedClasses);
            expect(typography.prop('noWrap')).toBeTruthy();
        });

        it('should render notification-button', () => {
            const notificationButton = renderWrapper.findWhere(
                (n) => n.prop('id') === 'notification-button'
            );
            const badge = renderWrapper.find(Badge);
            const notificationIcon = renderWrapper.find(NotificationsIcon);

            expect(notificationButton).toHaveLength(1);
            expect(notificationIcon).toHaveLength(1);
            expect(badge).toHaveLength(1);
            expect(badge.prop('badgeContent')).toEqual(4);
            expect(badge.prop('color')).toEqual('secondary');
        });

        it('should render a Drawer component', () => {
            const expectedClasses = { paper: testClass.drawerPaper };
            const drawer = renderWrapper.find(Drawer);

            expect(drawer).toHaveLength(1);
            expect(drawer.prop('variant')).toEqual('permanent');
            expect(drawer.prop('classes')).toEqual(expectedClasses);
            expect(drawer.prop('open')).toBeTruthy();
        });

        it('should render a drawer-close-button', () => {
            const closeButtonContainer = renderWrapper.findWhere(
                (n) => n.prop('id') === 'drawer-close-button-container'
            );
            const drawerCloseButton = renderWrapper.findWhere(
                (n) => n.prop('id') === 'drawer-close-button'
            );
            const menuListItem = renderWrapper.find(MenuListItem);

            expect(closeButtonContainer).toHaveLength(1);
            expect(closeButtonContainer.prop('className')).toEqual(
                testClass.toolbarIcon
            );

            expect(drawerCloseButton).toHaveLength(1);
            expect(renderWrapper.find(ChevronLeftIcon)).toHaveLength(1);

            expect(menuListItem).toHaveLength(1);
            expect(menuListItem.prop('userID')).toEqual(props.userID);
        });

        it('should render a main tag for children components', () => {
            const contentContainer = renderWrapper.findWhere(
                (n) => n.prop('id') === 'children-content'
            );
            expect(contentContainer).toHaveLength(1);
            expect(contentContainer.prop('className')).toEqual(
                testClass.content
            );
        });
    });

    describe('onClicks', () => {
        let onClickWrapper: ShallowWrapper;

        beforeEach(() => {
            jest.spyOn(component, 'handleDrawerOpen').mockImplementation();
            jest.spyOn(component, 'handleDrawerClose').mockImplementation();
            onClickWrapper = shallow(component.render());
        });

        it('should call handleDrawerOpen', () => {
            const drawerOpenButton = onClickWrapper.findWhere(
                (n) => n.prop('id') === 'drawer-open-button'
            );

            drawerOpenButton.simulate('click');
            expect(component.handleDrawerOpen).toHaveBeenCalled();
        });

        it('should call handleDrawerClose', () => {
            const drawerOpenButton = onClickWrapper.findWhere(
                (n) => n.prop('id') === 'drawer-close-button'
            );

            drawerOpenButton.simulate('click');
            expect(component.handleDrawerClose).toHaveBeenCalled();
        });
    });
});
