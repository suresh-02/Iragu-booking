import axios from "axios";

const createApiSession = axios.create({
  baseURL: "http://localhost:8082",
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

export default createApiSession;
