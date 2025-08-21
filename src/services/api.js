// API service for backend communication
const API_BASE_URL = 'http://localhost:5000/api';

class ApiService {
  // Get auth token from localStorage
  getAuthToken() {
    return localStorage.getItem('token');
  }

  // Generic request method
  async request(endpoint, options = {}) {
    const url = `${API_BASE_URL}${endpoint}`;
    const token = this.getAuthToken();
    
    const config = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    };

    // Add Authorization header if token exists
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }

    try {
      const response = await fetch(url, config);
      const data = await response.json();

      if (!response.ok) {
        const error = new Error(data.message || data.error || `HTTP error! status: ${response.status}`);
        error.status = response.status;
        throw error;
      }

      return data;
    } catch (error) {
      console.error('API request failed:', error);
      throw error;
    }
  }

  // Product endpoints
  async getProducts(params = {}) {
    const queryString = new URLSearchParams(params).toString();
    const endpoint = `/products${queryString ? `?${queryString}` : ''}`;
    return this.request(endpoint);
  }

  async getProduct(id) {
    return this.request(`/products/${id}`);
  }

  async getCategories() {
    return this.request('/products/categories');
  }

  // Order endpoints
  async createOrder(orderData) {
    return this.request('/orders', {
      method: 'POST',
      body: JSON.stringify(orderData)
    });
  }

  async getOrder(orderNumber) {
    return this.request(`/orders/${orderNumber}`);
  }

  async getAllOrders() {
    return this.request('/orders');
  }

  async updateOrderStatus(orderId, status) {
    return this.request(`/orders/${orderId}/status`, {
      method: 'PUT',
      body: JSON.stringify({ status })
    });
  }

  // Authentication endpoints
  async login(credentials) {
    return this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(credentials)
    });
  }

  async logout() {
    return this.request('/auth/logout', {
      method: 'POST'
    });
  }

  async getCurrentUser() {
    return this.request('/auth/me');
  }
}

// Create and export a singleton instance
const apiService = new ApiService();
export default apiService;
