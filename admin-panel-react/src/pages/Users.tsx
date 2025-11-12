export default function Users() {
  return (
    <div>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-800">Users Management</h1>
        <p className="text-gray-600 mt-2">Manage all users on the platform</p>
      </div>

      <div className="bg-white rounded-lg shadow-md p-8 text-center">
        <p className="text-gray-600">
          User management interface coming soon. Currently, users can be viewed through the Golang admin panel at{' '}
          <a
            href="http://localhost:9000/admin/users"
            target="_blank"
            rel="noopener noreferrer"
            className="text-primary hover:underline"
          >
            http://localhost:9000/admin/users
          </a>
        </p>
      </div>
    </div>
  );
}
