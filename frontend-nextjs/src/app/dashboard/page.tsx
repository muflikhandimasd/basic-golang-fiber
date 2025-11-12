'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Navbar from '@/components/Navbar';
import { useAuth } from '@/contexts/AuthContext';
import { apiService } from '@/services/api';
import { EnrollmentWithDetails } from '@/types';
import { FaBook, FaSpinner, FaClock, FaCheckCircle } from 'react-icons/fa';
import Link from 'next/link';

export default function DashboardPage() {
  const { isAuthenticated, loading: authLoading, user } = useAuth();
  const router = useRouter();
  const [enrollments, setEnrollments] = useState<EnrollmentWithDetails[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!authLoading && !isAuthenticated) {
      router.push('/login');
    } else if (isAuthenticated) {
      loadEnrollments();
    }
  }, [isAuthenticated, authLoading]);

  const loadEnrollments = async () => {
    try {
      const data = await apiService.getMyEnrollments();
      setEnrollments(data);
    } catch (error) {
      console.error('Failed to load enrollments:', error);
    } finally {
      setLoading(false);
    }
  };

  if (authLoading || loading) {
    return (
      <div className="min-h-screen bg-gray-50">
        <Navbar />
        <div className="flex justify-center items-center py-20">
          <FaSpinner className="animate-spin text-4xl text-primary" />
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />

      <div className="container mx-auto px-4 py-8">
        <div className="mb-8">
          <h1 className="text-4xl font-bold text-gray-800 mb-2">My Dashboard</h1>
          <p className="text-gray-600">Welcome back, {user?.full_name}!</p>
        </div>

        <div className="grid md:grid-cols-3 gap-6 mb-8">
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-600 text-sm">Enrolled Courses</p>
                <p className="text-3xl font-bold text-primary">{enrollments.length}</p>
              </div>
              <FaBook className="text-4xl text-primary opacity-20" />
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-600 text-sm">In Progress</p>
                <p className="text-3xl font-bold text-warning">
                  {enrollments.filter((e) => e.progress > 0 && e.progress < 100).length}
                </p>
              </div>
              <FaClock className="text-4xl text-warning opacity-20" />
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-600 text-sm">Completed</p>
                <p className="text-3xl font-bold text-success">
                  {enrollments.filter((e) => e.progress === 100).length}
                </p>
              </div>
              <FaCheckCircle className="text-4xl text-success opacity-20" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-2xl font-bold mb-6">My Courses</h2>

          {enrollments.length > 0 ? (
            <div className="space-y-4">
              {enrollments.map((enrollment) => (
                <div
                  key={enrollment.id}
                  className="border border-gray-200 rounded-lg p-4 hover:shadow-md transition"
                >
                  <div className="flex justify-between items-start mb-3">
                    <div>
                      <h3 className="text-lg font-semibold text-gray-800">{enrollment.course_name}</h3>
                      <p className="text-sm text-gray-600">
                        Enrolled on {new Date(enrollment.enrolled_at).toLocaleDateString()}
                      </p>
                    </div>
                    <Link
                      href={`/courses/${enrollment.course_id}`}
                      className="bg-primary text-white px-4 py-2 rounded hover:bg-blue-600 transition text-sm"
                    >
                      View Course
                    </Link>
                  </div>

                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div
                      className="bg-primary h-2 rounded-full transition-all"
                      style={{ width: `${enrollment.progress}%` }}
                    />
                  </div>
                  <p className="text-sm text-gray-600 mt-2">Progress: {enrollment.progress}%</p>
                </div>
              ))}
            </div>
          ) : (
            <div className="text-center py-12">
              <FaBook className="mx-auto text-6xl text-gray-300 mb-4" />
              <p className="text-gray-600 mb-4">You haven't enrolled in any courses yet.</p>
              <Link
                href="/courses"
                className="inline-block bg-primary text-white px-6 py-3 rounded-lg hover:bg-blue-600 transition"
              >
                Browse Courses
              </Link>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
