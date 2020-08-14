<template>
  <div>
    <ul>
      <li v-for="record in dateFormatedRecords" :key="record.start">start: {{ record.start }} end: {{ record.end }}</li>
    </ul>
    <b-link to="/">戻る</b-link>
  </div>
</template>

<script>
import { refreshAuthToken } from '../../../service/apiBase';
import { doListRecord } from '../../../service/recorde';
export default {
  data() {
    return { records: [] };
  },
  computed: {
    dateFormatedRecords() {
      if (!this.records.length) {
        return this.records;
      }
      return this.records.map((value, index, array) => {
        value.start = this.formatDate(new Date(value.start));
        value.end = this.formatDate(new Date(value.end));
        return value;
      });
    },
  },
  beforeMount() {
    this.listRecord();
  },
  methods: {
    listRecord() {
      doListRecord()
        .then(res => (this.records = res.data))
        .catch(error => {
          if (error.response.status === 401) {
            refreshAuthToken().then(() => {
              this.listRecord();
            });
          } else {
            alert('エラー: ' + JSON.stringify(error.response));
          }
        });
    },
    formatDate(date) {
      return (
        date.getFullYear() +
        '/' +
        ('0' + (date.getMonth() + 1)).slice(-2) +
        '/' +
        ('0' + date.getDate()).slice(-2) +
        ' ' +
        ('0' + date.getHours()).slice(-2) +
        ':' +
        ('0' + date.getMinutes()).slice(-2) +
        ':' +
        ('0' + date.getSeconds()).slice(-2) +
        '(JST)'
      );
    },
  },
};
</script>
