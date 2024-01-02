<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
        <label>应用名称</label>
        <el-input size="default" v-model="state.appName" placeholder="应用名称" style="max-width: 180px"> </el-input>
        <label class="ml10">执行端IP</label>
        <el-input size="default" v-model="state.appIp" placeholder="执行端IP" style="max-width: 180px"> </el-input>
        <label class="ml10">任务名称</label>
        <el-input size="default" v-model="state.taskName" placeholder="任务名称" style="max-width: 180px"> </el-input>
        <label class="ml10">任务组ID</label>
        <el-input size="default" v-model="state.taskGroupId" placeholder="任务组ID" style="max-width: 180px"> </el-input>
        <label class="ml10">任务ID</label>
        <el-input size="default" v-model="state.taskId" placeholder="任务ID" style="max-width: 180px"> </el-input>
        <label class="ml10">执行时间大于x毫秒的记录</label>
        <el-input size="default" v-model="state.searchUseTs" placeholder="执行时间大于x毫秒的记录" style="max-width: 180px"> </el-input>

				<el-button size="default" type="primary" class="ml10" @click="onQuery">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
        <el-table-column width="180px" label="TraceID" show-overflow-tooltip>
          <template #default="scope">
            <span @click="onDetail(scope.row)">{{scope.row.TraceIdN}}</span>
          </template>
        </el-table-column>
        <el-table-column width="200px" label="应用" show-overflow-tooltip>
          <template #default="scope">
            <el-tag size="mini">{{scope.row.AppName}} {{scope.row.AppIp}}</el-tag><br>
            <el-tag size="mini">{{scope.row.AppId}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column width="120px" prop="UseDesc" label="执行耗时" show-overflow-tooltip></el-table-column>
        <el-table-column prop="TaskName" label="任务名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="TaskGroupId" label="任务组ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="TaskId" label="任务ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="WatchKey" label="WatchKey" show-overflow-tooltip></el-table-column>
        <el-table-column width="300px" label="异常" show-overflow-tooltip>
          <template #default="scope">
            <el-tag size="mini" v-if="scope.row.Exception!=null">{{scope.row.Exception.ExceptionCallFile}}:{{scope.row.Exception.ExceptionCallLine}} {{scope.row.Exception.ExceptionCallFuncName}}</el-tag><br  v-if="scope.row.Exception!=null">
            <el-tag size="mini" v-if="scope.row.Exception!=null">{{scope.row.Exception.ExceptionMessage}}</el-tag>
            <el-tag size="mini" v-else>无</el-tag>
          </template>
        </el-table-column>
        <el-table-column width="180px" prop="CreateAt" label="请求时间" show-overflow-tooltip></el-table-column>
				<el-table-column label="操作" width="100">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onDetail(scope.row)">追踪</el-button>
					</template>
				</el-table-column>
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
    <detailDialog ref="detailDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="fopsTaskRunning">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";
import {friendlyJSONstringify} from "@intlify/shared";

// 引入 api 请求接口
const serverApi = fopsApi();
// 引入组件
const detailDialog = defineAsyncComponent(() => import('/src/views/fops/linkTrace/detailV2Dialog.vue'));


// 定义变量内容
const detailDialogRef = ref();
const state = reactive({
  keyWord:'',
  appName:'',
  appIp:'',
  taskName:'',
  taskGroupId:'',
  searchUseTs:0,
  taskId:'',
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

  var data={
    appName:state.appName,
    appIp:state.appIp,
    taskName:state.taskName,
    taskGroupId:state.taskGroupId.toString(),
    searchUseTs:state.searchUseTs.toString(),
    taskId:state.taskId.toString(),
    pageSize:state.tableData.param.pageSize.toString(),
    pageIndex:state.tableData.param.pageNum.toString(),
  }
  const params = new URLSearchParams(data).toString();
  // 请求接口
  serverApi.linkTraceFScheduleList(params).then(function (res){
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
  detailDialogRef.value.openDialog(row);
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
