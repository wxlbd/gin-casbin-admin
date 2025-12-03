-- PostgreSQL 版本 - 创建表结构
-- 移除MySQL特有语法，转换为PostgreSQL语法

-- 创建 attachment 表
CREATE TABLE IF NOT EXISTS attachment (
  id BIGSERIAL PRIMARY KEY,
  storage_mode varchar(20) NOT NULL DEFAULT 'local',
  origin_name varchar(255) DEFAULT NULL,
  object_name varchar(50) DEFAULT NULL,
  hash varchar(64) DEFAULT NULL,
  mime_type varchar(255) DEFAULT NULL,
  storage_path varchar(100) DEFAULT NULL,
  suffix varchar(20) DEFAULT NULL,
  size_byte bigint DEFAULT NULL,
  size_info varchar(50) DEFAULT NULL,
  url varchar(255) DEFAULT NULL,
  created_by bigint NOT NULL DEFAULT 0,
  updated_by bigint NOT NULL DEFAULT 0,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL,
  remark varchar(255) NOT NULL DEFAULT ''
);

-- 创建唯一索引
CREATE UNIQUE INDEX IF NOT EXISTS attachment_hash_unique ON attachment(hash);
CREATE INDEX IF NOT EXISTS attachment_storage_path_index ON attachment(storage_path);

-- 创建 casbin_rule 表
CREATE TABLE IF NOT EXISTS casbin_rule (
  id BIGSERIAL PRIMARY KEY,
  ptype varchar(100) DEFAULT NULL,
  v0 varchar(100) DEFAULT NULL,
  v1 varchar(100) DEFAULT NULL,
  v2 varchar(100) DEFAULT NULL,
  v3 varchar(100) DEFAULT NULL,
  v4 varchar(100) DEFAULT NULL,
  v5 varchar(100) DEFAULT NULL
);

-- 创建唯一索引
CREATE UNIQUE INDEX IF NOT EXISTS idx_casbin_rule ON casbin_rule(ptype, v0, v1, v2, v3, v4, v5);

-- 创建 dict_data 表
CREATE TABLE IF NOT EXISTS dict_data (
  id BIGSERIAL PRIMARY KEY,
  type_code varchar(30) NOT NULL,
  label varchar(20) NOT NULL,
  value varchar(255) NOT NULL,
  status integer NOT NULL DEFAULT 1,
  sort integer NOT NULL,
  remark varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT NULL,
  deleted_at integer DEFAULT NULL
);

-- 创建 dict_types 表
CREATE TABLE IF NOT EXISTS dict_types (
  id BIGSERIAL PRIMARY KEY,
  code varchar(30) NOT NULL,
  name varchar(20) NOT NULL,
  status integer NOT NULL DEFAULT 1,
  sort integer NOT NULL,
  remark varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT NULL,
  deleted_at integer DEFAULT NULL
);

-- 创建 roles 表
CREATE TABLE IF NOT EXISTS roles (
  id BIGSERIAL PRIMARY KEY,
  name varchar(30) NOT NULL,
  code varchar(100) NOT NULL,
  status smallint NOT NULL DEFAULT 1,
  sort smallint NOT NULL DEFAULT 0,
  created_by bigint NOT NULL DEFAULT 0,
  updated_by bigint NOT NULL DEFAULT 0,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL,
  remark varchar(255) NOT NULL DEFAULT ''
);

-- 创建唯一索引
CREATE UNIQUE INDEX IF NOT EXISTS role_code_unique ON roles(code);

-- 创建 role_menus 表
CREATE TABLE IF NOT EXISTS role_menus (
  id BIGSERIAL PRIMARY KEY,
  role_id bigint NOT NULL,
  menu_id bigint NOT NULL,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL
);

-- 创建 sys_menus 表
CREATE TABLE IF NOT EXISTS sys_menus (
  id BIGSERIAL PRIMARY KEY,
  parent_id bigint DEFAULT 0,
  menu_type smallint NOT NULL DEFAULT 1,
  title varchar(50) NOT NULL,
  name varchar(50) DEFAULT NULL,
  path varchar(200) DEFAULT '',
  component varchar(255) DEFAULT NULL,
  rank integer DEFAULT 99,
  redirect varchar(255) DEFAULT NULL,
  icon varchar(100) DEFAULT NULL,
  extra_icon varchar(100) DEFAULT NULL,
  enter_transition varchar(50) DEFAULT NULL,
  leave_transition varchar(50) DEFAULT NULL,
  active_path varchar(255) DEFAULT NULL,
  auths varchar(500) DEFAULT NULL,
  frame_src varchar(500) DEFAULT NULL,
  frame_loading smallint DEFAULT 1,
  keep_alive smallint DEFAULT 0,
  hidden_tag smallint DEFAULT 0,
  fixed_tag smallint DEFAULT 0,
  show_link smallint DEFAULT 1,
  show_parent smallint DEFAULT 0,
  status smallint DEFAULT 1,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL
);

-- 创建 users 表
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  username varchar(20) NOT NULL,
  password varchar(100) NOT NULL,
  user_type varchar(3) NOT NULL DEFAULT '100',
  nickname varchar(30) NOT NULL DEFAULT '',
  phone varchar(11) NOT NULL DEFAULT '',
  email varchar(50) NOT NULL DEFAULT '',
  avatar varchar(255) NOT NULL DEFAULT '',
  signed varchar(255) NOT NULL DEFAULT '',
  status smallint NOT NULL DEFAULT 1,
  login_ip varchar(45) NOT NULL DEFAULT '127.0.0.1',
  login_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  backend_setting jsonb DEFAULT NULL,
  created_by bigint NOT NULL DEFAULT 0,
  updated_by bigint NOT NULL DEFAULT 0,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL,
  remark varchar(255) NOT NULL DEFAULT ''
);

-- 创建唯一索引
CREATE UNIQUE INDEX IF NOT EXISTS idx_username ON users(username);

-- 创建 user_login_log 表
CREATE TABLE IF NOT EXISTS user_login_log (
  id BIGSERIAL PRIMARY KEY,
  username varchar(20) NOT NULL,
  ip varchar(45) DEFAULT NULL,
  os varchar(255) DEFAULT NULL,
  browser varchar(255) DEFAULT NULL,
  status smallint NOT NULL DEFAULT 1,
  message varchar(50) DEFAULT NULL,
  login_time timestamp NOT NULL
);

-- 创建索引
CREATE INDEX IF NOT EXISTS user_login_log_username_index ON user_login_log(username);

-- 创建 user_operation_log 表
CREATE TABLE IF NOT EXISTS user_operation_log (
  id BIGSERIAL PRIMARY KEY,
  username varchar(20) NOT NULL,
  method varchar(20) NOT NULL,
  router varchar(500) NOT NULL,
  service_name varchar(30) NOT NULL,
  ip varchar(45) DEFAULT NULL,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL,
  remark varchar(255) DEFAULT NULL
);

-- 创建索引
CREATE INDEX IF NOT EXISTS user_operation_log_username_index ON user_operation_log(username);

-- 创建 user_roles 表
CREATE TABLE IF NOT EXISTS user_roles (
  id BIGSERIAL PRIMARY KEY,
  user_id bigint NOT NULL,
  role_id bigint NOT NULL,
  created_at timestamp DEFAULT NULL,
  updated_at timestamp DEFAULT NULL
);