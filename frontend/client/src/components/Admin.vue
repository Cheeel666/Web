<template>
    <div class="hello">
    <h1>Добро пожаловать на страницу Администратора.</h1>
    <h1>Список пользователей:</h1>
        <table class="table table-hover">
          <thead>
            <tr>
              <th scope="col">Имя</th>
              <th scope="col">Почта</th>
              <th scope="col">Роль</th>
            </tr>
          </thead>
          <tbody v-if="myRole == 'admin'">
            <tr v-for="(user, index) in users" :key="index">
              <td>{{ user.username }}</td>
              <td>{{ user.email }}</td>
              <td> {{ user.role }}</td>
              <td v-if="user.role !='admin'">
                <button type="submit" @click="deleteUser(user.email)" class="btn btn-danger btn-sm">
                    Удалить
                </button>
              </td>
              <td v-if="user.role =='user'">
                <button type="submit" @click="makeMod(user.email)" class="btn btn-warning btn-sm">
                    Сделать модератором
                </button>
              </td>
            </tr>
          </tbody>
          <tbody v-if="myRole == 'mod'">
            <tr v-for="(user, index) in users" :key="index">
              <td>{{ user.username }}</td>
              <td>{{ user.email }}</td>
              <td> {{ user.role }}</td>
              <td v-if="user.role =='user'">
                <button type="submit" @click="deleteUser(user.email)" class="btn btn-danger btn-sm">
                    Удалить
                </button>
              </td>
            </tr>
          </tbody>
        </table>
    <tr></tr>
    </div>
</template>

<script>
import axios from 'axios';
/* eslint-disable */
  class User{
    constructor(username, email, role) {
      this.username = username;
      this.email = email;
      this.role = role;
    }
  }
    export default {
        data () {
            return {
                users: [],
                myRole: localStorage.getItem('role')
            }
        },
        methods: {
          getUsers() {
            const path = 'http://localhost:5005/api/v1/user/users';
            axios.get(path)
              .then((res) => {
                res.data.forEach(element => this.users.push(new User(element["username"], element["email"], element["role"])));
              })
            .catch((error) => {
          // eslint-disable-next-line
                console.error(error);
            });
          },
          deleteUser(email) {
          //   this.$http.delete('http://localhost:5005/delete_user', {
          //   email: email,
            this.$http.delete('http://localhost:5005/api/v1/user/'+email)
          // })
            
            this.$emit('userDeleted')
            this.$router.go()
          },
          makeMod(email) {
            this.$http.patch('http://localhost:5005/api/v1/user/change_role/' + email, {role:"mod"})
            this.$emit('RoleChanged')
            this.$router.go()
          },
        },
        created() {
            this.getUsers();
        },
    }
    
</script>
<style scoped>
   .topnav {
  background-color: #333;
  overflow: hidden;
}

/* Style the links inside the navigation bar */
.topnav a {
  float: left;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

/* Change the color of links on hover */
.topnav a:hover {
  background-color: #ddd;
  color: black;
}

/* Add a color to the active/current link */
.topnav a.active {
  background-color: #04AA6D;
  color: white;
}
* {
  padding:4px;
  margin:4px;
}
h1, h2 {
        font-weight: normal;
    }
ul {
        list-style-type: none;
        padding: 0;
    }
li {
        display: inline-block;
        margin: 0 10px;
    }
a {
        color: #42b983;
    }
</style>
