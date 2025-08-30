<template>
  <div class="cart">
    <div class="cart-header">
      <h1>Shopping Cart</h1>
      <button 
        v-if="cartItems.length > 0" 
        @click="clearCart" 
        class="clear-cart-btn"
      >
        Clear Cart
      </button>
    </div>

    <!-- Empty cart message -->
    <div v-if="cartItems.length === 0" class="empty-cart">
      <div class="empty-cart-icon">🛒</div>
      <h2>Your cart is empty</h2>
      <p>Add some products to get started!</p>
      <router-link to="/products" class="btn btn-primary">
        Browse Products
      </router-link>
    </div>

    <!-- Cart items -->
    <div v-else class="cart-content">
      <div class="cart-items">
        <div 
          v-for="item in cartItems" 
          :key="item.id" 
          class="cart-item"
        >
          <div class="item-image">
            <img :src="item.image" :alt="item.name" />
          </div>
          <div class="item-details">
            <h3>{{ item.name }}</h3>
            <p class="item-description">{{ item.description }}</p>
            <div class="item-price">${{ item.price }}</div>
          </div>
          <div class="item-quantity">
            <button 
              @click="updateQuantity(item.id, item.quantity - 1)"
              class="quantity-btn"
              :disabled="item.quantity <= 1"
            >
              -
            </button>
            <span class="quantity">{{ item.quantity }}</span>
            <button 
              @click="updateQuantity(item.id, item.quantity + 1)"
              class="quantity-btn"
            >
              +
            </button>
          </div>
          <div class="item-total">
            ${{ (item.price * item.quantity).toFixed(2) }}
          </div>
          <button 
            @click="removeFromCart(item.id)" 
            class="remove-btn"
          >
            ✕
          </button>
        </div>
      </div>

      <!-- Cart summary -->
      <div class="cart-summary">
        <h2>Order Summary</h2>
        <div class="summary-row">
          <span>Items ({{ itemCount }}):</span>
          <span>${{ subtotal.toFixed(2) }}</span>
        </div>
        <div class="summary-row">
          <span>Shipping:</span>
          <span>${{ shipping.toFixed(2) }}</span>
        </div>
        <div class="summary-row">
          <span>Tax:</span>
          <span>${{ tax.toFixed(2) }}</span>
        </div>
        <div class="summary-row total">
          <span>Total:</span>
          <span>${{ total.toFixed(2) }}</span>
        </div>
        <button @click="startCheckout" class="checkout-btn">
          Proceed to Checkout
        </button>
      </div>
    </div>

    <!-- Checkout Form Modal -->
    <div v-if="showCheckoutForm" class="checkout-modal">
      <div class="checkout-modal-content">
        <div class="checkout-header">
          <h2>Complete Your Order</h2>
          <button @click="resetCheckout" class="close-btn">✕</button>
        </div>
        
        <form @submit.prevent="processOrder" class="checkout-form">
          <div class="form-group">
            <label for="name">Full Name *</label>
            <input 
              id="name"
              v-model="checkoutForm.name" 
              type="text" 
              required 
              placeholder="Enter your full name"
            />
          </div>
          
          <div class="form-group">
            <label for="email">Email Address *</label>
            <input 
              id="email"
              v-model="checkoutForm.email" 
              type="email" 
              required 
              placeholder="Enter your email address"
            />
          </div>
          
          <div class="form-group">
            <label for="address">Shipping Address *</label>
            <textarea 
              id="address"
              v-model="checkoutForm.address" 
              required 
              placeholder="Enter your complete shipping address"
              rows="3"
            ></textarea>
          </div>
          
          <div class="order-summary">
            <h3>Order Summary</h3>
            <div class="summary-item">
              <span>Items ({{ itemCount }}):</span>
              <span>${{ subtotal.toFixed(2) }}</span>
            </div>
            <div class="summary-item">
              <span>Shipping:</span>
              <span>${{ shipping.toFixed(2) }}</span>
            </div>
            <div class="summary-item">
              <span>Tax:</span>
              <span>${{ tax.toFixed(2) }}</span>
            </div>
            <div class="summary-item total">
              <span>Total:</span>
              <span>${{ total.toFixed(2) }}</span>
            </div>
          </div>
          
          <div class="checkout-actions">
            <button 
              type="button" 
              @click="resetCheckout" 
              class="btn-secondary"
              :disabled="isProcessing"
            >
              Cancel
            </button>
            <button 
              type="submit" 
              class="btn-primary"
              :disabled="isProcessing"
            >
              {{ isProcessing ? 'Processing...' : 'Place Order' }}
            </button>
          </div>
          
          <div v-if="error" class="error-message">
            {{ error }}
          </div>
        </form>
      </div>
    </div>

    <!-- Order Confirmation Modal -->
    <div v-if="orderResult" class="order-confirmation-modal">
      <div class="order-confirmation-content">
        <div class="success-icon">✅</div>
        <h2>Order Placed Successfully!</h2>
        <div class="order-details">
          <p><strong>Order Number:</strong> {{ orderResult.orderNumber }}</p>
          <p><strong>Total Amount:</strong> ${{ orderResult.total.toFixed(2) }}</p>
          <p><strong>Status:</strong> {{ orderResult.status }}</p>
        </div>
        <p class="confirmation-message">
          Thank you for your order! We'll send you an email confirmation shortly.
          You can use your order number to track your order status.
        </p>
        <div class="confirmation-actions">
          <button @click="resetCheckout" class="btn-primary">
            Continue Shopping
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, ref } from 'vue'
import { useCartStore } from '../stores/cart'
import apiService from '../services/api'

