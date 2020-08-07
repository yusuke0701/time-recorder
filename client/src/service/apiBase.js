import axios from 'axios';

export { doGet, doPost, doPut, doDelete };

let axiosInstance;

async function makeAxiosInstance() {
  const token = await getLocalStorage('token');

  axiosInstance = axios.create({
    baseURL: 'https://us-central1-hoge-hoge-123456789.cloudfunctions.net',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

function getLocalStorage(key = null) {
  return new Promise(resolve => {
    chrome.storage.local.get(key, item => {
      key ? resolve(item[key]) : resolve(item);
    });
  });
}

async function doGet(url) {
  if (axiosInstance) {
    return axiosInstance.get(url);
  }
  await makeAxiosInstance();
  return axiosInstance.get(url);
}

function doPost(url, data) {
  // MEMO: doGetが先に呼ばれるので、axiosInstance のチェックはしない
  return axiosInstance.post(url, data);
}

function doPut(url, data) {
  // MEMO: doGetが先に呼ばれるので、axiosInstance のチェックはしない
  return axiosInstance.put(url, data);
}

function doDelete(url, config) {
  // MEMO: doGetが先に呼ばれるので、axiosInstance のチェックはしない
  return axiosInstance.delete(url, config);
}
