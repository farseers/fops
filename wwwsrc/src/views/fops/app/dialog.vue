<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="900px">
			<el-form ref="gitDialogFormRef" :model="state.ruleForm" size="default" label-width="120px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="应用名称">
							<el-input v-model="state.ruleForm.AppName" placeholder="请输入应用名称" style="max-width: 200px;margin-right: 5px"></el-input>
              名称需要与应用的AppName完全一致，才能检查健康状态
              <el-button v-if="state.dialog.type=='edit'" @click="onDelete" size="default" type="danger" style="margin-left: 5px;">删 除</el-button>
						</el-form-item>
					</el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="是否健康">
              <el-tag v-if="state.ruleForm.IsHealth" size="mini" type="success">健康</el-tag>
              <el-tag v-else size="mini" type="warning">不健康</el-tag>
            </el-form-item>
          </el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="镜像版本">
              <el-tag size="mini">{{state.ruleForm.DockerVer}}</el-tag>
						</el-form-item>
					</el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="集群版本">
              <el-tag size="mini">{{friendlyJSONstringify(state.ruleForm.ClusterVer)}}</el-tag>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="18" :md="18" :lg="18" :xl="18" class="mb20">
            <el-form-item label="Git主仓库" style="float: left">
              <el-tag size="mini">{{state.ruleForm.AppGitName}}</el-tag>
<!--              <el-input v-model="state.ruleForm.AppGit"  placeholder="请输入应用的源代码" clearable></el-input>-->
            </el-form-item>
            <el-button type="primary" @click="onOpenGit(2)" size="default" style="margin-left: 5px;">添加Git</el-button>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="依赖仓库" style="float: left">
            <el-button type="success" @click="onOpenGit(1)" size="default" style="margin-left: 5px;">添加依赖的仓库</el-button>
            <el-table :data="state.gitList" style="width: 100%">
              <el-table-column prop="Id" label="编号" width="60" />
              <el-table-column prop="Name" label="Git名称" show-overflow-tooltip width="120"></el-table-column>
              <el-table-column prop="Hub" label="托管地址" show-overflow-tooltip></el-table-column>
              <el-table-column label="操作" width="100">
                <template #default="scope">
                  <el-button size="small" text type="primary" @click="delGit(scope.row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="Shell脚本" >
              <el-input v-model="state.ruleForm.ShellScript" type="textarea" placeholder="请输入Shell脚本" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="Dockerfile内容" >
              <el-input v-model="state.ruleForm.Dockerfile" type="textarea"  placeholder="请输入Dockerfile内容" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="Dockerfile">
              <el-input v-model="state.ruleForm.DockerfilePath" placeholder="请输入Dockerfile路径" clearable></el-input>
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

    <el-dialog title="Git列表" v-model="state.gitDialogIsShow" width="700px;" height="300px;">
      <el-card shadow="hover" class="layout-padding-auto">
        <div class="system-user-search mb15">
          <el-button size="default" type="success" class="ml10" @click="SureCheck()">
            <el-icon>
            </el-icon>
            确认选择
          </el-button>
        </div>
        <el-table ref="multipleTable" :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%" :row-key="getRowKey" @selection-change="handleSelectionChange">
          <el-table-column type="selection" :reserve-selection="true" width="55"></el-table-column>
          <el-table-column prop="Id"  label="编号" width="60" />
          <el-table-column prop="Name" label="Git名称" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Hub" label="托管地址" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Branch" label="Git分支" show-overflow-tooltip></el-table-column>
        </el-table>
      </el-card>
    </el-dialog>

	</div>
</template>

<script setup lang="ts" name="fopsAppDialog">
import { reactive, ref,onMounted } from 'vue';
import {fopsApi} from "/@/api/fops";
import { ElMessageBox, ElMessage } from 'element-plus';
import {friendlyJSONstringify} from "@intlify/shared";
// 引入 api 请求接口
const serverApi = fopsApi();
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const gitDialogFormRef = ref();
const multipleTable = ref();
const state = reactive({
	ruleForm: {
    AppName:'', //应用名称
    DockerVer: '', // 镜像版本
    ShellScript: '', // Shell脚本
    ClusterVer: '', // 集群版本
    AppGit: 0, // 应用的源代码
    AppGitName: '', // 应用的源代码
    FrameworkGits:[], // 依赖的框架源代码
    Dockerfile: '', // Dockerfile内容
    DockerfilePath: '', // Dockerfile路径
    IsHealth:false, // 是否健康
	},
  gitList:[],
  SelectItem:[],
  gitDialogIsShow:false,
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 10,
    },
  },
  gitType:1,
});

