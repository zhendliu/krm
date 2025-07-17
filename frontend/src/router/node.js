// 用户相关路由都放在这里
const Layout = () => import('../view/layout/Layout.vue')

const nodeRoutes = 
    {
        path: "/node",
        component: Layout,
        redirect: "/node/list",
        children: [
            {
                path: 'list',
                component: ()=> import('../view/node/List.vue')
            },
        ]
    }


export default nodeRoutes