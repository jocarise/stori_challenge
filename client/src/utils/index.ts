export const checkLocalStorageValueExists = (key: string): boolean => {
  if (typeof window !== "undefined") {
    // Check if in browser
    const value = localStorage.getItem(key);
    return value !== null;
  }

  return false;
};

export const setLocalStorageValue = <T>(key: string, value: T): void => {
  if (typeof window !== "undefined") {
    // Check if in browser
    localStorage.setItem(key, JSON.stringify(value));
  }
};

export const getLocalStorageValue = <T>(key: string): T | null => {
  if (typeof window !== "undefined") {
    // Check if in browser
    const value = localStorage.getItem(key);
    return value ? JSON.parse(value) : null;
  }
  return null; // Return null if not in browser
};

export const setCookie = (
  name: string,
  value: string,
  days: number = 7,
  path: string = "/"
): void => {
  const expires = new Date();
  expires.setTime(expires.getTime() + days * 24 * 60 * 60 * 1000); // Set expiration time
  const expiresString = `expires=${expires.toUTCString()}`;

  document.cookie = `${name}=${value}; ${expiresString}; path=${path};`;
};

export function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0"); // Months are zero-based
  const day = String(date.getDate()).padStart(2, "0");

  return `${year}-${month}-${day}`;
}
