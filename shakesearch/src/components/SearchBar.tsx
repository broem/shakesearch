import React, { useState, useEffect } from 'react';
import { TextField, Autocomplete } from '@mui/material';
import api from '../api';

type SearchBarProps = {
  onSearch: (query: string) => void;
};

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [inputValue, setInputValue] = useState('')
  const [suggestions, setSuggestions] = useState<string[]>([]);

  useEffect(() => {
    const getSuggestions = async () => {
      if (inputValue.length > 0) {
        const suggestionsFromAPI = await api.suggest(inputValue);
        setSuggestions(suggestionsFromAPI);
      } else {
        setSuggestions([]);
      }
    };

    getSuggestions();
  }, [inputValue]);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    onSearch(inputValue);
  };

  return (
      <form onSubmit={handleSubmit} style={{ margin: '0 auto' }}>
        <Autocomplete
          options={suggestions}
          getOptionLabel={(option) => option}
          noOptionsText=""
          inputValue={inputValue}
          onInputChange={(event, newValue) => {
            setInputValue(newValue);
          }}
          onChange={(event, newValue) => {
            if (newValue !== null) {
              setInputValue(newValue);
              onSearch(newValue);
            }
          }}
          renderInput={(params) => (
            <TextField
              {...params}
              placeholder="Search..."
              variant="outlined"
              color="primary"
              fullWidth
              sx={{
                color: '#F5F5F5',
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
        />
      </form>
  );
};

export default SearchBar;