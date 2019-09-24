import WebLayout from './layouts/Web'
import AdminLayout from './layouts/Admin'

import Page404 from './pages/Page404'
import Home from './pages/Home'
import Login from '@/pages/Login'
import Register from '@/pages/Register'

export default [
  {
    path: '/404',
    component: Page404,
    hideInMenu: true,
  },
  {
    path: '/',
    component: WebLayout,
    routes: [
      {
        path: '/',
        name: 'Wellcome',
        component: Home,
      },
      {
        path: '/login',
        name: 'Login',
        component: Login,
      },
      {
        path: '/register',
        name: 'Register Your Account',
        component: Register,
      },
      {
        path: '/forgot',
        name: 'Forgot Password',
        component: Page404,
      },
      {
        path: '/reset',
        name: 'Reset Password',
        component: Page404,
      },
    ],
  },
  {
    path: '/admin',
    component: AdminLayout,
    routes: [
      {
        path: '/admin',
        name: 'Dashboard',
        component: Page404,
      },
    ],
  },
]
