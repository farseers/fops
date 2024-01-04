<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="buildDialogFormRef" :model="state.ruleForm" size="default" label-width="90px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="应用名称">
							<el-input v-model="state.ruleForm.AppName" placeholder="请输入应用名称"></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="集群信息">
							<el-input v-model="state.ruleForm.ClusterId" placeholder="请输入集群信息ID" clearable></el-input>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="fopsBuildDialog">
import { reactive, ref } from 'vue';
import {fopsApi} from "/@/api/fops";
import { ElMessageBox, ElMessage } from 'element-plus';
// 引入 api 请求接口
const serverApi = fopsApi();
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
// 定义变量内容
const buildDialogFormRef = ref();
const state = reactive({
	ruleForm: {
    AppName: '', // 应用名称
    ClusterId: 0, // 集群信息
	},
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: any) => {
  state.dialog.type=type
	if (type === 'edit') {
		state.ruleForm = row;
		state.dialog.title = '修改Git';
		state.dialog.submitTxt = '修 改';
    // 绑定数据
    state.ruleForm.AppName=row.AppName
    state.ruleForm.ClusterId=row.ClusterId

	} else {
		state.dialog.title = '新增Git';
		state.dialog.submitTxt = '新 增';

    state.ruleForm.AppName=''
    state.ruleForm.ClusterId=0
		// 清空表单，此项需加表单验证才能使用
		// nextTick(() => {
		// 	gitDialogFormRef.value.resetFields();
		// });
	}
	state.dialog.isShowDialog = true;
	getMenuData();
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = () => {
  // 提交数据
  var param={
    "AppName":state.ruleForm.AppName,
    "ClusterId":state.ruleForm.ClusterId,
  }

	if (state.dialog.type === 'add') {
    serverApi.buildAdd(param).then(async function(res){
      if(res.Status){
        ElMessage.success("添加成功")
        closeDialog();
        emit('refresh');
      }else{
        ElMessage.error(res.StatusMessage)
      }
    })
  }else if (state.dialog.type=='edit'){
    serverApi.gitEdit(param).then(async function(res){
      if(res.Status){
        ElMessage.success("修改成功")
        closeDialog();
        emit('refresh');
      }else{
        ElMessage.error(res.StatusMessage)
      }
    })
  }
};
// 初始化部门数据
const getMenuData = () => {
};

// 暴露变量
defineExpose({
	openDialog,
  onSubmit
});
</script>
