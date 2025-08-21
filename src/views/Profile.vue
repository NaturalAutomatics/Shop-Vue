<template>
  <div class="profile">
    <div class="profile-header">
      <h1>My Profile</h1>
      <button @click="handleLogout" class="logout-btn" :disabled="loading">
        {{ loading ? 'Signing out...' : 'Sign Out' }}
      </button>
    </div>

    <div v-if="user" class="profile-content">
      <div class="profile-card">
        <div class="profile-avatar">
          <div class="avatar-placeholder">
            {{ user.name.charAt(0).toUpperCase() }}
          </div>
        </div>
        
        <div class="profile-info">
          <h2>{{ user.name }}</h2>
          <div class="info-item">
            <span class="label">Username:</span>
            <span class="value">{{ user.username }}</span>
          </div>
          <div class="info-item">
            <span class="label">Email:</span>
            <span class="value">{{ user.email }}</span>
          </div>
          <div class="info-item">
            <span class="label">Role:</span>
            <span class="value role-badge" :class="user.role">
              {{ user.role }}
            </span>
          </div>
          <div class="info-item">
            <span class="label">User ID:</span>
            <span class="value">{{ user.id }}</span>
          </div>
        </div>
      </div>

      <div class="profile-actions">
        <div class="action-card">
          <h3>Account Actions</h3>
          <div class="action-buttons">
            <button class="action-btn primary">
              Edit Profile
            </button>
            <button class="action-btn secondary">
              Change Password
            </button>
            <button class="action-btn secondary">
              View Orders
            </button>
          </div>
        </div>

        <div v-if="isAdmin" class="action-card admin">
          <h3>Admin Panel</h3>
          <div class="action-buttons">
            <button class="action-btn admin-btn">
              Manage Products
            </button>
            <button class="action-btn admin-btn">
              View All Orders
            </button>
            <button class="action-btn admin-btn">
              User Management
            </button>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="loading" class="loading">
      <p>Loading profile...</p>
    </div>

    <div v-else class="error">
      <p>Failed to load profile. Please try again.</p>
    </div>
  </div>
</template>

<script>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

export default {
  name: 'Profile',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()

    const handleLogout = async () => {
      try {
        await authStore.logout()
        router.push('/login')
      } catch (error) {
        console.error('Logout error:', error)
      }
    }

    onMounted(async () => {
      if (!authStore.isAuthenticated) {
        router.push('/login')
        return
      }

      if (!authStore.user) {
        await authStore.fetchCurrentUser()
      }
    })

    return {
      user: authStore.user,
      loading: authStore.loading,
      isAdmin: authStore.isAdmin,
      handleLogout
    }
  }
}
</script>

<style scoped>
.profile {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

.profile-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  flex-wrap: wrap;
  gap: 1rem;
}

.profile-header h1 {
  color: #2c3e50;
  margin: 0;
}

.logout-btn {
  padding: 0.5rem 1rem;
  background: #e74c3c;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s ease;
}

.logout-btn:hover:not(:disabled) {
  background: #c0392b;
  transform: translateY(-1px);
}

.logout-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.profile-content {
  display: grid;
  gap: 2rem;
}

.profile-card {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 2rem;
  flex-wrap: wrap;
}

.profile-avatar {
  flex-shrink: 0;
}

.avatar-placeholder {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 2rem;
  font-weight: bold;
}

.profile-info {
  flex: 1;
  min-width: 250px;
}

.profile-info h2 {
  color: #2c3e50;
  margin: 0 0 1rem 0;
  font-size: 1.5rem;
}

.info-item {
  display: flex;
  margin-bottom: 0.75rem;
  align-items: center;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.label {
  font-weight: 600;
  color: #7f8c8d;
  min-width: 80px;
}

.value {
  color: #2c3e50;
}

.role-badge {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
}

.role-badge.admin {
  background: #e74c3c;
  color: white;
}

.role-badge.customer {
  background: #3498db;
  color: white;
}

.profile-actions {
  display: grid;
  gap: 1.5rem;
}

.action-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.action-card h3 {
  color: #2c3e50;
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
}

.action-buttons {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.action-btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s ease;
  text-decoration: none;
  display: inline-block;
}

.action-btn.primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.action-btn.primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.action-btn.secondary {
  background: #f8f9fa;
  color: #2c3e50;
  border: 1px solid #e1e8ed;
}

.action-btn.secondary:hover {
  background: #e9ecef;
  transform: translateY(-1px);
}

.action-btn.admin-btn {
  background: #e74c3c;
  color: white;
}

.action-btn.admin-btn:hover {
  background: #c0392b;
  transform: translateY(-1px);
}

.loading, .error {
  text-align: center;
  padding: 3rem;
  color: #7f8c8d;
}

.error {
  color: #e74c3c;
}

@media (max-width: 768px) {
  .profile-header {
    flex-direction: column;
    align-items: stretch;
  }
  
  .profile-card {
    flex-direction: column;
    text-align: center;
  }
  
  .info-item {
    justify-content: center;
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .action-btn {
    width: 100%;
  }
}
</style>
