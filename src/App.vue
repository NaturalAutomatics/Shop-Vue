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
import { computed } from 'vue'
import { useCartStore } from './stores/cart'

export default {
  name: 'App',
  setup() {
    // Using Pinia store (similar to Angular services but with reactive state)
    const cartStore = useCartStore()
    
    // Computed property (similar to Angular getters)
    const cartItemCount = computed(() => cartStore.itemCount)
    
    return {
      cartItemCount
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
