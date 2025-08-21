import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Products from '../views/Products.vue'
import Cart from '../views/Cart.vue'
import Login from '../views/Login.vue'
import Profile from '../views/Profile.vue'
import DatabaseAdmin from '../views/DatabaseAdmin.vue'

// Route configuration - similar to Angular's Routes array
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/products',
    name: 'Products',
    component: Products
  },
  {
    path: '/cart',
    name: 'Cart',
    component: Cart
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile
  },
  {
    path: '/admin/database',
    name: 'DatabaseAdmin',
    component: DatabaseAdmin
  }
]

// Create router instance
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
