import axios from "axios";
import firebase from "firebase";

export { doGet, doPost, doPut, doDelete, makeSearchParams };

const axiosInstance = axios.create({
  baseURL: process.env.VUE_APP_API_ORIGIN
});

function doGet(url) {
  return firebase
    .auth()
    .currentUser.getIdToken(true)
    .then(token => {
      return axiosInstance.get(url, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
    })
    .catch(e => {
      alert(e);
    });
}

function doPost(url, data) {
  return firebase
    .auth()
    .currentUser.getIdToken(true)
    .then(token => {
      return axiosInstance.post(url, data, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
    })
    .catch(e => {
      alert(e);
    });
}

function doPut(url, data) {
  return firebase
    .auth()
    .currentUser.getIdToken(true)
    .then(token => {
      return axiosInstance.put(url, data, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
    })
    .catch(e => {
      alert(e);
    });
}

function doDelete(url) {
  return firebase
    .auth()
    .currentUser.getIdToken(true)
    .then(token => {
      return axiosInstance.delete(url, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
    })
    .catch(e => {
      alert(e);
    });
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
