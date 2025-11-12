import { useEffect, useState } from 'react';
import { adminApi } from '../services/api';
import { CourseWithDetails } from '../types';
import { FaBook, FaUsers, FaChartLine, FaCheckCircle, FaSpinner } from 'react-icons/fa';
import { Link } from 'react-router-dom';

export default function Dashboard() {
  const [courses, setCourses] = useState<CourseWithDetails[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      const coursesData = await adminApi.getCourses({ limit: 1000 });
      setCourses(coursesData);
    } catch (error) {
      console.error('Failed to load data:', error);
    } finally {
      setLoading(false);
    }
  };

  const stats = {
    totalCourses: courses.length,
    publishedCourses: courses.filter((c) => c.status === 'published').length,
    draftCourses: courses.filter((c) => c.status === 'draft').length,
    totalEnrollments: courses.reduce((sum, c) => sum + c.enrollment_count, 0),
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-20">
        <FaSpinner className="animate-spin text-4xl text-primary" />
      </div>
    );
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-800">Dashboard</h1>
        <p className="text-gray-600 mt-2">Overview of your eCourses platform</p>
      </div>

      {/* Stats Grid */}
      <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm font-medium">Total Courses</p>
              <p className="text-3xl font-bold text-gray-800 mt-2">{stats.totalCourses}</p>
            </div>
            <div className="w-12 h-12 bg-primary bg-opacity-10 rounded-lg flex items-center justify-center">
              <FaBook className="text-2xl text-primary" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm font-medium">Published</p>
              <p className="text-3xl font-bold text-gray-800 mt-2">{stats.publishedCourses}</p>
            </div>
            <div className="w-12 h-12 bg-success bg-opacity-10 rounded-lg flex items-center justify-center">
              <FaCheckCircle className="text-2xl text-success" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm font-medium">Draft Courses</p>
              <p className="text-3xl font-bold text-gray-800 mt-2">{stats.draftCourses}</p>
            </div>
            <div className="w-12 h-12 bg-warning bg-opacity-10 rounded-lg flex items-center justify-center">
              <FaChartLine className="text-2xl text-warning" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm font-medium">Total Enrollments</p>
              <p className="text-3xl font-bold text-gray-800 mt-2">{stats.totalEnrollments}</p>
            </div>
            <div className="w-12 h-12 bg-danger bg-opacity-10 rounded-lg flex items-center justify-center">
              <FaUsers className="text-2xl text-danger" />
            </div>
          </div>
        </div>
      </div>

      {/* Recent Courses */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-xl font-bold text-gray-800">Recent Courses</h2>
          <Link
            to="/courses"
            className="text-primary hover:text-blue-700 font-medium text-sm"
          >
            View All
          </Link>
        </div>

        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Course
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Instructor
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
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {courses.slice(0, 5).map((course) => (
                <tr key={course.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4">
                    <p className="font-medium text-gray-900">{course.title}</p>
                  </td>
                  <td className="px-6 py-4 text-gray-600">{course.instructor_name}</td>
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
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
