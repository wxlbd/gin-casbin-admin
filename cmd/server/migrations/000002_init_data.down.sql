-- PostgreSQL 版本 - 清空初始数据

-- 由于有外键约束,需要先删除子表数据
DELETE FROM role_menus WHERE role_id = 1;
DELETE FROM user_roles WHERE id = 1;
DELETE FROM users WHERE id = 1;
DELETE FROM roles WHERE id = 1;
DELETE FROM sys_menus WHERE id BETWEEN 1 AND 23;