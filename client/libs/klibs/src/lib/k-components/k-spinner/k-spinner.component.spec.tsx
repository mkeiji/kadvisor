import React from 'react';
import { render } from '@testing-library/react';
import { KSpinner } from './k-spinner.component';

describe('KSpinner', () => {
    it('should render successfully', () => {
        const { baseElement } = render(<KSpinner />);
        expect(baseElement).toBeTruthy();
    });
});
