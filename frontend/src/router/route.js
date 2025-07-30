import {createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Pool from '../views/Pool.vue'
import ResultRegister from '../views/ResultRegister.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/pool', component: Pool },
  { path: '/result_register', component: ResultRegister}
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router