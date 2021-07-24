import React, { Component, Fragment } from 'react';
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
import {
    KComposedChartPropsType,
    KComposedChartState
} from '../dashboard.models';

export default class KComposedChartComponent extends Component<
    KComposedChartPropsType,
    KComposedChartState
> {
    service: ReportsApiService;
    viewModelService: ChartsViewModelService;
    destroy$ = new Subject<boolean>();

    constructor(readonly props: KComposedChartPropsType) {
        super(props);
        this.service = this.props.service
            ? this.props.service
            : new ReportsApiService(props.userID);
        this.viewModelService = this.props.viewModelService
            ? this.props.viewModelService
            : new ChartsViewModelService();

        this.state = {
            data: [],
            minDomain: 0,
            leftTicks: undefined,
            rightTicks: undefined
        };
    }

    componentDidMount() {
        this.loadData();
    }

    componentWillUnmount() {
        this.destroy$.next(true);
        this.destroy$.unsubscribe();
    }

    componentDidUpdate(prevProps: KComposedChartPropsType) {
        if (this.props.year !== prevProps.year) {
            this.loadData();
        }
    }

    render() {
        return (
            <Fragment>
                <ResponsiveContainer>
                    <ComposedChart
                        width={500}
                        height={400}
                        data={this.state.data}
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
                            domain={[this.state.minDomain, 'auto']}
                            ticks={this.state.leftTicks}
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
                            domain={[this.state.minDomain, 'auto']}
                            ticks={this.state.rightTicks}
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

    private loadData() {
        this.service
            .getYtdWithForecastReport(this.props.year)
            .pipe(takeUntil(this.destroy$))
            .subscribe(
                (m: MonthReport[]) => {
                    const minBalance = this.viewModelService.getMinBalance(m);

                    //NOTE: expenses need to be converted to positive values
                    m.map((obj) => (obj.expense = obj.expense * -1));

                    this.setState({
                        minDomain: minBalance > 0 ? 0 : minBalance,
                        data: m
                    });

                    const [
                        left,
                        right
                    ] = this.viewModelService.getTicksForNegativeBalance(m);
                    this.setState({
                        leftTicks: left,
                        rightTicks: right
                    });
                },
                () =>
                    this.setState({
                        data: this.viewModelService.getEmptyMonthReport()
                    })
            );
    }
}
