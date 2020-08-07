import { doGet, doPost } from './apiBase';
export { doCreateRecord, doGetLastRecord, doListRecord, doUpdateRecord };

function doCreateRecord() {
  return doPost(`/CreateRecord`);
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
