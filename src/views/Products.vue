<template>
  <div class="products">
    <div class="products-header">
      <h1>Our Products</h1>
      <div class="filters">
        <input 
          v-model="searchTerm" 
          type="text" 
          placeholder="Search products..."
          class="search-input"
        />
        <select v-model="selectedCategory" class="category-select">
          <option value="">All Categories</option>
          <option value="electronics">Electronics</option>
          <option value="clothing">Clothing</option>
          <option value="books">Books</option>
        </select>
      </div>
    </div>

    <!-- Products grid -->
    <div class="products-grid">
      <div 
        v-for="product in filteredProducts" 
        :key="product.id" 
        class="product-card"
      >
        <div class="product-image">
          <img :src="product.image" :alt="product.name" />
        </div>
        <div class="product-info">
          <h3>{{ product.name }}</h3>
          <p class="product-description">{{ product.description }}</p>
          <div class="product-price">${{ product.price }}</div>
          <button 
            @click="addToCart(product)" 
            class="add-to-cart-btn"
            :disabled="isInCart(product.id)"
          >
            {{ isInCart(product.id) ? 'In Cart' : 'Add to Cart' }}
          </button>
        </div>
      </div>
    </div>

    <!-- No products message -->
    <div v-if="filteredProducts.length === 0" class="no-products">
      <p>No products found matching your criteria.</p>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useCartStore } from '../stores/cart'
import apiService from '../services/api'

export default {
  name: 'Products',
  setup() {
    // Reactive state
    const products = ref([])
    const searchTerm = ref('')
    const selectedCategory = ref('')
    const loading = ref(false)
    const error = ref(null)
    
    // Store
    const cartStore = useCartStore()
    
    // Computed properties
    const filteredProducts = computed(() => {
      return products.value.filter(product => {
        const matchesSearch = product.name.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
                            product.description.toLowerCase().includes(searchTerm.value.toLowerCase())
        const matchesCategory = !selectedCategory.value || product.category === selectedCategory.value
        return matchesSearch && matchesCategory
      })
    })
    
    // Methods
    const addToCart = (product) => {
      cartStore.addToCart(product)
    }
    
    const isInCart = (productId) => {
      return cartStore.items.some(item => item.id === productId)
    }
    
    const fetchProducts = async () => {
      loading.value = true
      error.value = null
      
      try {
        const params = {}
        if (selectedCategory.value) params.category = selectedCategory.value
        if (searchTerm.value) params.search = searchTerm.value
        
        const response = await apiService.getProducts(params)
        products.value = response.data
      } catch (err) {
        error.value = 'Failed to load products. Please try again.'
        console.error('Error fetching products:', err)
      } finally {
        loading.value = false
      }
    }
    
    // Lifecycle hook - similar to Angular's ngOnInit
    onMounted(() => {
      fetchProducts()
    })
    
    return {
      products,
      searchTerm,
      selectedCategory,
      filteredProducts,
      loading,
      error,
      addToCart,
      isInCart,
      fetchProducts
    }
  }
}
</script>

<style scoped>
.products {
  max-width: 1200px;
  margin: 0 auto;
}

.products-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  flex-wrap: wrap;
  gap: 1rem;
}

.products-header h1 {
  color: #2c3e50;
  margin: 0;
}

.filters {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.search-input,
.category-select {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
}

.search-input {
  min-width: 200px;
}

.products-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
}

@media (min-width: 1024px) {
  .products-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

.product-card {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.product-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.product-image {
  height: 200px;
  overflow: hidden;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.product-card:hover .product-image img {
  transform: scale(1.05);
}

.product-info {
  padding: 1.5rem;
}

.product-info h3 {
  color: #2c3e50;
  margin: 0 0 0.5rem 0;
  font-size: 1.2rem;
}

.product-description {
  color: #7f8c8d;
  margin: 0 0 1rem 0;
  line-height: 1.5;
  font-size: 0.9rem;
}

.product-price {
  font-size: 1.5rem;
  font-weight: bold;
  color: #667eea;
  margin-bottom: 1rem;
}

.add-to-cart-btn {
  width: 100%;
  padding: 0.75rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 6px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.add-to-cart-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.add-to-cart-btn:disabled {
  background: #95a5a6;
  cursor: not-allowed;
  transform: none;
}

.no-products {
  text-align: center;
  padding: 3rem;
  color: #7f8c8d;
}

@media (max-width: 768px) {
  .products-header {
    flex-direction: column;
    align-items: stretch;
  }
  
  .filters {
    flex-direction: column;
  }
  
  .search-input {
    min-width: auto;
  }
  
  .products-grid {
    grid-template-columns: 1fr;
  }
}
</style>
