<template>
  <div class="container">
    <form id="add-todo" @submit.prevent="add">
      <input v-model="newTodo"/>
      <button type="submit">Add</button>
    </form>
    <ul class="todo-overview-list">
      <li v-for="todo in todos">
        <p v-bind:class="{ 'todo-done': todo.done }">
          <input type="checkbox" :checked="todo.done" @input="toggleDone(todo, $event)">
          {{todo.body}}
        </p>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data: () => ({
    newTodo: '',
    todos: [],
  }),
  created() {
    this.list();
  },
  methods: {
    toggleDone(item) {
      window.backend.TodoApp.Done(item.id, !item.done).then(result => {
        this.list();
      }).catch(() => alert("Error while fetching todos"));
    },
    list(e) {
      this.newTodo = '';
      window.backend.TodoApp.List().then(result => {
        this.todos = result;
      }).catch(() => alert("Error while fetching todos: " + ee));
    },
    add(e) {
      window.backend.TodoApp.Add(this.newTodo).then(result => {
        this.list();
      }).catch((ee) => alert("Error while fetching todos" + ee));
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  #add-todo {
    width: 100%;
  }
  #add-todo input {
    font-size: 2em;
    width : calc(100% - 110px);
  }
  #add-todo button:hover {
    width: 100px;
    font-size: 1.7em;
    border-color: blue;
    background-color: blue;
    color: white;
    border: 3px solid white;
    border-radius: 10px;
    padding: 9px;
    cursor: pointer;
    transition: 500ms;
  }
  #add-todo button {
    width: 100px;
    font-size: 1.7em;
    border-color: white;
    background-color: #121212;
    color: white;
    border: 3px solid white;
    border-radius: 10px;
    padding: 9px;
    cursor: pointer;
  }
  ul.todo-overview-list {
    padding: 0;
    list-style: none;
  }
  ul.todo-overview-list p {
    display: block;
    padding: 16px;
    border-bottom: 1px solid #ddd;
  }
  ul.todo-overview-list p.todo-done {
    background-color: #ddd;
  }
</style>
