# dbproxy
简单的数据库sql代理

1. 资源管理系统
2. 提供服务的API操作


    1. 提供简单的sql的操作和封装sql的操作
    2. 不同的数据库连接需要不同的命名
    
    GET|POST
    {
        db:mysql,
        type:raw|func_name,
        data:"select * from mysql",
        data:"{参数}",
    }
    
