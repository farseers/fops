<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
<!--				<el-input size="default" v-model="state.keyWord" placeholder="请输入任务名称" style="max-width: 180px"> </el-input>-->
        <el-input size="default" v-model="state.taskGroupId" placeholder="请输入任务组ID" style="max-width: 180px"  class="ml10"> </el-input>
<!--        <el-select v-model="state.enable" placeholder="请选择运行状态"  class="ml10" @change="onEnableChange">-->
<!--          <el-option label="全部" :value="-1"></el-option>-->
<!--          <el-option label="停止" :value="0"></el-option>-->
<!--          <el-option label="运行中" :value="1"></el-option>-->
<!--        </el-select>-->
        <el-select v-model="state.taskStatus" placeholder="请选择调度状态" class="ml10" @change="onStatusChange">
          <el-option label="全部" :value="-1"></el-option>
          <el-option label="未开始" :value="0"></el-option>
          <el-option label="调度中" :value="1"></el-option>
          <el-option label="调度失败" :value="2"></el-option>
          <el-option label="执行中" :value="3"></el-option>
          <el-option label="失败" :value="4"></el-option>
          <el-option label="成功" :value="5"></el-option>
        </el-select>
				<el-button size="default" type="primary" class="ml10" @click="onQuery">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
				<el-table-column prop="Id" label="序号" width="60" />
				<el-table-column prop="Name" label="任务组名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="Caption" label="任务组标题" show-overflow-tooltip></el-table-column>
				<el-table-column prop="StartAt" label="开始时间" show-overflow-tooltip></el-table-column>
				<el-table-column prop="RunAt" label="运行时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="RunSpeed" label="运行耗时" show-overflow-tooltip></el-table-column>
        <el-table-column prop="Progress" label="进度" show-overflow-tooltip></el-table-column>
        <el-table-column label="任务状态" show-overflow-tooltip>
          <template #default="scope">
            <el-tag v-if="scope.row.Status==0">未开始</el-tag>
            <el-tag v-if="scope.row.Status==1">调度中</el-tag>
            <el-tag v-if="scope.row.Status==2">调度失败</el-tag>
            <el-tag v-if="scope.row.Status==3">执行中</el-tag>
            <el-tag v-if="scope.row.Status==4">失败</el-tag>
            <el-tag v-if="scope.row.Status==5">成功</el-tag>
          </template>
        </el-table-column>
<!--				<el-table-column label="操作" width="100">-->
<!--					<template #default="scope">-->
<!--						<el-button size="small" text type="primary" @click="onDetail(scope.row)">详情信息</el-button>-->
<!--            <el-button size="small" text type="primary" @click="onEdit('edit',scope.row)">修改</el-button>-->
<!--            <el-button size="small" text type="primary" @click="onDel(scope.row)">删除</el-button>-->
<!--					</template>-->
<!--				</el-table-column>-->
			</el-table>
			<el-pagination
				@size-change="onHandleSizeChange"
				@current-change="onHandleCurrentChange"
				class="mt15"
				:pager-count="5"
				:page-sizes="[10, 20, 30]"
				v-model:current-page="state.tableData.param.pageNum"
				background
				v-model:page-size="state.tableData.param.pageSize"
				layout="total, sizes, prev, pager, next, jumper"
				:total="state.tableData.total"
			>
			</el-pagination>
		</el-card>
    <editDialog ref="editDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="fopsTask">
import { defineAsyncComponent, reactive, onMounted, ref,nextTick  } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";

// 引入 api 请求接口
const serverApi = fopsApi();

// 引入组件
const editDialog = defineAsyncComponent(() => import('/src/views/fops/task/editGroupDialog.vue'));


// 定义变量内容
const editDialogRef = ref();
const state = reactive({
  keyWord:'',
  enable:-1,
  taskStatus:-1,
  taskGroupId:'',
	tableData: {
		data: [],
		total: 0,
		loading: false,
		param: {
			pageNum: 1,
			pageSize: 10,
		},
	},
});

// 初始化表格数据
const getTableData = () => {
	state.tableData.loading = true;
  // if (state.clientId==""){
  //   state.clientId="0"
  // }

  const params = new URLSearchParams();
  params.append('taskStatus', state.taskStatus.toString());
  params.append('taskGroupId', state.taskGroupId);
  params.append('pageSize', state.tableData.param.pageSize.toString());
  params.append('pageIndex', state.tableData.param.pageNum.toString());
  // 请求接口
  serverApi.taskList(params.toString()).then(function (res){
    if (res.Status){
      state.tableData.data = res.Data.List;
      state.tableData.total = res.Data.RecordCount;
      setTimeout(() => {
        state.tableData.loading = false;
      }, 500);
    }else{
      state.tableData.data=[]
      setTimeout(() => {
        state.tableData.loading = false;
      }, 500);
    }

  })

};
const onDetail=(row: any)=>{

}
const onQuery=()=>{
  getTableData();
}
const onEdit=(type: string, row: any)=>{
  editDialogRef.value.openDialog(type, row);
}
// 删除用户
const onDel = (row: any) => {
	ElMessageBox.confirm(`此操作将永久删除：“${row.Name}”，是否继续?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
      // 删除逻辑
      serverApi.taskDel({"TaskGroupId":row.Id}).then(function (res){
        if (res.Status){
          getTableData();
          ElMessage.success('删除成功');
        }else{
          ElMessage.error(res.StatusMessage)
        }
      })
		})
		.catch(() => {});
};
// 分页改变
const onHandleSizeChange = (val: number) => {
	state.tableData.param.pageSize = val;
	getTableData();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
	state.tableData.param.pageNum = val;
	getTableData();
};
const onStatusChange=(value:number)=>{
  state.taskStatus=value
}
const onEnableChange=(value:number)=>{
  state.enable=value
}
// 页面加载时
onMounted(() => {
  // 等待下一次 DOM 更新后再执行代码
  nextTick(() => {
    getTableData();
  });
});
</script>

<style scoped lang="scss">
.system-user-container {
	:deep(.el-card__body) {
		display: flex;
		flex-direction: column;
		flex: 1;
		overflow: auto;
		.el-table {
			flex: 1;
		}
	}
}
</style>
