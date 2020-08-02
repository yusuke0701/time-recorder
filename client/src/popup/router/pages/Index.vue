<template>
  <div>
    <div v-if="startID === null"></div>
    <div v-else-if="startID === ''">
      <button @click="start">Start</button>
    </div>
    <div v-else>
      <button @click="end">End</button>
    </div>
    <div>
      <router-link to="/list">一覧へ</router-link>
    </div>
  </div>
</template>

<script>
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
          alert('エラー: ' + JSON.stringify(error.response));
        });
    },
    getLastRecord() {
      doGetLastRecord()
        .then(res => (this.startID = res.data))
        .catch(error => {
          if (error.response.status === 404) {
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
          alert('エラー: ' + JSON.stringify(error.response));
        });
    },
  },
};
</script>

<style lang="scss" scoped>
p {
  font-size: 20px;
}
</style>
