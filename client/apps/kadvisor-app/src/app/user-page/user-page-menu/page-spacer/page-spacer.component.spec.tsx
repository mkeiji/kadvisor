import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import PageSpacer from './page-spacer.component';
import { Container } from '@material-ui/core';
import { KCopyright } from '@client/klibs';

describe('PageSpacer', () => {
    let wrapper: ShallowWrapper;

    const testClasses = {};
    const testText = 'Hello Test';
    const testNode = (
        <div>
            <p>{testText}</p>
        </div>
    );

    beforeEach(() => {
        wrapper = shallow(
            <PageSpacer classes={testClasses}>{testNode}</PageSpacer>
        );
    });

    it('should render a Container', () => {
        expect(wrapper.find(Container)).toHaveLength(1);
    });

    it('should render passed content', () => {
        const expected = `${testText}<KCopyright />`;
        expect(wrapper.text()).toEqual(expected);
    });

    it('should render a KCopyright component', () => {
        expect(wrapper.find(KCopyright)).toHaveLength(1);
    });
});
