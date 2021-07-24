import React from 'react';
import { mount, ReactWrapper, shallow } from 'enzyme';
import { of } from 'rxjs';
import { mocked } from 'ts-jest/utils';
import BalanceCard from './balance-card.component';
import BalanceCardService from './balance-card.service';
import { Link, Typography } from '@material-ui/core';
import { KFormatUtil, UserBalance } from '@client/klibs';
jest.mock('./balance-card.service');

describe('', () => {
    let useEffect: any;
    let mockService: BalanceCardService;
    let testBalance: UserBalance;

    const testID = 1;
    const mockServiceMgr = mocked(BalanceCardService, true);
    const mockUseEffect = () => {
        useEffect.mockImplementationOnce((f: any) => f());
    };

    beforeEach(() => {
        mockServiceMgr.mockClear();
        testBalance = {
            userID: testID,
            balance: 10
        };
        mockService = ({
            getUserBalance: jest.fn(() => of(testBalance))
        } as unknown) as BalanceCardService;

        mockServiceMgr.mockImplementationOnce(() => mockService);
    });

    afterAll(() => {
        jest.unmock('./balance-card.service');
    });

    describe('create service', () => {
        beforeEach(() => {
            useEffect = jest.spyOn(React, 'useEffect');
            mockUseEffect();
        });

        it('should test', () => {
            shallow(<BalanceCard userID={testID}></BalanceCard>);
            expect(BalanceCardService).toHaveBeenCalled();
        });
    });

    describe('renders', () => {
        let wrapper: ReactWrapper;

        beforeEach(() => {
            wrapper = mount(<BalanceCard userID={testID}></BalanceCard>);
        });

        it('Balance Card', () => {
            const expectedBalance = '$10.00';

            expect(wrapper.find(Typography).first().text()).toEqual(
                'Balance Card'
            );
            expect(wrapper.find('#balanceCard').first().text()).toEqual(
                expectedBalance
            );
        });

        it('Balance date', () => {
            const expectedDate = `on ${KFormatUtil.dateDisplayFormat(
                new Date()
            )}`;
            expect(wrapper.find('#balanceDate').first().text()).toEqual(
                expectedDate
            );
        });

        it('Report Link', () => {
            expect(wrapper.find(Link)).toHaveLength(1);
        });
    });
});
