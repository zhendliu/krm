export const MENU_CONFIG = [
    // 集群资源管理的菜单配置
    {
        title: "集群资源",
        index: "/cluster",
        icon: "iconfont icon-jiqunguanli",
        items: [
         // 查看集群
         {
             index: "/cluster/list",
             title: "查看集群"
         },
        //  查看节点
        {
            index: "/node/list",
            title: "查看节点"
        },
        {
            index: "/namespace/list",
            title: "命名空间"
        },
        ]
    },
    // 调度资源管理的菜单配置
    {
        title: "调度资源",
        index: "scheduling",
        icon: "iconfont icon-tiaodu",
        subMenu: [
            // 先分子菜单
            {
                title: "Pod",
                index: "pod",
                icon: "iconfont icon-Pod",
                items: [
                    {
                        index: "/pod/list",
                        title: "查看"
                    },
                ]
            },
            // 有一个子分类
            {
                title: "Deployment",
                index: "deployment",
                icon: "iconfont icon-Deployment",
                items: [
                   // 添加产品
                    {
                        index: "/deployment/list",
                        title: "查看"
                    },
                    // 查看产品
                    {
                        index: "/deployment/create",
                        title: "创建"
                    } 
                ]
            },
            // 有一个子分类
            {
                title: "StatefulSet",
                index: "statefulset",
                icon: "iconfont icon-StatefulSet",
                items: [
                   // 添加产品
                    {
                        index: "/statefulset/list",
                        title: "查看"
                    },
                    // 查看产品
                    {
                        index: "/statefulset/create",
                        title: "创建"
                    } 
                ]
            },
            // 有一个子分类
            {
                title: "DaemonSet",
                index: "daemonset",
                icon: "iconfont icon-DaemonSet",
                items: [
                   // 添加产品
                    {
                        index: "/daemonset/list",
                        title: "查看"
                    },
                    // 查看产品
                    {
                        index: "/daemonset/create",
                        title: "创建"
                    } 
                ]
            },
            // 有一个子分类
            {
                title: "CronJob",
                index: "cronjob",
                icon: "iconfont icon-jihuarenwu",
                items: [
                   // 添加产品
                    {
                        index: "/cronjob/list",
                        title: "查看"
                    },
                    // 查看产品
                    {
                        index: "/cronjob/create",
                        title: "创建"
                    } 
                ]
            },
        ]
    },
    // 用户管理
    // {
    //    title: "用户管理",
    //    index: "/user",
    //    icon: "iconfont icon-yonghuguanli",
    //    items: [
    //     // 查看用户
    //     {
    //         index: "/user/list",
    //         title: "查看用户"
    //     },
    //    ]
    // },
    // 产品管理
    // {
    //     title: "产品管理",
    //     index: "/product",
    //     icon: "iconfont icon-chanpinguanli",
    //     subMenu: [
    //         // 先分子菜单
    //         {
    //             title: "水产品",
    //             index: "/product/aquatic",
    //             icon: "iconfont icon-yangzhiqujiancejinyong",
    //             items: [
    //                // 添加产品
    //                 {
    //                     index: "/user/aquatic/add",
    //                     title: "添加商品"
    //                 },
    //                 // 查看产品
    //                 {
    //                     index: "/user/aquatic/list",
    //                     title: "查看列表"
    //                 } 
    //             ]
    //         },
    //         // 有一个子分类
    //         {
    //             title: "电子产品",
    //             index: "/product/elec",
    //             icon: "iconfont icon-huiyishi",
    //             items: [
    //                // 添加产品
    //                 {
    //                     index: "/user/elec/add",
    //                     title: "添加商品"
    //                 },
    //                 // 查看产品
    //                 {
    //                     index: "/user/elec/list",
    //                     title: "查看列表"
    //                 } 
    //             ]
    //         }
    //     ]
    // },
    // 订单管理
    // {
    //     title: "订单管理",
    //     index: "/order",
    //     icon: "iconfont icon-dingdanguanli",
    //     subMenu: [
    //         // 先分子菜单
    //         {
    //             title: "水产品",
    //             index: "/order/aquatic",
    //             icon: "iconfont icon-yangzhiqujiancejinyong",
    //             items: [
    //                 {
    //                     index: "/order/aquatic/list",
    //                     title: "查看列表"
    //                 } 
    //             ]
    //         },
    //         // 有一个子分类
    //         {
    //             title: "电子产品",
    //             index: "/order/elec",
    //             icon: "iconfont icon-huiyishi",
    //             items: [
    //                 // 查看列表
    //                 {
    //                     index: "/order/elec/list",
    //                     title: "查看列表"
    //                 } 
    //             ]
    //         }
    //     ]
    // },
]