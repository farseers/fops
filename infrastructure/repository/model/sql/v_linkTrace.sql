CREATE VIEW {database}.{table}
    (
    trace_id             String default '' comment '上下文ID',
    app_id               String default '' comment '应用ID',
    app_name             String default '' comment '应用名称',
    app_ip               String default '' comment '应用IP',
    parent_app_name      String default '' comment '上游应用',
    start_ts             Int64  default 0 comment '调用开始时间戳（微秒）',

    end_ts               Int64  default 0 comment '调用结束时间戳（微秒）',
    use_ts               Int64  default 0 comment '总共使用时间（微秒）',
    use_desc             String default '' comment '总共使用时间（描述）',
    trace_type           Int64 comment '入口类型',
    exception            String comment '异常信息',
    web_domain           String default '' comment 'Web请求上下文',
    web_path             String default '' comment 'Web请求上下文',
    web_method           String default '' comment 'Web请求上下文',
    web_content_type     String default '' comment 'Web请求上下文',
    web_status_code      Int64  default 0 comment 'Web请求上下文',
    web_request_ip       String default '' comment 'Web请求上下文',
    consumer_server      String default '' comment '消费上下文',
    consumer_queue_name  String default '' comment '消费上下文',
    consumer_routing_key String default '' comment '消费上下文',
    task_name            String default '' comment '任务上下文',
    task_group_id        Int64  default 0 comment '任务上下文',
    task_id              Int64  default 0 comment '任务上下文',
    watch_key            String default '' comment 'Etcd上下文',
    create_at            DateTime64(3) comment '请求时间'
)
    AS
SELECT app_name, parent_app_name, app_ip, app_id,
       trace_id,
       start_ts, end_ts, use_ts, use_desc, trace_type,
       exception,
       web_domain,web_path,web_method,web_content_type,web_status_code,web_request_ip,
       consumer_server,consumer_queue_name,consumer_routing_key,
       task_name,task_group_id,task_id,
       watch_key,create_at
FROM {database}.link_trace
GROUP BY app_name, parent_app_name, app_ip, app_id,
    trace_id,
    start_ts, end_ts, use_ts, use_desc, trace_type,
    exception,
    web_domain,web_path,web_method,web_content_type,web_status_code,web_request_ip,
    consumer_server,consumer_queue_name,consumer_routing_key,
    task_name,task_group_id,task_id,
    watch_key,create_at