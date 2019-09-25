import WebLayout from './layouts/Web'

import Page404 from './pages/Page404'
import Home from './pages/Home'
import Login from '@/pages/Login'
import Logout from '@/pages/Logout'
import Register from '@/pages/Register'
import Dashboard from '@/pages/Dashboard'

export default [
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
        path: '/404',
        component: Page404,
        hideInMenu: true,
      },
      {
        path: '/login',
        name: 'Login',
        component: Login,
      },
      {
        path: '/logout',
        name: 'Logout',
        component: Logout,
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
      {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard,
      },
    ],
  }
]
