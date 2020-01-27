import React from 'react';
import { Component } from 'react';
import KRxios from 'klibs/krxios'
import logo from './logo.svg';
import './Main.css';
import Button from 'react-bootstrap/Button';

class Main extends Component {
    // TODO: DELETE this api call example.
    http = new KRxios('https://jsonplaceholder.typicode.com');
    users = this.http.get('/users');
    componentDidMount() {
        this.users.subscribe(
            response => {
                console.log(response);
            },
            err => {
                console.error(err);
            }
        );
    };
    
    componentWillUnmount() {
        this.users.unsubscribe();
    }

    render() {
        return (
            <div className="Main">
                <header className="Main-header">
                <img src={logo} className="Main-logo" alt="logo" />
                <p>
                    Edit <code>src/App.js</code> and save to reload.
                </p>
                <a
                    className="Main-link"
                    href="https://reactjs.org"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Learn React
                </a>
                <Button variant="success">Success</Button>
                </header>
            </div>
        );
    }
}

export default Main;
