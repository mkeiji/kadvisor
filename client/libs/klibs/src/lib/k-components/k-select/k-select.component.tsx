import React, { CSSProperties } from 'react';
import { FormControl, InputLabel, MenuItem, Select } from '@material-ui/core';
import { KSelectItem } from '@client/klibs';

interface KSelectPropsType {
    items: KSelectItem[];
    onValueChange: Function;
    value?: any;
    label?: string;
    class?: string;
    style?: CSSProperties;
    formVariant?: 'standard' | 'outlined' | 'filled';
}

export function KSelect(props: KSelectPropsType) {
    function getFormVariant(): 'standard' | 'outlined' | 'filled' {
        return props.formVariant ? props.formVariant : 'standard';
    }

    function getLabel(): JSX.Element {
        return <InputLabel>{props.label}</InputLabel>;
    }

    return (
        <FormControl
            variant={getFormVariant()}
            className={props.class}
            style={props.style}
        >
            {props.label ? getLabel() : null}
            <Select
                value={props.value}
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
