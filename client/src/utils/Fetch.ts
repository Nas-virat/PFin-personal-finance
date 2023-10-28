import axios from 'axios';

const BASE_URL = 'https://localhost:8000/api/';


export const fetch = axios.create({
    baseURL: BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});