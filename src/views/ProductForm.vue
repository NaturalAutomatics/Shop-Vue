<template>
  <div class="product-form-overlay" @click="closeForm">
    <div class="product-form" @click.stop>
      <div class="form-header">
        <h2>{{ isEdit ? 'Edit Product' : 'Add New Product' }}</h2>
        <button @click="closeForm" class="close-btn">&times;</button>
      </div>

      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="name">Product Name *</label>
          <input
            id="name"
            v-model="form.name"
            type="text"
            required
            placeholder="Enter product name"
          />
        </div>

        <div class="form-group">
          <label for="description">Description *</label>
          <textarea
            id="description"
            v-model="form.description"
            required
            placeholder="Enter product description"
            rows="3"
          ></textarea>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label for="price">Price *</label>
            <input
              id="price"
              v-model.number="form.price"
              type="number"
              step="0.01"
              min="0"
              required
              placeholder="0.00"
            />
          </div>

          <div class="form-group">
            <label for="stock">Stock *</label>
            <input
              id="stock"
              v-model.number="form.stock"
              type="number"
              min="0"
              required
              placeholder="0"
            />
          </div>
        </div>

        <div class="form-group">
          <label for="category">Category *</label>
          <select id="category" v-model="form.category" required>
            <option value="">Select category</option>
            <option value="electronics">Electronics</option>
            <option value="clothing">Clothing</option>
            <option value="books">Books</option>
            <option value="home">Home & Garden</option>
            <option value="sports">Sports</option>
            <option value="toys">Toys</option>
          </select>
        </div>

        <div class="form-group">
          <label for="image">Image URL</label>
          <input
            id="image"
            v-model="form.image"
            type="url"
            placeholder="https://example.com/image.jpg"
          />
          <div v-if="form.image" class="image-preview">
            <img :src="form.image" alt="Preview" @error="imageError = true" />
            <span v-if="imageError" class="image-error">Invalid image URL</span>
          </div>
        </div>

        <div class="form-actions">
          <button type="button" @click="closeForm" class="cancel-btn">
            Cancel
          </button>
          <button type="submit" :disabled="loading" class="submit-btn">
            {{ loading ? 'Saving...' : (isEdit ? 'Update Product' : 'Add Product') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import apiService from '../services/api.js'

const props = defineProps({
  product: {
    type: Object,
    default: null
  },
  isEdit: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'success'])

const loading = ref(false)
const imageError = ref(false)

const form = reactive({
  name: '',
  description: '',
  price: 0,
  stock: 0,
  category: '',
  image: ''
})

// Initialize form with product data if editing
if (props.isEdit && props.product) {
  Object.assign(form, props.product)
}

// Watch for image URL changes to reset error state
watch(() => form.image, () => {
  imageError.value = false
})

const closeForm = () => {
  emit('close')
}

const submitForm = async () => {
  loading.value = true
  
  try {
    const productData = {
      name: form.name,
      description: form.description,
      price: parseFloat(form.price),
      stock: parseInt(form.stock),
      category: form.category,
      image: form.image || 'https://via.placeholder.com/300x200?text=No+Image'
    }

    let response
    if (props.isEdit) {
      response = await apiService.request(`/admin/products/${props.product.id}`, {
        method: 'PUT',
        body: JSON.stringify(productData)
      })
    } else {
      response = await apiService.request('/admin/products', {
        method: 'POST',
        body: JSON.stringify(productData)
      })
    }

    if (response.success) {
      emit('success', response.data)
      closeForm()
    }
  } catch (error) {
    console.error('Error saving product:', error)
    alert('Failed to save product: ' + error.message)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.product-form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.product-form {
  background: white;
  border-radius: 10px;
  padding: 0;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.form-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #eee;
}

.form-header h2 {
  margin: 0;
  color: #2c3e50;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #999;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-btn:hover {
  color: #333;
}

form {
  padding: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #2c3e50;
}

.form-group input,
.form-group textarea,
.form-group select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  box-sizing: border-box;
}

.form-group input:focus,
.form-group textarea:focus,
.form-group select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
}

.image-preview {
  margin-top: 10px;
  text-align: center;
}

.image-preview img {
  max-width: 200px;
  max-height: 150px;
  border-radius: 6px;
  border: 1px solid #ddd;
}

.image-error {
  color: #dc3545;
  font-size: 12px;
  display: block;
  margin-top: 5px;
}

.form-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.cancel-btn,
.submit-btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
}

.cancel-btn {
  background: #6c757d;
  color: white;
}

.submit-btn {
  background: #28a745;
  color: white;
}

.submit-btn:disabled {
  background: #95a5a6;
  cursor: not-allowed;
}

.cancel-btn:hover:not(:disabled),
.submit-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

@media (max-width: 768px) {
  .product-form {
    width: 95%;
    margin: 20px;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .form-actions {
    flex-direction: column;
  }
}
</style>