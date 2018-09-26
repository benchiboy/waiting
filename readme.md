readme
========


# 用户登录接口
 整个登录过程采用JWT进行会话控制：
1、登录
 request: {login_name:$(login_name),passwd:$(passwd),login_time:$(login_time)}
 response:{status_code:$(result),status_msg:$(status_msg),token:$(token)}



 