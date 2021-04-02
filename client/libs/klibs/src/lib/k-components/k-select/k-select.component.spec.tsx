import React, { CSSProperties } from 'react';
import { shallow } from 'enzyme';
import { KSelect } from './k-select.component';
import { KSelectItem } from '../../k-models/gerneric-models';
import { FormControl, InputLabel, MenuItem, Select } from '@material-ui/core';

describe('KSelect', () => {
    const testItems = [
        {
            value: 1,
            displayValue: 'one'
        }
    ] as KSelectItem[];
    const testFn = jest.fn();

    describe('getFormVariant', () => {
        it('should default variant to standard', () => {
            const testVariant = 'standard';
            const wrapper = shallow(
                <KSelect items={testItems} onValueChange={testFn} />
            );
            expect(wrapper.find(FormControl).props().variant).toBe(testVariant);
        });

        it('should set form variant', () => {
            const testVariant = 'outlined';
            const wrapper = shallow(
                <KSelect
                    items={testItems}
                    onValueChange={testFn}
                    formVariant={testVariant}
                />
            );
            expect(wrapper.find(FormControl).props().variant).toBe(testVariant);
        });
    });

    describe('getLabel', () => {
        it('should render InputLabel', () => {
            const text = 'test';
            const wrapper = shallow(
                <KSelect
                    label={text}
                    items={testItems}
                    onValueChange={testFn}
                />
            );
            expect(wrapper.find(InputLabel).text()).toBe(text);
        });

        it('should not render InputLabel', () => {
            const wrapper = shallow(
                <KSelect items={testItems} onValueChange={testFn} />
            );
            expect(wrapper.find(InputLabel)).toHaveLength(0);
        });
    });

    describe('return', () => {
        it('should render FormControl', () => {
            const wrapper = shallow(
                <KSelect
                    label={'test'}
                    items={testItems}
                    onValueChange={testFn}
                />
            );
            expect(wrapper.find(FormControl)).toHaveLength(1);
        });

        it('should set FormControl props', () => {
            const testVariant = 'standard';
            const testClassName = 'testClass';
            const testStyle = { display: 'flex' } as CSSProperties;

            const wrapper = shallow(
                <KSelect
                    label={'test'}
                    items={testItems}
                    onValueChange={testFn}
                    formVariant={testVariant}
                    class={testClassName}
                    style={testStyle}
                />
            );

            expect(wrapper.find(FormControl).props().variant).toBe(testVariant);
            expect(wrapper.find(FormControl).props().className).toBe(
                testClassName
            );
            expect(wrapper.find(FormControl).props().style).toBe(testStyle);
        });

        it('should render Select', () => {
            const testValue = 1;
            const wrapper = shallow(
                <KSelect
                    items={testItems}
                    onValueChange={testFn}
                    value={testValue}
                />
            );

            expect(wrapper.find(Select)).toHaveLength(1);
            expect(wrapper.find(Select).props().value).toBe(testValue);
        });

        it('should render MenuItem', () => {
            const testValue = 1;
            const wrapper = shallow(
                <KSelect
                    items={testItems}
                    onValueChange={testFn}
                    value={testValue}
                />
            );

            expect(wrapper.find(MenuItem)).toHaveLength(1);
            expect(wrapper.find(MenuItem).props().value).toBe(testValue);
            expect(wrapper.find(MenuItem).text()).toBe(
                testItems[0].displayValue
            );
        });
    });
});
