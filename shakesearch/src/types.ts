export type SearchMatch = {
    documentId: string;
    position: number;
    context: string;
};

export interface SearchResponse {
    Results: SearchMatch[];
    TotalResults: number;
    Page: number;
    TotalPages: number;
};