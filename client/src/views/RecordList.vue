<template>
  <b-card
    class="text-center text-nowrap"
    :title="$route.params.selectedDate + '(JST)'"
  >
    <b-table hover :fields="tableHeaders" :items="tableValues"></b-table>
    <b-link to="/calendar">戻る</b-link>
  </b-card>
</template>

<script>
import { doListRecord } from "../services/record";
export default {
  data() {
    return {
      records: [],
      tableHeaders: ["category", "start", "end"]
    };
  },
  computed: {
    tableValues() {
      if (!this.records.length) {
        return [];
      }
      return this.records.map(value => {
        if (value.end === "0001-01-1") {
          return {
            category: value.category,
            start: this.formatDate(new Date(value.start_detail))
          };
        }
        return {
          category: value.category,
          start: this.formatDate(new Date(value.start_detail)),
          end: this.formatDate(new Date(value.end_detail))
        };
      });
    }
  },
  beforeMount() {
    this.listRecord();
  },
  methods: {
    listRecord() {
      const param = {
        start: this.$route.params.selectedDate
      };
      doListRecord(param)
        .then(res => (this.records = res.data))
        .catch(error => {
          if (error.response.status === 401) {
            this.listRecord();
          } else {
            alert("エラー: " + JSON.stringify(error.response));
          }
        });
    },
    formatDate(date) {
      return (
        ("0" + date.getHours()).slice(-2) +
        ":" +
        ("0" + date.getMinutes()).slice(-2) +
        ":" +
        ("0" + date.getSeconds()).slice(-2)
      );
    }
  }
};
</script>
