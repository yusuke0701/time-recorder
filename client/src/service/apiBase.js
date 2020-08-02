import axios from 'axios';

export { doGet, doPost, doPut, doDelete, makeSearchParams };

const axiosInstance = axios.create({
  baseURL: 'https://us-central1-hoge-hoge-123456789.cloudfunctions.net',
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

function doDelete(url, config) {
  return axiosInstance.delete(url, config);
}

function makeSearchParams(reqObj) {
  const searchParams = new URLSearchParams('');
  if (!reqObj) {
    return searchParams;
  }
  Object.keys(reqObj).forEach(key => {
    const property = reqObj[key];
    if (property === null || property === void 0) {
      return;
    }

    searchParams.set(key, property);
  });
  return `?${searchParams}`;
}
