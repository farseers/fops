<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="gitDialogFormRef" :model="state.ruleForm" size="default" label-width="100px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="任务名称">
							<el-input v-model="state.ruleForm.Name" placeholder="请输入任务名称" readonly></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="版本">
							<el-input v-model="state.ruleForm.Ver" placeholder="请输入版本" readonly></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="任务标题">
							<el-input v-model="state.ruleForm.Caption" placeholder="请输入任务标题" readonly></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="参数">
							<el-input v-model="state.ruleForm.Data.lotteryName" placeholder="请输入参数" readonly></el-input>
						</el-form-item>
					</el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="开始时间">
              <el-input v-model="state.ruleForm.StartAt" placeholder="开始时间" readonly></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="下次执行时间">
              <el-input v-model="state.ruleForm.NextAt" placeholder="下次执行时间" readonly></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="设置Cron">
              <el-input v-model="state.ruleForm.Cron" placeholder="请输入Cron时间格式" readonly></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="是否启用">
<!--              <el-table-column prop="status" label="部门状态" show-overflow-tooltip>-->
<!--              </el-table-column>-->
                <template #default="scope">
                  <el-tag type="success" v-if="state.ruleForm.IsEnable">启用</el-tag>
                  <el-tag type="info" v-else>禁用</el-tag>
                </template>
            </el-form-item>
          </el-col>

          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="激活时间">
              <el-input v-model="state.ruleForm.ActivateAt" placeholder="激活时间" readonly></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="最后运行时间">
              <el-input v-model="state.ruleForm.LastRunAt" placeholder="最后一次运行时间" readonly></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="平均时间">
              <el-input v-model="state.ruleForm.RunSpeedAvg" placeholder="请输入Cron时间格式" readonly></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="运行次数">
              <el-input v-model="state.ruleForm.RunCount" placeholder="请输入Cron时间格式" readonly></el-input>
            </el-form-item>
          </el-col>
<!--          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">-->
<!--            <el-form-item label="是否保存">-->
<!--              <template #default="scope">-->
<!--                <el-tag type="success" v-if="state.ruleForm.NeedSave">保存</el-tag>-->
<!--                <el-tag type="info" v-else>不保存</el-tag>-->
<!--              </template>-->
<!--            </el-form-item>-->
<!--          </el-col>-->
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="任务内容" >
              <el-tag size="small">ID：{{state.ruleForm.Task.Id}}</el-tag>
              <el-tag size="small">名称：{{state.ruleForm.Task.Name}}</el-tag>
              <el-tag size="small">版本：{{state.ruleForm.Task.Ver}}</el-tag>
              <el-tag size="small">标题：{{state.ruleForm.Task.Caption}}</el-tag>
            </el-form-item>
            <el-form-item label="运行数据" >
              <el-tag size="small">参数：{{JSON.stringify(state.ruleForm.Task.Data) }}</el-tag>
            </el-form-item>
            <el-form-item label="运行数据" >
              <el-tag size="small">运行耗时：{{state.ruleForm.Task.RunSpeed}}</el-tag>
              <el-tag size="small">进度：{{state.ruleForm.Task.Progress}}</el-tag>
            </el-form-item>
            <el-form-item label="任务时间" >
              <el-tag size="small">开始时间：{{state.ruleForm.Task.StartAt}}</el-tag>
              <el-tag size="small">运行时间：{{state.ruleForm.Task.RunAt}}</el-tag>
              <el-tag size="small">调度时间：{{state.ruleForm.Task.SchedulerAt}}</el-tag>
              <el-tag size="small">创建时间：{{state.ruleForm.Task.CreateAt}}</el-tag>
            </el-form-item>
            <el-form-item label="任务状态" >
              <el-tag size="small">状态：{{getStatusDesc(state.ruleForm.Task.Status)}}</el-tag>
            </el-form-item>
            <el-form-item label="客户端信息" >
              <el-tag size="small">客户端ID：{{state.ruleForm.Task.Client.Id}}</el-tag>
              <el-tag size="small">客户端名称：{{state.ruleForm.Task.Client.Name}}</el-tag>
              <el-tag size="small">客户端IP：{{state.ruleForm.Task.Client.Ip}}</el-tag>
              <el-tag size="small">客户端端口：{{state.ruleForm.Task.Client.Port}}</el-tag>
            </el-form-item>
          </el-col>
				</el-row>
			</el-form>

		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="fopsTaskDetailDialog">
import { reactive, ref } from 'vue';
import {fopsApi} from "/@/api/fops";
import { ElMessageBox, ElMessage } from 'element-plus';
// 引入 api 请求接口
const serverApi = fopsApi();
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const gitDialogFormRef = ref();
const state = reactive({
	ruleForm: {
    Name: '', // 任务名称
    Ver: 1, // 版本
    Caption: '', // 任务标题
    Data: {
      lotteryName:''
    }, // 仓库名称
    Task: {
      Id: 0,
      TaskGroupId: 0,
      Name: "",
      Ver: 0,
      Caption: "",
      StartAt: "",
      RunAt: "",
      RunSpeed: 0,
      Client: {
        Id: 0,
        Name: "",
        Ip: "",
        Port: 0
      },
      Progress: 0,
      Status: 0,
      SchedulerAt: "",
      Data: {},
      CreateAt: ""
    },
    StartAt: '', // 开始时间
    NextAt: '', // 下次执行时间
    Cron: '', // 时间
    IsEnable: false, // 是否启用
    ActivateAt:'',
    LastRunAt:'',
    RunSpeedAvg:0,
    RunCount:0,
    NeedSave:false,
	},
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (row2: any) => {
  //state.ruleForm = row;
  state.dialog.title = '任务组详情';
  state.dialog.submitTxt = '修 改';
  // 详情
  var url='/basicapi/taskGroup/info-'+row2.Name
  serverApi.taskGroupInfo(url).then(function (res){
    if (res.Status){
      var row=res.Data
      // 绑定数据
      state.ruleForm.Name=row.Name
      state.ruleForm.Ver=row.Ver
      state.ruleForm.Caption=row.Caption
      state.ruleForm.Data=row.Data
      state.ruleForm.StartAt=row.StartAt
      state.ruleForm.NextAt=row.NextAt
      state.ruleForm.Cron=row.Cron
      state.ruleForm.IsEnable=row.IsEnable
      state.ruleForm.ActivateAt=row.ActivateAt
      state.ruleForm.LastRunAt=row.LastRunAt
      state.ruleForm.RunSpeedAvg=row.RunSpeedAvg
      state.ruleForm.RunCount=row.RunCount
      state.ruleForm.NeedSave=row.NeedSave
      state.ruleForm.Task= row.Task
    }
  })
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};

const getStatusDesc=(status:number)=>{
  switch (status){
    case 0:
      return "未开始"
    case 1:
      return "调度中"
    case 2:
      return "调度失败"
    case 3:
      return "执行中"
    case 4:
      return "失败"
    case 5:
      return "成功"
  }
  return ""
}

// 暴露变量
defineExpose({
	openDialog,
});
</script>

<style>
textarea{
  height: 220px;
}
</style>
