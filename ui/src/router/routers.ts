// 常量路由
export const constantRoute = [
  {
    path: '/hello',
    component: () => import('@/views/HelloWorld.vue'),
    name: 'login',
    meta: {
      hidden: false,
    },
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
