import {createRouter,createWebHashHistory} from 'vue-router'
import {CONFIG} from '../config/index.js'
// import userRoutes from './user.js'
import clusterRoutes from './cluster.js'
import nodeRoutes from './node.js'
import namespaceRoutes from './namespace.js'
import {podRoutes, deploymentRoutes, statefulsetRoutes, daemonsetRoutes, cronjobRoutes} from './scheduling.js'
// const User = () => import('../views/User.vue')
const Login = () => import('../view/Login.vue')
// const Index = () => import('../view/Index.vue')
const Layout = () => import('../view/layout/Layout.vue')

// 定义用户相关的路由映射
// const userRoutes = []

// 定义路由映射
const routes = [
    // userRoutes,
    clusterRoutes,
    nodeRoutes,
    namespaceRoutes,
    podRoutes,
    deploymentRoutes,
    statefulsetRoutes,cronjobRoutes,daemonsetRoutes,
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/",
        component: Layout
    },
]

// 创建路由实例
const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

// 定义一个全局的守卫，去判断请求链接有灭有token字段
router.beforeEach(
  (to, from, next) => {
    // 1. 访问的是login，没有携带token ==> next()
    // 2. 访问的还是login, 携带了token ==> next("/")
    // 3. 访问的不是login，但是携带了token ==> next()
    // 4. 访问的不是login，没有携带token ==> next("/login")
    // 拿到访问路径
    const toPath = to.path
    const toLogin = toPath.indexOf("/login") // 0  没有找到-1
    // 判断本地是否有token
    const tokenStatus = window.localStorage.getItem(CONFIG.TOKEN_NAME)
    if (toLogin == 0 && tokenStatus) {
        next("/")
    } else if (toLogin == 0 && !tokenStatus) {
        // 放行
        next()
    } else if (tokenStatus) {
        // 放行
        next()
    } else {
        next("/login")
    }
  }
)
// router.push("/xxxx")
export default router