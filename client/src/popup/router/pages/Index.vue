<template>
  <div>
    <div v-if="startID === null"></div>
    <div v-else-if="startID === ''">
      <b-button @click="start">Start</b-button>
    </div>
    <div v-else>
      <b-button @click="end">End</b-button>
    </div>
    <div class="text-nowrap">
      <b-link to="/list">一覧へ</b-link>
    </div>
  </div>
</template>

<script>
import { refreshAuthToken } from '../../../service/apiBase';
import { doCreateRecord, doGetLastRecord, doUpdateRecord } from '../../../service/recorde';
export default {
  data() {
    return { startID: null };
  },
  beforeMount() {
    this.getLastRecord();
  },
  methods: {
    start() {
      doCreateRecord()
        .then(res => (this.startID = res.data))
        .catch(error => {
          if (error.response.status === 401) {
            refreshAuthToken().then(() => {
              this.start();
            });
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
        });
    },
    getLastRecord() {
      doGetLastRecord()
        .then(res => (this.startID = res.data))
        .catch(error => {
          if (error.response.status === 401) {
            refreshAuthToken().then(() => {
              this.getLastRecord();
            });
          } else if (error.response.status === 404) {
            this.startID = '';
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
        });
    },
    end() {
      doUpdateRecord(this.startID)
        .then(() => (this.startID = ''))
        .catch(error => {
          if (error.response.status === 401) {
            refreshAuthToken().then(() => {
              this.end();
            });
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
        });
    },
  },
};
</script>

<style lang="scss" scoped></style>
