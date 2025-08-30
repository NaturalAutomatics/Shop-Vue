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

export interface CartItem extends Product {
  quantity: number
}

export interface Order {
  id: string
  orderNumber: string
  customer: {
    name: string
    email: string
    address: string
  }
  items: CartItem[]
  subtotal: number
  shipping: number
  tax: number
  total: number
  status: string
  createdAt: Date
  updatedAt: Date
}

export interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
  error?: string
}