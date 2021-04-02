import React from 'react';
import { shallow } from 'enzyme';
import { KCopyright } from './k-copyright.component';
import { Link, Typography } from '@material-ui/core';

describe('KCopyright', () => {
    it('should render Typography', () => {
        const wrapper = shallow(<KCopyright />);
        expect(wrapper.find(Typography)).toHaveLength(1);
    });

    it('should render copyright text', () => {
        const wrapper = shallow(<KCopyright />);
        expect(wrapper.find(Typography).text()).toContain('Copyright');
        expect(wrapper.find(Typography).text()).toContain(
            `${new Date().getFullYear()}`
        );
    });

    it('should render Link', () => {
        const wrapper = shallow(<KCopyright />);
        expect(wrapper.find(Link)).toHaveLength(1);
        expect(wrapper.find(Link).text()).toBe('Kadvisor');
    });
});
