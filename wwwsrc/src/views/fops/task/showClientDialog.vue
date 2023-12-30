<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="gitDialogFormRef" :model="state.ruleForm" size="default" label-width="100px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
              <div>
                <span></span>
              </div>
					</el-col>
				</el-row>
			</el-form>
<!--			<template #footer>-->
<!--				<span class="dialog-footer">-->
<!--					<el-button @click="onCancel" size="default">取 消</el-button>-->
<!--					<el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>-->
<!--				</span>-->
<!--			</template>-->
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="fopsTaskEditDialog">
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
	},
  clientData:[],
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (data: any) => {
  state.dialog.title="客户端列表"
  state.clientData=data
	state.dialog.isShowDialog = true;
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

  // const originalObject = JSON.parse();
  // const modifiedObject = {};
  //
  // // 循环遍历原始对象，为每个键和值加上双引号并构建新的对象
  // for (const key in originalObject) {
  //   if (originalObject.hasOwnProperty(key)) {
  //     const quotedKey = `"${key}"`;
  //     const quotedValue = `${originalObject[key]}`;
  //     modifiedObject[quotedKey] = quotedValue;
  //   }
  // }

  // 提交数据
  var param={
    "Id":state.ruleForm.Id,
    "Name":state.ruleForm.Name,
    "Ver":state.ruleForm.Ver,
    "Caption":state.ruleForm.Caption,
    "Data":JSON.parse(state.ruleForm.Data),
    "StartAt":state.ruleForm.StartAt,
    "NextAt":state.ruleForm.NextAt,
    "Cron":state.ruleForm.Cron,
    "IsEnable":state.ruleForm.IsEnable
  }

	if (state.dialog.type === 'add') {
    // serverApi.taskUpdate(param).then(function (res){
    //   if(res.Status){
    //     ElMessage.success("添加成功")
    //     closeDialog();
    //     emit('refresh');
    //   }else{
    //     ElMessage.error(res.StatusMessage)
    //   }
    // })

  }else if (state.dialog.type=='edit'){
    serverApi.taskUpdate(param).then(function (res){
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

// 暴露变量
defineExpose({
	openDialog,
});
</script>
