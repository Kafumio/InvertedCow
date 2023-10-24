// 常量路由
export const constantRoute = [
  {
    path: '/',
    component: () => import('@/layout/index.vue'),
    name: 'root',
    redirect: '/home',
    children: [
      {
        // home
        path: '/home',
        name: 'home',
        component: () => import('@/views/home/index.vue'),
      },
      {
        path: '/signIn',
        component: () => import('@/views/sign_in/index.vue'),
        name: 'signIn',
        meta: {
          hidden: false,
        },
      },
      {
        path: '/signUp',
        component: () => import('@/views/sign_up/index.vue'),
        name: 'signUp',
        meta: {
          hidden: false,
        },
      },
    ],
  },
];

// 异步路由
export const asyncRoute = [];

// 任意路由
export const anyRoute = [
  {
    path: '/:pathMath(.*)*',
    redirect: '/404',
    name: 'Any',
    meta: {
      hidden: false,
    },
  },
];
