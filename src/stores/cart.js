import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

// Define store - similar to Angular services but with built-in reactivity
export const useCartStore = defineStore('cart', () => {
  // Reactive state - similar to BehaviorSubject in Angular
  const items = ref([])
  
  // Computed properties - similar to Angular getters
  const itemCount = computed(() => {
    return items.value.reduce((total, item) => total + item.quantity, 0)
  })
  
  const totalPrice = computed(() => {
    return items.value.reduce((total, item) => total + (item.price * item.quantity), 0)
  })
  
  // Actions - similar to Angular service methods
  const addToCart = (product) => {
    const existingItem = items.value.find(item => item.id === product.id)
    
    if (existingItem) {
      existingItem.quantity++
    } else {
      items.value.push({
        ...product,
        quantity: 1
      })
    }
  }
  
  const removeFromCart = (productId) => {
    const index = items.value.findIndex(item => item.id === productId)
    if (index > -1) {
      items.value.splice(index, 1)
    }
  }
  
  const updateQuantity = (productId, quantity) => {
    const item = items.value.find(item => item.id === productId)
    if (item) {
      if (quantity <= 0) {
        removeFromCart(productId)
      } else {
        item.quantity = quantity
      }
    }
  }
  
  const clearCart = () => {
    items.value = []
  }
  
  return {
    // State
    items,
    
    // Computed
    itemCount,
    totalPrice,
    
    // Actions
    addToCart,
    removeFromCart,
    updateQuantity,
    clearCart
  }
})
