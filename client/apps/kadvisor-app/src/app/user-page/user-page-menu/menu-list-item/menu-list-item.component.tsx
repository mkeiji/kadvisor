import React from 'react';
import { UserPageMenuListObject, UserPageMenuListTypes } from '../view-models';
import { Link } from 'react-router-dom';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListItem from '@material-ui/core/ListItem';
import DashboardIcon from '@material-ui/icons/Dashboard';
import AddIcon from '@material-ui/icons/Add';
import BarChartIcon from '@material-ui/icons/BarChart';
import LayersIcon from '@material-ui/icons/Layers';
import Divider from '@material-ui/core/Divider';
import AssignmentIcon from '@material-ui/icons/Assignment';
import { APP_PAGES, KRouterPathUtil } from '@client/klibs';

export default function MenuListItem(props: MenuListItemPropTypes) {
    const LIST_ITEMS = [
        {
            iconComponent: <Divider />,
            type: UserPageMenuListTypes.divider
        },
        {
            itemText: 'Dashboard',
            iconComponent: <DashboardIcon />,
            pagePath: APP_PAGES.dashboard,
            type: UserPageMenuListTypes.primary
        },
        {
            itemText: 'Entries',
            iconComponent: <AddIcon />,
            pagePath: APP_PAGES.entries,
            type: UserPageMenuListTypes.primary
        },
        {
            itemText: 'Settings',
            iconComponent: <LayersIcon />,
            pagePath: APP_PAGES.settings,
            type: UserPageMenuListTypes.primary
        },
        {
            itemText: 'Reports',
            iconComponent: <BarChartIcon />,
            pagePath: APP_PAGES.reports,
            type: UserPageMenuListTypes.primary
        },
        {
            iconComponent: <Divider />,
            type: UserPageMenuListTypes.divider
        },
        {
            itemText: 'Item 1',
            iconComponent: <AssignmentIcon />,
            pagePath: 'other',
            type: UserPageMenuListTypes.secondary
        }
    ] as UserPageMenuListObject[];

    function createListItem(
        itemID: number,
        userID: number,
        listItem: UserPageMenuListObject
    ): JSX.Element {
        if (listItem.type !== UserPageMenuListTypes.divider) {
            return (
                <ListItem
                    key={itemID}
                    button
                    component={Link}
                    to={KRouterPathUtil.getUserPage(userID, listItem.pagePath)}
                >
                    <ListItemIcon>{listItem.iconComponent}</ListItemIcon>
                    <ListItemText primary={listItem.itemText} />
                </ListItem>
            );
        } else {
            return <Divider key={itemID} />;
        }
    }

    function generateMenuListItems(userID: number): JSX.Element[] {
        const result = [] as JSX.Element[];
        LIST_ITEMS.map((obj: UserPageMenuListObject, i: number) =>
            result.push(createListItem(i, userID, obj))
        );
        return result;
    }

    return <div>{generateMenuListItems(props.userID)}</div>;
}

interface MenuListItemPropTypes {
    userID: number;
}
