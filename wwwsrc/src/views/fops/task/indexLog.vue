<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
				<el-input size="default" v-model="state.taskGroupId" placeholder="请输入任务组ID" style="max-width: 180px"> </el-input>
        <el-select v-model="state.logLevel" placeholder="请选择日志等级" clearable class="ml10">
          <el-option label="全部" :value="-1"></el-option>
          <el-option label="Trace" :value="0"></el-option>
          <el-option label="Debug" :value="1"></el-option>
          <el-option label="Info" :value="2"></el-option>
          <el-option label="Warning" :value="3"></el-option>
          <el-option label="Error" :value="4"></el-option>
          <el-option label="Critical" :value="5"></el-option>
        </el-select>
<!--        <el-select v-model="state.taskStatus" placeholder="请选择调度状态" clearable class="w100">-->
<!--          <el-option label="全部" value="-1"></el-option>-->
<!--          <el-option label="未开始" value="0"></el-option>-->
<!--          <el-option label="调度中" value="1"></el-option>-->
<!--          <el-option label="调度失败" value="2"></el-option>-->
<!--          <el-option label="执行中" value="3"></el-option>-->
<!--          <el-option label="失败" value="4"></el-option>-->
<!--          <el-option label="成功" value="5"></el-option>-->
<!--        </el-select>-->
<!--        <el-input size="default" v-model="state.clientId" placeholder="请输入客户端ID" style="max-width: 180px"> </el-input>-->
				<el-button size="default" type="primary" class="ml10" @click="onQuery">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
				<el-table-column prop="Id" label="序号" width="180" />
<!--				<el-table-column prop="TaskGroupId" label="任务组ID" show-overflow-tooltip></el-table-column>-->
				<el-table-column prop="Name" label="实现Job的特性名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="Caption" label="任务组标题" show-overflow-tooltip></el-table-column>
				<el-table-column prop="Ver" label="版本" show-overflow-tooltip></el-table-column>
        <el-table-column prop="TaskId" label="任务ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="LogLevel" label="日志级别" show-overflow-tooltip>
          <template #default="scope">
            <el-tag v-if="scope.row.LogLevel==0">Trace</el-tag>
            <el-tag v-if="scope.row.LogLevel==1">Debug</el-tag>
            <el-tag v-if="scope.row.LogLevel==2">Information</el-tag>
            <el-tag v-if="scope.row.LogLevel==3">Warning</el-tag>
            <el-tag v-if="scope.row.LogLevel==4">Error</el-tag>
            <el-tag v-if="scope.row.LogLevel==5">Critical</el-tag>
            <el-tag v-if="scope.row.LogLevel==6">NoneLevel</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="Content" label="日志内容" show-overflow-tooltip></el-table-column>
        <el-table-column prop="CreateAt" label="日志时间" show-overflow-tooltip></el-table-column>
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

<script setup lang="ts" name="fopsTaskRunning">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
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
  taskGroupId:'',
  logLevel:-1,
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

  const params = new URLSearchParams();
  params.append('logLevel', state.logLevel.toString());
  params.append('taskGroupId', state.taskGroupId.toString());
  params.append('pageSize', state.tableData.param.pageSize.toString());
  params.append('pageIndex', state.tableData.param.pageNum.toString());

  // 请求接口
  serverApi.taskLogList(params.toString()).then(function (res){
    if (res.Status){
      state.tableData.data = res.Data.List;
      state.tableData.total = res.Data.RecordCount;
        state.tableData.loading = false;
    }else{
      state.tableData.data=[]
        state.tableData.loading = false;
    }

  })

};
const onDetail=(row: any)=>{

}
const onEdit=(type: string,row:any)=>{
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
const onQuery=()=>{
  getTableData();
}
// 页面加载时
onMounted(() => {
	getTableData();
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