export default {
  name: 'Cart',
  setup() {
    const cartStore = useCartStore()
    
    // Reactive state for checkout
    const showCheckoutForm = ref(false)
    const checkoutForm = ref({
      name: '',
      email: '',
      address: ''
    })
    const isProcessing = ref(false)
    const orderResult = ref(null)
    const error = ref(null)
    
    // Computed properties
    const cartItems = computed(() => cartStore.items)
    const itemCount = computed(() => cartStore.itemCount)
    const subtotal = computed(() => cartStore.totalPrice)
    const shipping = computed(() => subtotal.value > 0 ? 5.99 : 0)
    const tax = computed(() => subtotal.value * 0.08) // 8% tax
    const total = computed(() => subtotal.value + shipping.value + tax.value)
    
    // Methods
    const updateQuantity = (productId, quantity) => {
      cartStore.updateQuantity(productId, quantity)
    }
    
    const removeFromCart = (productId) => {
      cartStore.removeFromCart(productId)
    }
    
    const clearCart = () => {
      cartStore.clearCart()
    }
    
    const startCheckout = () => {
      if (cartItems.value.length === 0) {
        alert('Your cart is empty!')
        return
      }
      showCheckoutForm.value = true
    }
    
    const processOrder = async () => {
      if (!checkoutForm.value.name || !checkoutForm.value.email || !checkoutForm.value.address) {
        alert('Please fill in all required fields.')
        return
      }
      
      isProcessing.value = true
      error.value = null
      
      try {
        const orderData = {
          customer: {
            name: checkoutForm.value.name,
            email: checkoutForm.value.email,
            address: checkoutForm.value.address
          },
          items: cartItems.value,
          subtotal: subtotal.value,
          shipping: shipping.value,
          tax: tax.value,
          total: total.value
        }
        
        const response = await apiService.createOrder(orderData)
        orderResult.value = response.data
        
        // Clear cart after successful order
        cartStore.clearCart()
        showCheckoutForm.value = false
        
      } catch (err) {
        error.value = 'Failed to process order. Please try again.'
        console.error('Error creating order:', err)
      } finally {
        isProcessing.value = false
      }
    }
    
    const resetCheckout = () => {
      showCheckoutForm.value = false
      orderResult.value = null
      error.value = null
      checkoutForm.value = {
        name: '',
        email: '',
        address: ''
      }
    }
    
    return {
      cartItems,
      itemCount,
      subtotal,
      shipping,
      tax,
      total,
      showCheckoutForm,
      checkoutForm,
      isProcessing,
      orderResult,
      error,
      updateQuantity,
      removeFromCart,
      clearCart,
      startCheckout,
      processOrder,
      resetCheckout
    }
  }
}
</script>

<style scoped>
.cart {
  max-width: 1200px;
  margin: 0 auto;
}

.cart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.cart-header h1 {
  color: #2c3e50;
  margin: 0;
}

.clear-cart-btn {
  padding: 0.5rem 1rem;
  background: #e74c3c;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: background-color 0.3s;
}

.clear-cart-btn:hover {
  background: #c0392b;
}

.empty-cart {
  text-align: center;
  padding: 4rem 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.empty-cart-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
}

