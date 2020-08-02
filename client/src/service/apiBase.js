import axios from 'axios';

export { doGet, doPost, doPut, doDelete };

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
