import Link from 'next/link';
import Navbar from '@/components/Navbar';
import { FaGraduationCap, FaChalkboardTeacher, FaUsers, FaCertificate } from 'react-icons/fa';

export default function Home() {
  return (
    <div className="min-h-screen">
      <Navbar />

      {/* Hero Section */}
      <section className="bg-gradient-to-r from-primary to-blue-700 text-white py-20">
        <div className="container mx-auto px-4 text-center">
          <h1 className="text-5xl font-bold mb-6">Welcome to eCourses</h1>
          <p className="text-xl mb-8 max-w-2xl mx-auto">
            Discover thousands of courses and learn at your own pace.
            Join millions of learners worldwide and unlock your potential.
          </p>
          <div className="flex justify-center space-x-4">
            <Link
              href="/courses"
              className="bg-white text-primary px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition"
            >
              Browse Courses
            </Link>
            <Link
              href="/register"
              className="bg-transparent border-2 border-white px-8 py-3 rounded-lg font-semibold hover:bg-white hover:text-primary transition"
            >
              Get Started
            </Link>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 bg-gray-50">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl font-bold text-center mb-12">Why Choose eCourses?</h2>
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div className="text-center">
              <div className="bg-primary text-white w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <FaGraduationCap className="text-3xl" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Quality Courses</h3>
              <p className="text-gray-600">
                Learn from expert instructors with real-world experience
              </p>
            </div>

            <div className="text-center">
              <div className="bg-success text-white w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <FaChalkboardTeacher className="text-3xl" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Expert Instructors</h3>
              <p className="text-gray-600">
                World-class instructors teaching practical skills
              </p>
            </div>

            <div className="text-center">
              <div className="bg-warning text-white w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <FaUsers className="text-3xl" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Community</h3>
              <p className="text-gray-600">
                Join a global community of passionate learners
              </p>
            </div>

            <div className="text-center">
              <div className="bg-danger text-white w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <FaCertificate className="text-3xl" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Certificates</h3>
              <p className="text-gray-600">
                Earn certificates upon course completion
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="bg-secondary text-white py-16">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-3xl font-bold mb-4">Ready to Start Learning?</h2>
          <p className="text-xl mb-8">Join thousands of students already learning on eCourses</p>
          <Link
            href="/register"
            className="bg-primary px-8 py-3 rounded-lg font-semibold hover:bg-blue-600 transition inline-block"
          >
            Sign Up Now
          </Link>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-800 text-white py-8">
        <div className="container mx-auto px-4 text-center">
          <p>&copy; 2024 eCourses. All rights reserved.</p>
          <p className="mt-2 text-gray-400">Built with Next.js and Golang</p>
        </div>
      </footer>
    </div>
  );
}
