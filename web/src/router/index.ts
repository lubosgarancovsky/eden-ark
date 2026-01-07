import { createRouter, createWebHistory } from 'vue-router'
import {LoginView, SignUpView, ForgotPasswordView, ResetPasswordView} from "@/views";


const routes = [
    { path: '/login', name: 'Login', component: LoginView },
    { path: '/signup', name: 'Sign up', component: SignUpView },
    { path: '/forgot-password', name: 'Forgot password', component: ForgotPasswordView },
    { path: '/reset-password', name: 'Reset password', component: ResetPasswordView },
]


const router = createRouter({
    history: createWebHistory(), // HTML5 history mode
    routes
})

export default router