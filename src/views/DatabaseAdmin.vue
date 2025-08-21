<template>
  <div class="database-admin">
    <div class="header">
      <h1>🗄️ Database Administration Panel</h1>
      <p>Manage PostgreSQL database for Vue Shop</p>
    </div>

    <div class="connection-status">
      <div class="status-card" :class="{ 'connected': dbConnected, 'disconnected': !dbConnected }">
        <div class="status-indicator"></div>
        <div class="status-text">
          <h3>Database Connection</h3>
          <p>{{ dbConnected ? 'Connected' : 'Disconnected' }}</p>
          <small>{{ dbConnected ? 'PostgreSQL is running and accessible' : 'Check PostgreSQL service and connection settings' }}</small>
        </div>
      </div>
    </div>

    <div class="admin-grid">
      <!-- Database Statistics -->
      <div class="stats-card">
        <h3>📊 Database Statistics</h3>
        <div class="stats-grid">
          <div class="stat-item">
            <span class="stat-number">{{ stats.products || 0 }}</span>
            <span class="stat-label">Products</span>
          </div>
          <div class="stat-item">
            <span class="stat-number">{{ stats.users || 0 }}</span>
            <span class="stat-label">Users</span>
          </div>
          <div class="stat-item">
            <span class="stat-number">{{ stats.orders || 0 }}</span>
            <span class="stat-label">Orders</span>
          </div>
          <div class="stat-item">
            <span class="stat-number">{{ stats.totalValue || 0 }}</span>
            <span class="stat-label">Total Value</span>
          </div>
        </div>
        <button @click="refreshStats" class="refresh-btn" :disabled="loading">
          {{ loading ? 'Refreshing...' : '🔄 Refresh Stats' }}
        </button>
      </div>

      <!-- Quick Actions -->
      <div class="actions-card">
        <h3>⚡ Quick Actions</h3>
        <div class="action-buttons">
          <button @click="seedDatabase" class="action-btn primary" :disabled="loading">
            🌱 Seed Database
          </button>
          <button @click="clearDatabase" class="action-btn danger" :disabled="loading">
            🗑️ Clear All Data
          </button>
          <button @click="exportData" class="action-btn secondary" :disabled="loading">
            📤 Export Data
          </button>
          <button @click="backupDatabase" class="action-btn secondary" :disabled="loading">
            💾 Backup DB
          </button>
        </div>
      </div>

      <!-- Connection Settings -->
      <div class="settings-card">
        <h3>🔧 Connection Settings</h3>
        <div class="setting-item">
          <label>Host:</label>
          <input v-model="connection.host" type="text" placeholder="localhost" />
        </div>
        <div class="setting-item">
          <label>Port:</label>
          <input v-model="connection.port" type="text" placeholder="5432" />
        </div>
        <div class="setting-item">
          <label>Database:</label>
          <input v-model="connection.database" type="text" placeholder="vueshop" />
        </div>
        <div class="setting-item">
          <label>Username:</label>
          <input v-model="connection.username" type="text" placeholder="vueshop" />
        </div>
        <div class="setting-item">
          <label>Password:</label>
          <input v-model="connection.password" type="password" placeholder="••••••••" />
        </div>
        <button @click="testConnection" class="test-btn" :disabled="loading">
          🧪 Test Connection
        </button>
      </div>

      <!-- pgAdmin Access -->
      <div class="pgadmin-card">
        <h3>🌐 pgAdmin Access</h3>
        <p>Web-based PostgreSQL administration tool</p>
        <div class="pgadmin-info">
          <div class="info-item">
            <strong>URL:</strong> <a href="http://localhost:5050" target="_blank">http://localhost:5050</a>
          </div>
          <div class="info-item">
            <strong>Email:</strong> admin@admin.com
          </div>
          <div class="info-item">
            <strong>Password:</strong> root
          </div>
        </div>
        <button @click="openPgAdmin" class="pgadmin-btn">
          🚀 Open pgAdmin
        </button>
      </div>
    </div>

    <!-- Data Tables -->
    <div class="data-section">
      <h3>📋 Data Management</h3>
      
      <!-- Products Table -->
      <div class="table-section">
        <h4>Products ({{ products.length }})</h4>
        <div class="table-actions">
          <button @click="addProduct" class="add-btn">➕ Add Product</button>
          <button @click="refreshProducts" class="refresh-btn">🔄 Refresh</button>
        </div>
        <div class="table-container">
          <table class="data-table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Category</th>
                <th>Price</th>
                <th>Stock</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="product in products" :key="product.id">
                <td>{{ product.id }}</td>
                <td>{{ product.name }}</td>
                <td>{{ product.category }}</td>
                <td>${{ product.price }}</td>
                <td>{{ product.stock }}</td>
                <td>
                  <button @click="editProduct(product)" class="edit-btn">✏️</button>
                  <button @click="deleteProduct(product.id)" class="delete-btn">🗑️</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Users Table -->
      <div class="table-section">
        <h4>Users ({{ users.length }})</h4>
        <div class="table-actions">
          <button @click="addUser" class="add-btn">➕ Add User</button>
          <button @click="refreshUsers" class="refresh-btn">🔄 Refresh</button>
        </div>
        <div class="table-container">
          <table class="data-table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Username</th>
                <th>Name</th>
                <th>Email</th>
                <th>Role</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="user in users" :key="user.id">
                <td>{{ user.id }}</td>
                <td>{{ user.username }}</td>
                <td>{{ user.name }}</td>
                <td>{{ user.email }}</td>
                <td>{{ user.role }}</td>
                <td>
                  <button @click="editUser(user)" class="edit-btn">✏️</button>
                  <button @click="deleteUser(user.id)" class="delete-btn">🗑️</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Notifications -->
    <div v-if="notification.show" class="notification" :class="notification.type">
      {{ notification.message }}
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import apiService from '../services/api.js'

