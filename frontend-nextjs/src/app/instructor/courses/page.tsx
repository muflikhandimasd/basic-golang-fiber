'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Navbar from '@/components/Navbar';
import { useAuth } from '@/contexts/AuthContext';
import { apiService } from '@/services/api';
import { CourseWithDetails } from '@/types';
import { FaBook, FaSpinner, FaPlus, FaEdit, FaTrash, FaUsers } from 'react-icons/fa';
import Link from 'next/link';

export default function InstructorCoursesPage() {
  const { isAuthenticated, loading: authLoading, isInstructor, isAdmin } = useAuth();
  const router = useRouter();
  const [courses, setCourses] = useState<CourseWithDetails[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!authLoading && (!isAuthenticated || (!isInstructor && !isAdmin))) {
      router.push('/login');
    } else if (isAuthenticated && (isInstructor || isAdmin)) {
      loadCourses();
    }
  }, [isAuthenticated, authLoading, isInstructor, isAdmin]);

  const loadCourses = async () => {
    try {
      const data = await apiService.getMyCourses();
      setCourses(data);
    } catch (error) {
      console.error('Failed to load courses:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (courseId: number) => {
    if (!confirm('Are you sure you want to delete this course?')) {
      return;
    }

    try {
      await apiService.deleteCourse(courseId);
      setCourses(courses.filter((c) => c.id !== courseId));
      alert('Course deleted successfully');
    } catch (error) {
      alert('Failed to delete course');
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
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-4xl font-bold text-gray-800">My Courses</h1>
          <Link
            href="/instructor/courses/create"
            className="bg-primary text-white px-6 py-3 rounded-lg hover:bg-blue-600 transition flex items-center space-x-2"
          >
            <FaPlus />
            <span>Create New Course</span>
          </Link>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-600 text-sm">Total Courses</p>
                <p className="text-3xl font-bold text-primary">{courses.length}</p>
              </div>
              <FaBook className="text-4xl text-primary opacity-20" />
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-600 text-sm">Published</p>
                <p className="text-3xl font-bold text-success">
                  {courses.filter((c) => c.status === 'published').length}
                </p>
              </div>
              <FaBook className="text-4xl text-success opacity-20" />
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-600 text-sm">Total Students</p>
                <p className="text-3xl font-bold text-warning">
                  {courses.reduce((sum, c) => sum + c.enrollment_count, 0)}
                </p>
              </div>
              <FaUsers className="text-4xl text-warning opacity-20" />
            </div>
          </div>
        </div>

        {courses.length > 0 ? (
          <div className="bg-white rounded-lg shadow-md overflow-hidden">
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Course
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Status
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Students
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Lessons
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Price
                  </th>
                  <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {courses.map((course) => (
                  <tr key={course.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4">
                      <Link
                        href={`/courses/${course.id}`}
                        className="font-medium text-gray-900 hover:text-primary"
                      >
                        {course.title}
                      </Link>
                    </td>
                    <td className="px-6 py-4">
                      <span
                        className={`px-2 py-1 rounded text-xs font-semibold ${
                          course.status === 'published'
                            ? 'bg-green-100 text-green-800'
                            : course.status === 'draft'
                            ? 'bg-yellow-100 text-yellow-800'
                            : 'bg-gray-100 text-gray-800'
                        }`}
                      >
                        {course.status}
                      </span>
                    </td>
                    <td className="px-6 py-4 text-gray-600">{course.enrollment_count}</td>
                    <td className="px-6 py-4 text-gray-600">{course.lesson_count}</td>
                    <td className="px-6 py-4 text-gray-600">
                      ${course.price === 0 ? 'Free' : course.price.toFixed(2)}
                    </td>
                    <td className="px-6 py-4 text-right">
                      <div className="flex justify-end space-x-2">
                        <button
                          onClick={() => router.push(`/instructor/courses/${course.id}/edit`)}
                          className="text-primary hover:text-blue-700 p-2"
                          title="Edit"
                        >
                          <FaEdit />
                        </button>
                        <button
                          onClick={() => handleDelete(course.id)}
                          className="text-danger hover:text-red-700 p-2"
                          title="Delete"
                        >
                          <FaTrash />
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ) : (
          <div className="bg-white rounded-lg shadow-md p-12 text-center">
            <FaBook className="mx-auto text-6xl text-gray-300 mb-4" />
            <p className="text-gray-600 mb-4">You haven't created any courses yet.</p>
            <Link
              href="/instructor/courses/create"
              className="inline-block bg-primary text-white px-6 py-3 rounded-lg hover:bg-blue-600 transition"
            >
              Create Your First Course
            </Link>
          </div>
        )}
      </div>
    </div>
  );
}
