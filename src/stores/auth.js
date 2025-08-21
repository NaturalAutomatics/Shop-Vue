import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import apiService from '../services/api'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref(null)
  const token = ref(localStorage.getItem('token') || null)
  const loading = ref(false)
  const error = ref(null)

  // Getters
  const isAuthenticated = computed(() => !!token.value && !!user.value)
  const isAdmin = computed(() => user.value?.role === 'admin')

  // Actions
  const login = async (credentials) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await apiService.login(credentials)
      user.value = response.data.user
      token.value = response.data.token
      
      // Store token in localStorage
      localStorage.setItem('token', token.value)
      
      return response
    } catch (err) {
      error.value = err.message || 'Login failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    loading.value = true
    
    try {
      if (token.value) {
        await apiService.logout()
      }
    } catch (err) {
      console.error('Logout error:', err)
    } finally {
      // Clear state regardless of API call success
      user.value = null
      token.value = null
      localStorage.removeItem('token')
      loading.value = false
    }
  }

  const fetchCurrentUser = async () => {
    if (!token.value) return
    
    loading.value = true
    error.value = null
    
    try {
      const response = await apiService.getCurrentUser()
      user.value = response.data
    } catch (err) {
      error.value = err.message || 'Failed to fetch user data'
      // If token is invalid, clear it
      if (err.status === 401) {
        logout()
      }
    } finally {
      loading.value = false
    }
  }

  const clearError = () => {
    error.value = null
  }

  return {
    // State
    user,
    token,
    loading,
    error,
    
    // Getters
    isAuthenticated,
    isAdmin,
    
    // Actions
    login,
    logout,
    fetchCurrentUser,
    clearError
  }
})
