<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15" style="display: none">
				<el-input size="default" v-model="state.keyWord" placeholder="请输入任务名称" style="max-width: 180px"> </el-input>
        <el-select v-model="state.enable" placeholder="请选择运行状态" clearable class="ml10">
          <el-option label="全部" value="-1"></el-option>
          <el-option label="停止" value="0"></el-option>
          <el-option label="运行中" value="1"></el-option>
        </el-select>
        <el-select v-model="state.taskStatus" placeholder="请选择调度状态" clearable class="ml10">
          <el-option label="全部" value="-1"></el-option>
          <el-option label="未开始" value="0"></el-option>
          <el-option label="调度中" value="1"></el-option>
          <el-option label="调度失败" value="2"></el-option>
          <el-option label="执行中" value="3"></el-option>
          <el-option label="失败" value="4"></el-option>
          <el-option label="成功" value="5"></el-option>
        </el-select>
        <el-input size="default" v-model="state.clientId" placeholder="请输入客户端ID" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
				<el-table-column prop="Id" label="序号" width="200" />
				<el-table-column prop="Name" label="服务器节点名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="Ip" label="IP" show-overflow-tooltip></el-table-column>
				<el-table-column prop="Port" label="端口" show-overflow-tooltip></el-table-column>
				<el-table-column prop="IsLeader" label="是否主节点" show-overflow-tooltip></el-table-column>
        <el-table-column prop="ActivateAt" label="激活时间" show-overflow-tooltip></el-table-column>
<!--				<el-table-column label="操作" width="100" style="display: none">-->
<!--					<template #default="scope">-->
<!--						<el-button size="small" text type="primary" @click="onDetail(scope.row)">详情信息</el-button>-->
<!--            <el-button size="small" text type="primary" @click="onEdit('edit',scope.row)">修改</el-button>-->
<!--            <el-button size="small" text type="primary" @click="onDel(scope.row)">删除</el-button>-->
<!--					</template>-->
<!--				</el-table-column>-->
			</el-table>
<!--			<el-pagination-->
<!--				@size-change="onHandleSizeChange"-->
<!--				@current-change="onHandleCurrentChange"-->
<!--				class="mt15"-->
<!--				:pager-count="5"-->
<!--				:page-sizes="[10, 20, 30]"-->
<!--				v-model:current-page="state.tableData.param.pageNum"-->
<!--				background-->
<!--				v-model:page-size="state.tableData.param.pageSize"-->
<!--				layout="total, sizes, prev, pager, next, jumper"-->
<!--				:total="state.tableData.total"-->
<!--			>-->
<!--			</el-pagination>-->
		</el-card>
    <editDialog ref="editDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="fopsTaskTimeOut">
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
  enable:-1,
  taskStatus:-1,
  clientId:'',
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

	var param={
    //name:state.keyWord,
    //enable:state.enable,
    //taskStatus:state.taskStatus,
    //clientId:state.clientId,
    // pageSize:state.tableData.param.pageSize,
    // pageIndex:state.tableData.param.pageNum
  }
  // 请求接口
  serverApi.serverNodeList(param).then(function (res){
    if (res.Status){
      state.tableData.data = res.Data;
      state.tableData.total = res.Data.length;
        state.tableData.loading = false;
    }else{
      state.tableData.data=[]
        state.tableData.loading = false;
    }

  })

};
const onDetail=(row: any)=>{

}
const onEdit=(row:any)=>{
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
