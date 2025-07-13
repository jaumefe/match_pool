import {createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Pool from '../views/Pool.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/pool', component: Pool}
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router