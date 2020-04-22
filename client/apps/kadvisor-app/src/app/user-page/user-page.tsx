import React, { useEffect, useState } from 'react';
import UserPageStore from './user-page.store';
import { connect } from 'react-redux';
import UserPageMenu from './user-page-menu/user-page-menu.component';
import { useStyles } from './user-page-menu/user-page-menu.style.hook';
import Dashboard from './user-page-menu/dashboard/dashboard.component';
import EntryTable from './user-page-menu/entry/entry.component';

function UserPage(props: UserHomePropTypes) {
    const styleClasses = useStyles();
    const paramID = Number(props.match.params.id);
    const paramPage = props.match.params.page;
    const login = props.getLoginStore ? props.getLoginStore.login : {};

    const [idMatch, setIdMatch] = useState(false);
    useEffect(() => {
        handleIdChange(paramID);

        function handleIdChange(newID: any) {
            setIdMatch(newID === login.userID);
        }
    });

    function getPage(idMatch: boolean, page: string): JSX.Element {
        if (idMatch) {
            switch (page) {
                case 'dashboard':
                    return <Dashboard classes={styleClasses} />;

                case 'entries':
                    return <EntryTable classes={styleClasses} />;

                default:
                    return (
                        <h1 style={{ paddingTop: '100px' }}>PAGE NOT FOUND</h1>
                    );
            }
        } else {
            return (
                <h1 style={{ paddingTop: '100px', color: 'red' }}>
                    PLEASE LOGIN
                </h1>
            );
        }
    }

    return (
        <div>
            <UserPageMenu
                userID={login.userID}
                title={'Dashboard'}
                classes={styleClasses}
            >
                {getPage(idMatch, paramPage)}
            </UserPageMenu>
        </div>
    );
}

interface UserHomePropTypes {
    match: any;
    getLoginStore: any;
}

const store = new UserPageStore();
export default connect(store.mapStateToProps)(UserPage);
