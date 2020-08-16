<template>
  <b-card class="text-center text-nowrap" :title="$route.params.selectedDate">
    <ul>
      <li v-for="record in dateFormatedRecords" :key="record.start">
        <div class="text-left" v-if="record.end === '0001-01-1'">start: {{ record.start_detail }}</div>
        <div class="text-left" v-else>start: {{ record.start_detail }} end: {{ record.end_detail }}</div>
      </li>
    </ul>
    <b-link to="/calendar">日付選択画面へ</b-link>
  </b-card>
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
        value.start_detail = this.formatDate(new Date(value.start_detail));
        value.end_detail = this.formatDate(new Date(value.end_detail));
        return value;
      });
    },
  },
  beforeMount() {
    this.listRecord();
  },
  methods: {
    listRecord() {
      const param = {
        start: this.$route.params.selectedDate,
      };
      doListRecord(param)
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
      return ('0' + date.getHours()).slice(-2) + ':' + ('0' + date.getMinutes()).slice(-2) + ':' + ('0' + date.getSeconds()).slice(-2) + '(JST)';
    },
  },
};
</script>
