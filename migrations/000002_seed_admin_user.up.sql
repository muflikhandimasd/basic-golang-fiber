-- Insert default admin user
-- Password: admin123 (hashed with bcrypt)
INSERT INTO users (email, password, full_name, role)
VALUES (
    'admin@ecourses.com',
    '$2a$10$xN7qZV8L9XqX8e3Z9yYz3.rY5mQYO9KZ9bC3L6qH8yV3Z8eZ8eZ8e',
    'Admin User',
    'admin'
) ON CONFLICT (email) DO NOTHING;

-- Insert sample categories
INSERT INTO categories (name, description, slug) VALUES
    ('Web Development', 'Learn web development technologies', 'web-development'),
    ('Mobile Development', 'Build mobile applications', 'mobile-development'),
    ('Data Science', 'Master data science and analytics', 'data-science'),
    ('DevOps', 'DevOps and cloud technologies', 'devops'),
    ('Design', 'UI/UX and graphic design', 'design')
ON CONFLICT (slug) DO NOTHING;
