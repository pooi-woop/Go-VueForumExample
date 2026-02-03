<template>
  <div class="create-post-container">
    <div class="create-post-card">
      <div class="create-post-header">
        <h2 class="create-post-title">发布新帖子</h2>
        <p class="create-post-subtitle">分享你的游戏心得和技巧</p>
      </div>
      <el-form :model="postForm" :rules="rules" ref="postFormRef" class="create-post-form">
        <el-form-item label="帖子标题" prop="title">
          <el-input v-model="postForm.title" placeholder="请输入帖子标题" />
        </el-form-item>
        <el-form-item label="帖子内容" prop="content">
          <el-input v-model="postForm.content" type="textarea" rows="8" placeholder="请输入帖子内容..." />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit" :loading="loading">发布帖子</el-button>
          <el-button @click="handleCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { postAPI } from '../services/api';

export default {
  name: 'CreatePost',
  data() {
    return {
      postForm: {
        title: '',
        content: ''
      },
      rules: {
        title: [
          { required: true, message: '请输入帖子标题', trigger: 'blur' },
          { min: 3, max: 100, message: '标题长度应在3-100个字符之间', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入帖子内容', trigger: 'blur' },
          { min: 10, message: '内容长度至少为10个字符', trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  methods: {
    handleSubmit() {
      this.$refs.postFormRef.validate(async (valid) => {
        if (valid) {
          this.loading = true;
          try {
            await postAPI.createPost(this.postForm.title, this.postForm.content);
            this.$message.success('帖子发布成功');
            this.$router.push('/');
          } catch (error) {
            this.$message.error(`发布失败: ${error.message}`);
          } finally {
            this.loading = false;
          }
        }
      });
    },
    handleCancel() {
      this.$router.push('/');
    }
  },
  mounted() {
    // 检查用户是否已登录
    const token = localStorage.getItem('auth_token');
    if (!token) {
      this.$message.warning('请先登录');
      this.$router.push('/login');
    }
  }
}
</script>

<style scoped>
.create-post-container {
  display: flex;
  justify-content: center;
  padding: 2rem;
}

.create-post-card {
  background: rgba(22, 33, 62, 0.9);
  border-radius: 12px;
  padding: 2.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  width: 100%;
  max-width: 800px;
  border: 1px solid rgba(233, 69, 96, 0.3);
}

.create-post-header {
  text-align: center;
  margin-bottom: 2rem;
}

.create-post-title {
  color: #e94560;
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
}

.create-post-subtitle {
  color: #aaa;
  font-size: 1rem;
}

.create-post-form {
  margin-bottom: 1.5rem;
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

.el-button--primary {
  background: linear-gradient(135deg, #e94560, #c7365f);
  border: none;
  transition: all 0.3s ease;
}

.el-button--primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(233, 69, 96, 0.4);
}

.el-button--primary.is-loading {
  background: linear-gradient(135deg, #e94560, #c7365f);
}

.el-button {
  margin-right: 1rem;
}
</style>