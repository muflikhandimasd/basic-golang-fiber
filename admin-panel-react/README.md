# eCourses Admin Panel - React + Vite

A modern, responsive admin panel for the eCourses platform built with React, Vite, TypeScript, and Tailwind CSS.

## Features

- **Admin Authentication**: Secure login restricted to admin users only
- **Dashboard**: Overview with key statistics and metrics
- **Course Management**: View, search, filter, and delete courses
- **Category Management**: Full CRUD operations for categories
- **User Management**: View all platform users
- **Responsive Design**: Works seamlessly on desktop and mobile devices
- **Modern UI**: Clean interface with Tailwind CSS
- **Type Safety**: Full TypeScript support
- **Fast Development**: Powered by Vite for instant HMR

## Tech Stack

- **Framework**: React 18
- **Build Tool**: Vite 5
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **Routing**: React Router v6
- **HTTP Client**: Axios
- **Icons**: React Icons
- **Charts**: Recharts

## Prerequisites

- Node.js 18+ (or compatible version)
- npm or yarn
- Running backend API at http://localhost:9000

## Getting Started

### Installation

1. Navigate to the admin panel directory:
```bash
cd admin-panel-react
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment variables (optional):
   - The `.env` file is already configured with:
```bash
VITE_API_URL=http://localhost:9000/api
```

4. Run the development server:
```bash
npm run dev
```

5. Open [http://localhost:3001](http://localhost:3001) in your browser

## Project Structure

```
admin-panel-react/
├── src/
│   ├── components/          # Reusable components
│   │   └── Layout.tsx       # Main layout with sidebar
│   ├── contexts/            # React Context providers
│   │   └── AuthContext.tsx  # Authentication context
│   ├── pages/               # Page components
│   │   ├── Login.tsx        # Login page
│   │   ├── Dashboard.tsx    # Dashboard with stats
│   │   ├── Courses.tsx      # Courses management
│   │   ├── CourseEdit.tsx   # Edit course (placeholder)
│   │   ├── Users.tsx        # Users management
│   │   └── Categories.tsx   # Categories CRUD
│   ├── services/            # API service layer
│   │   └── api.ts           # Admin API service
│   ├── types/               # TypeScript definitions
│   │   └── index.ts
│   ├── App.tsx              # Main app component
│   ├── main.tsx             # Entry point
│   └── index.css            # Global styles
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
└── tailwind.config.js
```

## Available Scripts

- `npm run dev` - Start development server on port 3001
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint

## Admin Login

**Default Credentials:**
- **Email**: admin@ecourses.com
- **Password**: admin123

**Important**: Only users with `admin` role can access this panel.

## Features by Page

### Dashboard (`/dashboard`)
- Total courses count
- Published courses count
- Draft courses count
- Total enrollments count
- Recent courses table

### Courses Management (`/courses`)
- View all courses in a table
- Search courses by title or instructor
- Filter by status (all, published, draft, archived)
- Delete courses
- Quick stats (ID, title, instructor, category, status, students, lessons, price)

### Categories Management (`/categories`)
- View all categories
- Add new category
- Edit existing category
- Delete category
- Modal-based form for create/edit

### Users Management (`/users`)
- Placeholder page
- Links to Golang admin panel for full user management

## API Integration

The admin panel communicates with the Golang backend API:

### Authentication
- `POST /api/auth/login` - Admin login
- `GET /api/auth/profile` - Get current admin profile

### Courses
- `GET /api/courses` - List all courses
- `GET /api/courses/:id` - Get course details
- `DELETE /api/courses/:id` - Delete course

### Categories
- `GET /api/categories` - List all categories
- `POST /api/categories` - Create category
- `PUT /api/categories/:id` - Update category
- `DELETE /api/categories/:id` - Delete category

### Token Management
- Tokens are stored in `localStorage` with key `admin_token`
- User data stored with key `admin_user`
- Automatic token injection in all API requests
- Logout clears both token and user data

## Styling

Custom color scheme defined in `tailwind.config.js`:

```javascript
colors: {
  primary: '#3498db',    // Blue
  secondary: '#2c3e50',  // Dark gray/blue
  success: '#2ecc71',    // Green
  danger: '#e74c3c',     // Red
  warning: '#f39c12',    // Orange
}
```

## Layout Structure

The admin panel uses a sidebar layout:

- **Desktop**: Fixed sidebar on the left, content area on the right
- **Mobile**: Hamburger menu, sliding sidebar
- **Sidebar includes**:
  - Navigation links
  - User profile info
  - Logout button

## Protected Routes

All routes except `/login` are protected:

- Redirects to login if not authenticated
- Checks for admin role on login
- Automatically validates token on page load

## Building for Production

1. Build the application:
```bash
npm run build
```

2. Preview the build:
```bash
npm run preview
```

3. Deploy the `dist` folder to your hosting platform

## Deployment

This React application can be deployed to:

- Vercel
- Netlify
- GitHub Pages
- Any static hosting service

**Environment Variable**: Ensure `VITE_API_URL` points to your production API.

## Development Notes

- Hot Module Replacement (HMR) enabled for fast development
- TypeScript strict mode enabled
- ESLint configured for code quality
- Tailwind CSS for utility-first styling

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Related Projects

- **Backend API**: `../` (Golang + Fiber + SQLite)
- **Student/Instructor Frontend**: `../frontend-nextjs` (Next.js)
- **Golang Admin Panel**: http://localhost:9000/admin

## Troubleshooting

### Cannot connect to API
- Ensure the backend is running on http://localhost:9000
- Check `.env` file for correct API_URL

### Login fails with "Admin access required"
- Only users with `role: 'admin'` can access this panel
- Use the default admin credentials or create an admin user via the backend

### Build errors
- Delete `node_modules` and `package-lock.json`
- Run `npm install` again
- Clear Vite cache: `rm -rf node_modules/.vite`

## Contributing

1. Make changes in a feature branch
2. Test thoroughly on development server
3. Build and verify production build
4. Submit a pull request

## License

MIT License

## Support

For issues or questions, please refer to the main project repository.
