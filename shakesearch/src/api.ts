const host = process.env.REACT_APP_API_HOST || 'https://bleach-shakesearch.onrender.com/';

const search = async (query: string, page: number = 1, selectedWorks: string[] = []): Promise<any> => {
  try {
    const worksQuery = selectedWorks.length > 0 ? `&works=${selectedWorks.join(',')}` : '';
    const response = await fetch(`${host}search?q=${query}&page=${page}${worksQuery}`);

    if(!response.ok) throw new Error('Error fetching search data');

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    return [];
  }
  };
  
  const suggest = async (query: string): Promise<any> => {
    try{
    const response = await fetch(`${host}suggest?q=${query}`);

    if(!response.ok) throw new Error('Error fetching suggest data');

    const data = await response.json();
    return data;
    } catch (error) {
        console.error(error);
        return [];
    }
  };

export const getWorks = async (): Promise<any> => {
  try{
    const response = await fetch(`${host}documents`);

    if(!response.ok) throw new Error('Error fetching getworks data');

    const data = await response.json();
    return data;
  } catch (error) {
      console.error(error);
      return [];
  }
};
  
  export default {
    search,
    suggest,
  };