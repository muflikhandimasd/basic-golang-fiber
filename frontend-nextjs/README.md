# eCourses Frontend - Next.js

A modern, responsive frontend for the eCourses e-learning platform built with Next.js 14, TypeScript, and Tailwind CSS.

## Features

- **User Authentication**: Login and registration with JWT tokens
- **Role-Based Access**: Different interfaces for students, instructors, and admins
- **Course Browsing**: Browse and search all available courses
- **Course Details**: View detailed course information and lessons
- **Student Dashboard**: Track enrolled courses and progress
- **Instructor Dashboard**: Create and manage courses
- **Responsive Design**: Mobile-friendly interface
- **Type Safety**: Full TypeScript support

## Tech Stack

- **Framework**: Next.js 14 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **HTTP Client**: Axios
- **Icons**: React Icons
- **State Management**: React Context API

## Prerequisites

- Node.js 18+ (or compatible version)
- npm or yarn
- Running backend API (see main README)

## Getting Started

### Installation

1. Navigate to the frontend directory:
```bash
cd frontend-nextjs
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment variables:
   - Copy `.env.local` or create it with:
```bash
NEXT_PUBLIC_API_URL=http://localhost:9000/api
```

4. Run the development server:
```bash
npm run dev
```

5. Open [http://localhost:3000](http://localhost:3000) in your browser

## Project Structure

```
frontend-nextjs/
├── src/
│   ├── app/                  # Next.js 14 App Router pages
│   │   ├── page.tsx          # Home page
│   │   ├── login/            # Login page
│   │   ├── register/         # Registration page
│   │   ├── courses/          # Course listing and details
│   │   ├── dashboard/        # Student dashboard
│   │   └── instructor/       # Instructor dashboard
│   ├── components/           # Reusable components
│   │   ├── Navbar.tsx
│   │   └── CourseCard.tsx
│   ├── contexts/             # React Context providers
│   │   └── AuthContext.tsx
│   ├── services/             # API service layer
│   │   └── api.ts
│   └── types/                # TypeScript type definitions
│       └── index.ts
├── public/                   # Static assets
├── package.json
├── tsconfig.json
├── tailwind.config.js
└── next.config.js
```

## Available Scripts

- `npm run dev` - Start development server on port 3000
- `npm run build` - Build for production
- `npm start` - Start production server
- `npm run lint` - Run ESLint

## Features by Role

### Students
- Browse and search courses
- View course details and lessons
- Enroll in courses
- Track learning progress
- View enrolled courses dashboard

### Instructors
- All student features
- Create new courses
- Manage existing courses
- View course statistics
- Track student enrollments

### Admins
- All instructor features
- Manage categories
- Full platform access

## Key Pages

### Home (`/`)
Landing page with platform overview and call-to-action

### Login (`/login`)
User authentication page with demo credentials

### Register (`/register`)
New user registration with role selection

### Courses (`/courses`)
Browse all available courses with filtering

### Course Details (`/courses/[id]`)
Detailed course view with enrollment option

### Dashboard (`/dashboard`)
Student dashboard showing enrolled courses and progress

### Instructor Courses (`/instructor/courses`)
Instructor dashboard for managing courses

### Create Course (`/instructor/courses/create`)
Form to create new courses

## API Integration

The frontend communicates with the Golang backend API at `http://localhost:9000/api`.

### Authentication Flow
1. User logs in via `/api/auth/login`
2. JWT token is stored in localStorage
3. Token is automatically included in API requests
4. User data is cached in AuthContext

### API Service
The `apiService` class (`src/services/api.ts`) provides methods for all API endpoints:

- **Auth**: `login()`, `register()`, `getProfile()`, `logout()`
- **Courses**: `getCourses()`, `getCourse()`, `createCourse()`, etc.
- **Lessons**: `getCourseLessons()`, `createLesson()`, etc.
- **Enrollments**: `enrollCourse()`, `getMyEnrollments()`, etc.
- **Categories**: `getCategories()`, `createCategory()`, etc.

## Authentication

### Login Credentials
Demo account:
- Email: `admin@ecourses.com`
- Password: `admin123`

### Token Management
- JWT tokens are stored in localStorage
- Tokens expire after 24 hours
- Automatic token inclusion in API requests
- Logout clears all stored data

## Styling

The application uses Tailwind CSS with custom color scheme:

```javascript
colors: {
  primary: '#3498db',    // Blue
  secondary: '#2c3e50',  // Dark blue/gray
  success: '#2ecc71',    // Green
  danger: '#e74c3c',     // Red
  warning: '#f39c12',    // Orange
}
```

## Environment Variables

```bash
NEXT_PUBLIC_API_URL=http://localhost:9000/api
```

Note: Variables must be prefixed with `NEXT_PUBLIC_` to be accessible in the browser.

## Building for Production

1. Build the application:
```bash
npm run build
```

2. Start the production server:
```bash
npm start
```

The production build will be optimized and ready for deployment.

## Deployment

This Next.js application can be deployed to:

- Vercel (recommended)
- Netlify
- Docker
- Any Node.js hosting platform

Ensure the `NEXT_PUBLIC_API_URL` environment variable points to your production API.

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Contributing

1. Make changes in a feature branch
2. Test thoroughly
3. Submit a pull request

## License

MIT License

## Support

For issues or questions, please refer to the main project repository.
