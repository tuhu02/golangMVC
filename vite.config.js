import { defineConfig } from 'vite'

export default defineConfig({
  css: {
    postcss: './postcss.config.js',
  },
  build: {
    outDir: 'views/dist',
    rollupOptions: {
      input: {
        main: 'views/js/main.js'
      },
      output: {
        entryFileNames: '[name].js',
        assetFileNames: '[name].[ext]'
      }
    },
    minify: 'esbuild'
  },
  server: {
    watch: {
      include: ['views/**/*']
    }
  }
})
