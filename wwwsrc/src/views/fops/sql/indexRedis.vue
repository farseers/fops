<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
        <label>应用名称</label>
        <el-input size="default" v-model="state.appName" placeholder="应用名称" style="max-width: 180px"> </el-input>
        <label class="ml10">执行端IP</label>
        <el-input size="default" v-model="state.appIp" placeholder="执行端IP" style="max-width: 180px"> </el-input>
        <label class="ml10">KEY</label>
        <el-input size="default" v-model="state.key" placeholder="KEY" style="max-width: 180px"> </el-input>
        <label class="ml10">Field</label>
        <el-input size="default" v-model="state.field" placeholder="Field" style="max-width: 180px"> </el-input>
        <label class="ml10">往前推N分钟的数据</label>
        <el-select v-model="state.startMin" placeholder="往前推N分钟的数据" clearable class="ml10">
          <el-option label="全部" value="0"></el-option>
          <el-option label="1小时耗时最高" value="60"></el-option>
          <el-option label="30分钟耗时最高" value="30"></el-option>
          <el-option label="10分钟耗时最高" value="10"></el-option>
          <el-option label="5分钟耗时最高" value="5"></el-option>
          <el-option label="1分钟耗时最高" value="1"></el-option>
        </el-select>
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
        <el-table-column prop="TraceId" label="跟踪ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="AppId" label="应用ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="AppName" label="应用名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="AppIp" label="应用IP" show-overflow-tooltip></el-table-column>
        <el-table-column prop="StartTs" label="开始时间（微秒）" show-overflow-tooltip></el-table-column>
        <el-table-column prop="EndTs" label="结束时间（微秒）" show-overflow-tooltip></el-table-column>
        <el-table-column prop="UseTs" label="耗时（微秒）" show-overflow-tooltip></el-table-column>
        <el-table-column prop="Key" label="Key" show-overflow-tooltip></el-table-column>
        <el-table-column prop="Field" label="Field" show-overflow-tooltip></el-table-column>
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
const detailDialog = defineAsyncComponent(() => import('/src/views/fops/linkTrace/detailDialog.vue'));


// 定义变量内容
const detailDialogRef = ref();
const state = reactive({
  appName:'',
  appIp:'',
  key:'',
  field:'',
  startMin:0,
  searchUseTs:0,
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

  var params={
    appName:state.appName,
    appIp:state.appIp,
    key:state.key,
    field:state.field,
    startMin:state.startMin,
    searchUseTs:state.searchUseTs,
    pageSize:state.tableData.param.pageSize,
    pageIndex:state.tableData.param.pageNum,
  }
  // 请求接口
  serverApi.slowHttp(params).then(function (res){
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
  detailDialog.value.openDialog(row);
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
