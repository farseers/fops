select
    app_name,app_ip,app_id,host_name,process_id,startup_at,cpu_usage,memory_usage
from
    register
where
    activate_at >= (NOW() - INTERVAL 30 SECOND)
group by
    app_name,app_ip,app_id,host_name,process_id,startup_at,cpu_usage,memory_usage