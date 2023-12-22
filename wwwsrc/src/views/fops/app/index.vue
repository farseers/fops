<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
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
			</div>
      <div class="list-adapt-container layout-pd">
      <div class="flex-warp" v-if="state.tableData.data.length > 0">
        <el-row :gutter="15">
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb15" >
            <div class="flex-warp-item" v-for="(v, k) in state.tableData.data" :key="k">
              <div class="flex-warp-item-box">
                <div class="appItem">应用ID：{{ v.AppId }}</div>
                <div class="appItem">应用名称：{{ v.AppName }}</div>
                <div class="appItem">容器版本：{{ v.DockerVer }}</div>
                <div class="appItem">容器文件：{{ v.Dockerfile }}</div>
                <div class="appItem">AppGit：
                  <el-tag v-for="(item, k) in v.AppGitStr">{{item}}</el-tag>
                </div>
                <div class="appItem">Git框架：
                  <el-tag v-for="(item, k) in v.FrameworkGitsStr">{{item}}</el-tag>
                </div>
                <div class="appItem">容器文件：{{ v.DockerfilePath }}</div>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
      <el-empty v-else description="暂无数据"></el-empty>
      <template v-if="state.tableData.data.length > 0">
<!--        <el-pagination-->
<!--            style="text-align: right"-->
<!--            background-->
<!--            @size-change="onHandleSizeChange"-->
<!--            @current-change="onHandleCurrentChange"-->
<!--            :page-sizes="[10, 20, 30]"-->
<!--            :current-page="state.tableData.param.pageNum"-->
<!--            :page-size="state.tableData.param.pageSize"-->
<!--            layout="total, sizes, prev, pager, next, jumper"-->
<!--            :total="state.tableData.total"-->
<!--        >-->
<!--        </el-pagination>-->
      </template>

      </div>
<!--			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">-->
<!--				<el-table-column prop="AppId" label="应用ID" width="60" />-->
<!--				<el-table-column prop="AppName" label="应用名称" show-overflow-tooltip></el-table-column>-->
<!--				<el-table-column prop="DockerVer" label="镜像版本" show-overflow-tooltip></el-table-column>-->
<!--&lt;!&ndash;				<el-table-column prop="ShellScript" label="Shell脚本" show-overflow-tooltip></el-table-column>&ndash;&gt;-->
<!--&lt;!&ndash;				<el-table-column prop="ClusterVer" label="集群版本" show-overflow-tooltip></el-table-column>&ndash;&gt;-->
<!--				<el-table-column prop="AppGit" label="应用的源代码" show-overflow-tooltip></el-table-column>-->
<!--&lt;!&ndash;        <el-table-column prop="FrameworkGits" label="依赖的框架源代码" show-overflow-tooltip></el-table-column>&ndash;&gt;-->
<!--        <el-table-column prop="PullAt" label="拉取时间" show-overflow-tooltip></el-table-column>-->
<!--&lt;!&ndash;        <el-table-column prop="Dockerfile" label="Dockerfile内容" show-overflow-tooltip></el-table-column>&ndash;&gt;-->
<!--        <el-table-column prop="DockerfilePath" label="Dockerfile路径" show-overflow-tooltip></el-table-column>-->
<!--				<el-table-column label="操作" width="100">-->
<!--					<template #default="scope">-->
<!--						<el-button size="small" text type="primary" @click="onOpenEdit('edit', scope.row)"-->
<!--							>修改</el-button-->
<!--						>-->
<!--						<el-button size="small" text type="primary" @click="onRowDel(scope.row)">删除</el-button>-->
<!--					</template>-->
<!--				</el-table-column>-->
<!--			</el-table>-->
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
		<appDialog ref="appDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="fopsApp">
import { defineAsyncComponent, reactive, onMounted, ref,nextTick } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import {fopsApi} from "/@/api/fops";

// 引入 api 请求接口
const serverApi = fopsApi();

// 引入组件
const appDialog = defineAsyncComponent(() => import('/@/views/fops/app/dialog.vue'));

// 定义变量内容
const appDialogRef = ref();
const state = reactive({
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
// 打开新增用户弹窗
const onOpenAdd = (type: string) => {
  appDialogRef.value.openDialog(type);
};
// 打开修改用户弹窗
const onOpenEdit = (type: string, row: any) => {
  appDialogRef.value.openDialog(type, row);
};
// 删除用户
const onRowDel = (row: any) => {
	ElMessageBox.confirm(`此操作将永久删除：“${row.AppName}”，是否继续?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
      // 删除逻辑
      serverApi.appsDel({"AppName":row.AppName}).then(function (res){
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
.flex-warp {
  display: flex;
  flex-wrap: wrap;
  align-content: flex-start;
  margin: 0 -5px;
  .flex-warp-item {
    padding: 5px;
    width: 300px;
    height: 360px;
    border: #666 1px solid;
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
</style>
