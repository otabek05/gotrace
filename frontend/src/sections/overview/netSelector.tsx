import React from 'react';
import {
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
  SxProps,
  Theme,
} from '@mui/material';
import { NetworkInterface } from 'src/types/net_interface';

interface NetSelectorProps {
  label: string;
  value: NetworkInterface | null;
  interfaces: NetworkInterface[];
  onChange: (netInterface: NetworkInterface) => void;
  size?: 'small' | 'medium';
  minWidth?: number;
  sx?: SxProps<Theme>;
}

export const NetSelector: React.FC<NetSelectorProps> = ({
  label,
  value,
  interfaces,
  onChange,
  size = 'small',
  minWidth = 220,
  sx = {},
}) => {
  const handleChange = (e: SelectChangeEvent<string>) => {
    const selected = interfaces.find(
      (i) => i.Name === e.target.value
    );
    if (selected) {
      onChange(selected);
    }
  };

  return (
    <FormControl size={size} sx={{ minWidth, ...sx }}>
      <InputLabel>{label}</InputLabel>
      <Select
        label={label}
        value={value?.Name ?? ''}
        onChange={handleChange}
      >
        {interfaces.map((net) => (
          <MenuItem key={net.Name} value={net.Name}>
            {net.Name}
            {net.Addresses.length > 0 && (
              <span style={{ marginLeft: 8, color: '#888' }}>
                ({net.Addresses[0].IP})
              </span>
            )}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};
