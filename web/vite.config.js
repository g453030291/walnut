import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    cors: true,
    proxy: {
        '/api': {
            target: 'https://chattest.top/walunt',
            changeOrigin: true,
            ws: true,
            secure: false,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
    }
  },
    // 修改编译后的输出路径dist->web
  build: {
    outDir: 'web',
  },
    base: '/web/',
})
