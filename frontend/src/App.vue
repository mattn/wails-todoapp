<template>
  <div id="app">
    <Todo/>
  </div>
</template>

<script>
import Todo from "./components/Todo.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Todo
  }
};
</script>
