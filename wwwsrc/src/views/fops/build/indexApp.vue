<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
        <el-select v-model="state.clusterId" placeholder="请选择集群" class="ml10" @change="onClusterChange">
          <el-option
              v-for="item in state.clusterData"
              :key="item.Id"
              :label="item.Name"
              :value="item.Id"
          ></el-option>
        </el-select>
<!--				<el-input size="default" placeholder="请输入用户名称" style="max-width: 180px"> </el-input>-->
<!--				<el-button size="default" type="primary" class="ml10">-->
<!--					<el-icon>-->
<!--						<ele-Search />-->
<!--					</el-icon>-->
<!--					查询-->
<!--				</el-button>-->
				<el-button size="default" type="success" class="ml10" @click="onOpenAdd('add')">
					<el-icon>
            <ele-FolderAdd />
					</el-icon>
					新增应用
				</el-button>
        <el-button size="default" type="danger" class="ml10" @click="onClearDockerImage('add')">
          <el-icon>
            <ele-Delete />
          </el-icon>
          清除None镜像
        </el-button>
			</div>
      <div class="flex-warp" style="background: #e0e0e0">
        <el-row style="float:left;width: 65%">
          <el-col style="float: left;margin: 10px;"  :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb15" v-if="state.tableData.data.length > 0">
            <div  style="background: #ffffff;width: 24%"  class="flex-warp-item" v-for="(v, k) in state.tableData.data" :key="k">
              <div class="flex-warp-item-box" @click="onShowBuildList(v)">
                <div class="appItem">
                  <el-tag size="small">{{ v.AppName }}</el-tag>
                  <el-tag v-if="v.IsHealth" size="small" type="success">健康</el-tag>
                  <el-tag v-else-if="length(v.ActiveInstance) > 0" size="small" type="warning">不健康</el-tag>
                  <el-tag v-else size="small" type="danger">未运行</el-tag>
                </div>
                <div class="appItem">容器版本：{{ v.DockerVer }}</div>
                <div class="appItem">集群版本：{{ v.ClusterVer }}</div>
                <div class="appItem">仓库：{{ v.AppGitName }}
                </div>
                <div class="appItem">容器文件路径：{{ v.DockerfilePath }}</div>
                <div class="appItem">
                  <el-button size="default" @click="onOpenEdit('edit', v)" type="warning">修改</el-button>
                  <el-button @click="onBuildAdd(v)" size="default" type="danger">构建</el-button>

                </div>
              </div>
            </div>
          </el-col>
          <el-empty v-else description="暂无数据"></el-empty>
        </el-row>
        <el-row style="width: 35%;float:left;">
          <el-col style="float: left;background: #ffffff;margin: 15px 10px 10px 0;padding:5px;width: 98%" :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb15" v-if="state.tableLogData.data.length > 0">
            <h3 style="padding: 5px;">构建队列</h3>
            <template v-if="state.tableLogData.data.length > 0">
              <el-table  :data="state.tableLogData.data" v-loading="state.tableLogData.loading" style="width: 100%;background: #ffffff;">
                <el-table-column prop="Id" label="编号" width="70" />
                <el-table-column prop="AppName" label="应用名称" ></el-table-column>
                <el-table-column label="状态" width="90" show-overflow-tooltip>
                  <template #default="scope">
                    <el-tag v-if="scope.row.Status==0" size="small" type="info">未开始</el-tag>
                    <el-tag v-else-if="scope.row.Status==1" size="small" type="warning">构建中</el-tag>
                    <el-tag v-if="scope.row.Status==2 && scope.row.IsSuccess == true" size="small" type="success">成功</el-tag>
                    <el-tag v-else-if="scope.row.Status==2 && scope.row.IsSuccess == false" size="small" type="danger">失败</el-tag>
                  </template>
                </el-table-column>
                <el-table-column prop="FinishAt" width="170" label="完成时间"></el-table-column>
                <el-table-column label="操作" width="80">
                  <template #default="scope">
                    <el-button v-if="scope.row.Status!=0" size="small" type="success" @click="showLog(scope.row)">日志</el-button>
                  </template>
                </el-table-column>
              </el-table>
              <el-pagination
                  @size-change="onHandleSizeLogChange"
                  @current-change="onHandleCurrentLogChange"
                  class="mt15"
                  :pager-count="5"
                  :page-sizes="[10, 20, 30]"
                  v-model:current-page="state.tableLogData.param.pageNum"
                  background
                  v-model:page-size="state.tableLogData.param.pageSize"
                  layout="total, sizes, prev, pager, next, jumper"
                  :total="state.tableLogData.total"
              >
              </el-pagination>
            </template>
          </el-col>
          <el-empty v-else description="暂无数据"></el-empty>
        </el-row>
      </div>
		</el-card>
		<appDialog ref="appDialogRef" @refresh="getTableData()" />
    <el-dialog title="构建日志" v-model="state.logDialogIsShow" style="width: 80%;height: 85%;top:20px;margin-bottom: 50px">
      <el-card shadow="hover" class="layout-padding-auto" style="background-color:#393d49;overflow: auto;">
        <pre style="color: #fff;background-color:#393d49;height: 100%;" v-html="state.logContent"></pre>
      </el-card>
    </el-dialog>
	</div>
