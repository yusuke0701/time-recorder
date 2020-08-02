<template>
  <div>
    <ul>
      <li v-for="record in records" :key="record.start">start: {{ record.start }}, end: {{ record.end }}</li>
    </ul>
    <router-link to="/">戻る</router-link>
  </div>
</template>

<script>
import { doGetListRecord } from '../../../service/recorde';
export default {
  data() {
    return { records: '' };
  },
  beforeMount() {
    this.listRecord();
  },
  methods: {
    listRecord() {
      doGetListRecord()
        .then(() => {})
        .catch(error => {
          // TODO: catch に入る理由が不明。。
          if (error.response.status === 200) {
            this.records = error.response.data;
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
        });
    },
  },
};
</script>
