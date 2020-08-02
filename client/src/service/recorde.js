import { doGet, doPost } from './apiBase';
export { doCreateRecord, doGetRecord, doGetLastRecord, doListRecord, doUpdateRecord };

function doCreateRecord() {
  return doPost(`/CreateRecord`);
}

function doGetRecord() {
  return doGet(`/GetRecord`);
}

function doGetLastRecord() {
  return doGet(`/GetLastRecord`);
}

function doListRecord() {
  return doGet(`/ListRecord`);
}

function doUpdateRecord(payload) {
  return doPost(`/UpdateRecord`, payload);
}
