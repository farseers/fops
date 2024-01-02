<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="gitDialogFormRef" :model="state.ruleForm" size="default" label-width="100px">
				<el-row :gutter="35">
          <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="日志ID">
              <el-input v-model="state.ruleForm.LogIdN" placeholder="日志ID" clearable></el-input>
            </el-form-item>
          </el-col>
					<el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="跟踪ID">
							<el-input v-model="state.ruleForm.TraceIdN" placeholder="跟踪ID" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="应用ID">
							<el-input v-model="state.ruleForm.AppId" placeholder="应用ID" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="应用名称">
							<el-input v-model="state.ruleForm.AppName" placeholder="应用名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="应用IP">
							<el-input v-model="state.ruleForm.AppIp" placeholder="应用IP" clearable></el-input>
						</el-form-item>
					</el-col>
          <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="日志类型">
              <el-input v-model="state.ruleForm.LogLevel" placeholder="日志类型" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="内容">
              <span>{{state.ruleForm.Content}}</span>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="组件">
              <el-tag size="mini">{{state.ruleForm.Component}}</el-tag>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="创建时间">
              <el-input v-model="state.ruleForm.CreateAt" placeholder="创建时间" clearable></el-input>
            </el-form-item>
          </el-col>
				</el-row>
			</el-form>

		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="fopsLogDetailDialog">
import { reactive, ref } from 'vue';
import {fopsApi} from "/@/api/fops";
import { ElMessageBox, ElMessage } from 'element-plus';
import {friendlyJSONstringify} from "@intlify/shared";
// 引入 api 请求接口
const serverApi = fopsApi();
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const gitDialogFormRef = ref();
const state = reactive({
	ruleForm: {
    TraceIdN:'',
    AppId:0,
    AppName:'',
    AppIp:'',
    LogIdN:'',
    Content:'',
    Component:'',
    LogLevel:'',
    CreateAt:'',
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
  state.dialog.title = '日志详情';
  //state.dialog.submitTxt = '修 改';
  // 详情
  serverApi.logInfo(row2.LogIdN).then(function (res){
    if (res.Status){
      // 绑定数据
      state.ruleForm=res.Data
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
// 取消
const onCancel = () => {
	closeDialog();
};

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
