global.browser = require('webextension-polyfill');

chrome.identity.getAuthToken({ interactive: true }, token => {
  if (chrome.runtime.lastError) {
    console.log(chrome.runtime.lastError.message);
  } else {
    chrome.storage.local.set({ token: token }, () => {
      console.log('saved token: ' + token);
    });
  }
});
