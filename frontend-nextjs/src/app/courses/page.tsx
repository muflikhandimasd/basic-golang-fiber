'use client';

import { useEffect, useState } from 'react';
import Navbar from '@/components/Navbar';
import CourseCard from '@/components/CourseCard';
import { apiService } from '@/services/api';
import { CourseWithDetails } from '@/types';
import { FaSpinner } from 'react-icons/fa';

export default function CoursesPage() {
  const [courses, setCourses] = useState<CourseWithDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [filter, setFilter] = useState('all');

  useEffect(() => {
    loadCourses();
  }, [filter]);

  const loadCourses = async () => {
    try {
      setLoading(true);
      const status = filter === 'all' ? '' : filter;
      const data = await apiService.getCourses({ status, limit: 100 });
      setCourses(data);
    } catch (error) {
      console.error('Failed to load courses:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />

      <div className="container mx-auto px-4 py-8">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-4xl font-bold text-gray-800">All Courses</h1>

          <select
            value={filter}
            onChange={(e) => setFilter(e.target.value)}
            className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
          >
            <option value="all">All Courses</option>
            <option value="published">Published</option>
            <option value="draft">Draft</option>
          </select>
        </div>

        {loading ? (
          <div className="flex justify-center items-center py-20">
            <FaSpinner className="animate-spin text-4xl text-primary" />
          </div>
        ) : courses.length > 0 ? (
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {courses.map((course) => (
              <CourseCard key={course.id} course={course} />
            ))}
          </div>
        ) : (
          <div className="text-center py-20">
            <p className="text-gray-600 text-lg">No courses found.</p>
          </div>
        )}
      </div>
    </div>
  );
}
