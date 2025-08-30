<template>
  <div class="product-detail" v-if="product">
    <!-- Breadcrumb -->
    <nav class="breadcrumb">
      <router-link to="/products" class="breadcrumb-link">Products</router-link>
      <span class="breadcrumb-separator">›</span>
      <span class="breadcrumb-current">{{ product.name }}</span>
    </nav>

    <div class="product-container">
      <!-- Product Image -->
      <div class="product-image-section">
        <div class="main-image" @click="openImageModal">
          <img :src="product.image" :alt="product.name" />
          <div class="zoom-overlay">
            <span class="zoom-icon">🔍</span>
          </div>
        </div>
      </div>

      <!-- Product Info -->
      <div class="product-info-section">
        <div class="product-header">
          <span class="category-badge">{{ product.category }}</span>
          <h1 class="product-title">{{ product.name }}</h1>
        </div>

        <div class="price-section">
          <span class="price">${{ product.price }}</span>
          <span class="stock-info" :class="stockClass">
            {{ stockText }}
          </span>
        </div>

        <div class="description-section">
          <h3>Description</h3>
          <p class="description">{{ product.description }}</p>
        </div>

        <div class="actions-section">
          <div class="quantity-selector">
            <label for="quantity">Quantity:</label>
            <div class="quantity-controls">
              <button @click="decreaseQuantity" :disabled="quantity <= 1" class="qty-btn">−</button>
              <input 
                id="quantity" 
                v-model.number="quantity" 
                type="number" 
                min="1" 
                :max="product.stock"
                class="qty-input"
              />
              <button @click="increaseQuantity" :disabled="quantity >= product.stock" class="qty-btn">+</button>
            </div>
          </div>

          <button 
            @click="addToCart" 
            :disabled="!canAddToCart"
            class="add-to-cart-btn"
          >
            {{ addToCartText }}
          </button>
        </div>

        <div class="product-features">
          <div class="feature">
            <span class="feature-icon">🚚</span>
            <span>Free shipping on orders over $50</span>
          </div>
          <div class="feature">
            <span class="feature-icon">↩️</span>
            <span>30-day return policy</span>
          </div>
          <div class="feature">
            <span class="feature-icon">🔒</span>
            <span>Secure payment</span>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div v-else-if="loading" class="loading">
    <div class="spinner"></div>
    <p>Loading product...</p>
  </div>

  <div v-else class="error">
    <h2>Product not found</h2>
    <p>The product you're looking for doesn't exist.</p>
    <router-link to="/products" class="back-link">← Back to Products</router-link>
  </div>

  <!-- Image Modal -->
  <div v-if="showImageModal" class="image-modal" @click="closeImageModal">
    <div class="modal-content" @click.stop>
      <button class="close-btn" @click="closeImageModal">&times;</button>
      <img :src="product?.image" :alt="product?.name" class="modal-image" />
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import apiService from '../services/api'

export default {
  name: 'ProductDetail',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const cartStore = useCartStore()
    
    const product = ref(null)
    const loading = ref(true)
    const quantity = ref(1)
    const showImageModal = ref(false)
    
    const stockClass = computed(() => ({
      'in-stock': product.value?.stock > 10,
      'low-stock': product.value?.stock > 0 && product.value?.stock <= 10,
      'out-of-stock': product.value?.stock === 0
    }))
    
    const stockText = computed(() => {
      if (!product.value) return ''
      if (product.value.stock === 0) return 'Out of stock'
      if (product.value.stock <= 10) return `Only ${product.value.stock} left`
      return 'In stock'
    })
    
    const canAddToCart = computed(() => {
      return product.value && product.value.stock > 0 && quantity.value <= product.value.stock
    })
    
    const addToCartText = computed(() => {
      if (!product.value) return 'Loading...'
      if (product.value.stock === 0) return 'Out of Stock'
      if (cartStore.items.some(item => item.id === product.value.id)) return 'Update Cart'
      return 'Add to Cart'
    })
    
    const fetchProduct = async () => {
      try {
        loading.value = true
        const response = await apiService.getProduct(route.params.id)
        product.value = response.data
      } catch (error) {
        console.error('Error fetching product:', error)
        product.value = null
      } finally {
        loading.value = false
      }
    }
    
    const addToCart = () => {
      if (canAddToCart.value) {
        for (let i = 0; i < quantity.value; i++) {
          cartStore.addToCart(product.value)
        }
      }
    }
    
    const increaseQuantity = () => {
      if (quantity.value < product.value.stock) {
        quantity.value++
      }
    }
    
    const decreaseQuantity = () => {
      if (quantity.value > 1) {
        quantity.value--
      }
    }
    
    const openImageModal = () => {
      showImageModal.value = true
    }
    
    const closeImageModal = () => {
      showImageModal.value = false
    }
    
    onMounted(() => {
      fetchProduct()
    })
    
    return {
      product,
      loading,
      quantity,
      stockClass,
      stockText,
      canAddToCart,
      addToCartText,
      addToCart,
      increaseQuantity,
      decreaseQuantity,
      showImageModal,
      openImageModal,
      closeImageModal
    }
  }
}
</script>

