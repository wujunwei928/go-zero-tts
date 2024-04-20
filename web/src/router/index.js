
/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import index from '@/pages/index.vue'
import tts from '@/pages/tts.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/pages/index.vue'),
    },
    {
      path: '/tts',
      name: 'tts',
      component: () => import('@/pages/tts.vue'),
    },
    // {
    //   path: '/:pathMatch(.*)*',
    //   name: 'NotFound',
    //   component: () => import('@/pages/NotFound.vue'),
    // },
  ],
})

export default router
