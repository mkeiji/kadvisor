import React, { PureComponent, Fragment } from 'react';
import {
    ComposedChart,
    Line,
    Area,
    Bar,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer
} from 'recharts';

const data = [
    {
        name: 'Jan',
        expense: 590,
        balance: 18000,
        income: 1400,
        cnt: 490
    },
    {
        name: 'Feb',
        expense: 868,
        balance: 19670,
        income: 1506,
        cnt: 590
    },
    {
        name: 'Mar',
        expense: 1397,
        balance: 10980,
        income: 989,
        cnt: 350
    },
    {
        name: 'Apr',
        expense: 1480,
        balance: 14000,
        income: 1228,
        cnt: 480
    },
    {
        name: 'May',
        expense: 1520,
        balance: 19080,
        income: 1100,
        cnt: 460
    },
    {
        name: 'Jun',
        expense: 1400,
        balance: 6800,
        income: 1700,
        cnt: 380
    }
];

export default class KComposedChartComponent extends PureComponent {
    render() {
        return (
            <Fragment>
                <ResponsiveContainer>
                    <ComposedChart
                        width={500}
                        height={400}
                        data={data}
                        margin={{
                            top: 20,
                            right: 20,
                            bottom: 20,
                            left: 20
                        }}
                    >
                        <CartesianGrid stroke="#f5f5f5" />
                        <XAxis dataKey="name" />
                        <YAxis
                            yAxisId="left"
                            orientation="left"
                            label={{
                                value: 'Inc / Exp',
                                angle: -90,
                                position: 'insideLeft'
                            }}
                        />
                        <YAxis
                            yAxisId="right"
                            orientation="right"
                            label={{
                                value: 'Balance',
                                angle: -90,
                                position: 'insideRight'
                            }}
                        />
                        <Tooltip />
                        <Legend />
                        <Area
                            yAxisId="left"
                            type="monotone"
                            dataKey="income"
                            fill="#8884d8"
                            stroke="#8884d8"
                        />
                        <Area
                            yAxisId="left"
                            type="monotone"
                            dataKey="expense"
                            fill="#fcada1"
                            stroke="#fcada1"
                        />
                        <Bar
                            yAxisId="right"
                            dataKey="balance"
                            barSize={20}
                            fill="#413ea0"
                        />
                        {/* <Scatter dataKey="cnt" fill="red" /> */}
                    </ComposedChart>
                </ResponsiveContainer>
            </Fragment>
        );
    }
}
