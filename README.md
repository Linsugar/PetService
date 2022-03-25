# 目前架构：单体架构-才用MVC模式

# 鉴权：采用gin-Jwt

# 目前模块：
1.注册用户

2.登录

3.获取用户列表

4.添加宠物

5.宠物详情

6.发布动态

7.通过微信三方接口-获取到微信公众号文章

8.通过ini文件进行的配置

9.路由实现单例-避免重复创建
![image](https://user-images.githubusercontent.com/51956983/160052582-93d638cf-bb77-45fa-b049-e37851b66d2e.png)


10.增加定时任务合集-Cron

# 集成swagger：
集成swagger-每次进行修改时需要swag init进行更新
![image](https://user-images.githubusercontent.com/51956983/160049855-ff321e5f-c80f-4b14-b047-f47be770e2c3.png)


# 集成Redis：
Addr="localhost:port"

# 集成Mysql：
Host = "localhost"

Port = "port"

Database = "tablename"

UserName = "xxxx"

PassWord = "xxxxx"

CharSet = "utf8"
