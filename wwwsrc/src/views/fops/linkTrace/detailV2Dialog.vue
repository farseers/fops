<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="70%">
			<el-form ref="gitDialogFormRef" size="default" label-width="100px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20" v-if="state.tableData.length>0">
            <div>应用名称：{{state.AppName}}， 应用ID：{{state.AppId}}， 应用IP：{{state.AppIp}}，整体耗时：{{state.totalTs/1000}} ms</div>
            <div class="mt10"><el-tag>{{state.Desc}}</el-tag>{{state.Caption}}<el-button style="margin-left: 20px" type="primary">查看请求</el-button></div>
            <div :style="{'width':getWidth(),'overflow-x': 'auto','white-space': 'nowrap'}">
            <ul class="custom-list mt20">
              <li style="height: 35px;">
                <span v-for="(item, index) in state.TsArray" :key="index" :style="{'width':200+'px','float':'left'}">
                  {{item.StartTs/1000}}ms
                </span>
              </li>
              <li style="clear: both" v-for="(item, index) in state.TsArray" :key="index">
                <div v-for="(info,i) in state.tableData.filter(t=>t.StartTs==item.StartTs)">
                  <span :style="{'padding-left':index*200+'px'}" :title="info.Desc"><el-tag>{{info.AppName}}</el-tag>{{info.Caption}}，耗时：{{info.UseTs/1000}}ms</span>
                  <span v-if="info.Exception!=null">异常：{{friendlyJSONstringify(info.Exception)}}</span>
                  <span v-else></span>
                </div>
              </li>
            </ul>
            </div>
					</el-col>
          <el-col style="text-align: center" v-else>
            <span style="width: 100%;">暂无数据</span>
          </el-col>
				</el-row>
			</el-form>

<!--      <showDialog ref="showDialogRef" @refresh="getTableData()" />-->
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="fopsLinkDetailV2Dialog">
import {defineAsyncComponent, reactive, ref} from 'vue';
import {fopsApi} from "/@/api/fops";
import { ElMessageBox, ElMessage } from 'element-plus';
import {friendlyJSONstringify} from "@intlify/shared";
// 引入 api 请求接口
const serverApi = fopsApi();
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
const showDialog = defineAsyncComponent(() => import('/src/views/fops/sql/httpDialog.vue'));
// 定义变量内容
const gitDialogFormRef = ref();
const showDialogRef = ref();
const state = reactive({
	ruleForm: {},
  tableData:[{
    Rgba:'',
    AppId:0,
    AppIp:'',
    AppName:'',
    StartTs:0,
    UseTs:0,
    Caption:'',
    Desc:'',
    UseDesc:'',
    Exception:'',
  }],
  TsArray:[
      {
        StartTs:0,
        DiffTs:0
      }
  ],
  TraceId:'',
  totalTs:0,
  Rgba:'',
  AppId:0,
  AppIp:'',
  AppName:'',
  Desc:'',
  Caption:'',
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
  state.dialog.title = '链路追踪详情(TraceId：'+row2.TraceIdN+')';
  //state.dialog.submitTxt = '修 改';
  //console.log(row2)
  state.TraceId=row2.TraceIdN
  // 详情
  serverApi.linkTraceInfo(row2.TraceIdN).then(function (res){
    console.log(friendlyJSONstringify(res))
    if (res.Status){
      // 绑定数据
      state.tableData=res.Data
      if (res.Data.length>0){
        state.totalTs=res.Data[0].UseTs
        state.Rgba=res.Data[0].Rgba
        state.AppId=res.Data[0].AppId
        state.AppIp=res.Data[0].AppIp
        state.AppName=res.Data[0].AppName
        state.Desc=res.Data[0].Desc
        state.Caption=res.Data[0].Caption
        state.TsArray=[]
        for (let i = 0; i < state.tableData.length; i++) {
          var item=state.tableData[i]
          var isContains=state.TsArray.filter(t=>t.StartTs==item.StartTs)
          if (isContains.length==0){
            var diff=0
            if(i>0){
              diff=item.StartTs-state.tableData[i-1].StartTs
            }
            state.TsArray.push({StartTs:item.StartTs,DiffTs:diff})
          }
        }

      }
    }
  })
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
const onShow=(row: any)=>{
  showDialogRef.value.openDialog(row);
}
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

const getWidth=()=>{
  var str="100%"
  var total=(state.TsArray.length+1)*200
  if(total>1400){
    return total+"px"
  }
  return str;
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

/* 基本样式 */
.custom-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

/* 每个列表项的样式 */
.custom-list li {
  padding: 10px 15px;
  margin-bottom: 8px;
  border-radius: 5px;
  background-color: #f2f2f2;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease;
}

/* 悬停时的样式 */
.custom-list li:hover {
  background-color: #e0e0e0;
}
/* 基本样式，根据需要自定义样式 */
.timeline {
  position: relative;
  padding: 20px 0;
}

.timeline-event {
  display: flex;
  align-items: flex-start;
  margin-bottom: 20px;
}

.timeline-event-marker {
  width: 20px;
  height: 20px;
  background-color: #333;
  border-radius: 50%;
  margin-right: 10px;
}

.timeline-event-content {
  border: 1px solid #ccc;
  padding: 10px;
  border-radius: 5px;
  background-color: #f9f9f9;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
.el-dialog__body{
  overflow:auto!important;
}
</style>
