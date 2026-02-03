<template>
  <div class="app-container">
    <header class="app-header">
      <div class="header-content">
        <h1 class="app-title">恶雨三角洲行动论坛</h1>
        <div class="header-nav">
          <router-link to="/" class="nav-link">首页</router-link>
          <router-link to="/login" class="nav-link">登录</router-link>
          <router-link to="/register" class="nav-link">注册</router-link>
        </div>
      </div>
    </header>
    <main class="app-main">
      <router-view></router-view>
    </main>
    <footer class="app-footer">
      <p>&copy; 2026 恶雨三角洲行动论坛 - EyuDeltaForceForum</p>
    </footer>
  </div>
</template>

<style scoped>
.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.app-header {
  background: linear-gradient(135deg, #0f3460 0%, #16213e 100%);
  padding: 1rem 0;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.app-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #e94560;
  text-shadow: 0 0 10px rgba(233, 69, 96, 0.5);
}

.header-nav {
  display: flex;
  gap: 1.5rem;
}

.nav-link {
  color: #fff;
  text-decoration: none;
  font-size: 1rem;
  transition: all 0.3s ease;
  padding: 0.5rem 1rem;
  border-radius: 20px;
}

.nav-link:hover {
  background: rgba(233, 69, 96, 0.2);
  color: #e94560;
}

.app-main {
  flex: 1;
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
  padding: 2rem;
}

.app-footer {
  background: #16213e;
  padding: 1rem 0;
  text-align: center;
  color: #aaa;
  margin-top: 2rem;
}
</style>