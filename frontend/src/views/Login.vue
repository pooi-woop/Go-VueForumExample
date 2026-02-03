<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h2 class="login-title">用户登录</h2>
        <p class="login-subtitle">欢迎回到恶雨三角洲行动论坛</p>
      </div>
      <el-form :model="loginForm" :rules="rules" ref="loginFormRef" class="login-form">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" placeholder="请输入用户名" prefix-icon="User" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="loginForm.password" type="password" placeholder="请输入密码" prefix-icon="Lock" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="login-button" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
      <div class="login-footer">
        <span>还没有账号？</span>
        <router-link to="/register" class="register-link">立即注册</router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleLogin() {
      this.$refs.loginFormRef.validate((valid) => {
        if (valid) {
          // 模拟登录请求
          this.$message.success('登录成功');
          this.$router.push('/');
        }
      });
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  padding: 2rem;
}

.login-card {
  background: rgba(22, 33, 62, 0.9);
  border-radius: 12px;
  padding: 2.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  width: 100%;
  max-width: 400px;
  border: 1px solid rgba(233, 69, 96, 0.3);
}

.login-header {
  text-align: center;
  margin-bottom: 2rem;
}

.login-title {
  color: #e94560;
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
}

.login-subtitle {
  color: #aaa;
  font-size: 1rem;
}

.login-form {
  margin-bottom: 1.5rem;
}

.login-button {
  width: 100%;
  padding: 0.8rem;
  font-size: 1rem;
  background: linear-gradient(135deg, #e94560, #c7365f);
  border: none;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(233, 69, 96, 0.4);
}

.login-footer {
  text-align: center;
  color: #aaa;
  font-size: 0.9rem;
}

.register-link {
  color: #e94560;
  text-decoration: none;
  margin-left: 0.5rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

.register-link:hover {
  color: #fff;
}
</style>