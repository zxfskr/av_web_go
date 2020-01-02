CREATE DATABASE `av_web`;

use `av_web`;

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `password_hash` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
    `user_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '权限  2：admin   1：普通用户',
    `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '盐',
    `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile` varchar(100) NOT NULL DEFAULT '' COMMENT '手机号',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态  0：禁用   1：正常',
    `dept_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '部门ID',
    `dept_name` varchar(32) NOT NULL DEFAULT '' COMMENT '部门名',
    `role_name` varchar(32) NOT NULL DEFAULT '' COMMENT '',
    `role_id_list` varchar(32) NOT NULL DEFAULT '' COMMENT '',
    `create_name` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_name` varchar(32) NOT NULL DEFAULT '' COMMENT '修改人',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`user_id`) USING BTREE,
    UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='系统用户';
