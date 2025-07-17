// 用户相关路由都放在这里
const Layout = () => import('../view/layout/Layout.vue')

export const podRoutes = 
    {
        path: "/pod",
        component: Layout,
        redirect: "/pod/list",
        children: [
            {
                path: 'list',
                component: ()=> import('../view/pod/List.vue')
            },
        ]
    }
export const deploymentRoutes = 
    {
        path: "/deployment",
        component: Layout,
        redirect: "/deployment/list",
        children: [
            {
                path: 'list',
                component: ()=> import('../view/deployment/List.vue')
            },
            {
                path: 'create',
                component: ()=> import('../view/deployment/Create.vue')
            }
        ]
    }
export const statefulsetRoutes = 
    {
        path: "/statefulset",
        component: Layout,
        redirect: "/statefulset/list",
        children: [
            {
                path: 'list',
                component: ()=> import('../view/statefulset/List.vue')
            },
            {
                path: 'create',
                component: ()=> import('../view/statefulset/Create.vue')
            }
        ]
    }
export const daemonsetRoutes = 
    {
        path: "/daemonset",
        component: Layout,
        redirect: "/daemonset/list",
        children: [
            {
                path: 'list',
                component: ()=> import('../view/daemonset/List.vue')
            },
            {
                path: 'create',
                component: ()=> import('../view/daemonset/Create.vue')
            }
        ]
    }
export const cronjobRoutes = 
    {
        path: "/cronjob",
        component: Layout,
        redirect: "/cronjob/list",
        children: [
            {
                path: 'list',
                component: ()=> import('../view/cronjob/List.vue')
            },
            {
                path: 'create',
                component: ()=> import('../view/cronjob/Create.vue')
            }
        ]
    }