</template>

<script setup lang="ts" name="fopsApp">
import {defineAsyncComponent, reactive, onMounted, ref, nextTick, watch, onUnmounted} from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";

// var idPre = document.getElementById('idPre');
// idPre.scrollIntoView(false); // 滚动到底部

// 引入 api 请求接口
const serverApi = fopsApi();

// 引入组件
const appDialog = defineAsyncComponent(() => import('/@/views/fops/app/dialog.vue'));

// 定义变量内容
const appDialogRef = ref();
const state = reactive({
  logDialogIsShow:false,
  logContent:'',
	tableData: {
		data: [],
		total: 0,
		loading: false,
		param: {
			pageNum: 1,
			pageSize: 12,
		},
	},tableLogData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 12,
    },

  },
  appName:"",
  logId:0,
  clusterId:0,
  clusterData:[],
});

// 初始化表格数据
const getTableData = () => {
	state.tableData.loading = true;
	const data = [];
  // 请求接口
  serverApi.appsList({}).then(function (res){
    if (res.Status){
      for (let i = 0; i < res.Data.length; i++) {
        var item=res.Data[i]
        item.FrameworkGitsStr=getGitArray(item.FrameworkGits)
        item.AppGitStr=getGit(item.AppGit)
        data.push(item)
      }
      setTimeout(() => {
        state.tableData.data =data;
        state.tableData.total = data.length;
        state.tableData.loading = false;
      }, 500);
    }else{
      state.tableData.data=[]
      setTimeout(() => {
        state.tableData.loading = false;
      }, 500);
    }
  })
};