<style scoped>
.product-detail {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.breadcrumb {
  margin-bottom: 30px;
  font-size: 14px;
  color: #666;
}

.breadcrumb-link {
  color: #667eea;
  text-decoration: none;
}

.breadcrumb-link:hover {
  text-decoration: underline;
}

.breadcrumb-separator {
  margin: 0 10px;
}

.breadcrumb-current {
  color: #333;
}

.product-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: start;
}

.product-image-section {
  position: sticky;
  top: 20px;
}

.main-image {
  position: relative;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.main-image img {
  width: 100%;
  height: 500px;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.zoom-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.zoom-icon {
  font-size: 2rem;
  color: white;
}

.main-image:hover .zoom-overlay {
  opacity: 1;
}

.main-image:hover img {
  transform: scale(1.05);
}

.image-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.9);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

.modal-content {
  position: relative;
  max-width: 90vw;
  max-height: 90vh;
  animation: zoomIn 0.3s ease;
}

.modal-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
  border-radius: 8px;
}

.close-btn {
  position: absolute;
  top: -40px;
  right: 0;
  background: none;
  border: none;
  color: white;
  font-size: 2rem;
  cursor: pointer;
  padding: 5px;
  transition: opacity 0.3s ease;
}

.close-btn:hover {
  opacity: 0.7;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes zoomIn {
  from { transform: scale(0.8); }
  to { transform: scale(1); }
}

.product-info-section {
  padding: 20px 0;
}

.product-header {
  margin-bottom: 30px;
}

.category-badge {
  display: inline-block;
  background: #f0f0f0;
  color: #666;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 15px;
}

.product-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
  margin: 0;
  line-height: 1.2;
}

.price-section {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #eee;
}

.price {
  font-size: 2rem;
  font-weight: 700;
  color: #667eea;
}

.stock-info {
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
}

.stock-info.in-stock {
  background: #d4edda;
  color: #155724;
}

.stock-info.low-stock {
  background: #fff3cd;
  color: #856404;
}

.stock-info.out-of-stock {
  background: #f8d7da;
  color: #721c24;
}

.description-section {
  margin-bottom: 40px;
}

.description-section h3 {
  font-size: 1.2rem;
  color: #2c3e50;
  margin-bottom: 15px;
}

.description {
  font-size: 1rem;
  line-height: 1.6;
  color: #666;
}

.actions-section {
  margin-bottom: 40px;
}

.quantity-selector {
  margin-bottom: 20px;
}

.quantity-selector label {
  display: block;
  font-weight: 500;
  margin-bottom: 10px;
  color: #2c3e50;
}

.quantity-controls {
  display: flex;
  align-items: center;
  gap: 0;
  width: fit-content;
  border: 2px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
}

.qty-btn {
  background: #f8f9fa;
  border: none;
  padding: 12px 16px;
  cursor: pointer;
  font-size: 18px;
  font-weight: 600;
  color: #666;
  transition: all 0.2s ease;
}

.qty-btn:hover:not(:disabled) {
  background: #e9ecef;
  color: #333;
}

.qty-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.qty-input {
  border: none;
  padding: 12px 16px;
  width: 60px;
  text-align: center;
  font-size: 16px;
  font-weight: 500;
  outline: none;
}

.add-to-cart-btn {
  width: 100%;
  padding: 16px 32px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.add-to-cart-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.add-to-cart-btn:disabled {
  background: #95a5a6;
  cursor: not-allowed;
  transform: none;
}

.product-features {
  border-top: 1px solid #eee;
  padding-top: 30px;
}

.feature {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 15px;
  font-size: 14px;
  color: #666;
}

.feature-icon {
  font-size: 16px;
}

.loading, .error {
  text-align: center;
  padding: 60px 20px;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 20px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.back-link {
  color: #667eea;
  text-decoration: none;
  font-weight: 500;
}

.back-link:hover {
  text-decoration: underline;
}

@media (max-width: 768px) {
  .product-container {
    grid-template-columns: 1fr;
    gap: 30px;
  }
  
  .product-title {
    font-size: 2rem;
  }
  
  .price {
    font-size: 1.5rem;
  }
  
  .main-image img {
    height: 300px;
  }
  
  .modal-content {
    max-width: 95vw;
    max-height: 95vh;
  }
  
  .close-btn {
    top: -30px;
    font-size: 1.5rem;
  }
}
</style>