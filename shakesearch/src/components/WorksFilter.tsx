import React, { useEffect, useState } from 'react';
import { Autocomplete, Chip, styled } from '@mui/material';
import { TextField } from '@mui/material';
import { getWorks } from '../api';

const Container = styled('div')(({ theme }) => ({
    maxHeight: '300px',
    overflowY: 'auto',
    '& .MuiAutocomplete-popper': {
      maxHeight: 'inherit !important',
    },
    '& .MuiAutocomplete-paper': {
      backgroundColor: theme.palette.background.paper,
      boxShadow: theme.shadows[2],
      borderRadius: theme.shape.borderRadius,
    },
  }));

export type WorksFilterProps = {
  onFilterChange: (selectedWorks: string[]) => void;
};

const WorksFilter: React.FC<WorksFilterProps> = ({ onFilterChange }) => {
  const [works, setWorks] = useState<string[]>([]);

  useEffect(() => {
    getWorks().then(setWorks);
  }, []);

  const handleChange = (event: React.ChangeEvent<{}>, value: string[]) => {
    onFilterChange(value);
  };

  return (
    <Container>
    <Autocomplete
      multiple
      options={works}
      onChange={handleChange}
      renderInput={(params) => (
        <TextField
          {...params}
          variant="outlined"
          label="Filter by work"
          placeholder="Select a work"
          sx={{
            '& .MuiAutocomplete-popper': {
                maxHeight: '300px',
                overflowY: 'auto',
              },
            '& .MuiInputLabel-root': {
                color: '#F5F5F5',
            },
            '& .MuiChip-root': {
              backgroundColor: '#282c34',
              color: '#F5F5F5',
              '&:hover': {
                backgroundColor: '#282c34',
              },
            },
            '& .MuiChip-label': {
              color: '#F5F5F5',
            },
            '& .MuiOutlinedInput-root': {
              color: '#F5F5F5',
              '& fieldset': {
                borderColor: '#F5F5F5',
              },
              '&:hover fieldset': {
                borderColor: '#F5F5F5',
              },
              '&.Mui-focused fieldset': {
                borderColor: '#F5F5F5',
              },
            },
          }}
        />
      )}
      renderTags={(value, getTagProps) =>
          value ? (
            <Chip
              label={value[0]}
              {...getTagProps({ index: 0 })}
              sx={{
                backgroundColor: '#282c34',
                color: '#F5F5F5',
              }}
            />
          ) : null
        }
    />
    </Container>
  );
};

export default WorksFilter;