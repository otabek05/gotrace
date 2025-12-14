import React from 'react';
import { FormControl, InputLabel, Select, MenuItem, InputAdornment, TextField, SelectChangeEvent } from '@mui/material';

interface CustomTextFieldProps<T> {
  label: string;
  value: T;
onChange: (e: SelectChangeEvent<string>) => void; 
  data: Record<string, string>; // Map of string to string (label-value pairs)
  placeholder?: string;
  adornment?: string; // Optional icon or currency symbol for text field
}

const CustomTextField = <T,>({
  label,
  value,
  onChange,
  data,
}: CustomTextFieldProps<T>) => {
  
  return (
    <FormControl sx={{ minWidth: 150 }}>
      <InputLabel>{label}</InputLabel>
      <Select
        value={value as unknown as string}  // Cast value to string for Select component
        label={label}
        onChange={onChange}
      >
        {Object.entries(data).map(([key, labelValue]) => (
          <MenuItem key={key} value={key}>
            {labelValue}  {/* Display the value of the map */}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};

export default CustomTextField;
