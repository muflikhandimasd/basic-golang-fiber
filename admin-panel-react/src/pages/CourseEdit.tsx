import { useParams, useNavigate } from 'react-router-dom';
import { FaArrowLeft } from 'react-icons/fa';

export default function CourseEdit() {
  const { id } = useParams();
  const navigate = useNavigate();

  return (
    <div>
      <button
        onClick={() => navigate('/courses')}
        className="flex items-center space-x-2 text-primary hover:text-blue-700 mb-6"
      >
        <FaArrowLeft />
        <span>Back to Courses</span>
      </button>

      <div className="bg-white rounded-lg shadow-md p-8">
        <h1 className="text-2xl font-bold text-gray-800 mb-4">Edit Course #{id}</h1>
        <p className="text-gray-600">
          Course editing interface coming soon. For now, use the API endpoints to update courses.
        </p>
      </div>
    </div>
  );
}
