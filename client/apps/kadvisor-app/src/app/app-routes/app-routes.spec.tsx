import React from 'react';
import { render } from '@testing-library/react';

import AppRoutes from './app-routes';

xdescribe(' AppRoutes', () => {
    it('should render successfully', () => {
        const { baseElement } = render(<AppRoutes />);
        expect(baseElement).toBeTruthy();
    });
});
