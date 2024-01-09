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
        <el-table-column label="名称" style="line-height: 45px;height: 45px">
          <template #default="scope">
            <div style="float: left;padding-right: 10px;padding-top: 5px">
              <el-tag size="small" cursor="cursor" @click="onIsEnable(scope.row)" v-if="scope.row.IsEnable">启用</el-tag>
              <el-tag size="small" cursor="cursor" @click="onIsEnable(scope.row)" v-else type="info">停用</el-tag>
              <!--                      <el-switch v-model="scope.row.IsEnable" @click="onIsEnable(scope.row)" inline-prompt active-text="启用" inactive-text="停用"></el-switch>-->
            </div>
            <div style="float: left;padding-right: 10px;padding-top: 5px">
              <el-tag size="small" v-if="scope.row.Task.Status==0" type="info">未开始</el-tag>
              <el-tag size="small" v-if="scope.row.Task.Status==1" type="success">调度中</el-tag>
              <el-tag size="small" style="color:red" v-if="scope.row.Task.Status==2" type="warning">调度失败</el-tag>
              <el-tag size="small" v-if="scope.row.Task.Status==3" type="success">执行中</el-tag>
              <el-tag size="small" v-if="scope.row.Task.Status==4" type="danger">失败</el-tag>
              <el-tag size="small" style="color:green" v-if="scope.row.Task.Status==5" type="success">成功</el-tag>
            </div>
            <div @click="onTaskList(scope.row)" style="float: left">
              <span>{{scope.row.Caption}}</span><br>
              <span>{{scope.row.Name}}（<span style="color:#4eb8ff">Ver:{{scope.row.Ver}}</span>）</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="StartAt" label="开始时间" width="170" show-overflow-tooltip>
          <template #default="scope">
            <span label="开始时间">{{scope.row.StartAt}}</span><br>
            <span label="上次运行时间">{{scope.row.LastRunAt}}</span>
          </template>
        </el-table-column>
        <el-table-column label="下次运行时间" width="170" show-overflow-tooltip>
          <template #default="scope">
            <span label="Cron表达式">{{scope.row.Cron}}</span><br>
            <span style="color:red" label="下次运行时间" v-if="compareTime(scope.row.NextAt)"> {{scope.row.NextAt}}</span>
            <span label="下次运行时间" v-else> {{scope.row.NextAt}}</span>
          </template>
        </el-table-column>
        <el-table-column label="运行情况" width="170" show-overflow-tooltip>
          <template #default="scope">
            <span>平均耗时: {{scope.row.RunSpeedAvg}}ms</span><br>
            <span>运行次数: {{scope.row.RunCount}}</span>
          </template>
        </el-table-column>
        <el-table-column label="数据">
          <template #default="scope">
            <span>{{friendlyJSONstringify(scope.row.Data)}}</span>
          </template>
        </el-table-column>
        <el-table-column label="客户端信息" width="180" show-overflow-tooltip>
          <template #default="scope">
            <div v-for="(item, index) in scope.row.Clients.slice(0, 3)" :key="index">
              <el-tag size="small">{{item.Name}} {{item.Ip}}:{{item.Port}}</el-tag>
              <span v-if="scope.row.Clients.length>3">更多</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button size="small" text type="primary" @click="onDetail(scope.row)">详情信息</el-button>
            <el-button size="small" text type="warning" @click="onEdit('edit',scope.row)">修改</el-button>
            <el-button size="small" text type="danger" @click="onLog(scope.row)">日志</el-button>
            <el-button size="small" text type="info" @click="onDel(scope.row)">删除</el-button>
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
    <editDialog ref="editDialogRef" @refresh="getTableData()" />
    <detailDialog ref="detailDialogRef" @refresh="getTableData()" />
    <taskDialog ref="taskDialogRef" @refresh="getTableData()" />
    <logDialog ref="logDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="fopsTaskTimeOut">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";
import {friendlyJSONstringify} from "@intlify/shared";

// 引入 api 请求接口
const serverApi = fopsApi();
// 引入组件
const editDialog = defineAsyncComponent(() => import('/src/views/fops/task/editGroupDialog.vue'));
const detailDialog = defineAsyncComponent(() => import('/src/views/fops/task/detailGroupDialog.vue'));
const taskDialog = defineAsyncComponent(() => import('/src/views/fops/task/taskDialog.vue'));
const logDialog = defineAsyncComponent(() => import('/src/views/fops/task/logDialog.vue'));


// 定义变量内容
const editDialogRef = ref();
const detailDialogRef = ref();
const taskDialogRef = ref();
const logDialogRef = ref();
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

  const params = new URLSearchParams();
  params.append('pageSize', state.tableData.param.pageSize.toString());
  params.append('pageIndex', state.tableData.param.pageNum.toString());
  // 请求接口
  serverApi.taskNoRunList(params.toString()).then(function (res){
    if (res.Status){
      state.tableData.data = res.Data.List;
      state.tableData.total = res.Data.RecordCount;
    }else{
      state.tableData.data=[]
    }
    state.tableData.loading = false;
  })
};

const compareTime=(nextAt:any)=>{
  var convertedTime = new Date(nextAt)
  return convertedTime.getTime() < new Date().getTime();
}
const onDetail=(row: any)=>{
  detailDialogRef.value.openDialog(row);
}
const onQuery=()=>{
  getTableData();
}
const onEdit=(type: string, row: any)=>{
  editDialogRef.value.openDialog(type, row);
}
const onTaskList=(row: any)=>{
  taskDialogRef.value.openDialog(row);
}
const onLog=(row: any)=>{
  logDialogRef.value.openDialog(row);
}
// 删除
const onDel = (row: any) => {
  ElMessageBox.confirm(`此操作将永久删除：“${row.Name}”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        // 删除逻辑
        serverApi.taskDel({"taskGroupName":row.Name}).then(function (res){
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
//启用停用
const onIsEnable=(row: any)=>{
  var setEnable=row.IsEnable
  var tips=""
  if(setEnable){
    setEnable=false
    tips="停用"
  }else{
    setEnable=true
    tips="启用"
  }

  ElMessageBox.confirm(`该任务即将：“${tips}”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        // 设置状态
        serverApi.taskGroupSetEnable({"taskGroupName":row.Name,"enable":setEnable}).then(function (res){
          if (res.Status){
            getTableData();
            if(setEnable){
              ElMessage.success('启用-成功');
            }else{
              ElMessage.success('停用-成功');
            }

          }else{
            ElMessage.error(res.StatusMessage)
          }
        })
      })
      .catch(() => {});
}
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

<style lang="scss">
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
.el-table tr td {
  /* 你的自定义样式 */
  padding: 0 0!important;
}
</style>
