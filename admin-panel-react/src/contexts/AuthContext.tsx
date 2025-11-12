import React, { createContext, useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { adminApi } from '../services/api';
import { User, LoginRequest } from '../types';

interface AuthContextType {
  user: User | null;
  loading: boolean;
  login: (data: LoginRequest) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
  isAdmin: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const savedUser = adminApi.getUser();
    const token = adminApi.getToken();

    if (savedUser && token) {
      setUser(savedUser);
      adminApi.getProfile().catch(() => {
        adminApi.logout();
        setUser(null);
      });
    }
    setLoading(false);
  }, []);

  const login = async (data: LoginRequest) => {
    try {
      const response = await adminApi.login(data);

      // Check if user is admin
      if (response.data.user.role !== 'admin') {
        throw new Error('Access denied. Admin role required.');
      }

      setUser(response.data.user);
      navigate('/dashboard');
    } catch (error) {
      throw error;
    }
  };

  const logout = () => {
    adminApi.logout();
    setUser(null);
    navigate('/login');
  };

  const value: AuthContextType = {
    user,
    loading,
    login,
    logout,
    isAuthenticated: !!user,
    isAdmin: user?.role === 'admin',
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
