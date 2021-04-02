import React from 'react';
import { render } from '@testing-library/react';
import { KSelect } from './k-select.component';
import { KSelectItem } from '../../k-models/gerneric-models';

describe('KSelect', () => {
    const testItems = [
        {
            value: 1,
            displayValue: 'one'
        }
    ] as KSelectItem[];
    const testFn = jest.fn();

    it('should render successfully', () => {
        const { baseElement } = render(
            <KSelect label={'test'} items={testItems} onValueChange={testFn} />
        );
        expect(baseElement).toBeTruthy();
    });
});
