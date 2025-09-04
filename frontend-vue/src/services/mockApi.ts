// Mock API service with database data
export interface Product {
  id: number
  name: string
  description: string
  price: number
  category: string
  image: string
  stock: number
}

export interface User {
  id: number
  username: string
  email: string
  name: string
  role: string
}

// Mock products from database
const mockProducts: Product[] = [
  {
    id: 1,
    name: "Vue.js T-Shirt",
    description: "Comfortable cotton t-shirt with Vue.js logo",
    price: 25.99,
    category: "clothing",
    image: "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=300&fit=crop&bg=white",
    stock: 50
  },
  {
    id: 2,
    name: "JavaScript Book",
    description: "Comprehensive guide to modern JavaScript",
    price: 39.99,
    category: "books",
    image: "https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400&h=300&fit=crop&bg=white",
    stock: 25
  },
  {
    id: 3,
    name: "Wireless Headphones",
    description: "High-quality wireless headphones with noise cancellation",
    price: 89.99,
    category: "electronics",
    image: "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400&h=300&fit=crop&bg=white",
    stock: 15
  },
  {
    id: 4,
    name: "Vue Hoodie",
    description: "Warm and cozy hoodie perfect for Vue developers",
    price: 45.99,
    category: "clothing",
    image: "https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=400&h=300&fit=crop&bg=white",
    stock: 30
  },
  {
    id: 5,
    name: "Laptop Stand",
    description: "Ergonomic laptop stand for better posture",
    price: 29.99,
    category: "electronics",
    image: "https://images.unsplash.com/photo-1586953208448-b95a79798f07?w=400&h=300&fit=crop&bg=white",
    stock: 20
  },
  {
    id: 6,
    name: "Vue.js Guide",
    description: "Complete Vue.js 3 tutorial and reference",
    price: 19.99,
    category: "books",
    image: "https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=400&h=300&fit=crop&bg=white",
    stock: 40
  },
  {
    id: 7,
    name: "Mechanical Keyboard",
    description: "Premium mechanical keyboard with RGB backlighting",
    price: 129.99,
    category: "electronics",
    image: "https://images.unsplash.com/photo-1541140532154-b024d705b90a?w=400&h=300&fit=crop&bg=white",
    stock: 10
  },
  {
    id: 8,
    name: "Developer Mug",
    description: "Ceramic coffee mug perfect for coding sessions",
    price: 12.99,
    category: "clothing",
    image: "https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400&h=300&fit=crop&bg=white",
    stock: 100
  },
  {
    id: 9,
    name: "React vs Vue Book",
    description: "In-depth comparison of modern JavaScript frameworks",
    price: 34.99,
    category: "books",
    image: "https://images.unsplash.com/photo-1589829085413-56de8ae18c73?w=400&h=300&fit=crop&bg=white",
    stock: 35
  }
]

// Mock users from database
const mockUsers: User[] = [
  { id: 1, username: "admin", email: "admin@vueshop.com", name: "Admin User", role: "admin" },
  { id: 2, username: "john", email: "john@example.com", name: "John Doe", role: "customer" },
  { id: 3, username: "jane", email: "jane@example.com", name: "Jane Smith", role: "customer" }
]

// Mock API functions
export const mockApi = {
  // Products
  getProducts: (params?: { category?: string; search?: string; sort?: string }) => {
    return new Promise<{ success: boolean; data: Product[]; total: number }>((resolve) => {
      setTimeout(() => {
        let filtered = [...mockProducts]
        
        if (params?.category && params.category !== 'all') {
          filtered = filtered.filter(p => p.category === params.category)
        }
        
        if (params?.search) {
          const search = params.search.toLowerCase()
          filtered = filtered.filter(p => 
            p.name.toLowerCase().includes(search) || 
            p.description.toLowerCase().includes(search)
          )
        }
        
        if (params?.sort) {
          switch (params.sort) {
            case 'price-asc':
              filtered.sort((a, b) => a.price - b.price)
              break
            case 'price-desc':
              filtered.sort((a, b) => b.price - a.price)
              break
            case 'name-asc':
              filtered.sort((a, b) => a.name.localeCompare(b.name))
              break
            case 'name-desc':
              filtered.sort((a, b) => b.name.localeCompare(a.name))
              break
          }
        }
        
        resolve({ success: true, data: filtered, total: filtered.length })
      }, 300)
    })
  },

  getProduct: (id: number) => {
    return new Promise<{ success: boolean; data?: Product; error?: string }>((resolve) => {
      setTimeout(() => {
        const product = mockProducts.find(p => p.id === id)
        if (product) {
          resolve({ success: true, data: product })
        } else {
          resolve({ success: false, error: 'Product not found' })
        }
      }, 200)
    })
  },

  getCategories: () => {
    return new Promise<{ success: boolean; data: string[] }>((resolve) => {
      setTimeout(() => {
        const categories = [...new Set(mockProducts.map(p => p.category))].sort()
        resolve({ success: true, data: categories })
      }, 200)
    })
  },

  // Auth
  login: (username: string, password: string) => {
    return new Promise<{ success: boolean; data?: { user: User; token: string }; error?: string }>((resolve) => {
      setTimeout(() => {
        const user = mockUsers.find(u => u.username === username)
        if (user && ((username === 'admin' && password === 'admin123') || password === 'password123' || password === 'password456')) {
          resolve({
            success: true,
            data: {
              user,
              token: 'mock-token-' + Date.now()
            }
          })
        } else {
          resolve({ success: false, error: 'Invalid credentials' })
        }
      }, 500)
    })
  },

  // Orders
  createOrder: (orderData: any) => {
    return new Promise<{ success: boolean; data: any }>((resolve) => {
      setTimeout(() => {
        const orderId = 'VUE-' + Math.floor(Math.random() * 10000)
        resolve({
          success: true,
          data: {
            orderId,
            orderNumber: orderId,
            total: orderData.total,
            status: 'pending'
          }
        })
      }, 800)
    })
  }
}