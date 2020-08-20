<template>
  <b-card>
    <b-form-select v-model="selectedCategory" :disabled="startID !== ''" :options="category">Please select a category</b-form-select>
    <b-form inline v-if="startID !== null">
      <b-button pill variant="primary" @click="start" :disabled="startID !== ''">Start</b-button>
      <b-button pill variant="primary" @click="end" :disabled="startID === ''">End</b-button>
    </b-form>
    <b-link class="text-nowrap" to="/calendar">日付選択画面へ</b-link>
  </b-card>
</template>

<script>
import storage from '../../../mixins/storage';
import { refreshAuthToken } from '../../../service/apiBase';
import { doCreateRecord, doGetLastRecord, doUpdateRecord } from '../../../service/record';

export default {
  mixins: [storage],
  data() {
    return {
      startID: null,
      category: [],
      selectedCategory: '',
    };
  },
  beforeMount() {
    this.getLastRecord();
    this.getCategory();
  },
  methods: {
    start() {
      const param = {
        category: this.selectedCategory,
      };
      doCreateRecord(param)
        .then(res => (this.startID = res.data.id))
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
        .then(res => {
          this.startID = res.data.id;
          this.selectedCategory = res.data.category;
        })
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
    getCategory() {
      this.getLocalStorage('category').then(value => {
        this.category = value;
        if (value.length) {
          this.selectedCategory = value[0].value;
        } else {
          this.selectedCategory = '';
        }
      });
    },
  },
};
</script>

<style lang="scss" scoped></style>
