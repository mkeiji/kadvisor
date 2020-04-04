import React, {Component} from "react";
import { Button } from 'react-bootstrap';
import logo from 'components/main/logo.svg';
import PropTypes from "prop-types";
import ReactPlaceholderStore from "./react-placeholder.store";
import {connect} from "react-redux";

class ReactPlaceholder extends Component {
    render() {
        return (
            <div className="Main-header">
                <img src={logo} className="Main-logo" alt="logo" />
                <p>This is <code>kadvisor</code> home-page.</p>
                <a className="Main-link"
                   href="https://reactjs.org"
                   target="_blank"
                   rel="noopener noreferrer">Built on React</a>
                <Button variant="outline-danger"
                        onClick={
                            () => this.props.testStoreFunc("HELLO TEST")
                        }>Add Store with Redux</Button>
            </div>
        );
    }
}

ReactPlaceholder.propTypes = {
    testStoreFunc: PropTypes.func
};

const store = new ReactPlaceholderStore();
export default connect(store.mapStateToProps, store.mapDispatchToProps)(ReactPlaceholder);