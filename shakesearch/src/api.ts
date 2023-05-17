import { env } from "process";

const host = env.REACT_APP_API_HOST || 'http://localhost:3001/';

const search = async (query: string, page: number = 1, selectedWorks: string[] = []): Promise<any> => {
    const worksQuery = selectedWorks.length > 0 ? `&works=${selectedWorks.join(',')}` : '';
    const response = await fetch(`${host}search?q=${query}&page=${page}${worksQuery}`);
    const data = await response.json();
    return data;
  };
  
  const suggest = async (query: string): Promise<any> => {
    const response = await fetch(`${host}suggest?q=${query}`);
    const data = await response.json();
    return data;
  };

export const getWorks = async (): Promise<any> => {
    const response = await fetch(`${host}documents`);
    const data = await response.json();
    return data;
    };
  
  export default {
    search,
    suggest,
  };