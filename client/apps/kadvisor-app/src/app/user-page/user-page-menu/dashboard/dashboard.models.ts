import { KSelectItem, MonthReport, ReportsApiService } from '@client/klibs';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import EntryService from '../entry/entry.service';
import ChartsViewModelService from './charts/charts-view-model.service';
import DashEntriesViewModelService from './entries/dash-entries-view-model.service';
import { DashEntryRow } from './entries/view-model';

export interface DashboardState {
    graphYear: number;
    showYearDropdown: boolean;
    yearMenuItems: KSelectItem[];
}

export interface DashboardPropsType {
    userID: number;
    classes: ClassNameMap<any>;
    reportsApiService?: ReportsApiService;
}

export interface KComposedChartPropsType {
    userID: number;
    year: number;
    service?: ReportsApiService;
    viewModelService?: ChartsViewModelService;
}

export interface KComposedChartState {
    data: MonthReport[];
    minDomain: number;
    leftTicks: number[] | undefined;
    rightTicks: number[] | undefined;
}

export interface DashboardEntriesPropsType {
    userID: number;
    service?: EntryService;
    viewModelService?: DashEntriesViewModelService;
}

export interface DashboardEntriesState {
    rows: DashEntryRow[];
}
