# Vue Shop - Learning Vue.js after Angular

Welcome to your Vue.js learning journey! This project is designed to help you transition from Angular to Vue.js by building a real e-commerce application.

## 🎯 What You'll Learn

This series of lessons covers:

1. **Vue 3 Composition API** - The modern way to write Vue components
2. **Pinia State Management** - Vue's official store library
3. **Vue Router 4** - Client-side routing
4. **Component Architecture** - Building reusable components
5. **Reactive Data** - Understanding Vue's reactivity system
6. **Template Syntax** - Vue's template directives and expressions

## 🚀 Getting Started

### Prerequisites
- Node.js (version 16 or higher)
- npm or yarn package manager

### Installation

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Start the development server:**
   ```bash
   npm run dev
   ```

3. **Open your browser:**
   Navigate to `http://localhost:3000`

## 📚 Lesson Structure

### Lesson 1: Project Setup and Basics ✅
- **Files:** `package.json`, `vite.config.js`, `index.html`, `src/main.js`
- **Concepts:** Project structure, Vue 3 setup, Vite bundler
- **Angular Comparison:** Similar to Angular CLI setup

### Lesson 2: Component Architecture ✅
- **Files:** `src/App.vue`, `src/views/Home.vue`
- **Concepts:** Single File Components (SFC), template syntax, Composition API
- **Angular Comparison:** 
  - `App.vue` ≈ `app.component.ts` + `app.component.html` + `app.component.css`
  - `setup()` function ≈ `ngOnInit()` + dependency injection

### Lesson 3: State Management with Pinia ✅
- **Files:** `src/stores/cart.js`
- **Concepts:** Pinia stores, reactive state, computed properties
- **Angular Comparison:** 
  - Pinia stores ≈ Angular services + NgRx
  - `ref()` ≈ `BehaviorSubject`
  - `computed()` ≈ Angular getters

### Lesson 4: Routing ✅
- **Files:** `src/router/index.js`
- **Concepts:** Vue Router, route configuration, navigation
- **Angular Comparison:** 
  - Vue Router ≈ Angular Router
  - `<router-view />` ≈ `<router-outlet>`

### Lesson 5: Product Management ✅
- **Files:** `src/views/Products.vue`
- **Concepts:** Lists, filtering, computed properties, event handling
- **Angular Comparison:** 
  - `v-for` ≈ `*ngFor`
  - `v-model` ≈ `[(ngModel)]`
  - `@click` ≈ `(click)`

### Lesson 6: Shopping Cart ✅
- **Files:** `src/views/Cart.vue`
- **Concepts:** Complex state management, conditional rendering, form handling
- **Angular Comparison:** 
  - `v-if` ≈ `*ngIf`
  - `v-else` ≈ `*ngIf` with else template

## 🔄 Key Differences: Angular vs Vue

| Concept | Angular | Vue |
|---------|---------|-----|
| **Component Definition** | Class with decorators | Object with `setup()` function |
| **Template Syntax** | `*ngIf`, `*ngFor`, `[(ngModel)]` | `v-if`, `v-for`, `v-model` |
| **State Management** | Services + NgRx | Pinia stores |
| **Reactivity** | Change detection + Zone.js | Automatic reactivity |
| **File Structure** | Separate `.ts`, `.html`, `.css` | Single `.vue` files |
| **Dependency Injection** | Constructor injection | Import statements |

## 🛠️ Development Commands

```bash
# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Lint code
npm run lint
```

## 📁 Project Structure

```
vue-shop/
├── public/                 # Static assets
├── src/
│   ├── components/         # Reusable components
│   ├── views/             # Page components
│   │   ├── Home.vue       # Home page
│   │   ├── Products.vue   # Products listing
│   │   └── Cart.vue       # Shopping cart
│   ├── stores/            # Pinia stores
│   │   └── cart.js        # Cart state management
│   ├── router/            # Vue Router configuration
│   │   └── index.js       # Route definitions
│   ├── App.vue            # Root component
│   ├── main.js            # Application entry point
│   └── style.css          # Global styles
├── index.html             # HTML template
├── package.json           # Dependencies and scripts
├── vite.config.js         # Vite configuration
└── README.md              # This file
```

## 🎨 Features Implemented

- ✅ **Responsive Design** - Mobile-first approach
- ✅ **Product Catalog** - Browse and search products
- ✅ **Shopping Cart** - Add, remove, and update quantities
- ✅ **State Management** - Centralized cart state
- ✅ **Routing** - Navigation between pages
- ✅ **Modern UI** - Beautiful gradients and animations

## 🔮 Next Steps

After completing this project, you can:

1. **Add TypeScript** - Convert to `.vue` files with TypeScript
2. **Implement Authentication** - Add user login/registration
3. **Add Backend Integration** - Connect to a real API
4. **Add Testing** - Unit tests with Vitest
5. **Add PWA Features** - Service workers and offline support

## 🤝 Contributing

This is a learning project, but feel free to:
- Report bugs
- Suggest improvements
- Add new features
- Improve documentation

## 📖 Additional Resources

- [Vue 3 Documentation](https://vuejs.org/)
- [Pinia Documentation](https://pinia.vuejs.org/)
- [Vue Router Documentation](https://router.vuejs.org/)
- [Vite Documentation](https://vitejs.dev/)

---

Happy coding! 🚀
