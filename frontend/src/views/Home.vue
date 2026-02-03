<template>
  <div class="home-container">
    <div class="banner">
      <img src="https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20delta%20force%20forum%20banner%20with%20cyberpunk%20elements%20and%20rain%20effect&image_size=landscape_16_9" alt="论坛 banner" />
      <h2 class="banner-title">欢迎来到恶雨三角洲行动论坛</h2>
    </div>
    
    <div class="post-list">
      <h3 class="section-title">最新帖子</h3>
      <div class="posts-grid">
        <div v-for="post in posts" :key="post.id" class="post-card">
          <div class="post-header">
            <div class="user-info">
              <img :src="post.avatar" alt="用户头像" class="user-avatar" />
              <span class="username">{{ post.username }}</span>
            </div>
            <span class="post-time">{{ post.time }}</span>
          </div>
          <h4 class="post-title">{{ post.title }}</h4>
          <p class="post-content">{{ post.content }}</p>
          <div class="post-footer">
            <span class="post-stats">{{ post.views }} 浏览 · {{ post.comments }} 评论</span>
            <router-link :to="`/post/${post.id}`" class="read-more">阅读更多</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      posts: [
        {
          id: 1,
          title: '三角洲行动新赛季攻略分享',
          content: '新赛季已经开始，分享一些个人的游戏心得和技巧，希望对大家有所帮助...',
          username: '战术大师',
          avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20male%20soldier%20avatar%20with%20cyberpunk%20elements&image_size=square',
          time: '2026-02-03 10:00',
          views: 1234,
          comments: 56
        },
        {
          id: 2,
          title: '游戏角色cosplay分享',
          content: '最近cos了游戏中的角色，大家看看效果如何...',
          username: '二次元战士',
          avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20female%20cosplayer%20avatar&image_size=square',
          time: '2026-02-03 09:30',
          views: 892,
          comments: 34
        },
        {
          id: 3,
          title: '寻找组队队友',
          content: '想找几个固定队友一起开黑，有意向的可以留言...',
          username: '孤独玩家',
          avatar: 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=anime%20style%20male%20gamer%20avatar&image_size=square',
          time: '2026-02-03 08:45',
          views: 567,
          comments: 23
        }
      ]
    }
  }
}
</script>

<style scoped>
.home-container {
  width: 100%;
}

.banner {
  position: relative;
  margin-bottom: 2rem;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
}

.banner img {
  width: 100%;
  height: 300px;
  object-fit: cover;
}

.banner-title {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: linear-gradient(transparent, rgba(0, 0, 0, 0.8));
  padding: 2rem;
  margin: 0;
  color: #fff;
  font-size: 1.8rem;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
}

.section-title {
  color: #e94560;
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
  border-bottom: 2px solid #e94560;
  padding-bottom: 0.5rem;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 1.5rem;
}

.post-card {
  background: rgba(22, 33, 62, 0.8);
  border-radius: 10px;
  padding: 1.5rem;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
  border: 1px solid rgba(233, 69, 96, 0.2);
}

.post-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 24px rgba(233, 69, 96, 0.3);
}

.post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.8rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e94560;
}

.username {
  color: #fff;
  font-weight: 500;
}

.post-time {
  color: #aaa;
  font-size: 0.8rem;
}

.post-title {
  color: #fff;
  font-size: 1.2rem;
  margin-bottom: 1rem;
  font-weight: 600;
}

.post-content {
  color: #ddd;
  font-size: 0.9rem;
  line-height: 1.6;
  margin-bottom: 1.5rem;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  padding-top: 1rem;
}

.post-stats {
  color: #aaa;
  font-size: 0.8rem;
}

.read-more {
  color: #e94560;
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

.read-more:hover {
  color: #fff;
}
</style>