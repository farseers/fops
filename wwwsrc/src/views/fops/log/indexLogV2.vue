<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
        <label>TraceId</label>
        <el-input size="default" v-model="state.traceId" placeholder="链路ID" style="max-width: 180px"> </el-input>
        <label class="ml10">应用名称</label>
        <el-input size="default" v-model="state.appName" placeholder="应用名称" style="max-width: 150px;padding-left: 5px"> </el-input>
        <label class="ml10">执行端IP</label>
        <el-input size="default" v-model="state.appIp" placeholder="执行端IP" style="max-width: 120px;padding-left: 5px"> </el-input>
        <label class="ml10">日志内容</label>
        <el-input size="default" v-model="state.logContent" placeholder="日志内容" style="max-width: 250px;padding-left: 5px"> </el-input>
        <label class="ml10">日志类型</label>
        <el-select v-model="state.logLevel" placeholder="日志类型" class="ml10" style="max-width: 110px;" size="mini">
          <el-option label="全部" :value="-1"></el-option>
          <el-option label="Trace" :value="0"></el-option>
          <el-option label="Debug" :value="1"></el-option>
          <el-option label="Info" :value="2"></el-option>
          <el-option label="Warning" :value="3"></el-option>
          <el-option label="Error" :value="4"></el-option>
          <el-option label="Critical" :value="5"></el-option>
        </el-select>
				<el-button size="default" type="primary" class="ml10" @click="onQuery">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
        <el-table-column width="180px" label="LogID" show-overflow-tooltip>
          <template #default="scope">
            <span @click="onDetail(scope.row)">{{scope.row.LogIdN}}</span>
          </template>
        </el-table-column>
        <el-table-column width="200px" label="应用" show-overflow-tooltip>
          <template #default="scope">
            <el-tag size="mini">{{scope.row.AppName}} {{scope.row.AppIp}}</el-tag><br>
            {{scope.row.AppId}}
          </template>
        </el-table-column>
        <el-table-column width="100px" prop="LogLevel" label="日志类型" show-overflow-tooltip>
          <template #default="scope">
            <el-tag v-if="scope.row.LogLevel == 'Info'" size="mini">{{scope.row.LogLevel}}</el-tag>
            <el-tag v-else-if="scope.row.LogLevel == 'Debug'" type="info" size="mini">{{scope.row.LogLevel}}</el-tag>
            <el-tag v-else-if="scope.row.LogLevel == 'Warn'" type="warning" size="mini">{{scope.row.LogLevel}}</el-tag>
            <el-tag v-else-if="scope.row.LogLevel == 'Error'" type="danger" size="mini">{{scope.row.LogLevel}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="内容" show-overflow-tooltip>
          <template #default="scope">
            <el-tag v-if="scope.row.TraceId > 0" type="info" size="mini">TraceID {{scope.row.TraceId}}</el-tag>
            <br v-if="scope.row.TraceId > 0" type="info" size="mini" />
            {{scope.row.Content}}
          </template>
        </el-table-column>
        <el-table-column width="180px" prop="Component" label="组件" show-overflow-tooltip></el-table-column>
        <el-table-column width="180px" prop="CreateAt" label="请求时间" show-overflow-tooltip></el-table-column>
				<el-table-column label="操作" width="100">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onDetail(scope.row)">详情</el-button>
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

<script setup lang="ts" name="fopsLogList">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";
import {friendlyJSONstringify} from "@intlify/shared";

// 引入 api 请求接口
const serverApi = fopsApi();
// 引入组件
const detailDialog = defineAsyncComponent(() => import('/src/views/fops/log/detailDialog.vue'));


// 定义变量内容
const detailDialogRef = ref();

const state = reactive({
  traceId:'',
  appName:'',
  appIp:'',
  logContent:'',
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

  var data={
    appName:state.appName,
    appIp:state.appIp,
    traceId:state.traceId,
    logContent:state.logContent,
    logLevel:state.logLevel.toString(),
    pageSize:state.tableData.param.pageSize.toString(),
    pageIndex:state.tableData.param.pageNum.toString(),
  }
  const params = new URLSearchParams(data).toString();
  // 请求接口
  serverApi.logList(params).then(function (res){
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
