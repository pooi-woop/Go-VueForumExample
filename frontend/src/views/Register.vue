<template>
  <div class="register-container">
    <div class="register-card">
      <div class="register-header">
        <h2 class="register-title">用户注册</h2>
        <p class="register-subtitle">加入恶雨三角洲行动论坛</p>
      </div>
      <el-form :model="registerForm" :rules="rules" ref="registerFormRef" class="register-form">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="registerForm.username" placeholder="请输入用户名" prefix-icon="User" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="registerForm.password" type="password" placeholder="请输入密码" prefix-icon="Lock" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="registerForm.confirmPassword" type="password" placeholder="请确认密码" prefix-icon="Check" show-password />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="registerForm.email" type="email" placeholder="请输入邮箱" prefix-icon="Message" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="register-button" @click="handleRegister" :loading="loading">注册</el-button>
        </el-form-item>
      </el-form>
      <div class="register-footer">
        <span>已有账号？</span>
        <router-link to="/login" class="login-link">立即登录</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { authAPI } from '../services/api';

export default {
  name: 'Register',
  data() {
    return {
      registerForm: {
        username: '',
        password: '',
        confirmPassword: '',
        email: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 20, message: '用户名长度应在3-20个字符之间', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, message: '密码长度至少为6个字符', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请确认密码', trigger: 'blur' },
          { validator: this.validateConfirmPassword, trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  methods: {
    validateConfirmPassword(rule, value, callback) {
      if (value !== this.registerForm.password) {
        callback(new Error('两次输入的密码不一致'));
      } else {
        callback();
      }
    },
    async handleRegister() {
      this.$refs.registerFormRef.validate(async (valid) => {
        if (valid) {
          this.loading = true;
          try {
            await authAPI.register(this.registerForm.username, this.registerForm.password, this.registerForm.email);
            this.$message.success('注册成功');
            this.$router.push('/login');
          } catch (error) {
            this.$message.error(`注册失败: ${error.message}`);
          } finally {
            this.loading = false;
          }
        }
      });
    }
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  padding: 2rem;
}

.register-card {
  background: rgba(22, 33, 62, 0.9);
  border-radius: 12px;
  padding: 2.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  width: 100%;
  max-width: 450px;
  border: 1px solid rgba(233, 69, 96, 0.3);
}

.register-header {
  text-align: center;
  margin-bottom: 2rem;
}

.register-title {
  color: #e94560;
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
}

.register-subtitle {
  color: #aaa;
  font-size: 1rem;
}

.register-form {
  margin-bottom: 1.5rem;
}

.register-button {
  width: 100%;
  padding: 0.8rem;
  font-size: 1rem;
}

.el-button--primary {
  background: linear-gradient(135deg, #e94560, #c7365f);
  border: none;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.el-button--primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(233, 69, 96, 0.4);
}

.el-button--primary.is-loading {
  background: linear-gradient(135deg, #e94560, #c7365f);
}

.el-form-item__label {
  color: #ddd;
}

.el-input__wrapper {
  background-color: rgba(0, 0, 0, 0.2);
  border-color: rgba(233, 69, 96, 0.3);
}

.el-input__wrapper:hover {
  box-shadow: 0 0 0 1px rgba(233, 69, 96, 0.5) inset;
}

.el-input__wrapper.is-focus {
  box-shadow: 0 0 0 1px rgba(233, 69, 96, 0.5) inset;
}

.el-input__inner {
  color: #fff;
}

.register-footer {
  text-align: center;
  color: #aaa;
  font-size: 0.9rem;
}

.login-link {
  color: #e94560;
  text-decoration: none;
  margin-left: 0.5rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

.login-link:hover {
  color: #fff;
}
</style>