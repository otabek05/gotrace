import React from 'react';
import { FormControl, InputLabel, Select, MenuItem, SelectChangeEvent, SxProps, Theme } from '@mui/material';

interface Option<T extends string | number> {
  label: string;
  value: T;
}

interface CustomSelectProps<T extends string | number> {
  label: string;
  value: T | '';
  options: Option<T>[];
  onChange: (value: T) => void;
  size?: 'small' | 'medium';
  minWidth?: number;
  sx?: SxProps<Theme>;
}

export function CustomSelect<T extends string | number>({
  label,
  value,
  options,
  onChange,
  size = 'small',
  minWidth = 120,
  sx={}
}: CustomSelectProps<T>) {
  return (
    <FormControl size={size}  sx={{ minWidth, ...sx }}>
      <InputLabel>{label}</InputLabel>
      <Select
        value={value !== '' ? String(value) : ''} // cast value to string
        label={label}
        onChange={(e: SelectChangeEvent<string>) => {
          const rawValue = e.target.value;
          const selectedValue =
            typeof options[0].value === 'number' ? Number(rawValue) : rawValue;

          onChange(selectedValue as T);
        }}
      >
        {options.map((opt) => (
          <MenuItem key={opt.value} value={String(opt.value)}>
            {opt.label}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}