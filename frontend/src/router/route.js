import {createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Pool from '../views/Pool.vue'
import ResultRegister from '../views/ResultRegister.vue'
import Points from '../views/Points.vue'
import MainLayout from '../layouts/MainLayout.vue'
import { getToken } from '../services/auth'
import UserRegister from '../views/UserRegister.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path:'/',
    component: MainLayout,
    children: [
      {path: 'pool', name: 'Pool', component: Pool},
      {path: 'result_register', name: 'Register', component: ResultRegister},
      {path: 'points', name: 'Points', component: Points},
      {path: 'register_user', name: 'Register User', component: UserRegister}
    ],
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = getToken()
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router