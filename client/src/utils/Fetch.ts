import axios from 'axios';

const BASE_URL = 'http://localhost:8000/api';


export const fetchdata = axios.create({
    baseURL: BASE_URL,
    timeout: 3000,
    headers: {
        'Content-Type': 'application/json',
    },
});