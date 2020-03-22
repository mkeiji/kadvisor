import React from "react";
import { Button } from 'react-bootstrap';
import logo from 'components/main/logo.svg';

export default function ReactPlaceholder() {
    return (
        <div className="Main-header">
            <img src={logo} className="Main-logo" alt="logo" />
            <p>This is <code>kadvisor</code> home-page.</p>
            <a className="Main-link"
                href="https://reactjs.org"
                target="_blank"
                rel="noopener noreferrer">Built on React</a>
            <Button variant="success">Success Btn</Button>
        </div>
    )
}