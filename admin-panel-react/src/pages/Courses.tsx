import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { adminApi } from '../services/api';
import { CourseWithDetails } from '../types';
import { FaEdit, FaTrash, FaPlus, FaSpinner, FaSearch } from 'react-icons/fa';

export default function Courses() {
  const [courses, setCourses] = useState<CourseWithDetails[]>([]);
  const [filteredCourses, setFilteredCourses] = useState<CourseWithDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const [statusFilter, setStatusFilter] = useState('all');

  useEffect(() => {
    loadCourses();
  }, []);

  useEffect(() => {
    filterCourses();
  }, [courses, searchTerm, statusFilter]);

  const loadCourses = async () => {
    try {
      const data = await adminApi.getCourses({ limit: 1000 });
      setCourses(data);
    } catch (error) {
      console.error('Failed to load courses:', error);
    } finally {
      setLoading(false);
    }
  };

  const filterCourses = () => {
    let filtered = courses;

    if (statusFilter !== 'all') {
      filtered = filtered.filter((c) => c.status === statusFilter);
    }

    if (searchTerm) {
      filtered = filtered.filter(
        (c) =>
          c.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
          c.instructor_name.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }

    setFilteredCourses(filtered);
  };

  const handleDelete = async (id: number) => {
    if (!confirm('Are you sure you want to delete this course?')) {
      return;
    }

    try {
      await adminApi.deleteCourse(id);
      setCourses(courses.filter((c) => c.id !== id));
      alert('Course deleted successfully');
    } catch (error) {
      alert('Failed to delete course');
    }
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
      <div className="flex justify-between items-center mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">Courses Management</h1>
          <p className="text-gray-600 mt-2">Manage all courses on the platform</p>
        </div>
      </div>

      {/* Filters */}
      <div className="bg-white rounded-lg shadow-md p-4 mb-6">
        <div className="flex flex-col md:flex-row gap-4">
          <div className="flex-1">
            <div className="relative">
              <FaSearch className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
              <input
                type="text"
                placeholder="Search courses..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary text-gray-800"
              />
            </div>
          </div>
          <select
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value)}
            className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary text-gray-800"
          >
            <option value="all">All Status</option>
            <option value="published">Published</option>
            <option value="draft">Draft</option>
            <option value="archived">Archived</option>
          </select>
        </div>
      </div>

      {/* Courses Table */}
      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  ID
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Course
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Instructor
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Category
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
              {filteredCourses.map((course) => (
                <tr key={course.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 text-gray-900">{course.id}</td>
                  <td className="px-6 py-4">
                    <p className="font-medium text-gray-900">{course.title}</p>
                    <p className="text-sm text-gray-500">{course.level}</p>
                  </td>
                  <td className="px-6 py-4 text-gray-600">{course.instructor_name}</td>
                  <td className="px-6 py-4 text-gray-600">{course.category_name || '-'}</td>
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
                      <Link
                        to={`/courses/${course.id}/edit`}
                        className="text-primary hover:text-blue-700 p-2"
                        title="Edit"
                      >
                        <FaEdit />
                      </Link>
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

        {filteredCourses.length === 0 && (
          <div className="text-center py-12 text-gray-500">
            No courses found matching your criteria
          </div>
        )}
      </div>
    </div>
  );
}
