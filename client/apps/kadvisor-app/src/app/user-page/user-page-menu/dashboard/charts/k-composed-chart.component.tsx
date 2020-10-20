import React, { Fragment, useEffect, useState } from 'react';
import {
    ComposedChart,
    Area,
    Bar,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer
} from 'recharts';
import { MonthReport, ReportsApiService } from '@client/klibs';
import { takeUntil } from 'rxjs/operators';
import { Subject } from 'rxjs';
import ChartsViewModelService from './charts-view-model.service';

interface ChartPropsType {
    userID: number;
    year: number;
}

export default function KComposedChartComponent(props: ChartPropsType) {
    const destroy$ = new Subject<boolean>();
    const viewModelService = new ChartsViewModelService();
    const service = new ReportsApiService(props.userID);
    const [data, setData] = useState<MonthReport[]>([]);
    const [minValue, setMinValue] = useState(0);

    useEffect(() => {
        service
            .getYtdWithForecastReport(props.year)
            .pipe(takeUntil(destroy$))
            .subscribe(
                (m: MonthReport[]) => {
                    const min = Math.min.apply(
                        Math,
                        m.map((obj) => obj.balance)
                    );

                    //NOTE: expenses need to be converted to positive values
                    m.map((obj) => (obj.expense = obj.expense * -1));

                    setMinValue(min > 0 ? 0 : min);
                    setData(m);
                },
                () => setData(viewModelService.getEmptyMonthReport())
            );

        return () => {
            destroy$.next(true);
            destroy$.unsubscribe();
        };
    }, [props.year]);

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
                    <XAxis dataKey="month" />
                    <YAxis
                        yAxisId="left"
                        orientation="left"
                        type="number"
                        domain={[minValue, 'auto']}
                        label={{
                            value: 'Inc / Exp',
                            angle: -90,
                            position: 'insideLeft'
                        }}
                    />
                    <YAxis
                        yAxisId="right"
                        orientation="right"
                        type="number"
                        domain={[minValue, 'auto']}
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
