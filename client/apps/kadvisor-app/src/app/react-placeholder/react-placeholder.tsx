import React, { Component, CSSProperties } from 'react';
import ReactPlaceholderStore from './react-placeholder.store';
import { connect } from 'react-redux';
import klogo from '../../assets/klogo.png';
import { Button } from 'react-bootstrap';

class ReactPlaceholder extends Component<ReactPlaceholderPropTypes> {
    showTestStoreButton = false;

    getTestStoreButton = () => {
        return (
            <Button
                variant={'outline-danger'}
                onClick={() => this.props.testStoreFunc('HELLO TEST')}
            >
                Add Store with Redux
            </Button>
        );
    };

    render() {
        return (
            <div className="Main-header">
                <img
                    src={klogo}
                    style={logoStyle}
                    className="Main-logo"
                    alt="logo"
                />
                <p>
                    This is <code>kadvisor</code> home-page.
                </p>
                {this.showTestStoreButton ? this.getTestStoreButton : null}
            </div>
        );
    }
}

interface ReactPlaceholderPropTypes {
    getTestFromStore: string;
    testStoreFunc: Function;
}

const logoStyle = {
    width: '10%',
    height: '10%'
} as CSSProperties;

const store = new ReactPlaceholderStore();
export default connect(
    store.mapStateToProps,
    store.mapDispatchToProps
)(ReactPlaceholder);