// 打开弹窗
const openDialog = (type: string, row: any) => {
  state.dialog.type=type
	if (type === 'edit') {
    state.ruleForm = row;
    state.dialog.title = '修改应用';
    state.dialog.submitTxt = '修 改';
    // 请求数据
    serverApi.appsDetail({AppName:row.AppName}).then(function (res){
      if (res.Status){
        row=res.Data
        // 绑定数据
        state.ruleForm.AppName=row.AppName
        state.ruleForm.DockerVer=row.DockerVer
        state.ruleForm.ShellScript=row.ShellScript
        state.ruleForm.ClusterVer=row.ClusterVer
        state.ruleForm.AppGit=row.AppGit
        state.ruleForm.FrameworkGits=row.FrameworkGits
        state.ruleForm.Dockerfile=row.Dockerfile
        state.ruleForm.DockerfilePath=row.DockerfilePath
        state.SelectItem=row.FrameworkGits
        state.ruleForm.IsHealth=row.IsHealth
        loadGitInfo(row.AppGit)
        // 加载git数据
        loadGit(row.FrameworkGits)
      }
    })
	} else {
		state.dialog.title = '新增应用';
		state.dialog.submitTxt = '新 增';
		// 清空表单，此项需加表单验证才能使用
    state.ruleForm.AppName=""
    state.ruleForm.DockerVer=""
    state.ruleForm.ShellScript=""
    state.ruleForm.ClusterVer=""
    state.ruleForm.AppGit=0
    state.ruleForm.FrameworkGits=[]
    state.ruleForm.Dockerfile=""
    state.ruleForm.DockerfilePath=""
	}
	state.dialog.isShowDialog = true;
};

const loadGit=(lst:any)=>{
  state.gitList=[]
  for (let i = 0; i < lst.length; i++) {
    console.log(lst[i])
    serverApi.gitInfo({"gitId":lst[i]}).then(function (res){
      if (res.Status){
        state.gitList.push(res.Data)
      }else{
        state.gitList=[]
      }
    })
  }
}
const loadGitInfo=(id:any)=>{
    serverApi.gitInfo({"gitId":id}).then(function (res){
      if (res.Status){
        state.ruleForm.AppGitName= res.Data.Name
      }else{
        state.ruleForm.AppGitName=""
      }
    })
}
const delGit=(row:any)=>{
  state.ruleForm.FrameworkGits = state.ruleForm.FrameworkGits.filter(number => number !== parseInt(row.Id));
  loadGit(state.ruleForm.FrameworkGits)
}

// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};

// 删除
const onDelete = () => {
  ElMessageBox.confirm(`此操作将永久删除应用：“${state.ruleForm.AppName}”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        // 删除逻辑
        serverApi.appsDel({"appName":state.ruleForm.AppName}).then(function (res){
          if (res.Status){
            closeDialog();
            ElMessage.success('删除成功');
          }else{
            ElMessage.error(res.StatusMessage)
          }
        })
      })
      .catch(() => {});
};
// 提交
const onSubmit = () => {
  // 提交数据
  var param={
    "AppName":state.ruleForm.AppName,
    "ShellScript":state.ruleForm.ShellScript,
    "AppGit":parseInt(state.ruleForm.AppGit),
    "FrameworkGits":state.ruleForm.FrameworkGits,
  }

	if (state.dialog.type === 'add') {
    var json=JSON.stringify(param)
    serverApi.appsAdd(json).then(function (res){
      if(res.Status){
        ElMessage.success("添加成功")
        closeDialog();
        emit('refresh');
      }else{
        ElMessage.error(res.StatusMessage)
      }
    })

  }else if (state.dialog.type=='edit'){
    serverApi.appsEdit(param).then(function (res){
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

const getTableData = () => {
  const data = [];
  // 请求接口
  serverApi.gitList({}).then(function (res){
    if (res.Status){
      state.tableData.data = res.Data;
      state.tableData.total = res.Data.length;
    }else{
      state.tableData.data=[]
    }
  })
};
const getRowKey=(row:any)=>{
  return row.Id;
}
const handleSelectionChange=(val:any)=> {
  console.log(val)
  if(val.length==0){return;}
  state.SelectItem=[] // 清空
  for (let i = 0; i < val.length; i++) {
    var item=val[i]
    state.SelectItem.push(item.Id)
  }
  console.log(state.SelectItem)
}
const onOpenGit=(type:any)=>{
  getTableData()
  state.gitType=type
  state.gitDialogIsShow=true
  if (type==1){
    state.tableData.data.forEach(function (item,index){
      var rowArray=state.ruleForm.FrameworkGits.filter(t=>t==item.Id);
      if(rowArray.length>0)
      {
        multipleTable.value.setCurrentRow(item) //选中已经选中的数据
      }else{
        multipleTable.value.setCurrentRow(item,false)
      }
    })
  }

}
// 确认选择
const SureCheck=()=>{
  if (state.gitType==1){
    state.ruleForm.FrameworkGits=state.SelectItem
    loadGit(state.ruleForm.FrameworkGits)
  }else{
    state.ruleForm.AppGit=state.SelectItem[0]
    loadGitInfo(state.ruleForm.AppGit)
  }
  state.gitDialogIsShow=false
}
// 页面加载时
onMounted(() => {
  //getTableData();
});
// 暴露变量
defineExpose({
	openDialog,
  delGit,
});
</script>

<style>
textarea{
  height: 220px;
}
</style>
