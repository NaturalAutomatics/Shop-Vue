<template>
  <div class="home">
    <div class="hero">
      <h1>Welcome to Vue Shop!</h1>
      <p class="subtitle">Learn Vue.js by building a real e-commerce application</p>
      <div class="hero-buttons">
        <router-link to="/products" class="btn btn-primary">
          Browse Products
        </router-link>
        <button @click="showLessonInfo = !showLessonInfo" class="btn btn-secondary">
          Learn More
        </button>
      </div>
    </div>

        <!-- Categories section -->
        <div class="categories">
      <h2>Shop by Category</h2>
      <div v-if="categoriesError" class="categories-error">{{ categoriesError }}</div>
      <div v-else class="categories-grid">
        <div v-if="loadingCategories" class="category-skeleton" v-for="i in 6" :key="i"></div>
        <div v-else v-for="cat in categories" :key="cat" class="category-card">
          <router-link :to="{ name: 'Products', query: { category: cat } }" class="category-link">{{ cat }}</router-link>
        </div>
      </div>
    </div>

    <!-- Lesson information section -->
    <div v-if="showLessonInfo" class="lesson-info">
      <h2>Vue.js vs Angular - Key Differences</h2>
      <div class="comparison-grid">
        <div class="comparison-card">
          <h3>Template Syntax</h3>
          <p><strong>Vue:</strong> Uses <code>v-</code> directives (v-if, v-for, v-model)</p>
          <p><strong>Angular:</strong> Uses <code>*ng</code> directives (*ngIf, *ngFor, [(ngModel)])</p>
        </div>
        
        <div class="comparison-card">
          <h3>State Management</h3>
          <p><strong>Vue:</strong> Pinia stores with reactive state</p>
          <p><strong>Angular:</strong> Services with RxJS observables</p>
        </div>
        
        <div class="comparison-card">
          <h3>Component API</h3>
          <p><strong>Vue:</strong> Composition API with setup() function</p>
          <p><strong>Angular:</strong> Class-based components with decorators</p>
        </div>
        
        <div class="comparison-card">
          <h3>Reactivity</h3>
          <p><strong>Vue:</strong> Automatic reactivity with ref() and reactive()</p>
          <p><strong>Angular:</strong> Change detection with Zone.js</p>
        </div>
      </div>
    </div>

    <!-- Features section -->
    <div class="features">
      <h2>What You'll Learn</h2>
      <div class="features-grid">
        <div class="feature-card">
          <h3>🎯 Vue 3 Composition API</h3>
          <p>Learn the modern way to write Vue components with better TypeScript support and logic reuse.</p>
        </div>
        <div class="feature-card">
          <h3>🛒 State Management</h3>
          <p>Master Pinia for state management, Vue's official store library.</p>
        </div>
        <div class="feature-card">
          <h3>🛣️ Routing</h3>
          <p>Implement client-side routing with Vue Router 4.</p>
        </div>
        <div class="feature-card">
          <h3>🎨 Styling</h3>
          <p>Learn Vue's scoped styling and CSS-in-JS approaches.</p>
        </div>
      </div>
    </div>


  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import apiService from '../services/api'

export default {
  name: 'Home',
  setup() {
    // Reactive state - similar to Angular's component properties
    const showLessonInfo = ref(false)
    const categories = ref([])
    const loadingCategories = ref(false)
    const categoriesError = ref('')

    const loadCategories = async () => {
      loadingCategories.value = true
      categoriesError.value = ''
      try {
        const res = await apiService.getCategories()
        categories.value = res.data || []
      } catch (e) {
        categoriesError.value = 'Failed to load categories'
        console.error('Categories load error:', e)
      } finally {
        loadingCategories.value = false
      }
    }

    onMounted(() => {
      loadCategories()
    })

    return {
      showLessonInfo,
      categories,
      loadingCategories,
      categoriesError
    }
  }
}
</script>

<style scoped>
.home {
  max-width: 1200px;
  margin: 0 auto;
}

.hero {
  text-align: center;
  padding: 4rem 2rem;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border-radius: 12px;
  margin-bottom: 3rem;
}

.hero h1 {
  font-size: 3rem;
  color: #2c3e50;
  margin-bottom: 1rem;
}

.subtitle {
  font-size: 1.2rem;
  color: #7f8c8d;
  margin-bottom: 2rem;
}

.hero-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  text-decoration: none;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.btn-secondary {
  background: white;
  color: #667eea;
  border: 2px solid #667eea;
}

.btn-secondary:hover {
  background: #667eea;
  color: white;
}

.lesson-info {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  margin-bottom: 3rem;
}

.comparison-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.comparison-card {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  border-left: 4px solid #667eea;
}

.comparison-card h3 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.comparison-card p {
  margin-bottom: 0.5rem;
  line-height: 1.6;
}

.comparison-card code {
  background: #e9ecef;
  padding: 0.2rem 0.4rem;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
}

.features {
  margin-top: 3rem;
}

.features h2 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 2rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.feature-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: transform 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-4px);
}

.feature-card h3 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.feature-card p {
  color: #7f8c8d;
  line-height: 1.6;
}

/* Categories */
.categories {
  margin-top: 3rem;
}

.categories h2 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 1rem;
}

.categories-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
}

@media (min-width: 640px) {
  .categories-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .categories-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

.category-card {
  background: white;
  border-radius: 10px;
  border: 1px solid #e9ecef;
  padding: 1rem;
  text-align: center;
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.category-link {
  display: block;
  text-decoration: none;
  color: #667eea;
  font-weight: 600;
}

.category-skeleton {
  height: 48px;
  border-radius: 10px;
  background: linear-gradient(90deg, #f0f0f0 25%, #e6e6e6 37%, #f0f0f0 63%);
  background-size: 400% 100%;
  animation: shimmer 1.4s ease infinite;
}

@keyframes shimmer {
  0% { background-position: 100% 0; }
  100% { background-position: 0 0; }
}

@media (max-width: 768px) {
  .hero h1 {
    font-size: 2rem;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: center;
  }
  
  .comparison-grid {
    grid-template-columns: 1fr;
  }
}
</style>
