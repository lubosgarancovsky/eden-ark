import axios from "axios";

const baseURL = import.meta.env.DEV ? 'http://localhost:9091' : undefined;

export const httpClient = axios.create(
    { baseURL }
);