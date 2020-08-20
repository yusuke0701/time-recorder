import { doGet, doPost, doPut, makeSearchParams } from './apiBase';
export { doCreateRecord, doGetLastRecord, doListRecord, doUpdateRecord };

const recordAPIPath = '/Records';

function doCreateRecord(param) {
  const searchParam = makeSearchParams(param);
  return doPost(`${recordAPIPath}${searchParam}`);
}

function doGetLastRecord() {
  return doGet(`/GetLastRecord`);
}

function doListRecord(param) {
  const searchParam = makeSearchParams(param);
  return doGet(`/ListRecord${searchParam}`);
}

function doUpdateRecord(payload) {
  return doPut(recordAPIPath, payload);
}
