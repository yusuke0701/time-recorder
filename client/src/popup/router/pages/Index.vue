<template>
  <div>
    <div v-if="startID === ''">
      <button @click="start">Start</button>
    </div>
    <div v-else>
      <button @click="end">End</button>
    </div>
  </div>
</template>

<script>
import { doPostStart, doPostEnd } from '../../../service/recorde';
export default {
  data() {
    return { startID: '' };
  },
  methods: {
    start() {
      doPostStart()
        .then(res => {
          this.startID = res;
        })
        .catch(error => {
          // TODO: catch に入る理由が不明。。
          if (error.response.status === 200) {
            this.startID = error.response.data;
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
        });
    },
    end() {
      doPostEnd(this.startID)
        .then(() => {
          this.startID = '';
        })
        .catch(error => {
          // TODO: catch に入る理由が不明。。
          if (error.response.status === 200) {
            this.startID = '';
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
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
