import React from "react";
import { Box, TextField, InputAdornment, IconButton } from "@mui/material";
import { Icon } from "@iconify/react";

interface SearchInputProps {
  value: string;
  onChange: (val: string) => void;
  onSearch?: () => void;
  placeholder?: string;
  disabled?: boolean;
  sx?: object;
}

const SearchInputWithButton: React.FC<SearchInputProps> = ({
  value,
  onChange,
  onSearch,
  placeholder = "PORT",
  disabled = false,
  sx = {},
}) => {
  return (
    <Box sx={{ display: "flex", alignItems: "center", gap: 1, ...sx }}>
      <TextField
        label={placeholder}
        inputMode="numeric"
        size="small"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === "Enter" && onSearch) onSearch();
        }}
        sx={{ width: 250 }}
        disabled={disabled}
        placeholder={placeholder}
        slotProps={
          onSearch
            ? {
                input: {
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        onClick={onSearch}
                        size="small"
                        disabled={disabled}
                      >
                        <Icon
                          icon="solar:add-square-bold"
                          width={25}
                          height={25}
                        />
                      </IconButton>
                    </InputAdornment>
                  ),
                },
              }
            : undefined
        }
      />
    </Box>
  );
};

export default SearchInputWithButton;
