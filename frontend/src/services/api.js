// API服务层，实现与后端的交互

const API_BASE_URL = 'http://localhost:8080/api';

// 获取存储的认证令牌
const getAuthToken = () => {
  return localStorage.getItem('auth_token');
};

// 设置认证令牌
const setAuthToken = (token) => {
  if (token) {
    localStorage.setItem('auth_token', token);
  } else {
    localStorage.removeItem('auth_token');
  }
};

// 通用请求方法
const request = async (url, options = {}) => {
  const token = getAuthToken();
  
  const headers = {
    'Content-Type': 'application/json',
    ...options.headers,
  };
  
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  
  try {
    const response = await fetch(`${API_BASE_URL}${url}`, {
      ...options,
      headers,
    });
    
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `请求失败: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('API请求错误:', error);
    throw error;
  }
};

// 认证相关API
export const authAPI = {
  // 登录
  login: async (username, password) => {
    const data = await request('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ username, password }),
    });
    setAuthToken(data.token);
    return data;
  },
  
  // 注册
  register: async (username, password, email) => {
    return request('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ username, password, email }),
    });
  },
  
  // 登出
  logout: () => {
    setAuthToken(null);
  },
};

// 帖子相关API
export const postAPI = {
  // 获取帖子列表
  getPosts: async () => {
    return request('/posts');
  },
  
  // 获取帖子详情
  getPost: async (id) => {
    return request(`/posts/${id}`);
  },
  
  // 创建帖子
  createPost: async (title, content) => {
    return request('/posts', {
      method: 'POST',
      body: JSON.stringify({ title, content }),
    });
  },
  
  // 更新帖子
  updatePost: async (id, title, content) => {
    return request(`/posts/${id}`, {
      method: 'PUT',
      body: JSON.stringify({ title, content }),
    });
  },
  
  // 删除帖子
  deletePost: async (id) => {
    return request(`/posts/${id}`, {
      method: 'DELETE',
    });
  },
};

// 评论相关API
export const commentAPI = {
  // 获取帖子的评论
  getComments: async (postId) => {
    return request(`/comments/post/${postId}`);
  },
  
  // 创建评论
  createComment: async (postId, content) => {
    return request('/comments', {
      method: 'POST',
      body: JSON.stringify({ post_id: postId, content }),
    });
  },
  
  // 删除评论
  deleteComment: async (id) => {
    return request(`/comments/${id}`, {
      method: 'DELETE',
    });
  },
};

// 用户相关API
export const userAPI = {
  // 获取当前用户信息
  getCurrentUser: async () => {
    return request('/users/me');
  },
};

// 导出所有API
export default {
  auth: authAPI,
  post: postAPI,
  comment: commentAPI,
  user: userAPI,
  getAuthToken,
  setAuthToken,
};
