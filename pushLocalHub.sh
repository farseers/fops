# 更新farseer-go框架
cd ../farseer-go && sh git-update.sh
# 更新fops
cd ../fops && git pull
# 将忽略文件复制到上下文根目录中
#\cp .dockerignore ../
# 编译
docker build -t hub.fsgit.cc/fops:dev --network=host -f ./Dockerfile ../
# 发到内网
docker push hub.fsgit.cc/fops:dev && docker rmi hub.fsgit.cc/fops:dev

# docker
docker service rm fops
docker service create --name fops --replicas 1 -d --network=net \
--constraint node.role==worker \
--mount type=bind,src=/etc/localtime,dst=/etc/localtime \
-e "Database_default=DataType=mysql,PoolMaxSize=25,PoolMinSize=1,ConnectionString=root:steden@123@tcp(192.168.1.8:3306)/fops2?charset=utf8&parseTime=True&loc=Local" \
-e "Redis_default=Server=192.168.1.8:6379,DB=13,Password=steden@123,ConnectTimeout=600000,SyncTimeout=10000,ResponseTimeout=10000" \
-l "traefik.http.routers.fops.rule=Host(\`fops.fsgit.cc\`)" \
-l "traefik.http.routers.fops.entrypoints=websecure" \
-l "traefik.http.routers.fops.tls=true" \
-l "traefik.http.services.fops.loadbalancer.server.port=8889" \
hub.fsgit.cc/fops:dev

#docker rmi $(docker images -f "dangling=true" -q)