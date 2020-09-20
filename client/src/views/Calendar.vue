<template>
  <div class="text-center">
    <b-calendar v-model="selectedDate" @context="onContext"></b-calendar>
    <b-link class="text-nowrap" to="/">戻る</b-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedDate: "",
      context: null
    };
  },
  methods: {
    onContext(ctx) {
      this.context = ctx;
      if (this.selectedDate !== "") {
        this.$router.push(`/list/${this.selectedDate}`);
      }
    }
  }
};
</script>
