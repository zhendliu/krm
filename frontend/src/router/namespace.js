// 用户相关路由都放在这里
const Layout = () => import('../view/layout/Layout.vue')

const namespaceRoutes = 
    {
        path: "/namespace",
        component: Layout,
        redirect: "/namespace/list",
        children: [
            {
                path: 'add',
                component: ()=> import('../view/namespace/Add.vue')
            },
            {
                path: 'list',
                component: ()=> import('../view/namespace/List.vue')
            },
        ]
    }


export default namespaceRoutes