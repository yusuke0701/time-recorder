import { doGet, doPost } from './apiBase';
export { doPostStart, doPostEnd, doGetListRecord };

function doPostStart() {
  return doPost(`/Start`);
}

function doPostEnd(payload) {
  return doPost(`/End`, payload);
}

function doGetListRecord() {
  return doGet(`/ListRecord`);
}
