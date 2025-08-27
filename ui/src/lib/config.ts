const appConfig: {
  baseUrl: string;
  mode: string;
} = {
  // https://vite.dev/guide/env-and-mode
  baseUrl: import.meta.env.MODE !== 'production' ? import.meta.env.VITE_BASE_URL as string : '',
  mode: import.meta.env.MODE,
};

export default appConfig;
