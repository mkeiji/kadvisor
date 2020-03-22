import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Component } from 'react';
import './Main.css';
import MainNavBar from './main-nav-bar/main-nav-bar.component';
import ReactPlaceholder from './react-placeholder/react-placeholder.component';

class Main extends Component {
    render() {
        return (
            <Router>
                <div className="Main">
                    <MainNavBar />
                    <Route exact path="/" component={ReactPlaceholder} />
                    <Route path="/about" render={ () => (
                        <React.Fragment>
                            <h1 style={{paddingTop: '100px'}}>ABOUT PAGE</h1>
                        </React.Fragment>
                    )} />
                </div>
            </Router>
        );
    }
}

export default Main;
