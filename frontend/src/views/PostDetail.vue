<template>
  <div class="post-detail-container">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="5" animated />
    </div>
    
    <!-- 错误状态 -->
    <div v-else-if="error" class="error-container">
      <el-alert
        :title="error"
        type="error"
        show-icon
        :closable="false"
      />
      <el-button type="primary" @click="fetchPostDetail" style="margin-top: 1rem">重试</el-button>
    </div>
    
    <!-- 帖子详情 -->
    <div v-else>
      <div class="post-detail-card">
        <div class="post-header">
          <div class="user-info">
            <img :src="post.avatar || 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20user%20avatar&image_size=square'" alt="用户头像" class="user-avatar" />
            <div class="user-details">
              <span class="username">{{ post.username }}</span>
              <span class="post-time">{{ formatTime(post.created_at) }}</span>
            </div>
          </div>
        </div>
        <h2 class="post-title">{{ post.title }}</h2>
        <div class="post-content">
          <p>{{ post.content }}</p>
        </div>
        <div class="post-stats">
          <span>{{ post.views }} 浏览 · {{ post.comment_count }} 评论</span>
        </div>
      </div>
      
      <div class="comments-section">
        <h3 class="section-title">评论区</h3>
        <div class="comment-form">
          <el-form :model="commentForm" :rules="commentRules" ref="commentFormRef">
            <el-form-item prop="content">
              <el-input v-model="commentForm.content" type="textarea" rows="3" placeholder="写下你的评论..." />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleComment" :loading="submittingComment">发表评论</el-button>
            </el-form-item>
          </el-form>
        </div>
        
        <!-- 评论加载状态 -->
        <div v-if="loadingComments" class="loading-container">
          <el-skeleton :rows="2" animated />
        </div>
        
        <!-- 评论错误状态 -->
        <div v-else-if="errorComments" class="error-container">
          <el-alert
            :title="errorComments"
            type="error"
            show-icon
            :closable="false"
          />
          <el-button type="primary" @click="fetchComments" style="margin-top: 1rem">重试</el-button>
        </div>
        
        <!-- 评论列表 -->
        <div v-else class="comments-list">
          <div v-for="comment in comments" :key="comment.id" class="comment-item">
            <div class="comment-header">
              <div class="comment-user">
                <img :src="comment.avatar || 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20user%20avatar&image_size=square'" alt="评论用户头像" class="comment-avatar" />
                <span class="comment-username">{{ comment.username }}</span>
              </div>
              <span class="comment-time">{{ formatTime(comment.created_at) }}</span>
            </div>
            <div class="comment-content">{{ comment.content }}</div>
          </div>
          
          <!-- 空评论状态 -->
          <div v-if="comments.length === 0" class="empty-container">
            <el-empty description="暂无评论" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { postAPI, commentAPI } from '../services/api';

export default {
  name: 'PostDetail',
  data() {
    return {
      post: {},
      comments: [],
      commentForm: {
        content: ''
      },
      commentRules: {
        content: [
          { required: true, message: '请输入评论内容', trigger: 'blur' }
        ]
      },
      loading: true,
      loadingComments: false,
      error: null,
      errorComments: null,
      submittingComment: false
    }
  },
  methods: {
    async fetchPostDetail() {
      this.loading = true;
      this.error = null;
      try {
        const postId = this.$route.params.id;
        const data = await postAPI.getPost(postId);
        this.post = data;
        this.fetchComments();
      } catch (error) {
        this.error = `获取帖子详情失败: ${error.message}`;
        console.error('获取帖子详情失败:', error);
      } finally {
        this.loading = false;
      }
    },
    
    async fetchComments() {
      this.loadingComments = true;
      this.errorComments = null;
      try {
        const postId = this.$route.params.id;
        const data = await commentAPI.getComments(postId);
        this.comments = data;
      } catch (error) {
        this.errorComments = `获取评论失败: ${error.message}`;
        console.error('获取评论失败:', error);
      } finally {
        this.loadingComments = false;
      }
    },
    
    async handleComment() {
      this.$refs.commentFormRef.validate(async (valid) => {
        if (valid) {
          this.submittingComment = true;
          try {
            const postId = this.$route.params.id;
            await commentAPI.createComment(postId, this.commentForm.content);
            this.commentForm.content = '';
            this.$message.success('评论发表成功');
            this.fetchComments();
          } catch (error) {
            this.$message.error(`评论发表失败: ${error.message}`);
          } finally {
            this.submittingComment = false;
          }
        }
      });
    },
    
    formatTime(timeString) {
      if (!timeString) return '';
      const date = new Date(timeString);
      return date.toLocaleString();
    }
  },
  mounted() {
    this.fetchPostDetail();
  },
  watch: {
    '$route.params.id': {
      handler() {
        this.fetchPostDetail();
      },
      immediate: true
    }
  }
}
</script>

<style scoped>
.post-detail-container {
  width: 100%;
}

.loading-container {
  background: rgba(22, 33, 62, 0.8);
  border-radius: 10px;
  padding: 2rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(233, 69, 96, 0.2);
}

.error-container {
  background: rgba(22, 33, 62, 0.8);
  border-radius: 10px;
  padding: 2rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(233, 69, 96, 0.2);
}

.empty-container {
  padding: 2rem;
  text-align: center;
}

.post-detail-card {
  background: rgba(22, 33, 62, 0.8);
  border-radius: 10px;
  padding: 2rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(233, 69, 96, 0.2);
}

.post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e94560;
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

.username {
  color: #fff;
  font-weight: 600;
  font-size: 1.1rem;
}

.post-time {
  color: #aaa;
  font-size: 0.9rem;
}

.post-title {
  color: #fff;
  font-size: 1.8rem;
  margin-bottom: 1.5rem;
  font-weight: 700;
}

.post-content {
  color: #ddd;
  font-size: 1rem;
  line-height: 1.8;
  margin-bottom: 2rem;
}

.post-stats {
  color: #aaa;
  font-size: 0.9rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  padding-top: 1rem;
}

.comments-section {
  background: rgba(22, 33, 62, 0.8);
  border-radius: 10px;
  padding: 2rem;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(233, 69, 96, 0.2);
}

.section-title {
  color: #e94560;
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
  border-bottom: 2px solid #e94560;
  padding-bottom: 0.5rem;
}

.comment-form {
  margin-bottom: 2rem;
}

.comments-list {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.comment-item {
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  border-left: 3px solid #e94560;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.comment-user {
  display: flex;
  align-items: center;
  gap: 0.8rem;
}

.comment-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
}

.comment-username {
  color: #fff;
  font-weight: 500;
}

.comment-time {
  color: #aaa;
  font-size: 0.8rem;
}

.comment-content {
  color: #ddd;
  font-size: 0.9rem;
  line-height: 1.6;
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
</style>