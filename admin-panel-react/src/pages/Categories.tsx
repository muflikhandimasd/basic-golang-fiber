import { useEffect, useState } from 'react';
import { adminApi } from '../services/api';
import { Category } from '../types';
import { FaEdit, FaTrash, FaPlus, FaSpinner, FaTimes } from 'react-icons/fa';

export default function Categories() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [editingCategory, setEditingCategory] = useState<Category | null>(null);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    slug: '',
  });

  useEffect(() => {
    loadCategories();
  }, []);

  const loadCategories = async () => {
    try {
      const data = await adminApi.getCategories();
      setCategories(data);
    } catch (error) {
      console.error('Failed to load categories:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (editingCategory) {
        await adminApi.updateCategory(editingCategory.id, formData);
        alert('Category updated successfully');
      } else {
        await adminApi.createCategory(formData);
        alert('Category created successfully');
      }
      setShowModal(false);
      setFormData({ name: '', description: '', slug: '' });
      setEditingCategory(null);
      loadCategories();
    } catch (error: any) {
      alert(error.response?.data?.message || 'Failed to save category');
    }
  };

  const handleEdit = (category: Category) => {
    setEditingCategory(category);
    setFormData({
      name: category.name,
      description: category.description,
      slug: category.slug,
    });
    setShowModal(true);
  };

  const handleDelete = async (id: number) => {
    if (!confirm('Are you sure you want to delete this category?')) {
      return;
    }

    try {
      await adminApi.deleteCategory(id);
      setCategories(categories.filter((c) => c.id !== id));
      alert('Category deleted successfully');
    } catch (error) {
      alert('Failed to delete category');
    }
  };

  const openCreateModal = () => {
    setEditingCategory(null);
    setFormData({ name: '', description: '', slug: '' });
    setShowModal(true);
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
          <h1 className="text-3xl font-bold text-gray-800">Categories Management</h1>
          <p className="text-gray-600 mt-2">Manage course categories</p>
        </div>
        <button
          onClick={openCreateModal}
          className="bg-primary text-white px-6 py-3 rounded-lg hover:bg-blue-600 transition flex items-center space-x-2"
        >
          <FaPlus />
          <span>Add Category</span>
        </button>
      </div>

      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  ID
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Name
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Slug
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Description
                </th>
                <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {categories.map((category) => (
                <tr key={category.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 text-gray-900">{category.id}</td>
                  <td className="px-6 py-4">
                    <p className="font-medium text-gray-900">{category.name}</p>
                  </td>
                  <td className="px-6 py-4 text-gray-600">{category.slug}</td>
                  <td className="px-6 py-4 text-gray-600 max-w-md truncate">
                    {category.description || '-'}
                  </td>
                  <td className="px-6 py-4 text-right">
                    <div className="flex justify-end space-x-2">
                      <button
                        onClick={() => handleEdit(category)}
                        className="text-primary hover:text-blue-700 p-2"
                        title="Edit"
                      >
                        <FaEdit />
                      </button>
                      <button
                        onClick={() => handleDelete(category.id)}
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

        {categories.length === 0 && (
          <div className="text-center py-12 text-gray-500">No categories found</div>
        )}
      </div>

      {/* Modal */}
      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-lg max-w-md w-full p-6">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-bold text-gray-800">
                {editingCategory ? 'Edit Category' : 'Add Category'}
              </h2>
              <button
                onClick={() => setShowModal(false)}
                className="text-gray-500 hover:text-gray-700"
              >
                <FaTimes />
              </button>
            </div>

            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label className="block text-gray-700 font-semibold mb-2">Name</label>
                <input
                  type="text"
                  value={formData.name}
                  onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary text-gray-800"
                  required
                />
              </div>

              <div>
                <label className="block text-gray-700 font-semibold mb-2">Slug</label>
                <input
                  type="text"
                  value={formData.slug}
                  onChange={(e) => setFormData({ ...formData, slug: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary text-gray-800"
                  required
                />
              </div>

              <div>
                <label className="block text-gray-700 font-semibold mb-2">Description</label>
                <textarea
                  value={formData.description}
                  onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary text-gray-800"
                  rows={3}
                />
              </div>

              <div className="flex space-x-4">
                <button
                  type="submit"
                  className="flex-1 bg-primary text-white py-2 rounded-lg hover:bg-blue-600 transition"
                >
                  {editingCategory ? 'Update' : 'Create'}
                </button>
                <button
                  type="button"
                  onClick={() => setShowModal(false)}
                  className="px-6 py-2 border border-gray-300 rounded-lg hover:bg-gray-50 transition text-gray-800"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
