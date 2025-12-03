SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 创建 attachment 表
CREATE TABLE IF NOT EXISTS `attachment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `storage_mode` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'local' COMMENT '存储模式:local=本地,oss=阿里云,qiniu=七牛云,cos=腾讯云',
  `origin_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原文件名',
  `object_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '新文件名',
  `hash` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件hash',
  `mime_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '资源类型',
  `storage_path` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储目录',
  `suffix` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件后缀',
  `size_byte` bigint(20) DEFAULT NULL COMMENT '字节数',
  `size_info` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件大小',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url地址',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `attachment_hash_unique` (`hash`),
  KEY `attachment_storage_path_index` (`storage_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='上传文件信息表';

-- 创建 casbin_rule 表
CREATE TABLE IF NOT EXISTS `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- 创建 dict_data 表
CREATE TABLE IF NOT EXISTS `dict_data` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type_code` varchar(30) CHARACTER SET latin1 NOT NULL COMMENT '字典类型编码',
  `label` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '字典key',
  `value` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '字典值',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '字典状态:1-正常,2-禁用',
  `sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='字典数据表';

-- 创建 dict_types 表
CREATE TABLE IF NOT EXISTS `dict_types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(30) NOT NULL COMMENT '字典类型编码',
  `name` varchar(20) NOT NULL COMMENT '字典key',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '字典状态:1-正常,2-禁用',
  `sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典类型表';

-- 创建 roles 表
CREATE TABLE IF NOT EXISTS `roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色代码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1=正常,2=停用',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_code_unique` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色信息表';

-- 创建 role_menus 表
CREATE TABLE IF NOT EXISTS `role_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单映射表';

-- 创建 sys_menus 表
CREATE TABLE IF NOT EXISTS `sys_menus` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `menu_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '菜单类型（1代表菜单、2代表iframe、3代表外链、4代表按钮）',
  `title` varchar(50) NOT NULL COMMENT '菜单名称',
  `name` varchar(50) DEFAULT NULL COMMENT '路由名称',
  `path` varchar(200) DEFAULT '' COMMENT '路由路径',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `rank` int(11) DEFAULT '99' COMMENT '显示排序',
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) DEFAULT NULL COMMENT '菜单图标',
  `extra_icon` varchar(100) DEFAULT NULL COMMENT '右侧图标',
  `enter_transition` varchar(50) DEFAULT NULL COMMENT '进场动画',
  `leave_transition` varchar(50) DEFAULT NULL COMMENT '离场动画',
  `active_path` varchar(255) DEFAULT NULL COMMENT '激活路由路径',
  `auths` varchar(500) DEFAULT NULL COMMENT '权限标识',
  `frame_src` varchar(500) DEFAULT NULL COMMENT '外链地址',
  `frame_loading` tinyint(1) DEFAULT '1' COMMENT '是否显示加载动画',
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '是否缓存',
  `hidden_tag` tinyint(1) DEFAULT '0' COMMENT '是否隐藏标签',
  `fixed_tag` tinyint(1) DEFAULT '0' COMMENT '是否固定标签',
  `show_link` tinyint(1) DEFAULT '1' COMMENT '是否显示',
  `show_parent` tinyint(1) DEFAULT '0' COMMENT '是否显示父级菜单',
  `status` tinyint(4) DEFAULT '1' COMMENT '菜单状态（0停用 1正常）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单权限表';

-- 创建 user 表
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID,主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `user_type` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '100' COMMENT '用户类型:100=系统用户',
  `nickname` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机',
  `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `signed` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '个人签名',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1=正常,2=停用',
  `login_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '127.0.0.1' COMMENT '最后登陆IP',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登陆时间',
  `backend_setting` json DEFAULT NULL COMMENT '后台设置数据',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 创建 user_login_log 表
CREATE TABLE IF NOT EXISTS `user_login_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP地址',
  `os` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '浏览器',
  `status` smallint(6) NOT NULL DEFAULT '1' COMMENT '登录状态 (1成功 2失败)',
  `message` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提示消息',
  `login_time` datetime NOT NULL COMMENT '登录时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_login_log_username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='登录日志表';

-- 创建 user_operation_log 表
CREATE TABLE IF NOT EXISTS `user_operation_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `method` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方式',
  `router` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求路由',
  `service_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务名称',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求IP地址',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_operation_log_username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- 创建 user_roles 表
CREATE TABLE IF NOT EXISTS `user_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色映射表';

SET FOREIGN_KEY_CHECKS = 1; 