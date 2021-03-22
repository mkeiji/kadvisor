import React from 'react';
import { render } from '@testing-library/react';

import ReactPlaceholder from './react-placeholder';

xdescribe(' ReactPlaceholder', () => {
    it('should render successfully', () => {
        const { baseElement } = render(<ReactPlaceholder />);
        expect(baseElement).toBeTruthy();
    });
});
