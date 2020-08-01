import { doPost } from './apiBase';
export { doPostStart, doPostEnd };

function doPostStart() {
  return doPost(`/Start`);
}

function doPostEnd(payload) {
  return doPost(`/End`, payload);
}
