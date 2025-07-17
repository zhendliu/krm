// 用户相关路由都放在这里
const Layout = () => import('../view/layout/Layout.vue')

const userRoutes = 
    {
        path: "/user",
        component: Layout,
        // 定义/user的子路由
        // 访问/user自动跳转到list
        redirect: "/user/list",
        children: [
            {
                path: 'add',
                component: ()=> import('../view/user/Add.vue')
            },
            {
                path: 'list',
                component: ()=> import('../view/user/List.vue')
            },
        ]
    }


export default userRoutes