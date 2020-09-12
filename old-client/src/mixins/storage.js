export default {
  methods: {
    setLocalStorage(obj) {
      return new Promise(resolve => {
        chrome.storage.local.set(obj, () => resolve());
      });
    },
    getLocalStorage(key = null) {
      return new Promise(resolve => {
        chrome.storage.local.get(key, item => {
          key ? resolve(item[key]) : resolve(item);
        });
      });
    },
  },
};
