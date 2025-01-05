import { createRouter, createWebHistory } from 'vue-router'
import Landing from '../views/Landing.vue'
import Languages from '../views/Languages.vue'
import NotFound from '../views/NotFound.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'landing',
            component: Landing
        },
        {
            path: '/languages',
            name: 'languages',
            component: Languages
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'not-found',
            component: NotFound
        }
    ]
})

export default router