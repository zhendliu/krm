<script setup>
import { onBeforeMount, reactive, toRefs,ref, onMounted } from 'vue';
import { ElMessage  } from 'element-plus'
import { object2List,deleteTableRow, list2Object } from '../../utils/utils.js';
import { updateNodeHandler as updateHandler } from '../../api/node.js'
const data = reactive({
  itemForm: {
    spec: {
      taints: []
    }
  },
  labelList: [],
  // 定义污点的选项
  taintOptions: [
    {
      label: "禁止调度",
      value: "NoSchedule"
    },
    {
      label: "驱逐",
      value: "NoExecute"
    },
    {
      label: "尽量不调度",
      value: "PreferNoSchedule"
    }
  ]
})
const props = defineProps({
  itemForm: {
    type: Object
  },
  clusterId: {
    type: String,
  }
})
const emits = defineEmits(["rollback"])
const itemFormRef = ref()
let loading = ref(false)

const submit = ()=> {
  loading.value = true
  // 把labelList转成对象
  const labelObj = list2Object(data.labelList)
  // console.log("labelObj:", labelObj)
  // 把label对象赋值给itemForm
  data.itemForm.metadata.labels = labelObj
  // console.log("更新后的数据:", data.itemForm)
  updateHandler(props.clusterId, data.itemForm.metadata.name,data.itemForm).then((response)=>{
        // console.log("response:", response)
        ElMessage({
                message: response.data.message,
                type: 'success',
        })
        loading.value = false
        emits('rollback',props.clusterId)
    })
}
// 
onMounted(()=>{
  // 仍有问题，传递的仍然是内存地址
  // data.itemForm = props.itemForm
  // 把对象转成一个字符串
  // const jsonString = JSON.stringify(props.itemForm)
  // 再把字符串转成JSON，最后赋值给data.itemForm
  // data.itemForm = JSON.parse(jsonString)
  data.itemForm = JSON.parse(JSON.stringify(props.itemForm))
  // 把节点标签的对象转成list
  data.labelList = object2List(data.itemForm.metadata.labels)
  console.log("转换后的列表:", data.labelList)
})
const addLabelRow = ()=> {
  const labelRow = {
    key: "",
    value: ""
  }
  // 在列表后面添加数据
  // data.labelList.push(labelRow)
  // 在列表前面添加数据
  data.labelList.unshift(labelRow)
}
const addTaintRow = ()=> {
  const taintRow = {
    key: "",
    value: "",
    effect: "NoSchedule"
  }
  // 在列表后面添加数据
  // data.labelList.push(labelRow)
  // 在列表前面添加数据
  data.itemForm.spec.taints.unshift(taintRow)
}
const {itemForm,labelList,taintOptions} = toRefs(data)
</script>

<template>
  <el-form 
  ref="itemFormRef" 
  :model="itemForm" 
  label-width="140px"
  >
  <div>
    <el-tabs>
      <el-tab-pane label="标签配置">
        <!-- 标签配置的表格 -->
        <el-table 
        border
        :data="labelList" 
        style="width: 100%"
        height="320px"
        >
          <el-table-column align="center" prop="key" label="Key">
            <template #default="scope">
              <el-input v-model="scope.row.key"  placeholder="请输入标签的Key" />
            </template>
          </el-table-column>
          <el-table-column align="center" prop="value" label="Value">
            <template #default="scope">
              <el-input v-model="scope.row.value" placeholder="请输入标签的Value" />
            </template>
          </el-table-column>
          <el-table-column align="center" fixed="right" width="140px">
            <template #header>
              <el-button link type="primary" @click="addLabelRow">添加</el-button>
            </template>
            <template #default="scope">
              <el-button link type="warning" @click="deleteTableRow(labelList, scope.$index)">删除</el-button>
            </template>
          </el-table-column>
      </el-table>
      </el-tab-pane>
      <el-tab-pane label="污点配置">
        <el-table 
        border
        :data="itemForm.spec.taints" 
        style="width: 100%"
        height="320px"
        >
          <el-table-column align="center" prop="key" label="Key">
            <template #default="scope">
              <el-input v-model="scope.row.key"  placeholder="请输入污点的Key" />
            </template>
          </el-table-column>
          <el-table-column align="center" prop="value" label="Value">
            <template #default="scope">
              <el-input v-model="scope.row.value" placeholder="请输入污点的Value" />
            </template>
          </el-table-column>
          <el-table-column align="center" prop="effect" label="Effect">
            <template #default="scope">
              <el-select v-model="scope.row.effect" placeholder="Select">
                <el-option
                  v-for="item in taintOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="center" fixed="right" width="140px">
            <template #header>
              <el-button link type="primary" @click="addTaintRow">添加</el-button>
            </template>
            <!-- scope.$index -->
            <template #default="scope">
              <el-button link type="warning" @click="deleteTableRow(itemForm.spec.taints, scope.$index)">删除</el-button>
            </template>
          </el-table-column>
      </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
  <span style="display: block; margin-top: 30px;" class="span-button">
    <el-button @click="submit()" type="warning">
      更新
    </el-button>
  </span>
</el-form>
</template>

<style scoped>
/* from */
.form-item{
  width: 420px;
  margin: 0 auto 20px 0;
}
.cluster-info{
  display: flex;
  flex-direction: row;
}
</style>
