
import axios from 'axios';

const http = axios.create();

http.defaults.timeout = 5000;

export default http;