// Reactive state
const loading = ref(false)
const dbConnected = ref(false)
const products = ref([])
const users = ref([])
const stats = reactive({
  products: 0,
  users: 0,
  orders: 0,
  totalValue: 0
})

const connection = reactive({
  host: 'localhost',
  port: '5432',
  database: 'vueshop',
  username: 'vueshop',
  password: 'vueshop123'
})

const notification = reactive({
  show: false,
  message: '',
  type: 'info'
})

// Methods
const showNotification = (message, type = 'info') => {
  notification.message = message
  notification.type = type
  notification.show = true
  setTimeout(() => {
    notification.show = false
  }, 5000)
}

const testConnection = async () => {
  loading.value = true
  try {
    const response = await apiService.request('/admin/test-connection', 'POST', connection)
    if (response.success) {
      dbConnected.value = true
      showNotification('Database connection successful!', 'success')
    } else {
      dbConnected.value = false
      showNotification('Database connection failed: ' + response.message, 'error')
    }
  } catch (error) {
    dbConnected.value = false
    showNotification('Connection test failed: ' + error.message, 'error')
  } finally {
    loading.value = false
  }
}

const refreshStats = async () => {
  loading.value = true
  try {
    const response = await apiService.request('/admin/stats', 'GET')
    if (response.success) {
      Object.assign(stats, response.data)
    }
  } catch (error) {
    showNotification('Failed to fetch stats: ' + error.message, 'error')
  } finally {
    loading.value = false
  }
}

const refreshProducts = async () => {
  try {
    const response = await apiService.request('/products', 'GET')
    if (response.success) {
      products.value = response.data
    }
  } catch (error) {
    showNotification('Failed to fetch products: ' + error.message, 'error')
  }
}

const refreshUsers = async () => {
  try {
    const response = await apiService.request('/admin/users', 'GET')
    if (response.success) {
      users.value = response.data
    }
  } catch (error) {
    showNotification('Failed to fetch users: ' + error.message, 'error')
  }
}

const seedDatabase = async () => {
  if (!confirm('This will seed the database with sample data. Continue?')) return
  
  loading.value = true
  try {
    const response = await apiService.request('/admin/seed', 'POST')
    if (response.success) {
      showNotification('Database seeded successfully!', 'success')
      await refreshStats()
      await refreshProducts()
      await refreshUsers()
    }
  } catch (error) {
    showNotification('Failed to seed database: ' + error.message, 'error')
  } finally {
    loading.value = false
  }
}

const clearDatabase = async () => {
  if (!confirm('⚠️ This will DELETE ALL DATA from the database! This action cannot be undone. Continue?')) return
  
  loading.value = true
  try {
    const response = await apiService.request('/admin/clear', 'POST')
    if (response.success) {
      showNotification('Database cleared successfully!', 'success')
      await refreshStats()
      await refreshProducts()
      await refreshUsers()
    }
  } catch (error) {
    showNotification('Failed to clear database: ' + error.message, 'error')
  } finally {
    loading.value = false
  }
}

const exportData = async () => {
  try {
    const response = await apiService.request('/admin/export', 'GET')
    if (response.success) {
      // Create and download file
      const blob = new Blob([JSON.stringify(response.data, null, 2)], { type: 'application/json' })
      const url = window.URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = `vueshop-export-${new Date().toISOString().split('T')[0]}.json`
      a.click()
      window.URL.revokeObjectURL(url)
      showNotification('Data exported successfully!', 'success')
    }
  } catch (error) {
    showNotification('Failed to export data: ' + error.message, 'error')
  }
}

const backupDatabase = async () => {
  showNotification('Backup functionality requires pgAdmin or manual pg_dump', 'info')
}

const openPgAdmin = () => {
  window.open('http://localhost:5050', '_blank')
}

const addProduct = () => {
  showNotification('Add product functionality coming soon!', 'info')
}

const editProduct = (product) => {
  showNotification(`Edit product ${product.name} - coming soon!`, 'info')
}

