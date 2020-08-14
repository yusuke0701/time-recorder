import { doGet, doPost, makeSearchParams } from './apiBase';
export { doCreateRecord, doGetLastRecord, doListRecord, doUpdateRecord };

function doCreateRecord() {
  return doPost(`/CreateRecord`);
}

function doGetLastRecord() {
  return doGet(`/GetLastRecord`);
}

function doListRecord(param) {
  const searchParam = makeSearchParams(param);
  return doGet(`/ListRecord${searchParam}`);
}

function doUpdateRecord(payload) {
  return doPost(`/UpdateRecord`, payload);
}
