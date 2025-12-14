import React, { useState, useRef } from "react";
import { Button, Popper, Paper, Box, ClickAwayListener, Typography, IconButton } from "@mui/material";
import { Service } from "src/constansts/capture";

export interface ServiceSelectProps {
  addService: (service:Service) => void;
  selectedServices: Service[];
  options: Service[];
}

export const ServiceSelect: React.FC<ServiceSelectProps> = ({ options, addService,selectedServices }) => {
  //const [selectedServices, setSelectedServices] = useState<Service[]>([]);
  const [open, setOpen] = useState(false);
  const anchorRef = useRef<HTMLButtonElement | null>(null);

  const handleClickAway = () => setOpen(false);

  return (
    <ClickAwayListener onClickAway={handleClickAway}>
      <Box>
        <Button
          ref={anchorRef}
          variant="contained"
          onClick={() => setOpen((prev) => !prev)}
        >
          Select Services
        </Button>

        <Popper open={open} anchorEl={anchorRef.current} placement="bottom-start" style={{ zIndex: 1300 }}>
          <Paper style={{ maxHeight: 250, overflowY: "auto", marginTop: 4, minWidth: 150 }}>
            <Box display="flex" flexDirection="column">
              {options.map((service) => {
                const isSelected = selectedServices.some((s) => s.value === service.value);
                return (
                  <Box
                    key={service.value}
                    onClick={() => addService(service)}
                    sx={{
                      padding: 1,
                      cursor: "pointer",
                      backgroundColor: isSelected ? "primary.main" : "transparent",
                      color: isSelected ? "white" : "text.primary",
                      "&:hover": {
                        backgroundColor: isSelected ? "primary.dark" : "action.hover",
                      },
                    }}
                  >
                    <Typography variant="body1">{service.label}</Typography>
                  </Box>
                );
              })}
            </Box>
          </Paper>
        </Popper>
      </Box>
    </ClickAwayListener>
  );
};

