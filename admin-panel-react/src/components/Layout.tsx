import { Outlet, Link, useLocation } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import {
  FaHome,
  FaBook,
  FaUsers,
  FaTags,
  FaSignOutAlt,
  FaUser,
  FaBars,
  FaTimes,
} from 'react-icons/fa';
import { useState } from 'react';

export default function Layout() {
  const { user, logout } = useAuth();
  const location = useLocation();
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const navigation = [
    { name: 'Dashboard', href: '/dashboard', icon: FaHome },
    { name: 'Courses', href: '/courses', icon: FaBook },
    { name: 'Users', href: '/users', icon: FaUsers },
    { name: 'Categories', href: '/categories', icon: FaTags },
  ];

  const isActive = (path: string) => location.pathname === path;

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Sidebar for desktop */}
      <aside className="fixed inset-y-0 left-0 z-50 w-64 bg-secondary text-white transform transition-transform duration-300 ease-in-out lg:translate-x-0">
        <div className="flex flex-col h-full">
          {/* Logo */}
          <div className="flex items-center justify-between px-6 py-4 border-b border-gray-700">
            <h1 className="text-xl font-bold">eCourses Admin</h1>
            <button
              onClick={() => setSidebarOpen(false)}
              className="lg:hidden text-white"
            >
              <FaTimes />
            </button>
          </div>

          {/* Navigation */}
          <nav className="flex-1 px-4 py-6 space-y-2">
            {navigation.map((item) => {
              const Icon = item.icon;
              return (
                <Link
                  key={item.name}
                  to={item.href}
                  className={`flex items-center space-x-3 px-4 py-3 rounded-lg transition ${
                    isActive(item.href)
                      ? 'bg-primary text-white'
                      : 'text-gray-300 hover:bg-gray-700 hover:text-white'
                  }`}
                >
                  <Icon />
                  <span>{item.name}</span>
                </Link>
              );
            })}
          </nav>

          {/* User info */}
          <div className="border-t border-gray-700 p-4">
            <div className="flex items-center space-x-3 mb-3">
              <div className="w-10 h-10 rounded-full bg-primary flex items-center justify-center">
                <FaUser />
              </div>
              <div className="flex-1 min-w-0">
                <p className="text-sm font-medium truncate">{user?.full_name}</p>
                <p className="text-xs text-gray-400 truncate">{user?.email}</p>
              </div>
            </div>
            <button
              onClick={logout}
              className="w-full flex items-center justify-center space-x-2 px-4 py-2 bg-danger rounded-lg hover:bg-red-600 transition"
            >
              <FaSignOutAlt />
              <span>Logout</span>
            </button>
          </div>
        </div>
      </aside>

      {/* Mobile sidebar backdrop */}
      {sidebarOpen && (
        <div
          className="fixed inset-0 z-40 bg-black bg-opacity-50 lg:hidden"
          onClick={() => setSidebarOpen(false)}
        />
      )}

      {/* Main content */}
      <div className="lg:ml-64">
        {/* Mobile header */}
        <header className="lg:hidden bg-white shadow-sm">
          <div className="flex items-center justify-between px-4 py-3">
            <button
              onClick={() => setSidebarOpen(true)}
              className="text-gray-600"
            >
              <FaBars className="text-xl" />
            </button>
            <h1 className="text-lg font-semibold">eCourses Admin</h1>
            <div className="w-6" /> {/* Spacer */}
          </div>
        </header>

        {/* Page content */}
        <main className="p-6">
          <Outlet />
        </main>
      </div>
    </div>
  );
}
