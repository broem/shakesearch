import React, { useEffect, useState } from 'react';
import Header from './components/Header';
import './App.css';
import SearchBar from './components/SearchBar';
import SearchResults from './components/SearchResults';
import api from './api';
import { SearchMatch } from './types';
import { Grid } from '@mui/material';
import WorksFilter from './components/WorksFilter';

function App() {
  const [results, setResults] = useState<SearchMatch[]>([]);
  const [query, setQuery] = useState<string>('');
  const [totalResults, setTotalResults] = useState<number>(0);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [totalPages, setTotalPages] = useState<number>(0);
  const [selectedWorks, setSelectedWorks] = useState<string[]>([]);

  useEffect(() => {
    if (query) {
      api.search(query, currentPage, selectedWorks).then((data) => {
        console.log(data);
        setResults(data.results);
        setTotalResults(data.totalResults);
        setCurrentPage(data.page);
        setTotalPages(data.totalPages);
      });
    }
  }, [query, currentPage, selectedWorks]);

  const handleSearch = async (query: string) => {
    setQuery(query);
    setCurrentPage(1);
  };

  const handleFilterChange = (selectedWorks: string[]) => {  // Added handleFilterChange
    setSelectedWorks(selectedWorks);
    setCurrentPage(1);
  };

  return (
    <div className="App" style={{ display: 'flex', flexDirection: 'column', height: '100vh' }}>
      <Header />
      <Grid container spacing={2}>
        <Grid item xs={12} sm={12}>
      <SearchBar onSearch={handleSearch} />
      </Grid>
      <Grid item xs={12} sm={12}>
          <WorksFilter onFilterChange={handleFilterChange} />
        </Grid>
      </Grid>
      <SearchResults 
        results={results}
        query={query}
        totalResults={totalResults}
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}
      />
    </div>
  );
}

export default App;
