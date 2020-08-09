import axios from 'axios';
import env from 'env';

export { doGet, doPost, doPut, doDelete, refreshAuthToken };

const axiosInstance = axios.create({
  baseURL: env.apiBaseURL,
});

function getAuthToken() {
  return new Promise(resolve => {
    chrome.identity.getAuthToken({ interactive: true }, token => {
      resolve(token);
    });
  });
}

function refreshAuthToken() {
  return new Promise(resolve => {
    getAuthToken().then(token => {
      chrome.identity.removeCachedAuthToken({ token: token }, () => {
        resolve();
      });
    });
  });
}

function doGet(url) {
  return getAuthToken().then(token => {
    return axiosInstance.get(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  });
}

function doPost(url, data) {
  return getAuthToken().then(token => {
    return axiosInstance.post(url, data, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  });
}

function doPut(url, data) {
  return getAuthToken().then(token => {
    return axiosInstance.put(url, data, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  });
}

function doDelete(url) {
  return getAuthToken().then(token => {
    return axiosInstance.delete(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  });
}
