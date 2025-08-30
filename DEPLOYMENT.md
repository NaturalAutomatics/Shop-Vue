# Firebase Deployment Setup

## Prerequisites
1. Firebase account with a project created
2. GitHub repository with the code

## Setup Steps

### 1. Firebase Project Setup
```bash
# Install Firebase CLI
npm install -g firebase-tools

# Login to Firebase
firebase login

# Initialize Firebase in your project
firebase init
```

### 2. GitHub Secrets Configuration
Add these secrets to your GitHub repository (Settings > Secrets and variables > Actions):

- `FIREBASE_TOKEN`: Get token with `firebase login:ci`
- `FIREBASE_PROJECT_ID`: Your Firebase project ID

### 3. Update Firebase Configuration
Edit `.firebaserc` and replace `your-firebase-project-id` with your actual Firebase project ID.

### 4. Deploy
Push to main branch to trigger automatic deployment via GitHub Actions.

## Manual Deployment
```bash
# Build frontend
cd frontend-vue
npm run build

# Build functions
cd ../functions
npm run build

# Deploy to Firebase
cd ..
firebase deploy
```

## Project Structure
- Frontend: Deployed to Firebase Hosting
- Backend API: Deployed as Firebase Functions
- Database: Use Firebase Firestore or external PostgreSQL

## Environment Variables
For production, update API endpoints in frontend to use Firebase Functions URLs.