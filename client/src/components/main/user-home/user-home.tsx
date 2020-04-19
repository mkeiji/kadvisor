import React, {useEffect, useState} from 'react';
import UserHomeStore from "./user-home.store";
import {connect} from "react-redux";

function UserHome(props: UserHomePropTypes) {
    const paramID = Number(props.match.params.id);
    const login = props.getLoginStore ? props.getLoginStore.login : {};
    let content;

    const [idMatch, setIdMatch] = useState(false);
    useEffect(() => {
        handleIdChange(paramID);

        function handleIdChange(newID: any) {
            setIdMatch(newID === login.userID);
        }
    });

    content = idMatch ?
        <h1 style={{paddingTop: '100px'}}>USER HOME PAGE</h1> :
        <h1 style={{paddingTop: '100px', color: 'red'}}>PLEASE LOGIN</h1>;

    return (
        <div>
            {content}
        </div>
    );
}

interface UserHomePropTypes {
    match: any;
    getLoginStore: any;
}

const store = new UserHomeStore();
export default connect(store.mapStateToProps)(UserHome);