<template>
  <div>
    <b-list-group>
      <b-list-group-item v-for="item in category" :key="item.value" class="d-flex justify-content-between align-items-center">
        {{ item.text }}
        <!-- <b-badge variant="primary" pill>14</b-badge> -->
        <!-- 削除ボタンと、変更ボタン -->
      </b-list-group-item>
    </b-list-group>
    <b-form inline>
      <b-form-input v-model="newCategoryName" placeholder="Enter category name"></b-form-input>
      <b-button :disabled="newCategoryName === ''" block @click="addCategory">カテゴリの追加</b-button>
    </b-form>
  </div>
</template>

<script>
import storage from '../mixins/storage';

export default {
  name: 'App',
  mixins: [storage],
  data() {
    return {
      newCategoryName: '',
      category: [],
    };
  },
  beforeMount() {
    this.getCategory();
  },
  methods: {
    getCategory() {
      this.getLocalStorage('category').then(value => {
        this.category = value;
      });
    },
    addCategory() {
      // TODO: category name が同じものが入ってないかチェックする
      this.category.push({ value: this.newCategoryName, text: this.newCategoryName });
      this.newCategoryName = '';

      // TODO: 他のプロパティが存在した場合、どうなるのか？
      const tmp = {
        category: this.category,
      };
      this.setLocalStorage(tmp).then(value => {
        this.setResult = value;
      });
    },
  },
};
</script>

<style scoped></style>
