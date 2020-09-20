import axios from "axios";
import env from "env";

export { doGet, doPost, doPut, doDelete, makeSearchParams };

const axiosInstance = axios.create({
  baseURL: env.apiBaseURL
});

function doGet(url) {
  return axiosInstance.get(url);
}

function doPost(url, data) {
  return axiosInstance.post(url, data);
}

function doPut(url, data) {
  return axiosInstance.put(url, data);
}

function doDelete(url) {
  return axiosInstance.delete(url);
}

function makeSearchParams(reqObj) {
  const searchParams = new URLSearchParams("");
  if (!reqObj) {
    return searchParams;
  }
  Object.keys(reqObj).forEach(key => {
    const property = reqObj[key];
    if (property === null || property === undefined) {
      return;
    }

    searchParams.set(key, property);
  });
  return `?${searchParams}`;
}
