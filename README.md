# DouYin
新建项目 
字节青训营抖音项目开发,希望项目成功完成！


# 文件夹说明
1.config配置文件，用于连接数据库等
2.处理流程
dao层连接数据库等
router层 接收url  -->  controller层 路由执行的函数  -->  logic层 可能查多个表，数据整合，逻辑判断  -->  model层 数据库的增删改查

router 接收url
controller 路由执行函数，并返回响应消息 response
// 现在删除dao层，实际整合到model层，调用数据库处理函数是通过model层中dao的对象


