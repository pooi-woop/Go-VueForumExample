<template>
  <div class="post-detail-container">
    <div class="post-detail-card">
      <div class="post-header">
        <div class="user-info">
          <img :src="post.avatar" alt="用户头像" class="user-avatar" />
          <div class="user-details">
            <span class="username">{{ post.username }}</span>
            <span class="post-time">{{ post.time }}</span>
          </div>
        </div>
      </div>
      <h2 class="post-title">{{ post.title }}</h2>
      <div class="post-content">
        <p>{{ post.content }}</p>
      </div>
      <div class="post-stats">
        <span>{{ post.views }} 浏览 · {{ post.comments }} 评论</span>
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
            <el-button type="primary" @click="handleComment">发表评论</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="comments-list">
        <div v-for="comment in comments" :key="comment.id" class="comment-item">
          <div class="comment-header">
            <div class="comment-user">
              <img :src="comment.avatar" alt="评论用户头像" class="comment-avatar" />
              <span class="comment-username">{{ comment.username }}</span>
            </div>
            <span class="comment-time">{{ comment.time }}</span>
          </div>
          <div class="comment-content">{{ comment.content }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PostDetail',
  data() {
    return {
      post: {
        id: 1,
        title: '三角洲行动新赛季攻略分享',
        content: '新赛季已经开始，分享一些个人的游戏心得和技巧，希望对大家有所帮助。\n\n1. 武器选择：建议携带突击步枪和冲锋枪的组合，适应不同场景。\n2. 战术策略：利用地形优势，合理使用烟雾弹和闪光弹。\n3. 团队配合：与队友保持沟通，分工明确，互相掩护。\n4. 装备搭配：根据个人玩法风格选择合适的配件和技能。\n\n祝大家在新赛季取得好成绩！',
        username: '战术大师',
        avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20male%20soldier%20avatar%20with%20cyberpunk%20elements&image_size=square',
        time: '2026-02-03 10:00',
        views: 1234,
        comments: 56
      },
      comments: [
        {
          id: 1,
          username: '游戏新手',
          avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20young%20gamer%20avatar&image_size=square',
          content: '非常感谢分享，对我很有帮助！',
          time: '2026-02-03 10:30'
        },
        {
          id: 2,
          username: '老玩家',
          avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20veteran%20gamer%20avatar&image_size=square',
          content: '补充一点，狙击手在远距离作战中也很重要。',
          time: '2026-02-03 11:00'
        }
      ],
      commentForm: {
        content: ''
      },
      commentRules: {
        content: [
          { required: true, message: '请输入评论内容', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleComment() {
      this.$refs.commentFormRef.validate((valid) => {
        if (valid) {
          // 模拟评论提交
          const newComment = {
            id: this.comments.length + 1,
            username: '当前用户',
            avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20user%20avatar&image_size=square',
            content: this.commentForm.content,
            time: new Date().toLocaleString()
          };
          this.comments.push(newComment);
          this.commentForm.content = '';
          this.$message.success('评论发表成功');
        }
      });
    }
  }
}
</script>

<style scoped>
.post-detail-container {
  width: 100%;
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
</style>