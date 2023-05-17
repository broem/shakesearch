import React from 'react';
import { Card, CardContent, Pagination, Typography } from '@mui/material';
import Highlighter from "react-highlight-words";
import { SearchMatch } from '../types';

type SearchResultsProps = {
  results: SearchMatch[];
  query: string;
  totalResults: number;
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
};

const SearchResults: React.FC<SearchResultsProps> = ({ 
    results, 
    query,
    totalResults,
    currentPage,
    totalPages,
    onPageChange,    
}) => {
    const handleChange = (event: React.ChangeEvent<unknown>, value: number) => {
        onPageChange(value);
      };

    const resultsPerPage = Math.ceil(totalResults / totalPages);
    const firstResultNumber = (currentPage - 1) * resultsPerPage + 1;
    const lastResultNumber = firstResultNumber + results.length - 1;

  return (
    <div style={{ overflow: 'auto', marginTop: '1em' }}>
      {totalResults > 0 && (
        <Typography variant="h6" color="#F5F5F5">
          {`Showing results ${firstResultNumber}-${lastResultNumber} of ${totalResults}`}
        </Typography>
      )}
      <Pagination
        count={totalPages}
        page={currentPage}
        onChange={handleChange}
        sx={{
          display: 'flex',
          justifyContent: 'center',
          padding: '1em',
          '& .MuiPaginationItem-root': {
            color: '#F5F5F5',
          },
          '& .Mui-selected': {
            color: '#F5F5F5',
          },
        }}
      />
      {Array.isArray(results) && results.map((result, index) => (
        <Card key={index} style={{ margin: '1em' }}>
          <CardContent>
          <Typography variant="h6" color="text.secondary">
            {result.documentId}
          </Typography>
          <Highlighter
            highlightClassName="YourHighlightClass"
            searchWords={[query]}
            autoEscape={true}
            textToHighlight={result.context}
            />
          </CardContent>
        </Card>
      ))}
    </div>
  );
};

export default SearchResults;