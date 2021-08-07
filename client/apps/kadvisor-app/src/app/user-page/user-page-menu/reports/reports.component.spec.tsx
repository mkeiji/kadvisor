import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import Reports from './reports.component';
import PageSpacer from '../page-spacer/page-spacer.component';

describe('Reports', () => {
    let wrapper: ShallowWrapper;

    const testClasses = {};
    const testID = 1;

    beforeEach(() => {
        wrapper = shallow(<Reports userID={testID} classes={testClasses} />);
    });

    it('should render a PageSpacer', () => {
        expect(wrapper.find(PageSpacer)).toHaveLength(1);
    });
});
