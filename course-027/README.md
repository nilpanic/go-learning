# HTTP响应格式处理

- 没有标准的答案

## 接口设计建议：

1. 接口定义要有版本号。

```shell
/v1/user
/v2/user
/v3/user
/v1/user/info <- account + metadata + order
@GET /v1/user
@POST add、update、delte
```

1. 接口设计是否需要遵循Restful风格?
    1. 参考文档：https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design#what-is-rest
    2. restful应用：https://docs.github.com/en/rest/guides/getting-started-with-the-rest-api?apiVersion=2022-11-28

   实际设计中Restful风格的一些问题：资源命名管理会比较复杂，数据验证处理不好统一，权限系统一定要做适配。

2. 接口响应内容设计：
    1. 状态码如何维护？放在内容里还是通过http状态码？
        1. 状态码到底是放在内容中还是使用http_code（http status_code局限性）
            1. https://api.aliyun.com/document/Ecs/2014-05-26/DescribeInstances
            2. https://api.aliyun.com/document/Dysmsapi/2017-05-25/SendSms
    2. 协议就是：http, grpc；数据格式：xml, JSON，JSONP，protobuf, GET,POST
3. 框架里如何设计：
    1. 公共函数包装
    2. 统一错误处理包装
    3. fork框架，定制版（扩展性强，但有维护成本）
4. 接口安全
    1. 鉴权（应用id+秘钥）
    2. 数据校验：
        1. 对于请求零信任，预防csrf，xss攻击
        2. 签名，
        3. 时间戳，重放限制
        4. 流控、IP白名单
    3. 数据加密(返回数据都是加密：)


### 说明
