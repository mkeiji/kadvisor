import React, { CSSProperties } from 'react';
import { FormControl, InputLabel, MenuItem, Select } from '@material-ui/core';
import { KSelectItem } from '@client/klibs';

interface KSelectPropsType {
    label: string;
    items: KSelectItem[];
    onValueChange: Function;
    initialValue?: any;
    class?: string;
    style?: CSSProperties;
}

export function KSelect(props: KSelectPropsType) {
    return (
        <FormControl className={props.class} style={props.style}>
            <InputLabel>{props.label}</InputLabel>
            <Select
                value={props.initialValue}
                onChange={(event) =>
                    props.onValueChange(event.target.value as number)
                }
            >
                {props.items.map((i) => (
                    <MenuItem key={i.value} value={i.value}>
                        {i.displayValue}
                    </MenuItem>
                ))}
            </Select>
        </FormControl>
    );
}
