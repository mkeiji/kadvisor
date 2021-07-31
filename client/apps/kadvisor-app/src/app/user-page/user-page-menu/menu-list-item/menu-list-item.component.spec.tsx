import React from 'react';
import { shallow, ShallowWrapper } from 'enzyme';
import MenuListItem from './menu-list-item.component';
import { Divider } from '@material-ui/core';
import DashboardIcon from '@material-ui/icons/Dashboard';
import AddIcon from '@material-ui/icons/Add';
import LayersIcon from '@material-ui/icons/Layers';
import BarChartIcon from '@material-ui/icons/BarChart';
import AssignmentIcon from '@material-ui/icons/Assignment';

describe('MenuListItem', () => {
    let wrapper: ShallowWrapper;

    const testID = 1;
    const props = {
        userID: testID
    };

    beforeEach(() => {
        wrapper = shallow(<MenuListItem userID={props.userID} />);
    });

    it('rendes divider components', () => {
        expect(wrapper.find(Divider)).toHaveLength(2);
    });

    it('rendes dashboard icon component', () => {
        expect(wrapper.find(DashboardIcon)).toHaveLength(1);
    });

    it('rendes addIcon component', () => {
        expect(wrapper.find(AddIcon)).toHaveLength(1);
    });

    it('rendes layers icon component', () => {
        expect(wrapper.find(LayersIcon)).toHaveLength(1);
    });

    it('rendes barchart icon component', () => {
        expect(wrapper.find(BarChartIcon)).toHaveLength(1);
    });

    it('rendes assignment icon component', () => {
        expect(wrapper.find(AssignmentIcon)).toHaveLength(1);
    });
});