const getTableLogData = () => {
  state.tableLogData.loading = true;
  const data = {
    appName:"",//state.appName
    pageIndex:state.tableLogData.param.pageNum,
    pageSize:state.tableLogData.param.pageSize,
  };
  // 请求接口
  serverApi.buildList(data).then(function (res){
    if (res.Status){
      state.tableLogData.data = res.Data.List;
      state.tableLogData.total = res.Data.RecordCount;
    }else{
      state.tableLogData.data=[]
    }
    state.tableLogData.loading = false;
  })
};
const getTableClusterData = () => {
  state.tableData.loading = true;
  const data = [];
  // 请求接口
  serverApi.clusterList({}).then(function (res){
    if (res.Status){
      var lst=[]
      for (let i = 0; i < res.Data.length; i++) {
        var item=res.Data[i]
        if (i==0){
          state.clusterId=item.Id;
        }
        item.Name=item.Name+" - "+item.DockerName
        lst.push(item)
      }
      state.clusterData = lst;
    }else{
      state.tableData.data=[]
    }
  })

};
const onClusterChange=(value:number)=>{
  state.clusterId=value
}
// 打开新增用户弹窗
const onOpenAdd = (type: string) => {
  appDialogRef.value.openDialog(type);
};
// 打开修改用户弹窗
const onOpenEdit = (type: string, row: any) => {
  appDialogRef.value.openDialog(type, row);
};
// 清除镜像
const onClearDockerImage = () => {
  ElMessageBox.confirm(`此操作将永久清除：“None镜像”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        // 删除逻辑
        serverApi.dockerClearImage().then(function (res){
          if (res.Status){
            ElMessage.success('清除成功');
          }else{
            ElMessageBox.alert(res.StatusMessage,'Warning',{ type: 'warning',dangerouslyUseHTMLString: true})
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
const onHandleSizeLogChange = (val: number) => {
  state.tableLogData.param.pageSize = val;
  getTableLogData();
};
const onShowBuildList=(row: any)=>{
  state.appName=row.AppName
  state.tableLogData.param.pageNum=1
  state.tableLogData.param.pageSize=10
  getTableLogData();
}
// 分页改变
const onHandleCurrentLogChange = (val: number) => {
  state.tableLogData.param.pageNum = val;
  getTableLogData();
};
// 定义定时器
let intervalId = null;
// 使用 watch 监听 state 中 count 属性的变化
watch(() => state.logDialogIsShow, (newValue, oldValue) => {
  //console.log(`count 从 ${oldValue} 变为 ${newValue}`);
  if(!newValue){
    clearInterval(intervalId);
  }else {
    intervalId = setInterval(onShowLog, 1000);
  }
});

const showLog=(row:any)=>{
  state.logId=row.Id
  serverApi.buildLog(state.logId.toString()).then(function (res){
    state.logContent=res
    state.logDialogIsShow=true
  })
}
const onShowLog=()=>{
  serverApi.buildLog(state.logId.toString()).then(function (res){
    console.log(res)
    state.logContent=res
  })
}
const onBuildAdd = (row:any) => {
  // 提交数据
  var param={
    "AppName":row.AppName,
    "ClusterId":state.clusterId,
  }
  serverApi.buildAdd(param).then(async function(res){
    if(res.Status){
      ElMessage.success("添加成功")
      // 刷新构建日志
      getTableLogData()
    }else{
      ElMessage.error(res.StatusMessage)
    }
  })
};
const getGitArray=(lst:[])=>{
  var array=[]
  for (let i = 0; i < lst.length; i++) {
    serverApi.gitInfo({"gitId":lst[i]}).then(function (res){
      if (res.Status){
        array.push(res.Data.Name)
      }
    })
  }
  return array
}
const getGit=(val:number)=>{
  var array=[]
  serverApi.gitInfo({"gitId":val}).then(function (res){
      if (res.Status){
        array.push(res.Data.Name)
      }
    })
  return array
}
let intervalLogId = null;
// 页面加载时
onMounted(() => {
	getTableData();
  getTableLogData();
  getTableClusterData();
  intervalLogId = setInterval(getTableLogData, 1000);
});
// 页面注销的时候
onUnmounted(()=>{
  clearInterval(intervalLogId);
})
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
.flex-warp {
  display: flex;
  flex-wrap: wrap;
  align-content: flex-start;
  margin: 0 -5px;
  .flex-warp-item {
    padding: 5px;
    width: 298px;
    min-height: 360px;
    //border: #666 1px solid;
    .flex-warp-item-box {
      border: 1px solid var(--next-border-color-light);
      width: 100%;
      height: 100%;
      border-radius: 2px;
      display: flex;
      flex-direction: column;
      transition: all 0.3s ease;

      .item-img {
        width: 100%;
        height: 215px;
        overflow: hidden;
        img {
          transition: all 0.3s ease;
          width: 100%;
          height: 100%;
        }
      }
      .item-txt {
        flex: 1;
        padding: 15px;
        display: flex;
        flex-direction: column;
        overflow: hidden;
        .item-txt-title {
          margin: 10px!important;
          text-overflow: ellipsis;
          overflow: hidden;
          -webkit-line-clamp: 2;
          -webkit-box-orient: vertical;
          display: -webkit-box;
          color: #666666;
          transition: all 0.3s ease;
          &:hover {
            color: var(--el-color-primary);
            text-decoration: underline;
            transition: all 0.3s ease;
          }
        }
        .item-txt-other {
          flex: 1;
          align-items: flex-end;
          display: flex;
          .item-txt-msg {
            font-size: 12px;
            color: #8d8d91;
          }
          .item-txt-price {
            display: flex;
            justify-content: space-between;
            align-items: center;
            .font-price {
              color: #ff5000;
              .font {
                font-size: 22px;
              }
            }
          }
        }
      }
    }
  }
}
.appItem{
  margin: 10px;
}
.el-row{
  margin: 0!important;
  display: block!important;

}
.flex-warp {
  width: 100%!important;
}
.flex-warp-item{
  float: left;
  margin: 5px;
}
.el-dialog__body{
  height: 100%!important;
}
</style>
