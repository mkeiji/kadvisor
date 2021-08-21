import {
    ForecastTableState,
    KSelectItem,
    ReportsApiService
} from '@client/klibs';
import { ClassNameMap } from '@material-ui/core/styles/withStyles';
import ForecastTableViewModelService from './forecast-table-view-model.service';
import ForecastTableService from './forecast-table.service';

export interface ForecastTablePropsType {
    userID: number;
    classes: ClassNameMap<any>;
    service?: ForecastTableService;
    reportsService?: ReportsApiService;
    viewModelService?: ForecastTableViewModelService;
}

export interface ForecastTableComponentState {
    yearMenuItems: KSelectItem[];
    loading: boolean;
    hasForecast: boolean;
    selectedYear: number;
    forecastYear: number | string;
    table: ForecastTableState;
}
