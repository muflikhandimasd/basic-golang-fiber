'use client';

import Link from 'next/link';
import { useAuth } from '@/contexts/AuthContext';
import { FaGraduationCap, FaUser, FaSignOutAlt, FaBook, FaChalkboardTeacher } from 'react-icons/fa';

export default function Navbar() {
  const { user, isAuthenticated, logout, isInstructor, isAdmin } = useAuth();

  return (
    <nav className="bg-secondary text-white shadow-lg">
      <div className="container mx-auto px-4">
        <div className="flex justify-between items-center h-16">
          <Link href="/" className="flex items-center space-x-2 text-xl font-bold">
            <FaGraduationCap className="text-2xl" />
            <span>eCourses</span>
          </Link>

          <div className="flex items-center space-x-6">
            <Link href="/courses" className="hover:text-primary transition">
              Courses
            </Link>

            {isAuthenticated ? (
              <>
                <Link href="/dashboard" className="hover:text-primary transition">
                  Dashboard
                </Link>

                {(isInstructor || isAdmin) && (
                  <Link href="/instructor/courses" className="hover:text-primary transition flex items-center space-x-1">
                    <FaChalkboardTeacher />
                    <span>My Courses</span>
                  </Link>
                )}

                <div className="flex items-center space-x-4">
                  <div className="flex items-center space-x-2">
                    <FaUser className="text-primary" />
                    <span className="text-sm">{user?.full_name}</span>
                  </div>
                  <button
                    onClick={logout}
                    className="flex items-center space-x-1 bg-danger px-4 py-2 rounded hover:bg-red-600 transition"
                  >
                    <FaSignOutAlt />
                    <span>Logout</span>
                  </button>
                </div>
              </>
            ) : (
              <>
                <Link
                  href="/login"
                  className="hover:text-primary transition"
                >
                  Login
                </Link>
                <Link
                  href="/register"
                  className="bg-primary px-4 py-2 rounded hover:bg-blue-600 transition"
                >
                  Sign Up
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}
