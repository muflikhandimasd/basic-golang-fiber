'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import Navbar from '@/components/Navbar';
import { apiService } from '@/services/api';
import { CourseWithDetails, Lesson } from '@/types';
import { useAuth } from '@/contexts/AuthContext';
import { FaBook, FaUser, FaClock, FaPlay, FaLock, FaSpinner } from 'react-icons/fa';

export default function CourseDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { isAuthenticated } = useAuth();
  const [course, setCourse] = useState<CourseWithDetails | null>(null);
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [loading, setLoading] = useState(true);
  const [enrolling, setEnrolling] = useState(false);

  useEffect(() => {
    loadCourse();
  }, [params.id]);

  const loadCourse = async () => {
    try {
      const courseId = Number(params.id);
      const [courseData, lessonsData] = await Promise.all([
        apiService.getCourse(courseId),
        apiService.getCourseLessons(courseId),
      ]);
      setCourse(courseData);
      setLessons(lessonsData);
    } catch (error) {
      console.error('Failed to load course:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleEnroll = async () => {
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }

    try {
      setEnrolling(true);
      await apiService.enrollCourse(Number(params.id));
      alert('Successfully enrolled in the course!');
      router.push('/dashboard');
    } catch (error: any) {
      alert(error.response?.data?.message || 'Failed to enroll in course');
    } finally {
      setEnrolling(false);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50">
        <Navbar />
        <div className="flex justify-center items-center py-20">
          <FaSpinner className="animate-spin text-4xl text-primary" />
        </div>
      </div>
    );
  }

  if (!course) {
    return (
      <div className="min-h-screen bg-gray-50">
        <Navbar />
        <div className="container mx-auto px-4 py-20 text-center">
          <p className="text-xl text-gray-600">Course not found</p>
        </div>
      </div>
    );
  }

  const getLevelColor = (level: string) => {
    switch (level) {
      case 'beginner':
        return 'bg-green-100 text-green-800';
      case 'intermediate':
        return 'bg-yellow-100 text-yellow-800';
      case 'advanced':
        return 'bg-red-100 text-red-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />

      <div className="bg-gradient-to-r from-primary to-blue-700 text-white py-12">
        <div className="container mx-auto px-4">
          <div className="max-w-4xl">
            <span className={`inline-block px-3 py-1 rounded text-sm font-semibold mb-4 ${getLevelColor(course.level)}`}>
              {course.level}
            </span>
            <h1 className="text-4xl font-bold mb-4">{course.title}</h1>
            <p className="text-xl mb-6">{course.description}</p>

            <div className="flex items-center space-x-6 text-sm">
              <div className="flex items-center">
                <FaUser className="mr-2" />
                <span>{course.instructor_name}</span>
              </div>
              <div className="flex items-center">
                <FaBook className="mr-2" />
                <span>{course.lesson_count} lessons</span>
              </div>
              <div className="flex items-center">
                <FaClock className="mr-2" />
                <span>{course.enrollment_count} students enrolled</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-4 py-8">
        <div className="grid lg:grid-cols-3 gap-8">
          <div className="lg:col-span-2">
            <div className="bg-white rounded-lg shadow-md p-6 mb-6">
              <h2 className="text-2xl font-bold mb-4">Course Content</h2>
              <div className="space-y-3">
                {lessons.map((lesson, index) => (
                  <div
                    key={lesson.id}
                    className="flex items-center justify-between p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition"
                  >
                    <div className="flex items-center space-x-4">
                      {lesson.is_free ? (
                        <FaPlay className="text-primary" />
                      ) : (
                        <FaLock className="text-gray-400" />
                      )}
                      <div>
                        <h3 className="font-semibold">{lesson.title}</h3>
                        <p className="text-sm text-gray-600">
                          {Math.floor(lesson.duration / 60)} minutes
                        </p>
                      </div>
                    </div>
                    {lesson.is_free && (
                      <span className="text-xs bg-green-100 text-green-800 px-2 py-1 rounded">
                        FREE
                      </span>
                    )}
                  </div>
                ))}
              </div>
            </div>
          </div>

          <div className="lg:col-span-1">
            <div className="bg-white rounded-lg shadow-md p-6 sticky top-4">
              <div className="text-center mb-6">
                <div className="text-4xl font-bold text-primary mb-2">
                  ${course.price === 0 ? 'Free' : course.price.toFixed(2)}
                </div>
                {course.price === 0 && <p className="text-green-600">This course is free!</p>}
              </div>

              <button
                onClick={handleEnroll}
                disabled={enrolling}
                className="w-full bg-primary text-white py-3 rounded-lg font-semibold hover:bg-blue-600 transition disabled:opacity-50"
              >
                {enrolling ? 'Enrolling...' : 'Enroll Now'}
              </button>

              {course.category_name && (
                <div className="mt-6 pt-6 border-t border-gray-200">
                  <h3 className="font-semibold mb-2">Category</h3>
                  <p className="text-gray-600">{course.category_name}</p>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
