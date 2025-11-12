-- Remove sample categories
DELETE FROM categories WHERE slug IN (
    'web-development',
    'mobile-development',
    'data-science',
    'devops',
    'design'
);

-- Remove default admin user
DELETE FROM users WHERE email = 'admin@ecourses.com';
