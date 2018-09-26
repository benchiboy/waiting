// waiting project doc.go

/*
waiting document
*/
package main

//2017-10-31
//register
//{"login_name":"admin","login_pwd":"123456","user_name":"张三","user_phone":"130555555","check_code":"946695"}
//login
//{"login_name":"admin","login_pwd":"123456"}

//2017-11-01
//forget password
//{"login_name":"admin","user_phone":"13049857690","check_code":"790874","login_pwd":"123456"}

//send check code
//{"user_phone":"13049857690","code_type":"r"}
//{"user_phone":"13049857690","code_type":"p"}

//{"user_phone":"13049857692","code_type":"r"}
//{"login_name":"admin2","login_pwd":"123456","user_name":"张三","user_phone":"13049857692","check_code":"946695"}

/////////////////////////////////////////////////////////////////////////////////////////////////////
/*
CREATE TABLE `wait_checkcodes` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) DEFAULT '' COMMENT '用户ID',
  `user_phone` varchar(21) DEFAULT '' COMMENT '电话号码',
  `check_code` varchar(8) DEFAULT '' COMMENT '短信校验码',
  `check_type` varchar(8) DEFAULT '' COMMENT '短信校验码类型',
  `send_code` char(1) DEFAULT '' COMMENT '短信验证码发送状态',
  `send_msg` char(50) DEFAULT '' COMMENT '短信验证码错误信息',
  `code_status` char(1) DEFAULT 'e' COMMENT 'e：可使用，d：已经使用',
  `verify_times` int(11) DEFAULT '0' COMMENT '短信校验次数',
  `valid_btime` datetime DEFAULT NULL COMMENT '有效开始时间',
  `valid_etime` datetime DEFAULT NULL COMMENT '有效结束时间',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`),
  KEY `idx_user_phone` (`user_phone`) USING BTREE,
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8;



CREATE TABLE `wait_users` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) DEFAULT '' COMMENT '用户ID',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `nick_name` varchar(50) DEFAULT '' COMMENT '用户昵称',
  `login_pwd` varchar(50) DEFAULT '' COMMENT '密码',
  `user_type` tinyint(4) DEFAULT '0' COMMENT '用户类型：1：内部，2：WEIBO,3：QQ',
  `user_name` varchar(100) DEFAULT '' COMMENT '用户姓名',
  `user_idno` varchar(30) DEFAULT '' COMMENT '证件号码',
  `user_phone` varchar(21) DEFAULT '' COMMENT '电话号码',
  `user_sex` char(1) DEFAULT '' COMMENT '用户性别',
  `is_admin` tinyint(4) DEFAULT '0' COMMENT '是否管理员',
  `user_islock` tinyint(4) DEFAULT '0' COMMENT '是否锁定',
  `pwderr_count` tinyint(4) DEFAULT '0' COMMENT '密码错误次数',
  `last_check_code` varchar(10) DEFAULT '' COMMENT '最新的图形验证码',
  `last_login_time` datetime DEFAULT NULL COMMENT '上次登录时间',
  `pic_full` varchar(100) DEFAULT '' COMMENT '生活照图片',
  `pic_head` varchar(100) DEFAULT '' COMMENT '头像',
  `user_memo` varchar(150) DEFAULT '' COMMENT '用户说明',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`),
  KEY `idx_login_name` (`login_name`) USING BTREE,
  KEY `idx_phone_no` (`user_phone`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8;

CREATE TABLE `wait_users_charge_orders` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) DEFAULT '' COMMENT '用户ID',
  `mct_type` varchar(10) DEFAULT '' COMMENT '支付商户类型',
  `trxn_type` varchar(10) DEFAULT '' COMMENT '交易类型',
  `old_order_no` varchar(50) DEFAULT '' COMMENT '原交易订单号',
  `order_no` varchar(50) DEFAULT '' COMMENT '支付订单号',
  `order_type` varchar(10) DEFAULT '' COMMENT '订单类型',
  `order_amt` decimal(15,2) DEFAULT '0.00' COMMENT '订单金额',
  `mct_order_no` varchar(50) DEFAULT '' COMMENT '支付机构订单号',
  `order_date` datetime DEFAULT NULL COMMENT '订单支付成功',
  `status_code` varchar(5) DEFAULT '' COMMENT '支付状态代码',
  `status_msg` varchar(50) DEFAULT '' COMMENT '支付失败错误原因',
  `memo` varchar(50) DEFAULT '' COMMENT '交易备注',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`),
  KEY `idx_order_no` (`order_no`),
  KEY `idx_trxn_type` (`trxn_type`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;

CREATE TABLE `wait_users_charge_orders_detail` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) DEFAULT '' COMMENT '用户ID',
  `order_no` varchar(50) DEFAULT '' COMMENT '支付订单号',
  `goods_no` varchar(50) DEFAULT '' COMMENT '商品编号',
  `goods_amt` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '商品价格',
  `goods_type` char(1) DEFAULT '' COMMENT '商品类型 V:虚拟商品 E:实体商品',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`),
  KEY `idx_order_no` (`order_no`),
  KEY `idx_goods_no` (`goods_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;




CREATE TABLE `wait_users_chat_msg` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) CHARACTER SET latin1 DEFAULT '' COMMENT '用户ID',
  `recv_user_id` varchar(50) DEFAULT '' COMMENT '接手消息用户ID',
  `chat_type` tinyint(4) DEFAULT '0' COMMENT '消息类型',
  `chat_msg` varchar(1000) DEFAULT '' COMMENT '消息内容',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;

CREATE TABLE `wait_users_detail` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) CHARACTER SET latin1 DEFAULT '' COMMENT '用户ID',
  `blood_type` char(255) DEFAULT NULL,
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `pic_head` varchar(100) DEFAULT '' COMMENT '头像',
  `pic_full` varchar(100) DEFAULT '' COMMENT '生活照图片',
  `version` int(11) DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;

CREATE TABLE `wait_users_pics` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) CHARACTER SET latin1 DEFAULT '' COMMENT '原图片用户ID',
  `pic_id` bigint(20) DEFAULT '0' COMMENT '图片ID',
  `pic_type` tinyint(4) DEFAULT '0' COMMENT '评论类型',
  `pic_name` varchar(100) DEFAULT '' COMMENT '评论内容',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;

CREATE TABLE `wait_users_pics_comments` (
  `auto_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) CHARACTER SET latin1 DEFAULT '' COMMENT '原图片用户ID',
  `comment_user_id` varchar(50) DEFAULT '' COMMENT '发出评论用户ID',
  `pic_id` bigint(20) DEFAULT '0' COMMENT '图片ID',
  `comment_type` tinyint(4) DEFAULT '0' COMMENT '评论类型',
  `comment_msg` varchar(1000) DEFAULT '' COMMENT '评论内容',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;


3.16      427367   423229     3.16   已还本金利息， 3.26号的；

*/