const deleteProduct = async (id) => {
  if (!confirm('Are you sure you want to delete this product?')) return
  
  try {
    const response = await apiService.request(`/admin/products/${id}`, 'DELETE')
    if (response.success) {
      showNotification('Product deleted successfully!', 'success')
      await refreshProducts()
      await refreshStats()
    }
  } catch (error) {
    showNotification('Failed to delete product: ' + error.message, 'error')
  }
}

const addUser = () => {
  showNotification('Add user functionality coming soon!', 'info')
}

const editUser = (user) => {
  showNotification(`Edit user ${user.username} - coming soon!`, 'info')
}

const deleteUser = async (id) => {
  if (!confirm('Are you sure you want to delete this user?')) return
  
  try {
    const response = await apiService.request(`/admin/users/${id}`, 'DELETE')
    if (response.success) {
      showNotification('User deleted successfully!', 'success')
      await refreshUsers()
      await refreshStats()
    }
  } catch (error) {
    showNotification('Failed to delete user: ' + error.message, 'error')
  }
}

// Lifecycle
onMounted(async () => {
  await testConnection()
  if (dbConnected.value) {
    await refreshStats()
    await refreshProducts()
    await refreshUsers()
  }
})
</script>

<style scoped>
.database-admin {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  text-align: center;
  margin-bottom: 30px;
}

.header h1 {
  color: #2c3e50;
  margin-bottom: 10px;
}

.header p {
  color: #7f8c8d;
  font-size: 18px;
}

.connection-status {
  margin-bottom: 30px;
}

.status-card {
  display: flex;
  align-items: center;
  padding: 20px;
  border-radius: 10px;
  background: #f8f9fa;
  border: 2px solid #e9ecef;
}

.status-card.connected {
  border-color: #28a745;
  background: #d4edda;
}

.status-card.disconnected {
  border-color: #dc3545;
  background: #f8d7da;
}

.status-indicator {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  margin-right: 15px;
  background: #dc3545;
}

.status-card.connected .status-indicator {
  background: #28a745;
}

.status-text h3 {
  margin: 0 0 5px 0;
  color: #2c3e50;
}

.status-text p {
  margin: 0 0 5px 0;
  font-weight: bold;
}

.status-text small {
  color: #6c757d;
}

.admin-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  margin-bottom: 40px;
}

.stats-card, .actions-card, .settings-card, .pgadmin-card {
  background: white;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.stats-card h3, .actions-card h3, .settings-card h3, .pgadmin-card h3 {
  margin: 0 0 20px 0;
  color: #2c3e50;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 15px;
  margin-bottom: 20px;
}

.stat-item {
  text-align: center;
  padding: 15px;
  background: #f8f9fa;
  border-radius: 8px;
}

.stat-number {
  display: block;
  font-size: 24px;
  font-weight: bold;
  color: #2c3e50;
}

.stat-label {
  color: #6c757d;
  font-size: 14px;
}

.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.action-btn {
  padding: 12px 20px;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.action-btn.primary {
  background: #007bff;
  color: white;
}

.action-btn.secondary {
  background: #6c757d;
  color: white;
}

.action-btn.danger {
  background: #dc3545;
  color: white;
}

.action-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.action-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.setting-item {
  margin-bottom: 15px;
}

.setting-item label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #2c3e50;
}

.setting-item input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.test-btn {
  width: 100%;
  padding: 12px;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  margin-top: 10px;
}

.pgadmin-info {
  margin: 15px 0;
}

.info-item {
  margin-bottom: 8px;
  font-size: 14px;
}

.info-item a {
  color: #007bff;
  text-decoration: none;
}

.pgadmin-btn {
  width: 100%;
  padding: 12px;
  background: #17a2b8;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.data-section {
  background: white;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.data-section h3 {
  margin: 0 0 20px 0;
  color: #2c3e50;
}

.table-section {
  margin-bottom: 30px;
}

.table-section h4 {
  margin: 0 0 15px 0;
  color: #2c3e50;
}

.table-actions {
  display: flex;
  gap: 10px;
  margin-bottom: 15px;
}

.add-btn, .refresh-btn {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.add-btn {
  background: #28a745;
  color: white;
}

.refresh-btn {
  background: #17a2b8;
  color: white;
}

.table-container {
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.data-table th,
.data-table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.data-table th {
  background: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.data-table tr:hover {
  background: #f8f9fa;
}

.edit-btn, .delete-btn {
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 5px;
  font-size: 12px;
}

.edit-btn {
  background: #ffc107;
  color: #212529;
}

.delete-btn {
  background: #dc3545;
  color: white;
}

.notification {
  position: fixed;
  top: 20px;
  right: 20px;
  padding: 15px 20px;
  border-radius: 6px;
  color: white;
  font-weight: 500;
  z-index: 1000;
  animation: slideIn 0.3s ease;
}

.notification.success {
  background: #28a745;
}

.notification.error {
  background: #dc3545;
}

.notification.info {
  background: #17a2b8;
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@media (max-width: 768px) {
  .admin-grid {
    grid-template-columns: 1fr;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
  }
  
  .table-container {
    font-size: 12px;
  }
  
  .data-table th,
  .data-table td {
    padding: 8px 6px;
  }
}
</style>
