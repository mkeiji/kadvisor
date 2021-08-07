import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import Settings from './settings.component';
import PageSpacer from '../page-spacer/page-spacer.component';
import ForecastTable from './forecast-table/forecast-table.component';
import ClassesTable from './classes-table/classes-table.component';

describe('Settings', () => {
    let wrapper: ShallowWrapper;

    const testClasses = {};
    const testID = 1;

    beforeEach(() => {
        wrapper = shallow(<Settings userID={testID} classes={testClasses} />);
    });

    it('should render a PageSpacer', () => {
        expect(wrapper.find(PageSpacer)).toHaveLength(1);
    });

    it('should render a ForecastTable', () => {
        expect(wrapper.find(ForecastTable)).toHaveLength(1);
    });

    it('should render a ClassesTable', () => {
        expect(wrapper.find(ClassesTable)).toHaveLength(1);
    });
});
