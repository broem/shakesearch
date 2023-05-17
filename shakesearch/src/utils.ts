export function escapeRegExp(string: string) {
    return string.replace(/[.*+\-?^${}()|[\]\\]/g, "\\$&"); // escape special characters
  }
  
  export function highlightQuery(result: string, query: string) {
    const regex = new RegExp(`(${escapeRegExp(query)})`, "gi");
    return result.replace(regex, "<mark>$1</mark>");
  }