import { createRouter, createWebHistory } from 'vue-router';
import useGlobalStore from './store/global';
import { ISpaceDetail } from '../types/index';

const routes = [
  {
    path: '/',
    name: 'home',
    redirect: () => {
      // 访问首页，默认调到服务管理列表页
      const { spaceList } = useGlobalStore();
      const firstHasPermSpace = spaceList.find((item: ISpaceDetail) => item.permission);
      const spaceId = firstHasPermSpace ? firstHasPermSpace.space_id : spaceList[0]?.space_id;
      return { name: 'service-mine', params: { spaceId } };
    },
  },
  {
    path: '/space/:spaceId',
    name: 'space',
    component: () => import('./views/space/index.vue'),
    children: [
      {
        path: 'service',
        children: [
          {
            path: 'mine',
            name: 'service-mine',
            meta: {
              navModule: 'service',
            },
            component: () => import('./views/space/service/list/index.vue'),
          },
          {
            path: 'all',
            name: 'service-all',
            meta: {
              navModule: 'service',
            },
            component: () => import('./views/space/service/list/index.vue'),
          },
          {
            path: ':appId(\\d+)',
            component: () => import('./views/space/service/detail/index.vue'),
            children: [
              {
                path: 'config',
                name: 'service-config',
                meta: {
                  navModule: 'service',
                },
                component: () => import('./views/space/service/detail/config/index.vue'),
              },
              {
                path: 'script',
                name: 'init-script',
                meta: {
                  navModule: 'service',
                },
                component: () => import('./views/space/service/detail/init-script/index.vue'),
              },
            ],
          },
        ],
      },
      {
        path: 'groups',
        name: 'groups-management',
        meta: {
          navModule: 'groups',
        },
        component: () => import('./views/space/groups/index.vue'),
      },
      {
        path: 'variables',
        name: 'variables-management',
        meta: {
          navModule: 'variables',
        },
        component: () => import('./views/space/variables/index.vue'),
      },
      {
        path: 'templates',
        meta: {
          navModule: 'templates',
        },
        children: [
          {
            path: 'list/:templateSpaceId?/:packageId?',
            name: 'templates-list',
            meta: {
              navModule: 'templates',
            },
            component: () => import('./views/space/templates/list/index.vue'),
          },
          {
            path: ':templateSpaceId/:packageId/version_manage/:templateId',
            name: 'template-version-manage',
            meta: {
              navModule: 'templates',
            },
            component: () => import('./views/space/templates/version-manage/index.vue'),
          },
        ],
      },
      {
        path: 'scripts',
        name: 'scripts-management',
        meta: {
          navModule: 'scripts',
        },
        component: () => import('./views/space/scripts/index.vue'),
        children: [
          {
            path: 'list',
            name: 'script-list',
            meta: {
              navModule: 'scripts',
            },
            component: () => import('./views/space/scripts/list/script-list.vue'),
          },
          {
            path: 'version_manage/:scriptId',
            name: 'script-version-manage',
            meta: {
              navModule: 'scripts',
            },
            component: () => import('./views/space/scripts/version-manage/index.vue'),
          },
        ],
      },
      {
        path: 'credentials',
        name: 'credentials-management',
        meta: {
          navModule: 'credentials',
        },
        component: () => import('./views/space/credentials/index.vue'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: () => import('./views/404.vue'),
  },
];

const router = createRouter({
  history: createWebHistory((window as any).SITE_URL),
  routes,
});

// 路由切换时，取消无权限页面
router.afterEach(() => {
  const globalStore = useGlobalStore();
  globalStore.$patch((state) => {
    state.showPermApplyPage = false;
  });
});

export default router;
