import yaml from 'js-yaml'
export const object2List = (obj) => {
    let list = []
    if (obj == null) {
        return list
    }
    for (let [key, value] of Object.entries(obj)) {
        const o = {
            key: key,
            value: value
        }
        // {key: "dot", value: "dotbalo1"}
        list.push(o)
    }
    return list
}
export const list2Object = (list) => {
    // console.log("接收的list:", list)
    let obj = {}
    if (list == null || list == undefined) {
        return obj
    }
    for(let i=0;i<list.length;i++) {
        const item = list[i]
        // {key: "xxx", value: "xxxx"} ==> { xxx: xxxx }
        if (item.key == "") {continue}
        obj[item.key] = item.value
    }
    return obj
}
export const deleteTableRow = (list , index) => {
    list.splice(index, 1)
}

export const obj2Yaml = (obj) => {
    console.log("需要转换的数据:", obj)
    let yamlData = null
    try {
        yamlData = yaml.dump(obj)
    } catch (error) {
        yamlData = `转换yaml失败: ${error.message}`
    }
    return yamlData
}