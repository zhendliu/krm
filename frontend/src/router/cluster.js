// 用户相关路由都放在这里
const Layout = () => import('../view/layout/Layout.vue')

const clusterRoutes = 
    {
        path: "/cluster",
        component: Layout,
        redirect: "/cluster/list",
        children: [
            {
                path: 'add',
                component: ()=> import('../view/cluster/Add.vue')
            },
            {
                path: 'list',
                component: ()=> import('../view/cluster/List.vue')
            },
        ]
    }


export default clusterRoutes