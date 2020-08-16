<template>
  <div>
    <b-form class="form-inline" v-if="startID !== null">
      <b-button pill variant="primary" @click="start" :disabled="startID !== ''">Start</b-button>
      <b-button pill variant="primary" @click="end" :disabled="startID === ''">End</b-button>
    </b-form>
    <b-link class="text-nowrap" to="/calendar">日付選択画面へ</b-link>
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
