# docker测试环境

1. 启动mysql

```
cd ${PROJECT}/
docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d -v `pwd`/db:/docker-entrypoint-initdb.d -p 3306:3306 mysql:5.7
```

```
docker run --net=host --name nginx -d -v `pwd`/dist:/dist -v `pwd`/conf:/etc/nginx/conf.d nginx
```

