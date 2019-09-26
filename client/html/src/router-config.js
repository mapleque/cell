import WebLayout from './layouts/Web'

import Page404 from './pages/Page404'
import Home from './pages/Home'
import Login from '@/pages/Login'
import Logout from '@/pages/Logout'
import Register from '@/pages/Register'
import Dashboard from '@/pages/Dashboard'
import Document from '@/pages/Document'
import Contribution from '@/pages/Contribution'
import Deployment from '@/pages/Deployment'

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
        name: 'Page Not Found',
        component: Page404,
      },
      {
        path: '/document',
        name: 'Document',
        component: Document,
      },
      {
        path: '/deployment',
        name: 'Deployment',
        component: Deployment,
      },
      {
        path: '/contribution',
        name: 'Contribution',
        component: Contribution,
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
