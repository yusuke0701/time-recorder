<template>
  <b-card>
    <!-- TODO: サインアウトの機能 -->
    <b-form-select
      v-model="selectedCategory"
      :disabled="startID !== ''"
      :options="category"
      >Please select a category</b-form-select
    >
    <b-form inline v-if="startID !== null">
      <b-button pill variant="primary" @click="start" :disabled="startID !== ''"
        >Start</b-button
      >
      <b-button pill variant="primary" @click="end" :disabled="startID === ''"
        >End</b-button
      >
    </b-form>
    <b-link class="text-nowrap" to="/calendar">過去の記録を見る</b-link>
  </b-card>
</template>

<script>
import {
  doCreateRecord,
  doGetLastRecord,
  doUpdateRecord
} from "../services/record";

export default {
  data() {
    return {
      startID: "",
      // category をストレージに保存する
      category: [
        { text: "work", value: "work" },
        { text: "study", value: "study" },
        { text: "game", value: "game" },
        { text: "anime", value: "anime" },
        { text: "movie", value: "movie" }
      ],
      selectedCategory: ""
    };
  },
  beforeMount() {
    this.getLastRecord();
    // this.getCategory();
  },
  methods: {
    start() {
      const param = {
        category: this.selectedCategory
      };
      doCreateRecord(param)
        .then(res => (this.startID = res.data.id))
        .catch(error => {
          if (error.response.status === 401) {
            this.start();
          } else {
            alert("エラー: " + JSON.stringify(error.response));
          }
        });
    },
    getLastRecord() {
      doGetLastRecord()
        .then(res => {
          this.startID = res.data.id;
          this.selectedCategory = res.data.category;
        })
        .catch(error => {
          if (error.response.status === 401) {
            this.getLastRecord();
          } else if (error.response.status === 404) {
            this.startID = "";
          } else {
            alert("エラー: " + JSON.stringify(error.response));
          }
        });
    },
    end() {
      doUpdateRecord(this.startID)
        .then(() => (this.startID = ""))
        .catch(error => {
          if (error.response.status === 401) {
            this.end();
          } else {
            alert("エラー: " + JSON.stringify(error.response));
          }
        });
    },
    getCategory() {
      this.getLocalStorage("category").then(value => {
        this.category = value;
        if (value.length) {
          this.selectedCategory = value[0].value;
        } else {
          this.selectedCategory = "";
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped></style>
