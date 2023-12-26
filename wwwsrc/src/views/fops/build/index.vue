<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
				<el-input size="default" placeholder="请输入应用名称" v-model="state.keyWord" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10" @click="onQuery">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button size="default" type="success" class="ml10" @click="onOpenAdd('add')">
					<el-icon>
            <ele-FolderAdd />
					</el-icon>
					添加构建
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
				<el-table-column prop="Id" label="编号" width="60" />
        <el-table-column prop="AppName" label="应用名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="ClusterId" label="集群信息" show-overflow-tooltip></el-table-column>
				<el-table-column prop="BuildNumber" label="构建号" show-overflow-tooltip></el-table-column>
				<el-table-column prop="Status" label="状态" show-overflow-tooltip></el-table-column>
				<el-table-column prop="IsSuccess" label="是否成功" show-overflow-tooltip></el-table-column>
				<el-table-column prop="CreateAt" label="开始时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="FinishAt" label="完成时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="BuildServerId" label="构建的服务端id" show-overflow-tooltip></el-table-column>
				<el-table-column label="操作" width="100">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="showLog(scope.row)">构建日志</el-button>
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
		<BuildDialog ref="buildDialogRef" @refresh="getTableData()" />


  <el-dialog title="构建日志" v-model="state.logDialogIsShow" width="700px;" height="300px;">
    <el-card shadow="hover" class="layout-padding-auto">
      <div>
        {{state.logContent}}
      </div>
    </el-card>
  </el-dialog>

  </div>
</template>

<script setup lang="ts" name="fopsBuild">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";

// 引入 api 请求接口
const serverApi = fopsApi();

// 引入组件
const BuildDialog = defineAsyncComponent(() => import('/@/views/fops/build/dialog.vue'));

// 定义变量内容
const buildDialogRef = ref();
const state = reactive({
  logDialogIsShow:false,
  logContent:'',
  keyWord:"",
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
	const data = {
    appName:state.keyWord,
    pageIndex:state.tableData.param.pageNum,
    pageSize:state.tableData.param.pageSize,
  };
  // 请求接口
  serverApi.buildList(data).then(function (res){
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
// 打开新增用户弹窗
const onOpenAdd = (type: string) => {
  buildDialogRef.value.openDialog(type);
};
// 打开修改用户弹窗
const onOpenEdit = (type: string, row: any) => {
  buildDialogRef.value.openDialog(type, row);
};

const showLog=(row:any)=>{
  serverApi.buildLog({"buildId":row.Id}).then(function (res){
      state.logContent=res
      ElMessage.success('删除成功');
      state.logDialogIsShow=true
  })
}

// 删除用户
const onRowDel = (row: any) => {
	ElMessageBox.confirm(`此操作将永久删除：“${row.Name}”，是否继续?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
      // 删除逻辑
      serverApi.gitDel({"GitId":row.Id}).then(function (res){
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
