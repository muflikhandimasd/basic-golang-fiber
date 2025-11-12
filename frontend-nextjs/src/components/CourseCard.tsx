import Link from 'next/link';
import { CourseWithDetails } from '@/types';
import { FaUser, FaBook, FaClock } from 'react-icons/fa';

interface CourseCardProps {
  course: CourseWithDetails;
}

export default function CourseCard({ course }: CourseCardProps) {
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
    <Link href={`/courses/${course.id}`}>
      <div className="bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300 overflow-hidden cursor-pointer h-full">
        <div className="h-48 bg-gradient-to-r from-primary to-blue-600 flex items-center justify-center">
          <FaBook className="text-6xl text-white opacity-50" />
        </div>
        <div className="p-6">
          <div className="flex justify-between items-start mb-2">
            <h3 className="text-xl font-bold text-gray-800 line-clamp-2">{course.title}</h3>
            <span className={`px-2 py-1 rounded text-xs font-semibold ${getLevelColor(course.level)}`}>
              {course.level}
            </span>
          </div>
          <p className="text-gray-600 text-sm mb-4 line-clamp-2">{course.description}</p>

          <div className="flex items-center text-sm text-gray-500 mb-3">
            <FaUser className="mr-2" />
            <span>{course.instructor_name}</span>
          </div>

          <div className="flex justify-between items-center text-sm text-gray-600">
            <div className="flex items-center">
              <FaBook className="mr-1" />
              <span>{course.lesson_count} lessons</span>
            </div>
            <div className="text-2xl font-bold text-primary">
              ${course.price === 0 ? 'Free' : course.price.toFixed(2)}
            </div>
          </div>

          {course.category_name && (
            <div className="mt-3 pt-3 border-t border-gray-200">
              <span className="text-xs text-gray-500">{course.category_name}</span>
            </div>
          )}
        </div>
      </div>
    </Link>
  );
}
