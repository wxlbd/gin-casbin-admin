-- 插入初始角色数据
INSERT INTO `roles` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`,
                    `remark`)
VALUES (1, '超级管理员', 'SuperAdmin', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '');

-- 插入初始用户数据
INSERT INTO `users` (`id`, `username`, `password`, `user_type`, `nickname`, `phone`, `email`, `avatar`, `signed`,
                    `status`, `login_ip`, `login_time`, `created_by`, `updated_by`, `created_at`, `updated_at`,
                    `remark`)
VALUES (1, 'admin', '$2y$10$T3Po5Ufu1pKiKczWqp.dbOOjmeZ4H3Oj0daATqlqXsZOvrRW2s2IS', '100', '创始人', '16858888988',
        'admin@adminmine.com', '', '广阔天地，大有所为', 1, '127.0.0.1', '2025-02-08 09:06:29', 0, 0,
        '2025-01-15 11:22:58', '2025-02-08 09:06:29', '');

-- 插入初始用户角色数据
INSERT INTO `user_roles` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`)
VALUES (1, 1, 1, '2025-01-15 11:22:58', '2025-01-15 11:22:58');

-- 插入初始菜单数据
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (1, 0, 1, 'menus.pureSysManagement', 'PureSystem', '/system', '', 13, '', 'ri:settings-3-line', '', '', '', '',
        '', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:49:41', '2025-02-05 14:38:47');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (2, 1, 1, 'menus.pureSystemMenu', 'SystemMenu', '/system/menu/index', '', 99, '', 'ep:menu', '', '', '', '', '',
        '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:50:12', '2025-01-24 18:50:12');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (3, 1, 1, 'menus.pureRole', 'SystemRole', '/system/role/index', '', 99, '', 'ri:admin-fill', '', '', '', '', '',
        '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:50:40', '2025-01-24 18:50:40');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (4, 1, 1, 'menus.pureUser', 'SystemUser', '/system/user/index', '', 99, '', 'ri:admin-line', '', '', '', '', '',
        '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:51:03', '2025-01-24 18:51:03');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (5, 2, 4, '菜单列表', '', '', '', 99, '', '', '', '', '', '', 'system:menu:list', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 10:50:14', '2025-01-25 10:50:14');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (6, 2, 4, '菜单添加', NULL, '', NULL, 99, NULL, NULL, NULL, NULL, NULL, NULL, 'system:menu:create', NULL, 1, 0,
        0, 0, 1, 0, 1, '2025-01-25 10:56:06', '2025-01-25 10:56:10');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (7, 2, 4, '菜单修改', NULL, '', NULL, 99, NULL, NULL, NULL, NULL, NULL, NULL, 'system:menu:update', NULL, 1, 0,
        0, 0, 1, 0, 1, '2025-01-25 10:57:14', NULL);
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (8, 2, 4, '菜单删除', '', '', '', 99, '', '', '', '', '', '', 'system:menu:delete', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 10:58:12', '2025-01-25 10:58:12');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (9, 3, 4, '角色列表', '', '', '', 99, '', '', '', '', '', '', 'system:role:list', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 10:58:57', '2025-01-25 10:58:57');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (10, 3, 4, '角色添加', '', '', '', 99, '', '', '', '', '', '', 'system:role:create', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:00:29', '2025-01-25 11:00:29');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (11, 3, 4, '角色修改', '', '', '', 99, '', '', '', '', '', '', 'system:role:update', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:00:59', '2025-01-25 11:00:59');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (12, 3, 4, '角色删除', '', '', '', 99, '', '', '', '', '', '', 'system:role:delete', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:09:18', '2025-01-25 11:09:18');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (13, 4, 4, '用户列表', '', '', '', 99, '', '', '', '', '', '', 'system:user:list', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:09:56', '2025-01-25 11:09:56');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (14, 4, 4, '用户添加', '', '', '', 99, '', '', '', '', '', '', 'system:user:create', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:12:23', '2025-01-25 11:12:23');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (15, 4, 4, '用户修改', '', '', '', 99, '', '', '', '', '', '', 'system:user:update', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:12:50', '2025-01-25 11:12:50');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (16, 4, 4, '用户删除', '', '', '', 99, '', '', '', '', '', '', 'system:user:delete', '', 1, 0, 0, 0, 1, 0, 1,
        '2025-01-25 11:13:25', '2025-01-25 11:13:25');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (17, 3, 4, '获取角色权限', '', '', '', 99, '', '', '', '', '', '', 'system:role:get:menus', '', 1, 0, 0, 0, 1, 0,
        1, '2025-01-25 15:13:28', '2025-01-25 15:13:28');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (18, 3, 4, '赋予角色权限', '', '', '', 99, '', '', '', '', '', '', 'system:role:set:menus', '', 1, 0, 0, 0, 1, 0,
        1, '2025-01-25 15:14:01', '2025-01-25 15:14:01');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (19, 4, 4, '获取用户角色', '', '', '', 99, '', '', '', '', '', '', 'system:user:get:roles', '', 1, 0, 0, 0, 1, 0,
        1, '2025-01-25 15:15:20', '2025-01-25 15:15:20');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (20, 4, 4, '赋予用户角色', '', '', '', 99, '', '', '', '', '', '', 'system:user:set:roles', '', 1, 0, 0, 0, 1, 0,
        1, '2025-01-25 15:15:45', '2025-01-25 15:15:45');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (21, 0, 1, 'menus.pureHome', 'Home', '/welcome', '', 1, '', 'ep:home-filled', '', '', '', '', '', '', 1, 0, 0, 0,
        1, 0, 1, '2025-02-06 09:54:19', '2025-02-06 09:55:24');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (22, 1, 1, '字典管理', 'sysDict', '/system/dict/index', '', 99, '', 'ep:memo', '', '', '', '', '', '', 1, 0, 0,
        0, 1, 0, 1, '2025-02-07 09:33:47', '2025-02-08 10:52:21');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`,
                         `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`,
                         `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`,
                         `show_parent`, `status`, `created_at`, `updated_at`)
VALUES (23, 1, 1, '字典数据', 'dictData', '/system/dict/dictData', '', 99, '', '', '', '', '', '', '', '', 1, 0, 0, 0,
        0, 0, 0, '2025-02-08 09:12:52', '2025-02-08 14:33:09');