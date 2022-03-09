<template>
  <div class="main">
    <table-main></table-main>
  </div>
</template>

<script>
import TableMain from './components/TableMain.vue'

export default {
  name: 'App',
  components: {
    TableMain
  },
  mixins: {

  }
}
</script>

<style>
.main {
  width: 90%;
  margin: auto;
  margin-top: 2%;
  background-color: #ffffff;
  border: 1px solid #e2e6e6;
  border-radius: 2px;
  box-shadow: 5px 5px #f1f1f1;
}
</style>
