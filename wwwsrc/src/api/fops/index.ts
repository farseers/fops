import request from '/@/utils/request';
import requestFS from '/@/utils/requestFS';
import requestFSGet from '/@/utils/requestFSGet';

/**
 * （不建议写成 request.post(xxx)，因为这样 post 时，无法 params 与 data 同时传参）
 *
 * 登录api接口集合
 * @method gitList git列表
 * @method gitAdd git添加
 * @method gitEdit git修改
 * @method gitDel git删除
 */
export function fopsApi() {
	return {
		gitList: (param: object) => {
			return request({
				url: '/git/list',
				method: 'post',
				data:param,
			});
		},
		gitAdd: (param: object) => {
			return request({
				url: '/git/add',
				method: 'post',
				data:param,
			});
		},
		gitInfo: (param: object) => {
			return request({
				url: '/git/info',
				method: 'post',
				data:param,
			});
		},
		gitEdit: (param: object) => {
			return request({
				url: '/git/update',
				method: 'post',
				data:param,
			});
		},
		gitDel: (param: object) => {
			return request({
				url: '/git/delete',
				method: 'post',
				data:param,
			});
		},appsList: (param: object) => {
			return request({
				url: '/apps/list',
				method: 'post',
				data:param,
			});
		},
		appsAdd: (param: object) => {
			return request({
				url: '/apps/add',
				method: 'post',
				data:param,
			});
		},
		appsEdit: (param: object) => {
			return request({
				url: '/apps/update',
				method: 'post',
				data:param,
			});
		},
		appsDel: (param: object) => {
			return request({
				url: '/apps/delete',
				method: 'post',
				data:param,
			});
		},appsDetail: (param: object) => {
			return request({
				url: '/apps/info',
				method: 'post',
				data:param,
			});
		},clusterAdd: (param: object) => {
			return request({
				url: '/cluster/add',
				method: 'post',
				data:param,
			});
		},clusterEdit: (param: object) => {
			return request({
				url: '/cluster/update',
				method: 'post',
				data:param,
			});
		},clusterList: (param: object) => {
			return request({
				url: '/cluster/list',
				method: 'post',
				data:param,
			});
		},clusterDel: (param: object) => {
			return request({
				url: '/cluster/delete',
				method: 'post',
				data:param,
			});
		},buildList: (param: object) => {
			return request({
				url: '/apps/build/list',
				method: 'post',
				data:param,
			});
		},buildAdd: (param: object) => {
			return request({
				url: '/apps/build/add',
				method: 'post',
				data:param,
			});
		},buildLog: (param: object) => {
			return request({
				url: '/apps/build/view',
				method: 'get',
				data:param,
			});
		},taskGroupList: (param: string) => {
			return requestFSGet({
				url: '/basicapi/taskGroup/list?'+param,
				method: 'get',
			});
		},taskList: (param: string) => {
			return requestFSGet({
				url: '/basicapi/task/list?'+param,
				method: 'get',
			});
		},taskGroupInfo: (url:string,param: object) => {
			return requestFSGet({
				url: url,
				method: 'get',
			});
		},taskGroupSetEnable: (param: object) => {
			return requestFS({
				url: '/basicapi/taskGroup/setEnable',
				method: 'post',
				data:param,
			});
		},taskUpdate: (param: object) => {
			return requestFS({
				url: '/basicapi/taskGroup/update',
				method: 'post',
				data:param,
			});
		},taskDel: (param: object) => {
			return requestFS({
				url: '/basicapi/taskGroup/delete',
				method: 'post',
				data:param,
			});
		},taskCount: (param: object) => {
			return requestFS({
				url: '/basicapi/taskGroup/count',
				method: 'get',
				data:param,
			});
		},taskNoRunCount: (param: object) => {
			return requestFS({
				url: '/basicapi/taskGroup/unRunCount',
				method: 'get',
				data:param,
			});
		},taskNoRunList: (param: string) => {
			return requestFSGet({
				url: '/basicapi/taskGroup/unRunList?'+param,
				method: 'get',
			});
		},taskRunningList: (param: string) => {
			return requestFSGet({
				url: '/basicapi/taskGroup/schedulerWorkingList?'+param,
				method: 'get',
			});
		},taskLogList: (param: string) => {
			return requestFSGet({
				url: '/basicapi/log/list?'+param,
				method: 'get',
			});
		},taskFinishedList: (param: object) => {
			return requestFSGet({
				url: '/basicapi/task/list',
				method: 'get',
				data:param,
			});
		},taskTodayFailCount: (param: object) => {
			return requestFS({
				url: '/basicapi/task/todayFailCount',
				method: 'get',
				data:param,
			});
		},serverNodeList: (param: object) => {
			return requestFS({
				url: '/basicapi/server/list',
				method: 'get',
				data:param,
			});
		},clientList: (param: object) => {
			return requestFS({
				url: '/basicapi/client/list',
				method: 'get',
				data:param,
			});
		},linkTraceWebApi:(param: object) => {
			return requestFS({
				url: '/linkTrace/webApiList',
				method: 'get',
				data:param,
			});
		},linkTraceTask:(param: object) => {
			return requestFS({
				url: '/linkTrace/taskList',
				method: 'get',
				data:param,
			});
		},linkTraceConsumerList:(param: object) => {
			return requestFS({
				url: '/linkTrace/consumerList',
				method: 'get',
				data:param,
			});
		},linkTraceFScheduleList:(param: object) => {
			return requestFS({
				url: '/linkTrace/fScheduleList',
				method: 'get',
				data:param,
			});
		},linkTraceInfo:(traceId: object) => {
			return requestFS({
				url: '/linkTrace/info/'+traceId,
				method: 'get',
				//data:param,
			});
		},slowSql:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowDbList',
				method: 'get',
				data:param,
			});
		},slowEs:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowEsList',
				method: 'get',
				data:param,
			});
		},slowEtcd:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowEtcdList',
				method: 'get',
				data:param,
			});
		},slowHand:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowHandList',
				method: 'get',
				data:param,
			});
		},slowHttp:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowHttpList',
				method: 'get',
				data:param,
			});
		},slowMq:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowMqList',
				method: 'get',
				data:param,
			});
		},slowRedis:(param: object) => {
			return requestFS({
				url: '/linkTrace/slowRedisList',
				method: 'get',
				data:param,
			});
		}
	};
}