.empty-cart h2 {
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.empty-cart p {
  color: #7f8c8d;
  margin-bottom: 2rem;
}

.cart-content {
  display: grid;
  grid-template-columns: 1fr 350px;
  gap: 2rem;
}

.cart-items {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.cart-item {
  display: grid;
  grid-template-columns: 100px 1fr auto auto auto;
  gap: 1rem;
  padding: 1.5rem;
  border-bottom: 1px solid #ecf0f1;
  align-items: center;
}

.cart-item:last-child {
  border-bottom: none;
}

.item-image {
  width: 100px;
  height: 80px;
  overflow: hidden;
  border-radius: 8px;
}

.item-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-details h3 {
  color: #2c3e50;
  margin: 0 0 0.5rem 0;
  font-size: 1.1rem;
}

.item-description {
  color: #7f8c8d;
  margin: 0 0 0.5rem 0;
  font-size: 0.9rem;
}

.item-price {
  color: #667eea;
  font-weight: 600;
}

.item-quantity {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.quantity-btn {
  width: 30px;
  height: 30px;
  border: 1px solid #ddd;
  background: white;
  border-radius: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  transition: all 0.3s;
}

.quantity-btn:hover:not(:disabled) {
  background: #667eea;
  color: white;
  border-color: #667eea;
}

.quantity-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.quantity {
  min-width: 30px;
  text-align: center;
  font-weight: 600;
}

.item-total {
  font-weight: bold;
  color: #2c3e50;
  font-size: 1.1rem;
}

.remove-btn {
  width: 30px;
  height: 30px;
  border: none;
  background: #e74c3c;
  color: white;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.8rem;
  transition: background-color 0.3s;
}

.remove-btn:hover {
  background: #c0392b;
}

.cart-summary {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  height: fit-content;
}

.cart-summary h2 {
  color: #2c3e50;
  margin: 0 0 1.5rem 0;
}

.summary-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;
  color: #7f8c8d;
}

.summary-row.total {
  border-top: 1px solid #ecf0f1;
  padding-top: 1rem;
  margin-top: 1rem;
  font-weight: bold;
  color: #2c3e50;
  font-size: 1.2rem;
}

.checkout-btn {
  width: 100%;
  padding: 1rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 6px;
  font-weight: 600;
  font-size: 1.1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 1rem;
}

.checkout-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

@media (max-width: 768px) {
  .cart-content {
    grid-template-columns: 1fr;
  }
  
  .cart-item {
    grid-template-columns: 80px 1fr;
    gap: 1rem;
  }
  
  .item-quantity,
  .item-total,
  .remove-btn {
    grid-column: 2;
    justify-self: start;
  }
  
  .item-quantity {
    margin-top: 0.5rem;
  }
  
  .item-total {
    margin-top: 0.5rem;
  }
  
  .remove-btn {
    justify-self: end;
    margin-top: 0.5rem;
  }
}

/* Checkout Modal Styles */
.checkout-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.checkout-modal-content {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  max-width: 500px;
  width: 90%;
  max-height: 90vh;
  overflow-y: auto;
}

.checkout-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.checkout-header h2 {
  margin: 0;
  color: #2c3e50;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #7f8c8d;
  padding: 0.5rem;
}

.close-btn:hover {
  color: #e74c3c;
}

.checkout-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-weight: 600;
  color: #2c3e50;
}

.form-group input,
.form-group textarea {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #667eea;
}

.order-summary {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  margin: 1rem 0;
}

.order-summary h3 {
  margin: 0 0 1rem 0;
  color: #2c3e50;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  color: #7f8c8d;
}

.summary-item.total {
  border-top: 1px solid #e9ecef;
  padding-top: 0.5rem;
  margin-top: 0.5rem;
  font-weight: bold;
  color: #2c3e50;
  font-size: 1.1rem;
}

.checkout-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}

.error-message {
  background: #fee;
  color: #c53030;
  padding: 1rem;
  border-radius: 6px;
  border: 1px solid #fed7d7;
}

/* Order Confirmation Modal */
.order-confirmation-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.order-confirmation-content {
  background: white;
  border-radius: 12px;
  padding: 3rem 2rem;
  max-width: 500px;
  width: 90%;
  text-align: center;
}

.success-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
}

.order-confirmation-content h2 {
  color: #2c3e50;
  margin-bottom: 1.5rem;
}

.order-details {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  margin: 1.5rem 0;
  text-align: left;
}

.order-details p {
  margin: 0.5rem 0;
  color: #2c3e50;
}

.confirmation-message {
  color: #7f8c8d;
  line-height: 1.6;
  margin-bottom: 2rem;
}

.confirmation-actions {
  display: flex;
  justify-content: center;
}

@media (max-width: 768px) {
  .checkout-modal-content,
  .order-confirmation-content {
    width: 95%;
    padding: 1.5rem;
  }
  
  .checkout-actions {
    flex-direction: column;
  }
}
</style>
