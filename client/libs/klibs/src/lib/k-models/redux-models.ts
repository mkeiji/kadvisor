import { Action } from 'redux';

export interface StoreAction extends Action {
    type: string;
    payload: any;
}
