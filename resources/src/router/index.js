import { createRouter, createWebHashHistory } from 'vue-router'
import helpers from '../utils/helpers'
import Cookies from 'js-cookie'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
    {
        path: '/',
        name: "Index",
        component: () => import('../views/Index.vue'),
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/Login.vue'),
    },
    {
        path: '/test',
        name: 'Test',
        component: () => import('../views/Test.vue'),
    },
    {
        path: '/color',
        name: 'Color',
        component: () => import('../views/Color.vue'),
    },
    {
        path: '/hospital',
        component: () => import('../views/Layout.vue'),
        children: [
        {
            path: '',
            name: "Hospital",
            component: () => import('../views/hospital/Index.vue')
        },
        {
            path: 'empty',
            component: () => import('../views/hospital/Empty.vue')
        },
        ]
    },
    ]
})

router.beforeEach((to, from, next) => {
    console.log('路由拦截')
    var user_token = helpers.getUrlKey('user_token')
    if (!(!user_token && typeof(user_token)!="undefined" && user_token!=0)){
        localStorage.setItem('user_token',user_token)
        Cookies.set('user_token',user_token)
        parent.location.reload();
        window.location.href="/"
    }
    var isAuthenticated = false

    user_token = localStorage.getItem('user_token')
    if (!user_token && typeof(user_token)!="undefined" && user_token!=0) {
        user_token = Cookies.get('user_token')
    }


    console.log('beforeEach_user_token',user_token)
    if(!(typeof user_token === 'undefined' || user_token === null || user_token === "")){
        isAuthenticated = true
    }
    if (to.name !== 'Login' && !isAuthenticated){
        next({ name: 'Login' })
    } else{
        if (to.name == "Login" && isAuthenticated) {
            next({name:"Index"})
        }else{
            next()
        }     
    }
})

export default router