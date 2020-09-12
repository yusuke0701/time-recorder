<template>
  <div>
    <b-list-group>
      <b-list-group-item v-for="item in category" :key="item.value" class="d-flex justify-content-between align-items-center">
        {{ item.text }}
        <b-button @click="deleteCategory(item.value)">削除</b-button>
      </b-list-group-item>
    </b-list-group>
    <b-form-input v-model="newCategoryName" placeholder="Enter category name"></b-form-input>
    <b-button :disabled="newCategoryName === ''" block @click="addCategory">カテゴリの追加</b-button>
    <b-button block @click="clearCategory">カテゴリの初期化</b-button>
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
      // index.vue の表示の仕方に合わせて保存しているが、正直、微妙
      // ただ、name だけにすると、配列表示が微妙になる。key の指定的なあれで
      const newCategory = { value: this.newCategoryName, text: this.newCategoryName };
      this.newCategoryName = '';

      if (this.category.filter(v => v.value === newCategory.value).length) {
        return;
      }
      this.category.push(newCategory);
      this.saveCategory();
    },
    deleteCategory(value) {
      this.category = this.category.filter(v => v.value !== value);
      this.saveCategory();
    },
    saveCategory() {
      // TODO: 他のプロパティが存在した場合、どうなるのか？
      if (!this.category) {
        this.category = this.getDefaultCategory();
      }
      const tmp = {
        category: this.category,
      };
      this.setLocalStorage(tmp).then(value => {
        this.setResult = value;
      });
    },
    clearCategory() {
      this.category = this.getDefaultCategory();
      this.saveCategory();
    },
    getDefaultCategory() {
      return [
        { text: 'work', value: 'work' },
        { text: 'study', value: 'study' },
        { text: 'game', value: 'game' },
        { text: 'anime', value: 'anime' },
        { text: 'movie', value: 'movie' },
      ];
    },
  },
};
</script>

<style scoped></style>
