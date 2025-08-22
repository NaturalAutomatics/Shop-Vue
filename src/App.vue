<template>
  <!-- Vue template syntax - similar to Angular template but with v- directives -->
  <div id="app">
    <!-- Navigation header -->
    <nav class="navbar">
      <div class="nav-container">
        <router-link to="/" class="nav-brand">
          🛍️ Vue Shop
        </router-link>
        <div class="nav-links">
          <router-link to="/" class="nav-link">Home</router-link>
          <router-link to="/products" class="nav-link">Products</router-link>
          <router-link to="/cart" class="nav-link">
            Cart ({{ cartItemCount }})
          </router-link>
          <div v-if="isAuthenticated" class="auth-links">
            <router-link to="/profile" class="nav-link">Profile</router-link>
            <button @click="handleLogout" class="nav-link logout-btn">
              Logout
            </button>
            <router-link v-if="authStore.isAdmin" to="/admin/database" class="nav-link admin-link">🗄️ Database Admin</router-link>
          </div>
          <div v-else class="auth-links">
            <router-link to="/login" class="nav-link">Login</router-link>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main content area -->
    <main class="main-content">
      <!-- Router outlet - similar to <router-outlet> in Angular -->
      <router-view />
    </main>

    <!-- Footer -->
    <footer class="footer">
      <p>&copy; 2024 Vue Shop - Learning Vue.js after Angular</p>
    </footer>
  </div>
</template>

<script>
// Vue 3 Composition API - similar to Angular's dependency injection but more flexible
import { computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from './stores/cart'
import { useAuthStore } from './stores/auth'

export default {
  name: 'App',
  setup() {
    const router = useRouter()
    const cartStore = useCartStore()
    const authStore = useAuthStore()
    
    // Computed property (similar to Angular getters)
    const cartItemCount = computed(() => cartStore.itemCount)
    const isAuthenticated = computed(() => authStore.isAuthenticated)
    
    const handleLogout = async () => {
      try {
        await authStore.logout()
        router.push('/login')
      } catch (error) {
        console.error('Logout error:', error)
      }
    }
    
    // Check authentication status on app start
    onMounted(async () => {
      console.log('App mounted, auth state:', { 
        token: authStore.token, 
        user: authStore.user, 
        isAuthenticated: authStore.isAuthenticated 
      })
      if (authStore.token && !authStore.user) {
        await authStore.fetchCurrentUser()
      }
    })
    
    // Watch for authentication state changes
    watch(() => authStore.isAuthenticated, (newValue, oldValue) => {
      console.log('Auth state changed:', { oldValue, newValue, token: authStore.token, user: authStore.user })
    })
    
    return {
      cartItemCount,
      isAuthenticated,
      handleLogout,
      authStore
    }
  }
}
</script>

<style scoped>
/* Scoped styles - only apply to this component */
.navbar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 1rem 0;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-brand {
  font-size: 1.5rem;
  font-weight: bold;
  color: white;
  text-decoration: none;
}

.nav-links {
  display: flex;
  gap: 2rem;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav-link:hover {
  background-color: rgba(255,255,255,0.1);
}

.auth-links {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.logout-btn {
  background: none;
  border: none;
  cursor: pointer;
  font-size: inherit;
  font-family: inherit;
}

.main-content {
  min-height: calc(100vh - 140px);
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.footer {
  background-color: #f8f9fa;
  text-align: center;
  padding: 1rem;
  border-top: 1px solid #e9ecef;
}
</style>
