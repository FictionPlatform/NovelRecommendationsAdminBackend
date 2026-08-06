/*
 Navicat Premium Dump SQL

 Source Server         : my-server
 Source Server Type    : MySQL
 Source Server Version : 80040 (8.0.40)
 Source Host           : 127.0.0.1:3306
 Source Schema         : bitxxadmin

 Target Server Type    : MySQL
 Target Server Version : 80040 (8.0.40)
 File Encoding         : 65001

 Date: 23/12/2024 22:08:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_sys_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_api`;
CREATE TABLE `admin_sys_api` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '功能描述',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址',
  `api_type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '接口类型',
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=150 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='接口管理';

-- ----------------------------
-- Records of admin_sys_api
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (1, 'admin-获取字典类型分页列表', '/admin-api/v1/admin/sys/sys-dict/type', '1', 'GET', '2024-12-13 20:37:31', '2024-12-14 14:30:12', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (2, 'admin-获取字典类型全部列表', '/admin-api/v1/admin/sys/sys-dict/type/option-select', '1', 'GET', '2024-12-13 20:37:31', '2024-12-14 14:58:06', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (3, 'admin-导出字典类型', '/admin-api/v1/admin/sys/sys-dict/type/export', '1', 'GET', '2024-12-13 20:37:32', '2024-12-14 14:32:48', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (4, 'admin-获取字典类型详情', '/admin-api/v1/admin/sys/sys-dict/type/:id', '1', 'GET', '2024-12-13 20:37:32', '2024-12-14 14:32:25', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (5, 'admin-获取字典数据分页列表', '/admin-api/v1/admin/sys/sys-dict/data', '1', 'GET', '2024-12-13 20:37:32', '2024-12-14 14:36:52', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (6, 'admin-获取字典数据全部列表', '/admin-api/v1/admin/sys/sys-dict/data/select', '1', 'GET', '2024-12-13 20:37:32', '2024-12-15 22:40:03', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (7, 'admin-获取字典数据详情', '/admin-api/v1/admin/sys/sys-dict/data/:id', '1', 'GET', '2024-12-13 20:37:32', '2024-12-14 14:38:35', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (9, 'admin-获取部门管理�?, '/admin-api/v1/admin/sys/sys-dept/dept-tree', '1', 'GET', '2024-12-13 20:37:33', '2024-12-14 14:18:07', 0, 1, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (10, 'admin-根据角色获取部门', '/admin-api/v1/admin/sys/sys-dept/role-dept-tree-select/:roleId', '1', 'GET', '2024-12-13 20:37:33', '2024-12-14 14:05:10', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (11, 'admin-获取部门管理详情', '/admin-api/v1/admin/sys/sys-dept/:id', '1', 'GET', '2024-12-13 20:37:33', '2024-12-14 14:18:07', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (12, 'admin-获取菜单管理�?, '/admin-api/v1/admin/sys/sys-menu', '1', 'GET', '2024-12-13 20:37:33', '2024-12-14 14:13:21', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (13, 'admin-根据角色获取菜单', '/admin-api/v1/admin/sys/sys-menu/menu-role', '1', 'GET', '2024-12-13 20:37:33', '2024-12-15 22:39:47', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (14, 'admin-获取全部菜单以及选中的菜单编�?, '/admin-api/v1/admin/sys/sys-menu/role-menu-tree-select/:roleId', '1', 'GET', '2024-12-13 20:37:33', '2024-12-14 14:00:05', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (15, 'admin-获取菜单管理详情', '/admin-api/v1/admin/sys/sys-menu/:id', '1', 'GET', '2024-12-13 20:37:34', '2024-12-14 14:12:33', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (16, 'admin-获取服务器信�?, '/admin-api/v1/admin/sys/sys-monitor', '1', 'GET', '2024-12-13 20:37:34', '2024-12-14 15:12:52', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (17, 'admin-普罗米监�?, '/admin-api/v1/admin/sys/sys-monitor/prom', '1', 'GET', '2024-12-13 20:37:34', '2024-12-15 22:38:58', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (18, 'admin-ping', '/admin-api/v1/admin/sys/sys-monitor/ping', '1', 'GET', '2024-12-13 20:37:34', '2024-12-15 22:39:10', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (20, 'admin-获取表管理分页列�?, '/admin-api/v1/admin/sys/sys-table', '1', 'GET', '2024-12-13 20:37:35', '2024-12-13 21:45:39', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (21, 'admin-表管理下载代�?, '/admin-api/v1/admin/sys/sys-table/gen/download/:id', '1', 'GET', '2024-12-13 20:37:35', '2024-12-14 14:52:44', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (22, 'admin-表管理中生成菜单数据', '/admin-api/v1/admin/sys/sys-table/gen/db/:id', '1', 'GET', '2024-12-13 20:37:35', '2024-12-14 14:53:37', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (23, 'admin-生成表管理的代码', '/admin-api/v1/admin/sys/sys-table/gen/:id', '1', 'GET', '2024-12-13 20:37:35', '2024-12-14 14:52:39', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (24, 'admin-获取表管理的DB表分页列�?, '/admin-api/v1/admin/sys/sys-table/db-tables', '1', 'GET', '2024-12-13 20:37:35', '2024-12-15 22:16:01', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (25, 'admin-预览表管理的代码页面', '/admin-api/v1/admin/sys/sys-table/preview/:id', '1', 'GET', '2024-12-13 20:37:35', '2024-12-14 14:47:12', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (26, 'admin-获取表管理详�?, '/admin-api/v1/admin/sys/sys-table/:id', '1', 'GET', '2024-12-13 20:37:36', '2024-12-14 14:58:06', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (27, 'admin-获取接口管理分页列表', '/admin-api/v1/admin/sys/sys-api', '1', 'GET', '2024-12-13 20:37:36', '2024-12-14 13:06:50', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (28, 'admin-同步接口数据', '/admin-api/v1/admin/sys/sys-api/sync', '1', 'GET', '2024-12-13 20:37:36', '2024-12-14 13:10:41', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (30, 'admin-获取接口管理全部列表', '/admin-api/v1/admin/sys/sys-api/list', '1', 'GET', '2024-12-13 20:37:36', '2024-12-14 14:13:21', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (31, 'admin-导出接口管理', '/admin-api/v1/admin/sys/sys-api/export', '1', 'GET', '2024-12-13 20:37:36', '2024-12-14 14:22:31', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (32, 'admin-获取接口管理详情', '/admin-api/v1/admin/sys/sys-api/:id', '1', 'GET', '2024-12-13 20:37:37', '2024-12-14 13:07:22', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (33, 'admin-获取配置管理分页列表', '/admin-api/v1/admin/sys/sys-config', '1', 'GET', '2024-12-13 20:37:37', '2024-12-14 14:33:31', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (34, 'admin-导出配置管理', '/admin-api/v1/admin/sys/sys-config/export', '1', 'GET', '2024-12-13 20:37:37', '2024-12-14 14:34:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (36, 'admin-根据Key获取配置�?, '/admin-api/v1/admin/sys/sys-config/key/:configKey', '1', 'GET', '2024-12-13 20:37:37', '2024-12-15 22:38:41', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (37, 'admin-获取配置管理详情', '/admin-api/v1/admin/sys/sys-config/:id', '1', 'GET', '2024-12-13 20:37:37', '2024-12-14 14:34:00', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (38, 'admin-获取岗位管理分页列表', '/admin-api/v1/admin/sys/sys-post', '1', 'GET', '2024-12-13 20:37:38', '2024-12-14 14:24:08', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (39, 'admin-获取岗位管理全部列表', '/admin-api/v1/admin/sys/sys-post/list', '1', 'GET', '2024-12-13 20:37:38', '2024-12-14 13:50:46', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (40, 'admin-导出岗位管理', '/admin-api/v1/admin/sys/sys-post/export', '1', 'GET', '2024-12-13 20:37:38', '2024-12-13 20:37:38', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (41, 'admin-获取岗位管理详情', '/admin-api/v1/admin/sys/sys-post/:id', '1', 'GET', '2024-12-13 20:37:38', '2024-12-14 14:28:12', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (42, 'admin-获取系统用户管理分页列表', '/admin-api/v1/admin/sys/sys-user', '1', 'GET', '2024-12-13 20:37:38', '2024-12-14 13:14:41', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (43, 'admin-退出系�?, '/admin-api/v1/admin/sys/sys-user/logout', '1', 'GET', '2024-12-13 20:37:38', '2024-12-15 22:38:26', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (44, 'admin-获取系统登录用户信息', '/admin-api/v1/admin/sys/sys-user/profile', '1', 'GET', '2024-12-13 20:37:39', '2024-12-15 22:37:54', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (45, 'admin-获取系统用户管理详情', '/admin-api/v1/admin/sys/sys-user/:id', '1', 'GET', '2024-12-13 20:37:39', '2024-12-14 13:50:46', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (46, 'admin-获取登录日志分页列表', '/admin-api/v1/admin/sys/sys-login-log', '1', 'GET', '2024-12-13 20:37:39', '2024-12-14 14:35:28', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (47, 'admin-导出登录日志', '/admin-api/v1/admin/sys/sys-login-log/export', '1', 'GET', '2024-12-13 20:37:39', '2024-12-14 14:36:01', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (48, 'admin-获取登录日志详情', '/admin-api/v1/admin/sys/sys-login-log/:id', '1', 'GET', '2024-12-13 20:37:39', '2024-12-15 22:38:15', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (49, 'admin-获取操作日志分页列表', '/admin-api/v1/admin/sys/sys-oper-log', '1', 'GET', '2024-12-13 20:37:40', '2024-12-14 14:39:13', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (50, 'admin-导出操作日志', '/admin-api/v1/admin/sys/sys-oper-log/export', '1', 'GET', '2024-12-13 20:37:40', '2024-12-13 21:41:12', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (51, 'admin-获取操作日志详情', '/admin-api/v1/admin/sys/sys-oper-log/:id', '1', 'GET', '2024-12-13 20:37:40', '2024-12-15 22:37:36', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (52, 'admin-获取角色管理分页列表', '/admin-api/v1/admin/sys/sys-role', '1', 'GET', '2024-12-13 20:37:40', '2024-12-14 13:52:04', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (53, 'admin-获取角色管理全部列表', '/admin-api/v1/admin/sys/sys-role/list', '1', 'GET', '2024-12-13 20:37:40', '2024-12-14 13:50:46', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (54, 'admin-获取角色管理详情', '/admin-api/v1/admin/sys/sys-role/:id', '1', 'GET', '2024-12-13 20:37:40', '2024-12-14 14:05:10', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (55, 'app-获取用户管理分页列表', '/admin-api/v1/app/user/user', '3', 'GET', '2024-12-13 20:37:41', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (56, 'app-获取国家区号管理分页列表', '/admin-api/v1/app/user/user-country-code', '3', 'GET', '2024-12-13 20:37:41', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (57, 'app-导出国家区号管理', '/admin-api/v1/app/user/user-country-code/export', '3', 'GET', '2024-12-13 20:37:41', '2024-12-14 12:51:44', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (58, 'app-获取国家区号管理详情', '/admin-api/v1/app/user/user-country-code/:id', '3', 'GET', '2024-12-13 20:37:41', '2024-12-14 12:51:23', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (59, 'app-获取用户配置管理分页列表', '/admin-api/v1/app/user/user-conf', '3', 'GET', '2024-12-13 20:37:41', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (60, 'app-获取用户配置管理详情', '/admin-api/v1/app/user/user-conf/:id', '3', 'GET', '2024-12-13 20:37:41', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (61, 'app-获取账变记录分页列表', '/admin-api/v1/app/user/user-account-log', '3', 'GET', '2024-12-13 20:37:42', '2024-12-14 12:52:00', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (62, 'app-导出账变记录', '/admin-api/v1/app/user/user-account-log/export', '3', 'GET', '2024-12-13 20:37:42', '2024-12-14 12:52:11', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (63, 'app-获取账变记录详情', '/admin-api/v1/app/user/user-account-log/:id', '3', 'GET', '2024-12-13 20:37:42', '2024-12-15 22:37:25', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (64, 'app-获取用户等级管理分页列表', '/admin-api/v1/app/user/user-level', '3', 'GET', '2024-12-13 20:37:42', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (65, 'app-导出用户等级管理', '/admin-api/v1/app/user/user-level/export', '3', 'GET', '2024-12-13 20:37:42', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (66, 'app-获取用户等级管理详情', '/admin-api/v1/app/user/user-level/:id', '3', 'GET', '2024-12-13 20:37:42', '2024-12-14 12:49:29', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (67, 'app-获取用户操作日志分页列表', '/admin-api/v1/app/user/user-oper-log', '3', 'GET', '2024-12-13 20:37:43', '2024-12-14 12:50:33', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (68, 'app-导出用户操作日志', '/admin-api/v1/app/user/user-oper-log/export', '3', 'GET', '2024-12-13 20:37:43', '2024-12-14 12:50:48', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (69, 'app-获取用户操作日志详情', '/admin-api/v1/app/user/user-oper-log/:id', '3', 'GET', '2024-12-13 20:37:43', '2024-12-15 22:37:16', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (70, 'app-导出用户管理', '/admin-api/v1/app/user/user/export', '3', 'GET', '2024-12-13 20:37:43', '2024-12-13 20:56:34', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (71, 'app-获取用户管理详情', '/admin-api/v1/app/user/user/:id', '3', 'GET', '2024-12-13 20:37:43', '2024-12-14 12:48:02', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (72, 'plugins-获取公告管理分页列表', '/admin-api/v1/plugins/content/content-announcement', '2', 'GET', '2024-12-13 20:37:44', '2024-12-14 13:03:46', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (73, 'plugins-导出公告管理', '/admin-api/v1/plugins/content/content-announcement/export', '2', 'GET', '2024-12-13 20:37:44', '2024-12-13 21:06:33', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (74, 'plugins-获取公告管理详情', '/admin-api/v1/plugins/content/content-announcement/:id', '2', 'GET', '2024-12-13 20:37:44', '2024-12-14 13:04:08', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (75, 'plugins-获取文章管理分页列表', '/admin-api/v1/plugins/content/content-article', '2', 'GET', '2024-12-13 20:37:44', '2024-12-14 13:02:56', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (76, 'plugins-导出文章管理', '/admin-api/v1/plugins/content/content-article/export', '2', 'GET', '2024-12-13 20:37:44', '2024-12-13 21:04:27', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (77, 'plugins-获取文章管理详情', '/admin-api/v1/plugins/content/content-article/:id', '2', 'GET', '2024-12-13 20:37:44', '2024-12-14 13:03:18', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (78, 'plugins-获取内容分类管理分页列表', '/admin-api/v1/plugins/content/content-category', '2', 'GET', '2024-12-13 20:37:45', '2024-12-14 13:01:47', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (79, 'plugins-导出内容分类管理', '/admin-api/v1/plugins/content/content-category/export', '2', 'GET', '2024-12-13 20:37:45', '2024-12-14 13:02:34', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (80, 'plugins-获取内容分类管理详情', '/admin-api/v1/plugins/content/content-category/:id', '2', 'GET', '2024-12-13 20:37:45', '2024-12-14 13:02:14', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (81, 'plugins-获取APP管理分页列表', '/admin-api/v1/plugins/filemgr/filemgr-app', '2', 'GET', '2024-12-13 20:37:45', '2024-12-14 13:04:41', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (82, 'plugins-导出APP管理', '/admin-api/v1/plugins/filemgr/filemgr-app/export', '2', 'GET', '2024-12-13 20:37:45', '2024-12-13 21:08:56', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (83, 'plugins-获取APP管理详情', '/admin-api/v1/plugins/filemgr/filemgr-app/:id', '2', 'GET', '2024-12-13 20:37:45', '2024-12-15 22:15:04', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (84, 'plugins-获取验证码管理分页列�?, '/admin-api/v1/plugins/msg/msg-code', '2', 'GET', '2024-12-13 20:37:46', '2024-12-14 13:05:37', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (85, 'plugins-获取验证码管理详�?, '/admin-api/v1/plugins/msg/msg-code/:id', '2', 'GET', '2024-12-13 20:37:46', '2024-12-15 22:36:18', 0, 1, '未使�?);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (86, 'admin-获取图形验证�?, '/admin-api/v1/captcha', '1', 'GET', '2024-12-13 20:37:46', '2024-12-15 22:36:26', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (87, '静态文�?, '/admin-api/files/*filepath', '1', 'GET', '2024-12-13 20:37:46', '2024-12-15 22:36:57', 0, 1, '静态路由，该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (88, '静态文�?', '/static/*filepath', '1', 'GET', '2024-12-13 20:37:46', '2024-12-15 22:37:05', 0, 1, '静态路由，该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (89, 'admin-新增字典数据', '/admin-api/v1/admin/sys/sys-dict/data', '1', 'POST', '2024-12-13 20:37:46', '2024-12-14 14:38:04', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (90, 'admin-新增字典类型', '/admin-api/v1/admin/sys/sys-dict/type', '1', 'POST', '2024-12-13 20:37:47', '2024-12-14 14:31:05', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (91, 'admin-添加部门管理', '/admin-api/v1/admin/sys/sys-dept', '1', 'POST', '2024-12-13 20:37:47', '2024-12-14 14:17:17', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (92, 'admin-新增系统用户管理', '/admin-api/v1/admin/sys/sys-user', '1', 'POST', '2024-12-13 20:37:47', '2024-12-14 13:43:52', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (93, 'admin-更新系统登录用户头像', '/admin-api/v1/admin/sys/sys-user/profile/avatar', '1', 'POST', '2024-12-13 20:37:47', '2024-12-15 22:35:54', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (94, 'admin-新增配置管理', '/admin-api/v1/admin/sys/sys-config', '1', 'POST', '2024-12-13 20:37:47', '2024-12-14 14:33:44', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (95, 'admin-新增表管�?, '/admin-api/v1/admin/sys/sys-table', '1', 'POST', '2024-12-13 20:37:48', '2024-12-15 22:16:01', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (96, 'admin-新增菜单管理', '/admin-api/v1/admin/sys/sys-menu', '1', 'POST', '2024-12-13 20:37:48', '2024-12-14 14:13:21', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (97, 'admin-新增岗位管理', '/admin-api/v1/admin/sys/sys-post', '1', 'POST', '2024-12-13 20:37:48', '2024-12-14 14:27:41', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (98, 'admin-新增角色管理', '/admin-api/v1/admin/sys/sys-role', '1', 'POST', '2024-12-13 20:37:48', '2024-12-14 13:56:24', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (99, 'app-新增用户管理', '/admin-api/v1/app/user/user', '3', 'POST', '2024-12-13 20:37:48', '2024-12-14 12:47:50', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (100, 'app-新增国家区号管理', '/admin-api/v1/app/user/user-country-code', '3', 'POST', '2024-12-13 20:37:48', '2024-12-14 12:51:13', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (101, 'app-新增用户等级管理', '/admin-api/v1/app/user/user-level', '3', 'POST', '2024-12-13 20:37:49', '2024-12-14 12:49:18', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (102, 'plugins-新增公告管理', '/admin-api/v1/plugins/content/content-announcement', '2', 'POST', '2024-12-13 20:37:49', '2024-12-14 13:03:58', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (103, 'plugins-新增文章管理', '/admin-api/v1/plugins/content/content-article', '2', 'POST', '2024-12-13 20:37:49', '2024-12-14 13:03:09', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (104, 'plugins-新增内容分类管理详情', '/admin-api/v1/plugins/content/content-category', '2', 'POST', '2024-12-13 20:37:49', '2024-12-14 13:02:03', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (105, 'plugins-新增APP管理', '/admin-api/v1/plugins/filemgr/filemgr-app', '2', 'POST', '2024-12-13 20:37:49', '2024-12-15 22:14:53', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (106, 'plugins-上传APP', '/admin-api/v1/plugins/filemgr/filemgr-app/upload', '2', 'POST', '2024-12-13 20:37:49', '2024-12-15 22:15:04', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (107, 'admin-登录系统', '/admin-api/v1/login', '1', 'POST', '2024-12-13 20:37:50', '2024-12-15 22:35:12', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (108, 'admin-更新系统登录用户信息', '/admin-api/v1/admin/sys/sys-user/profile', '1', 'PUT', '2024-12-13 20:37:50', '2024-12-15 22:35:35', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (109, 'admin-更新系统登录用户密码', '/admin-api/v1/admin/sys/sys-user/profile/pwd', '1', 'PUT', '2024-12-13 20:37:50', '2024-12-15 22:35:39', 0, 1, '该接口无需角色校验');
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (110, 'admin-重置系统用户密码', '/admin-api/v1/admin/sys/sys-user/pwd-reset', '1', 'PUT', '2024-12-13 20:37:50', '2024-12-14 13:46:43', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (111, 'admin-更新系统用户状�?, '/admin-api/v1/admin/sys/sys-user/update-status', '1', 'PUT', '2024-12-13 20:37:50', '2024-12-14 13:50:46', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (112, 'admin-更新系统用户管理', '/admin-api/v1/admin/sys/sys-user/:id', '1', 'PUT', '2024-12-13 20:37:50', '2024-12-14 13:50:46', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (113, 'admin-更新字典数据', '/admin-api/v1/admin/sys/sys-dict/data/:id', '1', 'PUT', '2024-12-13 20:37:51', '2024-12-14 14:38:35', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (114, 'admin-更新字典类型', '/admin-api/v1/admin/sys/sys-dict/type/:id', '1', 'PUT', '2024-12-13 20:37:51', '2024-12-14 14:32:25', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (115, 'admin-更新部门管理', '/admin-api/v1/admin/sys/sys-dept/:id', '1', 'PUT', '2024-12-13 20:37:51', '2024-12-14 14:18:07', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (116, 'admin-更新角色管理状�?, '/admin-api/v1/admin/sys/sys-role/role-status', '1', 'PUT', '2024-12-13 20:37:51', '2024-12-14 14:00:05', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (117, 'admin-更新角色管理数据权限', '/admin-api/v1/admin/sys/sys-role/role-data-scope', '1', 'PUT', '2024-12-13 20:37:51', '2024-12-14 14:05:10', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (118, 'admin-更新角色管理', '/admin-api/v1/admin/sys/sys-role/:id', '1', 'PUT', '2024-12-13 20:37:51', '2024-12-14 14:00:05', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (119, 'admin-更新接口管理', '/admin-api/v1/admin/sys/sys-api/:id', '1', 'PUT', '2024-12-13 20:37:52', '2024-12-14 13:07:22', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (120, 'admin-更新配置管理', '/admin-api/v1/admin/sys/sys-config/:id', '1', 'PUT', '2024-12-13 20:37:52', '2024-12-14 14:34:00', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (121, 'admin-更新表管�?, '/admin-api/v1/admin/sys/sys-table/:id', '1', 'PUT', '2024-12-13 20:37:52', '2024-12-14 14:58:06', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (122, 'admin-更新菜单管理', '/admin-api/v1/admin/sys/sys-menu/:id', '1', 'PUT', '2024-12-13 20:37:52', '2024-12-14 14:12:33', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (123, 'admin-更新岗位管理', '/admin-api/v1/admin/sys/sys-post/:id', '1', 'PUT', '2024-12-13 20:37:52', '2024-12-14 14:28:12', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (124, 'app-更新用户配置管理', '/admin-api/v1/app/user/user-conf/:id', '3', 'PUT', '2024-12-13 20:37:52', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (125, 'app-更新国家区号管理', '/admin-api/v1/app/user/user-country-code/:id', '3', 'PUT', '2024-12-13 20:37:53', '2024-12-14 12:51:23', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (126, 'app-更新用户等级管理', '/admin-api/v1/app/user/user-level/:id', '3', 'PUT', '2024-12-13 20:37:53', '2024-12-14 12:49:29', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (127, 'app-更新用户管理', '/admin-api/v1/app/user/user/:id', '3', 'PUT', '2024-12-13 20:37:53', '2024-12-14 12:48:02', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (128, 'plugins-更新公告管理', '/admin-api/v1/plugins/content/content-announcement/:id', '2', 'PUT', '2024-12-13 20:37:53', '2024-12-14 13:04:08', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (129, 'plugins-更新文章管理', '/admin-api/v1/plugins/content/content-article/:id', '2', 'PUT', '2024-12-13 20:37:53', '2024-12-14 13:03:18', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (130, 'plugins-更新内容分类管理', '/admin-api/v1/plugins/content/content-category/:id', '2', 'PUT', '2024-12-13 20:37:54', '2024-12-14 13:02:14', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (131, 'plugins-更新APP管理', '/admin-api/v1/plugins/filemgr/filemgr-app/:id', '2', 'PUT', '2024-12-13 20:37:54', '2024-12-15 22:15:04', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (132, 'admin-删除字典数据', '/admin-api/v1/admin/sys/sys-dict/data', '1', 'DELETE', '2024-12-13 20:37:54', '2024-12-14 14:38:51', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (133, 'admin-删除字典类型', '/admin-api/v1/admin/sys/sys-dict/type', '1', 'DELETE', '2024-12-13 20:37:54', '2024-12-14 14:32:36', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (134, 'admin-删除部门管理', '/admin-api/v1/admin/sys/sys-dept', '1', 'DELETE', '2024-12-13 20:37:54', '2024-12-14 14:18:35', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (135, 'admin-删除接口管理', '/admin-api/v1/admin/sys/sys-api', '1', 'DELETE', '2024-12-13 20:37:54', '2024-12-14 13:07:36', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (136, 'admin-删除配置管理', '/admin-api/v1/admin/sys/sys-config', '1', 'DELETE', '2024-12-13 20:37:55', '2024-12-14 14:34:11', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (137, 'admin-删除表管�?, '/admin-api/v1/admin/sys/sys-table', '1', 'DELETE', '2024-12-13 20:37:55', '2024-12-14 14:53:10', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (138, 'admin-删除登录日志', '/admin-api/v1/admin/sys/sys-login-log', '1', 'DELETE', '2024-12-13 20:37:55', '2024-12-13 21:38:09', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (139, 'admin-删除菜单管理', '/admin-api/v1/admin/sys/sys-menu', '1', 'DELETE', '2024-12-13 20:37:55', '2024-12-14 14:12:48', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (140, 'admin-删除操作日志', '/admin-api/v1/admin/sys/sys-oper-log', '1', 'DELETE', '2024-12-13 20:37:55', '2024-12-14 14:40:07', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (141, 'admin-删除岗位管理', '/admin-api/v1/admin/sys/sys-post', '1', 'DELETE', '2024-12-13 20:37:55', '2024-12-13 21:32:47', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (142, 'admin-删除角色管理', '/admin-api/v1/admin/sys/sys-role', '1', 'DELETE', '2024-12-13 20:37:56', '2024-12-14 14:00:31', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (143, 'admin-删除系统用户管理', '/admin-api/v1/admin/sys/sys-user', '1', 'DELETE', '2024-12-13 20:37:56', '2024-12-14 13:50:52', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (144, 'app-删除国家区号管理', '/admin-api/v1/app/user/user-country-code', '3', 'DELETE', '2024-12-13 20:37:56', '2024-12-14 15:16:20', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (145, 'app-删除用户等级管理', '/admin-api/v1/app/user/user-level', '3', 'DELETE', '2024-12-13 20:37:56', '2024-12-14 12:49:38', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (146, 'plugins-删除公告管理', '/admin-api/v1/plugins/content/content-announcement', '2', 'DELETE', '2024-12-13 20:37:56', '2024-12-13 21:06:13', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (147, 'plugins-删除文章管理', '/admin-api/v1/plugins/content/content-article', '2', 'DELETE', '2024-12-13 20:37:56', '2024-12-14 13:03:29', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (148, 'plugins-删除内容分类管理', '/admin-api/v1/plugins/content/content-category', '2', 'DELETE', '2024-12-13 20:37:57', '2024-12-14 13:02:24', 0, 0, NULL);
INSERT INTO `admin_sys_api` (`id`, `description`, `path`, `api_type`, `method`, `created_at`, `updated_at`, `create_by`, `update_by`, `remark`) VALUES (149, 'plugins-删除APP管理', '/admin-api/v1/plugins/filemgr/filemgr-app', '2', 'DELETE', '2024-12-13 20:37:57', '2024-12-13 21:08:39', 0, 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_casbin_rule`;
CREATE TABLE `admin_sys_casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  UNIQUE KEY `idx_admin_sys_casbin_rule` (`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of admin_sys_casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user', 'GET', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-conf', 'GET', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-conf/:id', 'GET', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-conf/:id', 'PUT', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-country-code', 'DELETE', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-country-code', 'GET', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-level', 'GET', '', '', '');
INSERT INTO `admin_sys_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'test', '/admin-api/v1/app/user/user-level/export', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_config
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_config`;
CREATE TABLE `admin_sys_config` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigType',
  `is_frontend` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark',
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置管理';

-- ----------------------------
-- Records of admin_sys_config
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '管理-皮肤样式', 'admin_sys_index_skinName', 'skin-green', '1', '1', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿�?skin-green、紫�?skin-purple、红�?skin-red、黄�?skin-yellow', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:02');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '管理-侧栏主题', 'admin_sys_index_sideTheme', 'theme-dark', '1', '1', '主框架页-侧边栏主�?深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:06');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '管理-系统名称', 'admin_sys_app_name', 'go-admin后台管理系统', '1', '1', '', 1, 1, '2021-03-17 08:52:06', '2023-03-11 23:16:19');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '管理-系统logo', 'admin_sys_app_logo', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', '1', '', 1, 1, '2021-03-17 08:53:19', '2023-03-11 23:16:15');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '管理-单次excel导出数据�?, 'admin_sys_max_export_size', '10000', '1', '1', '', 0, 1, '2021-07-28 16:53:48', '2023-03-11 23:15:56');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '插件-文件管理-App OSS Bucket', 'plugin_filemgr_app_oss_bucket', '请自行配�?, '2', '2', '', 0, 1, '2021-08-13 14:36:23', '2023-03-11 23:14:45');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '插件-文件管理-App OSS AccessKeyId', 'plugin_filemgr_app_oss_access_key_id', '请自行配�?, '2', '2', '', 0, 1, '2021-08-13 14:37:15', '2023-03-11 23:14:41');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '插件-文件管理-App OSS AccessKeySecret', 'plugin_filemgr_app_oss_access_key_secret', '请自行配�?, '2', '2', '', 0, 1, '2021-08-13 14:38:00', '2023-03-11 23:14:33');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '插件-文件管理-App OSS Endpoint', 'plugin_filemgr_app_oss_endpoint', '请自行配�?, '2', '2', '', 0, 1, '2021-08-13 14:38:50', '2023-03-11 23:14:28');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '插件-文件管理-App OSS 根目�?, 'plugin_filemgr_app_oss_root_path', 'testfile/', '2', '2', '', 0, 1, '2021-08-13 14:39:31', '2023-03-11 23:14:22');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '管理-用户-默认头像', 'admin_sys_user_default_avatar', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '3', '2', '', 1, 1, '2023-03-10 18:07:03', '2023-03-10 18:07:03');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dept`;
CREATE TABLE `admin_sys_dept` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `leader` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='部门管理';

-- ----------------------------
-- Records of admin_sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 0, '0,', 'Admin', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:25');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 1, '0,1,', '研发�?, 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 1, '0,1,', '运维�?, 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2024-12-13 20:17:47');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 1, '0,1,', '客服�?, 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2024-12-13 20:17:47');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 1, '0,1,', '人力资源', 3, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:53');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 1, '0,1,', '市场', 10, 'admin', '', '', 1, 1, 1, '2021-12-02 10:13:38', '2021-12-02 10:13:38');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dict_data`;
CREATE TABLE `admin_sys_dict_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `dict_sort` int DEFAULT NULL,
  `dict_label` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `css_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `list_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `default_val` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='字典数据管理';

-- 字典数据�?dict_type 查询是主路径
ALTER TABLE `admin_sys_dict_data` ADD INDEX `idx_admin_sys_dict_data_dict_type` (`dict_type`);

-- ----------------------------
-- Records of admin_sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 0, '正常', '2', 'admin_sys_normal_disable', '', '', '', '0', '', '系统正常', 1, 1, '2021-05-13 19:56:38', '2022-04-25 00:42:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 0, '停用', '1', 'admin_sys_normal_disable', '', '', '', '0', '', '系统停用', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 0, '�?, '1', 'admin_sys_user_sex', '', '', '', '0', '', '性别�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 0, '�?, '2', 'admin_sys_user_sex', '', '', '', '0', '', '性别�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 0, '未知', '3', 'admin_sys_user_sex', '', '', '', '0', '', '性别未知', 1, 1, '2021-05-13 19:56:38', '2023-03-05 12:03:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 0, '显示', '2', 'admin_sys_menu_show_hide', '', '', '', '0', '', '显示菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, 0, '隐藏', '1', 'admin_sys_menu_show_hide', '', '', '', '0', '', '隐藏菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, 0, '�?, '1', 'admin_sys_yes_no', '', '', '', '0', '', '系统默认�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, 0, '�?, '2', 'admin_sys_yes_no', '', '', '', '0', '', '系统默认�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, 0, '通知', '1', 'admin_sys_notice_type', '', '', '', '0', '', '通知', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, 0, '公告', '2', 'admin_sys_notice_type', '', '', '', '0', '', '公告', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, 0, '正常', '2', 'admin_sys_common_status', '', '', '', '0', '', '正常状�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, 0, '关闭', '1', 'admin_sys_common_status', '', '', '', '0', '', '关闭状�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, 0, '新增', '1', 'admin_sys_oper_type', '', '', '', '0', '', '新增操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, 0, '修改', '2', 'admin_sys_oper_type', '', '', '', '0', '', '修改操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, 0, '删除', '3', 'admin_sys_oper_type', '', '', '', '0', '', '删除操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, 0, '授权', '4', 'admin_sys_oper_type', '', '', '', '0', '', '授权操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, 0, '导出', '5', 'admin_sys_oper_type', '', '', '', '0', '', '导出操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, 0, '导入', '6', 'admin_sys_oper_type', '', '', '', '0', '', '导入操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, 0, '强退', '7', 'admin_sys_oper_type', '', '', '', '0', '', '强退操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, 0, '生成代码', '8', 'admin_sys_oper_type', '', '', '', '0', '', '生成操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, 0, '清空数据', '9', 'admin_sys_oper_type', '', '', '', '0', '', '清空操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, 0, '成功', '1', 'admin_sys_notice_status', '', '', '', '0', '', '成功状�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, 0, '失败', '2', 'admin_sys_notice_status', '', '', '', '0', '', '失败状�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, 0, '登录', '10', 'admin_sys_oper_type', '', '', '', '0', '', '登录操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, 0, '退�?, '11', 'admin_sys_oper_type', '', '', '', '0', '', '', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, 0, '获取验证�?, '12', 'admin_sys_oper_type', '', '', '', '0', '', '获取验证�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, 0, '正常', '1', 'admin_sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:01', '2021-07-09 11:40:01');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, 0, '停用', '2', 'admin_sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:14', '2021-07-09 11:40:14');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, 0, '安卓', '1', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:39', '2021-08-13 13:35:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, 0, 'IOS', '2', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:51', '2021-08-13 13:35:51');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, 0, '类型1', '1', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:07', '2021-08-13 13:37:07');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, 0, '类型2', '2', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:19', '2021-08-13 13:37:19');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (34, 0, '类型3', '3', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:39', '2021-08-13 13:37:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (35, 0, '本地', '1', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (36, 0, '外链', '2', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (37, 0, 'OSS', '3', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:33', '2021-08-13 14:02:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (38, 0, '已发�?, '2', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:47', '2021-12-09 12:42:47');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (39, 0, '待发�?, '1', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:54', '2021-12-09 12:42:54');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (42, 0, 'GET', 'GET', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:26', '2022-04-26 00:03:26');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (43, 0, 'POST', 'POST', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:40', '2022-04-26 00:03:40');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (44, 0, 'DELETE', 'DELETE', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:49', '2022-04-26 00:03:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (45, 0, 'PUT', 'PUT', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:04:06', '2022-04-26 00:04:06');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (46, 0, 'HEAD', 'HEAD', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:07:02', '2022-04-26 00:07:02');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (47, 0, '管理', '1', 'admin_sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:23', '2024-12-13 20:11:07');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (48, 0, '插件', '2', 'admin_sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:32', '2023-03-01 11:05:32');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (49, 0, '应用', '3', 'admin_sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:42', '2023-03-01 11:05:42');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (50, 0, '展示', '1', 'admin_sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:49', '2023-03-01 11:07:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (51, 0, '隐藏', '2', 'admin_sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:56', '2023-03-01 11:07:56');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (52, 0, '登录', '1', 'admin_sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:04', '2023-03-01 14:43:04');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (53, 0, '退�?, '2', 'admin_sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:10', '2023-03-01 14:43:10');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (55, 0, '全部数据权限', '1', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:36', '2023-03-04 13:29:36');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (56, 0, '自定数据权限', '2', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:43', '2023-03-04 13:29:43');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (57, 0, '本部门数据权�?, '3', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:49', '2023-03-04 13:29:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (58, 0, '本部门及以下数据权限', '4', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:56', '2023-03-04 13:29:56');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (59, 0, '仅本人数据权�?, '5', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:30:04', '2023-03-04 13:30:04');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (60, 0, 'int64', 'int64', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:08:26', '2023-03-07 10:08:26');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (61, 0, 'int', 'int', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:12:42', '2023-03-07 10:12:42');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (62, 0, 'string', 'string', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:05', '2023-03-07 10:13:05');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (63, 0, 'decimal', 'decimal.Decimal', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:16', '2023-03-07 10:13:29');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (64, 0, 'time', '*time.Time', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:43', '2023-03-07 10:13:43');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (65, 0, '=', 'EQ', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:20:53', '2023-03-07 10:20:53');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (66, 0, '!=', 'NE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:06', '2023-03-07 10:21:06');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (67, 0, '>', 'GT', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:20', '2023-03-07 10:21:20');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (68, 0, '>=', 'GTE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:33', '2023-03-07 10:21:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (69, 0, '<', 'LT', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:45', '2023-03-07 10:21:45');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (70, 0, '<=', 'LTE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:57', '2023-03-07 10:21:57');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (71, 0, 'LIKE', 'LIKE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:22:08', '2023-03-07 10:22:08');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (72, 0, '文本�?, 'input', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:39', '2023-03-07 10:23:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (73, 0, '下拉�?, 'select', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:49', '2023-03-07 10:23:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (74, 0, '单选框', 'radio', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:59', '2023-03-07 10:23:59');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (75, 0, '文本�?, 'textarea', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:24:08', '2023-03-07 10:24:08');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (76, 0, '目录', '1', 'admin_sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:00', '2023-03-08 10:42:14');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (77, 0, '菜单', '2', 'admin_sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:10', '2023-03-08 10:42:10');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (78, 0, '按钮', '3', 'admin_sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:22', '2023-03-08 10:42:22');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (79, 0, '类型1', '1', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:55:57', '2023-03-08 11:55:57');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (80, 0, '类型2', '2', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:56:02', '2023-03-08 11:56:02');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (81, 0, '数字文本�?, 'numInput', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:12:33', '2023-03-09 20:12:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (82, 0, 'CNY', '1', 'app_money_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:24:26', '2023-03-09 20:24:26');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (83, 0, '类型1', '1', 'app_account_change_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:27:45', '2023-03-09 20:27:45');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (84, 0, '允许用户登录', '1', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:01', '2023-03-11 14:08:01');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (85, 0, '禁止用户登录', '2', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:10', '2023-03-11 14:08:10');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (86, 0, '后台用户', '1', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:41', '2023-03-11 14:14:41');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (87, 0, '前台用户', '2', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:59', '2023-03-11 14:14:59');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (88, 0, '发送成�?, '1', 'plugin_msg_sendstatus', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:22', '2023-09-26 10:42:22');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (89, 0, '发送失�?, '2', 'plugin_msg_sendstatus', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:31', '2023-09-26 10:42:31');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (90, 0, '邮箱', '1', 'plugin_msg_code_type', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:58', '2023-09-26 10:42:58');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (91, 0, '短信', '2', 'plugin_msg_code_type', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:43:04', '2023-09-26 10:43:04');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (92, 0, '待发�?, '1', 'plugin_filemgr_publish_status', '', '', '', '1', '', '', 1, 1, '2024-12-01 23:20:36', '2024-12-01 23:20:36');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (93, 0, '已发�?, '2', 'plugin_filemgr_publish_status', '', '', '', '1', '', '', 1, 1, '2024-12-01 23:20:45', '2024-12-01 23:20:45');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (94, 0, '数字文本�?, 'numInput', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:39', '2023-03-07 10:23:39');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dict_type`;
CREATE TABLE `admin_sys_dict_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='字典类型管理';

-- ----------------------------
-- Records of admin_sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '管理-开�?, 'admin_sys_normal_disable', '0', '系统开关列�?, 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:35');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '管理-用户性别', 'admin_sys_user_sex', '0', '用户性别列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:06');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '管理-菜单状�?, 'admin_sys_menu_show_hide', '0', '菜单状态列�?, 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:02');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '管理-是否', 'admin_sys_yes_no', '0', '系统是否列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:58');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '管理-通知类型', 'admin_sys_notice_type', '0', '通知类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:53');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '管理-状�?, 'admin_sys_common_status', '0', '登录状态列�?, 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:49');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '管理-操作类型', 'admin_sys_oper_type', '0', '操作类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:42');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '管理-通知状�?, 'admin_sys_notice_status', '0', '通知状态列�?, 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:39');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '管理-基本状�?, 'admin_sys_status', '0', '基本通用状�?, 1, 1, '2021-07-09 11:39:21', '2023-03-11 23:21:23');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '插件-文件管理-App发布状�?, 'plugin_filemgr_publish_status', '2', '', 1, 1, '2021-12-09 12:42:31', '2023-03-11 23:20:01');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '插件-文件管理-App系统平台', 'plugin_filemgr_app_platform', '0', 'App系统平台', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:17');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, '插件-文件管理-App类型', 'plugin_filemgr_app_type', '0', 'app属�?, 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:13');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, '插件-文件管理-App下载类型', 'plugin_filemgr_app_download_type', '0', '', 1, 1, '2021-08-13 14:02:03', '2023-03-11 23:20:06');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '管理-接口-请求方法', 'admin_sys_api_method', '0', '', 1, 1, '2022-04-26 00:03:11', '2023-03-01 21:56:41');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '管理-配置-类型', 'admin_sys_config_type', '1', '1-管理 2-插件 3-应用', 1, 1, '2023-03-01 11:04:56', '2024-12-13 20:10:17');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '管理-配置-是否前台展示', 'admin_sys_config_is_frontend', '1', '1-展示 2-隐藏', 1, 1, '2023-03-01 11:06:28', '2023-03-01 11:08:07');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '管理-登录日志-日志状�?, 'admin_sys_loginlog_status', '1', '1-登录 2-退�?, 1, 1, '2023-03-01 14:42:56', '2023-03-01 14:42:56');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '管理-角色-数据范围', 'admin_sys_role_data_scope', '1', '1-全部数据权限 2- 自定义数据权�?3-本部门数据权�?4-本部门及以下数据权限 5-仅本人数据权�?, 1, 1, '2023-03-04 13:29:21', '2023-03-04 13:29:21');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '管理-模板-go类型', 'admin_sys_gen_go_type', '1', '', 1, 1, '2023-03-07 10:08:07', '2023-03-07 10:08:07');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '管理-模板-查询类型', 'admin_sys_gen_query_type', '1', '', 1, 1, '2023-03-07 10:20:19', '2023-03-07 10:20:19');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '管理-模板-显示类型', 'admin_sys_gen_html_type', '1', '', 1, 1, '2023-03-07 10:23:23', '2023-03-07 10:23:23');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '管理-菜单-类型', 'admin_sys_menu_type', '1', '', 1, 1, '2023-03-08 10:33:32', '2023-03-08 10:33:32');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, '应用-用户-等级', 'app_user_level_type', '1', '', 1, 1, '2023-03-08 11:44:48', '2023-03-08 11:44:48');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, '应用-用户-资产-资金类型', 'app_money_type', '1', '1-CNY', 1, 1, '2023-03-09 20:24:17', '2023-03-11 14:06:46');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, '应用-用户-资产-账变类型', 'app_account_change_type', '1', '1-类型1', 1, 1, '2023-03-09 20:27:33', '2023-03-11 14:06:38');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, '应用-用户-行为类型', 'app_user_action_type', '1', '', 1, 1, '2023-03-11 14:06:29', '2023-03-11 14:06:29');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, '应用-用户-用户更新类型', 'app_user_by_type', '1', '', 1, 1, '2023-03-11 14:14:06', '2023-03-11 14:14:27');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, '插件-消息-验证码类�?, 'plugin_msg_code_type', '1', '1-邮箱 2-短信', 1, 1, '2023-03-12 12:12:30', '2023-03-12 12:15:20');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, '插件-消息-验证码发送状�?, 'plugin_msg_sendstatus', '1', '', 1, 1, '2023-03-12 12:14:56', '2023-03-12 13:23:37');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_gen_column
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_gen_column`;
CREATE TABLE `admin_sys_gen_column` (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_id` int DEFAULT NULL,
  `column_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `column_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `column_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `go_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `go_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `json_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_pk` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_required` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '该值是否参与新增或者编�?,
  `is_list` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '列表',
  `is_query` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `query_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `html_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=288 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='db字段管理';

-- ----------------------------
-- Records of admin_sys_gen_column
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 1, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 1, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 1, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 1, 'json_result', '返回数据', 'varchar(255)', 'string', 'JsonResult', 'jsonResult', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (5, 1, 'latency_time', '耗时', 'varchar(128)', 'string', 'LatencyTime', 'latencyTime', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (6, 1, 'oper_ip', '客户端ip', 'varchar(128)', 'string', 'OperIp', 'operIp', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (7, 1, 'oper_location', '访问位置', 'varchar(128)', 'string', 'OperLocation', 'operLocation', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (8, 1, 'oper_time', '操作时间', 'timestamp', '*time.Time', 'OperTime', 'operTime', '2', '1', '2', '2', 'EQ', 'datetime', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (9, 1, 'oper_url', '访问地址', 'varchar(255)', 'string', 'OperUrl', 'operUrl', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (10, 1, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (11, 1, 'request_method', '请求方式', 'varchar(128)', 'string', 'RequestMethod', 'requestMethod', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (12, 1, 'status', '操作状�?, 'varchar(4)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (13, 1, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (14, 1, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (15, 1, 'user_agent', 'ua', 'varchar(255)', 'string', 'UserAgent', 'userAgent', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (16, 1, 'user_id', '操作�?, 'int', 'int64', 'UserId', 'userId', '2', '2', '2', '2', 'EQ', 'input', '', 16, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (17, 2, 'avatar', '头像路径', 'varchar(1000)', 'string', 'Avatar', 'avatar', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (18, 2, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (19, 2, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (20, 2, 'email', '电子邮箱', 'varchar(300)', 'string', 'Email', 'email', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (21, 2, 'id', '用户编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (22, 2, 'level_id', '用户等级编号', 'int', 'int64', 'LevelId', 'levelId', '2', '1', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (23, 2, 'mobile', '手机号码', 'varchar(100)', 'string', 'Mobile', 'mobile', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (24, 2, 'mobile_title', '用户手机号国家前缀', 'varchar(255)', 'string', 'MobileTitle', 'mobileTitle', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (25, 2, 'money', '余额', 'decimal(30,18)', 'decimal.Decimal', 'Money', 'money', '2', '1', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (26, 2, 'parent_id', '父级编号', 'int', 'int64', 'ParentId', 'parentId', '2', '1', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (27, 2, 'parent_ids', '所有父级编�?, 'varchar(1000)', 'string', 'ParentIds', 'parentIds', '2', '1', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (28, 2, 'pay_pwd', '提现密码', 'varchar(100)', 'string', 'PayPwd', 'payPwd', '2', '1', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (29, 2, 'pwd', '登录密码', 'varchar(100)', 'string', 'Pwd', 'pwd', '2', '1', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (30, 2, 'ref_code', '推荐�?, 'varchar(255)', 'string', 'RefCode', 'refCode', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (31, 2, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (32, 2, 'status', '状�?1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 16, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (33, 2, 'tree_leaf', '是否最末级', 'char(1)', 'string', 'TreeLeaf', 'treeLeaf', '2', '1', '2', '2', 'EQ', 'input', '', 17, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (34, 2, 'tree_level', '层次级别', 'int', 'int64', 'TreeLevel', 'treeLevel', '2', '1', '2', '2', 'EQ', 'input', '', 18, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (35, 2, 'tree_sort', '本级排序号（升序�?, 'int', 'int64', 'TreeSort', 'treeSort', '2', '1', '2', '2', 'EQ', 'input', '', 19, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (36, 2, 'tree_sorts', '所有级别排序号', 'varchar(1000)', 'string', 'TreeSorts', 'treeSorts', '2', '1', '2', '2', 'EQ', 'input', '', 20, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (37, 2, 'true_name', '真实姓名', 'varchar(100)', 'string', 'TrueName', 'trueName', '2', '1', '2', '2', 'EQ', 'input', '', 21, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (38, 2, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 22, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (39, 2, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 23, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (40, 2, 'user_name', '用户昵称', 'varchar(100)', 'string', 'UserName', 'userName', '2', '1', '2', '2', 'EQ', 'input', '', 24, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (41, 3, 'after_money', '账变后金�?, 'decimal(30,18)', 'decimal.Decimal', 'AfterMoney', 'afterMoney', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (42, 3, 'before_money', '账变前金�?, 'decimal(30,18)', 'decimal.Decimal', 'BeforeMoney', 'beforeMoney', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (43, 3, 'change_money', '账变金额', 'decimal(10,2)', 'decimal.Decimal', 'ChangeMoney', 'changeMoney', '2', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (44, 3, 'change_type', '帐变类型(1-类型1)', 'varchar(30)', 'string', 'ChangeType', 'changeType', '2', '1', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (45, 3, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (46, 3, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (47, 3, 'id', '账变编号', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (48, 3, 'money_type', '金额类型 1:余额 ', 'char(10)', 'string', 'MoneyType', 'moneyType', '2', '1', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (49, 3, 'remarks', '备注信息', 'varchar(500)', 'string', 'Remarks', 'remarks', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (50, 3, 'status', '状态（1正常 2-异常�?, 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (51, 3, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (52, 3, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (53, 3, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '1', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (54, 4, 'can_login', '1-允许登陆�?-不允许登�?, 'char(1)', 'string', 'CanLogin', 'canLogin', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (55, 4, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (56, 4, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (57, 4, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (58, 4, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (59, 4, 'status', '状态（1-正常 2-异常）\n', 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (60, 4, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (61, 4, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (62, 4, 'user_id', '用户id', 'int', 'int64', 'UserId', 'userId', '2', '1', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (63, 5, 'code', '区号', 'varchar(12)', 'string', 'Code', 'code', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (64, 5, 'country', '国家或地�?, 'varchar(64)', 'string', 'Country', 'country', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (65, 5, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (66, 5, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (67, 5, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (68, 5, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (69, 5, 'status', '状�?1-可用 2-停用)', 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (70, 5, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (71, 5, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (72, 6, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (73, 6, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (74, 6, 'id', '主键', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (75, 6, 'level', '等级', 'int', 'int64', 'Level', 'level', '2', '1', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (76, 6, 'level_type', '等级类型', 'varchar(10)', 'string', 'LevelType', 'levelType', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (77, 6, 'name', '等级名称', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (78, 6, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (79, 6, 'status', '状�?1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (80, 6, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (81, 6, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (82, 7, 'action_type', '用户行为类型', 'char(2)', 'string', 'ActionType', 'actionType', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (83, 7, 'by_type', '更新用户类型 1-app用户 2-后台用户', 'char(2)', 'string', 'ByType', 'byType', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (84, 7, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (85, 7, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (86, 7, 'id', '日志编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (87, 7, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (88, 7, 'status', '状�?1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (89, 7, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (90, 7, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (91, 7, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '1', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (92, 8, 'content', '内容', 'text', 'string', 'Content', 'content', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (93, 8, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (94, 8, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (95, 8, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (96, 8, 'num', '阅读次数', 'int', 'int64', 'Num', 'num', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (97, 8, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (98, 8, 'status', '状态（0正常 1删除 2停用 3冻结�?, 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (99, 8, 'title', '标题', 'varchar(255)', 'string', 'Title', 'title', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (100, 8, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (101, 8, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (102, 9, 'cate_id', '分类编号', 'int', 'int64', 'CateId', 'cateId', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (103, 9, 'content', '内容', 'text', 'string', 'Content', 'content', '2', '2', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (104, 9, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (105, 9, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (106, 9, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (107, 9, 'name', '名称', 'varchar(255)', 'string', 'Name', 'name', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (108, 9, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (109, 9, 'status', '状态（1-正常 2-异常�?, 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (110, 9, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (111, 9, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (112, 10, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (113, 10, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (114, 10, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (115, 10, 'name', '名称', 'varchar(255)', 'string', 'Name', 'name', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (116, 10, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (117, 10, 'status', '状态（1-正常 2-异常�?, 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (118, 10, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (119, 10, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (120, 11, 'app_type', '版本(1-默认)', 'char(1)', 'string', 'AppType', 'appType', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (121, 11, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (122, 11, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (123, 11, 'download_type', '下载类型(1-本地 2-外链 3-oss )', 'char(1)', 'string', 'DownloadType', 'downloadType', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (124, 11, 'download_url', '下载地址(download_type=1使用)', 'varchar(255)', 'string', 'DownloadUrl', 'downloadUrl', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (125, 11, 'id', '主键', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (126, 11, 'local_address', '本地地址', 'varchar(255)', 'string', 'LocalAddress', 'localAddress', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (127, 11, 'platform', '平台 (1-安卓 2-苹果)', 'char(1)', 'string', 'Platform', 'platform', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (128, 11, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (129, 11, 'status', '状态（1-已发�?2-待发布）\n', 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (130, 11, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (131, 11, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (132, 11, 'version', '版本�?, 'varchar(100)', 'string', 'Version', 'version', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (133, 12, 'code', '验证�?, 'varchar(12)', 'string', 'Code', 'code', '2', '1', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (134, 12, 'code_type', '验证码类�?1-邮箱�?-短信', 'char(1)', 'string', 'CodeType', 'codeType', '2', '1', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (135, 12, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (136, 12, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (137, 12, 'id', '验证码编�?, 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (138, 12, 'remark', '备注异常', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (139, 12, 'status', '验证码状�?1-发送成�?2-发送失�?, 'char(1)', 'string', 'Status', 'status', '2', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (140, 12, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '1', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (141, 12, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '1', '2', '2', 'EQ', 'datetime', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (142, 12, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '1', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (143, 13, 'api_type', '接口类型', 'varchar(16)', 'string', 'ApiType', 'apiType', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (144, 13, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (145, 13, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (146, 13, 'description', '功能描述', 'varchar(256)', 'string', 'Description', 'description', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (147, 13, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (148, 13, 'method', '请求类型', 'varchar(32)', 'string', 'Method', 'method', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (149, 13, 'path', '地址', 'varchar(128)', 'string', 'Path', 'path', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (150, 13, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (151, 13, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (152, 14, 'config_key', 'ConfigKey', 'varchar(128)', 'string', 'ConfigKey', 'configKey', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (153, 14, 'config_name', 'ConfigName', 'varchar(128)', 'string', 'ConfigName', 'configName', '2', '2', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (154, 14, 'config_type', 'ConfigType', 'varchar(64)', 'string', 'ConfigType', 'configType', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (155, 14, 'config_value', 'ConfigValue', 'varchar(255)', 'string', 'ConfigValue', 'configValue', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (156, 14, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (157, 14, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (158, 14, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (159, 14, 'is_frontend', '是否前台', 'char(1)', 'string', 'IsFrontend', 'isFrontend', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (160, 14, 'remark', 'Remark', 'varchar(128)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (161, 14, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (162, 14, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (163, 15, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (164, 15, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (165, 15, 'dept_name', '', 'varchar(128)', 'string', 'DeptName', 'deptName', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (166, 15, 'email', '', 'varchar(80)', 'string', 'Email', 'email', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (167, 15, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (168, 15, 'leader', '', 'varchar(128)', 'string', 'Leader', 'leader', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (169, 15, 'parent_id', '', 'int', 'int64', 'ParentId', 'parentId', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (170, 15, 'parent_ids', '', 'varchar(255)', 'string', 'ParentIds', 'parentIds', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (171, 15, 'phone', '', 'varchar(20)', 'string', 'Phone', 'phone', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (172, 15, 'sort', '', 'int', 'int64', 'Sort', 'sort', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (173, 15, 'status', '', 'tinyint', 'int64', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (174, 15, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (175, 15, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (176, 16, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (177, 16, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (178, 16, 'css_class', '', 'varchar(128)', 'string', 'CssClass', 'cssClass', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (179, 16, 'default_val', '', 'varchar(8)', 'string', 'Default', 'default', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (180, 16, 'dict_label', '', 'varchar(128)', 'string', 'DictLabel', 'dictLabel', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (181, 16, 'dict_sort', '', 'int', 'int64', 'DictSort', 'dictSort', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (182, 16, 'dict_type', '', 'varchar(64)', 'string', 'DictType', 'dictType', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (183, 16, 'dict_value', '', 'varchar(255)', 'string', 'DictValue', 'dictValue', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (184, 16, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (185, 16, 'is_default', '', 'varchar(8)', 'string', 'IsDefault', 'isDefault', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (186, 16, 'list_class', '', 'varchar(128)', 'string', 'ListClass', 'listClass', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (187, 16, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (188, 16, 'status', '', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (189, 16, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (190, 16, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (191, 17, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (192, 17, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (193, 17, 'dict_name', '', 'varchar(128)', 'string', 'DictName', 'dictName', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (194, 17, 'dict_type', '', 'varchar(128)', 'string', 'DictType', 'dictType', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (195, 17, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (196, 17, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (197, 17, 'status', '', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (198, 17, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (199, 17, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (200, 18, 'business_name', '业务�?, 'varchar(255)', 'string', 'BusinessName', 'businessName', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (201, 18, 'class_name', '类名', 'varchar(255)', 'string', 'ClassName', 'className', '2', '2', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (202, 18, 'create_by', '创建�?, 'bigint', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (203, 18, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (204, 18, 'function_author', '作�?, 'varchar(255)', 'string', 'FunctionAuthor', 'functionAuthor', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (205, 18, 'function_name', '功能描述', 'varchar(255)', 'string', 'FunctionName', 'functionName', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (206, 18, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (207, 18, 'is_plugin', '是否插件 1-�?2-�?, 'char(1)', 'string', 'IsPlugin', 'isPlugin', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (208, 18, 'module_name', '接口�?, 'varchar(255)', 'string', 'ModuleName', 'moduleName', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (209, 18, 'package_name', '应用�?, 'varchar(255)', 'string', 'PackageName', 'packageName', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (210, 18, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (211, 18, 'table_comment', '表描�?, 'varchar(255)', 'string', 'TableComment', 'tableComment', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (212, 18, 'table_name', '表名', 'varchar(255)', 'string', 'TableName', 'tableName', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (213, 18, 'update_by', '更新�?, 'bigint', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (214, 18, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (215, 19, 'agent', '代理', 'varchar(255)', 'string', 'Agent', 'agent', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (216, 19, 'browser', '浏览�?, 'varchar(255)', 'string', 'Browser', 'browser', '2', '2', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (217, 19, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (218, 19, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (219, 19, 'id', '主键编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (220, 19, 'ipaddr', 'ip地址', 'varchar(255)', 'string', 'Ipaddr', 'ipaddr', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (221, 19, 'login_location', '归属�?, 'varchar(255)', 'string', 'LoginLocation', 'loginLocation', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (222, 19, 'login_time', '登录时间', 'timestamp', '*time.Time', 'LoginTime', 'loginTime', '2', '1', '2', '2', 'EQ', 'datetime', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (223, 19, 'os', '系统', 'varchar(255)', 'string', 'Os', 'os', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (224, 19, 'platform', '固件', 'varchar(255)', 'string', 'Platform', 'platform', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (225, 19, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (226, 19, 'status', '状�?1-登录 2-退�?, 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (227, 19, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (228, 19, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (229, 19, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (230, 20, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (231, 20, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (232, 20, 'element', '', 'varchar(255)', 'string', 'Element', 'element', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (233, 20, 'icon', '', 'varchar(128)', 'string', 'Icon', 'icon', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (234, 20, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (235, 20, 'is_affix', '是否固定 1-�?2-�?, 'char(1)', 'string', 'IsAffix', 'isAffix', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (236, 20, 'is_frame', '是否内嵌 1-�?2-�?, 'char(1)', 'string', 'IsFrame', 'isFrame', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (237, 20, 'is_hidden', '是否隐藏 1-�?2-�?, 'char(1)', 'string', 'IsHidden', 'isHidden', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (238, 20, 'is_keep_alive', '是否缓存 1-�?2-�?, 'char(1)', 'string', 'IsKeepAlive', 'isKeepAlive', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (239, 20, 'menu_type', '', 'char(1)', 'string', 'MenuType', 'menuType', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (240, 20, 'parent_id', '上级菜单id', 'int', 'int64', 'ParentId', 'parentId', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (241, 20, 'parent_ids', '上级菜单id集合', 'varchar(255)', 'string', 'ParentIds', 'parentIds', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (242, 20, 'path', '', 'varchar(255)', 'string', 'Path', 'path', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (243, 20, 'permission', '', 'varchar(255)', 'string', 'Permission', 'permission', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (244, 20, 'redirect', '', 'varchar(255)', 'string', 'Redirect', 'redirect', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (245, 20, 'sort', '', 'int', 'int64', 'Sort', 'sort', '2', '2', '2', '2', 'EQ', 'input', '', 16, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (246, 20, 'title', '', 'varchar(128)', 'string', 'Title', 'title', '2', '2', '2', '2', 'EQ', 'input', '', 17, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (247, 20, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 18, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (248, 20, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 19, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (249, 21, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (250, 21, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (251, 21, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (252, 21, 'post_code', '', 'varchar(128)', 'string', 'PostCode', 'postCode', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (253, 21, 'post_name', '', 'varchar(128)', 'string', 'PostName', 'postName', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (254, 21, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (255, 21, 'sort', '', 'tinyint', 'int64', 'Sort', 'sort', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (256, 21, 'status', '', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (257, 21, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (258, 21, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (259, 22, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (260, 22, 'created_at', '创建时间', 'datetime(3)', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (261, 22, 'data_scope', '', 'varchar(128)', 'string', 'DataScope', 'dataScope', '2', '2', '2', '2', 'EQ', 'input', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (262, 22, 'id', '', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (263, 22, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (264, 22, 'role_key', '', 'varchar(128)', 'string', 'RoleKey', 'roleKey', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (265, 22, 'role_name', '', 'varchar(128)', 'string', 'RoleName', 'roleName', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (266, 22, 'role_sort', '', 'bigint', 'int64', 'RoleSort', 'roleSort', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (267, 22, 'status', '状�?1-正常 2-停用', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (268, 22, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (269, 22, 'updated_at', '最后更新时�?, 'datetime(3)', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (270, 23, 'avatar', '头像', 'varchar(255)', 'string', 'Avatar', 'avatar', '2', '2', '2', '2', 'EQ', 'input', '', 1, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (271, 23, 'create_by', '创建�?, 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', 'EQ', 'input', '', 2, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (272, 23, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 3, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (273, 23, 'dept_id', '部门', 'int', 'int64', 'DeptId', 'deptId', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (274, 23, 'email', '邮箱', 'varchar(128)', 'string', 'Email', 'email', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (275, 23, 'id', '编码', 'int', 'int64', 'Id', 'id', '1', '1', '2', '2', 'EQ', 'input', '', 6, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (276, 23, 'nick_name', '昵称', 'varchar(128)', 'string', 'NickName', 'nickName', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (277, 23, 'password', '密码', 'varchar(128)', 'string', 'Password', 'password', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (278, 23, 'phone', '手机�?, 'varchar(11)', 'string', 'Phone', 'phone', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (279, 23, 'post_id', '岗位', 'int', 'int64', 'PostId', 'postId', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (280, 23, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (281, 23, 'role_id', '角色ID', 'int', 'int64', 'RoleId', 'roleId', '2', '2', '2', '2', 'EQ', 'input', '', 12, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (283, 23, 'sex', '性别', 'char(1)', 'string', 'Sex', 'sex', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (284, 23, 'status', '状�?, 'varchar(4)', 'string', 'Status', 'status', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (285, 23, 'update_by', '更新�?, 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', 'EQ', 'input', '', 16, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (286, 23, 'updated_at', '最后更新时�?, 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', 'EQ', 'datetime', '', 17, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (287, 23, 'username', '用户�?, 'varchar(64)', 'string', 'Username', 'username', '2', '2', '2', '2', 'EQ', 'input', '', 18, '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_gen_table`;
CREATE TABLE `admin_sys_gen_table` (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '表名',
  `table_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '表描�?,
  `class_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '类名',
  `package_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '应用�?,
  `module_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '接口�?,
  `function_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '功能描述',
  `function_author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '作�?,
  `business_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '业务�?,
  `is_plugin` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '1' COMMENT '是否插件 1-�?2-�?,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  `create_by` bigint DEFAULT NULL COMMENT '创建�?,
  `update_by` bigint DEFAULT NULL COMMENT '更新�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='表管�?;

-- ----------------------------
-- Records of admin_sys_gen_table
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 'admin_sys_oper_log', '操作日志', 'SysOperLog', 'admin', 'sys-oper-log', '操作日志', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 'app_user', '用户管理', 'User', 'app', 'user', '用户管理', 'Jason', 'user', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 'app_user_account_log', '账变记录', 'UserAccountLog', 'app', 'user-account-log', '账变记录', 'Jason', 'user', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 'app_user_conf', '用户配置', 'UserConf', 'app', 'user-conf', '用户配置', 'Jason', 'user', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (5, 'app_user_country_code', '国家区号', 'UserCountryCode', 'app', 'user-country-code', '国家区号', 'Jason', 'user', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (6, 'app_user_level', '用户等级', 'UserLevel', 'app', 'user-level', '用户等级', 'Jason', 'user', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (7, 'app_user_oper_log', '用户关键行为日志�?, 'UserOperLog', 'app', 'user-oper-log', '用户关键行为日志�?, 'Jason', 'user', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (8, 'plugins_content_announcement', '公告管理', 'ContentAnnouncement', 'plugins', 'content-announcement', '公告管理', 'Jason', 'content', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (9, 'plugins_content_article', '文章管理', 'ContentArticle', 'plugins', 'content-article', '文章管理', 'Jason', 'content', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (10, 'plugins_content_category', '文章分类管理', 'ContentCategory', 'plugins', 'content-category', '文章分类管理', 'Jason', 'content', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (11, 'plugins_filemgr_app', 'app升级管理', 'FilemgrApp', 'plugins', 'filemgr-app', 'app升级管理', 'Jason', 'filemgr', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (12, 'plugins_msg_code', '验证码记�?, 'MsgCode', 'plugins', 'msg-code', '验证码记�?, 'Jason', 'msg', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (13, 'admin_sys_api', '接口管理', 'SysApi', 'admin', 'sys-api', '接口管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (14, 'admin_sys_config', '配置管理', 'SysConfig', 'admin', 'sys-config', '配置管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (15, 'admin_sys_dept', '部门管理', 'SysDept', 'admin', 'sys-dept', '部门管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (16, 'admin_sys_dict_data', '字典数据管理', 'SysDictData', 'admin', 'sys-dict-data', '字典数据管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (17, 'admin_sys_dict_type', '字典类型管理', 'SysDictType', 'admin', 'sys-dict-type', '字典类型管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (18, 'admin_sys_gen_table', '表管�?, 'SysGenTable', 'admin', 'sys-gen-table', '表管�?, 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (19, 'admin_sys_login_log', '登录日志', 'SysLoginLog', 'admin', 'sys-login-log', '登录日志', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (20, 'admin_sys_menu', '菜单管理', 'SysMenu', 'admin', 'sys-menu', '菜单管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (21, 'admin_sys_post', '岗位管理', 'SysPost', 'admin', 'sys-post', '岗位管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (22, 'admin_sys_role', '角色管理', 'SysRole', 'admin', 'sys-role', '角色管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
INSERT INTO `admin_sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (23, 'admin_sys_user', '系统用户管理', 'SysUser', 'admin', 'sys-user', '系统用户管理', 'Jason', 'sys', '1', '', '2024-12-13 19:56:40', '2024-12-13 19:56:40', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_login_log`;
CREATE TABLE `admin_sys_login_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `user_id` int DEFAULT NULL COMMENT '用户编号',
  `ipaddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '归属�?,
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '浏览�?,
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '系统',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '代理',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登录时间',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状�?1-登录 2-退�?,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='登录日志';

-- 登录日志按用户查�?统计主路�?ALTER TABLE `admin_sys_login_log` ADD INDEX `idx_admin_sys_login_log_user_id_created_at` (`user_id`, `created_at`);

-- ----------------------------
-- Records of admin_sys_login_log
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_login_log` (`id`, `user_id`, `ipaddr`, `login_location`, `browser`, `os`, `agent`, `platform`, `login_time`, `status`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 1, '::1', '----', 'Edge 131.0.0.0', 'Intel Mac OS X 10_15_7', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 'Macintosh', '2024-12-13 19:58:24', '1', '登录操作', '2024-12-13 19:58:24', '2024-12-13 19:58:24', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_menu`;
CREATE TABLE `admin_sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `element` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL COMMENT '上级菜单id',
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '上级菜单id集合',
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_keep_alive` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否缓存 1-�?2-�?,
  `is_affix` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否固定 1-�?2-�?,
  `is_hidden` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否隐藏 1-�?2-�?,
  `is_frame` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否内嵌 1-�?2-�?,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=135 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='菜单管理';

-- ----------------------------
-- Records of admin_sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统管理', 'SettingFilled', '/admin/sys', '', '/admin/sys/sys-api', '', 300, 0, '0,', '1', '', '', '2', '', 1, 1, '2021-05-20 21:58:46', '2024-12-14 15:04:42');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '用户管理', 'HeartFilled', '/admin/sys/sys-user', '/admin/sys/sys-user/index', '', '', 10, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:10:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '新增系统用户管理', 'AppstoreOutlined', '', '', '', 'admin:sys-user:add', 1, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 13:43:52');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '获取系统用户管理分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-user:query', 0, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 13:14:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '更新系统用户管理', 'AppstoreOutlined', '', '', '', 'admin:sys-user:edit', 2, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 13:50:46');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '删除系统用户管理', 'AppstoreOutlined', '', '', '', 'admin:sys-user:del', 3, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 13:50:52');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '菜单管理', 'BookFilled', '/admin/sys/sys-menu', '/admin/sys/sys-menu/index', '', '', 30, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:10:42');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '角色管理', 'FlagFilled', '/admin/sys/sys-role', '/admin/sys/sys-role/index', '', '', 20, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:10:26');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '部门管理', 'BuildFilled', '/admin/sys/sys-dept', '/admin/sys/sys-dept/index', '', '', 40, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:10:54');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '岗位管理', 'DribbbleCircleFilled', '/admin/sys/sys-post', '/admin/sys/sys-post/index', '', '', 50, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:11:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '字典管理', 'ContainerFilled', '/admin/sys/sys-dicttype', '/admin/sys/sys-dicttype/index', '', '', 60, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:11:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '字典数据', 'ControlFilled', '/admin/sys/sys-dictdata', '/admin/sys/sys-dictdata/index', '', '', 100, 1, '0,1,', '2', '2', '2', '1', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:12:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '参数管理', 'GoldenFilled', '/admin/sys/sys-config', '/admin/sys/sys-config/index', '', '', 70, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:11:54');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, '登录日志', 'FileTextFilled', '/admin/sys/sys-loginlog', '/admin/sys/sys-loginlog/index', '', '', 90, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:12:07');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, '操作日志', 'BugFilled', '/admin/sys/sys-operalog', '/admin/sys/sys-operlog/index', '', '', 120, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:12:31');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, '新增菜单管理', 'AppstoreOutlined', '', '', '', 'admin:sys-menu:add', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:13:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '更新菜单管理', 'AppstoreOutlined', '', '', '', 'admin:sys-menu:edit', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:12:33');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '查询菜单管理列表', 'AppstoreOutlined', '', '', '', 'admin:sys-menu:query', 0, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:02:45');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '删除菜单管理', 'AppstoreOutlined', '', '', '', 'admin:sys-menu:del', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:12:48');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '新增角色管理', 'AppstoreOutlined', '', '', '', 'admin:sys-role:add', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 13:56:24');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '获取角色管理全部列表', 'AppstoreOutlined', '', '', '', 'admin:sys-role:query', 0, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 13:52:04');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '更新角色管理', 'AppstoreOutlined', '', '', '', 'admin:sys-role:edit', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:00:05');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '删除角色管理', 'AppstoreOutlined', '', '', '', 'admin:sys-role:del', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:00:31');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '获取部门管理列表', 'AppstoreOutlined', '', '', '', 'admin:sys-dept:query', 0, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 14:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '新增部门管理', 'AppstoreOutlined', '', '', '', 'admin:sys-dept:add', 1, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 14:17:17');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, '更新部门管理', 'AppstoreOutlined', '', '', '', 'admin:sys-dept:edit', 2, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 14:18:07');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, '删除部门管理', 'AppstoreOutlined', '', '', '', 'admin:sys-dept:del', 3, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 14:18:35');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, '获取岗位管理分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-post:query', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:24:08');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, '新增岗位管理', 'AppstoreOutlined', '', '', '', 'admin:sys-post:add', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:27:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, '修改岗位管理', 'AppstoreOutlined', '', '', '', 'admin:sys-post:edit', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:28:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, '删除岗位', 'AppstoreOutlined', '', '', '', 'admin:sys-post:del', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-13 21:32:47');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, '获取字典类型分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-type:query', 0, 11, '0,1,11,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:30:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, '新增字典类型', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-type:add', 0, 11, '0,1,11,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:31:05');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (34, '更新字典类型', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-type:edit', 0, 11, '0,1,11,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:32:25');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (35, '删除字典类型', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-type:del', 0, 11, '0,1,11,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:32:36');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (36, '获取字典数据分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-data:query', 0, 12, '0,1,12,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:36:52');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (37, '新增字典数据', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-data:add', 0, 12, '0,1,12,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:38:04');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (38, '更新字典数据', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-data:edit', 0, 12, '0,1,12,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:38:35');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (39, '删除字典数据', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-data:del', 0, 12, '0,1,12,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:38:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (40, '获取配置管理分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-config:query', 0, 13, '0,1,13,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:33:31');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (41, '新增配置管理', 'AppstoreOutlined', '', '', '', 'admin:sys-config:add', 0, 13, '0,1,13,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:33:44');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (42, '更新配置管理', 'AppstoreOutlined', '', '', '', 'admin:sys-config:edit', 0, 13, '0,1,13,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:34:00');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (43, '删除配置管理', 'AppstoreOutlined', '', '', '', 'admin:sys-config:del', 0, 13, '0,1,13,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:34:11');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (44, '获取登录日志分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-login-log:query', 0, 14, '0,1,14,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:35:28');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (45, '删除登录日志', 'AppstoreOutlined', '', '', '', 'admin:sys-login-log:del', 0, 14, '0,1,14,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-13 21:38:09');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (46, '获取操作日志分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-oper-log:query', 0, 15, '0,1,15,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:39:13');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (47, '删除操作日志', 'AppstoreOutlined', '', '', '', 'admin:sys-oper-log:del', 0, 15, '0,1,15,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-14 14:40:07');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (48, '代码生成', 'BugFilled', '/admin/sys/sys-tools/sys-gen', '/admin/sys/sys-tools/sys-gen/index', '', '', 20, 54, '0,1,54,', '2', '2', '2', '2', '1', 1, 1, '2020-04-11 15:52:48', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (49, '代码生成修改', 'ClockCircleFilled', '/admin/sys/sys-tools/sys-gen/edit', '/admin/sys/sys-tools/sys-gen/edit/index', '', '', 100, 54, '0,1,54,', '2', '2', '2', '1', '1', 1, 1, '2020-04-11 15:52:48', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (50, '服务监控', 'ExclamationCircleFilled', '/admin/sys/sys-tools/sys-monitor', '/admin/sys/sys-tools/sys-monitor/index', '', '', 0, 54, '0,1,54,', '2', '2', '2', '2', '1', 1, 1, '2020-04-14 00:28:19', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (51, '接口管理', 'ThunderboltFilled', '/admin/sys/sys-api', '/admin/sys/sys-api/index', '', '', 0, 1, '0,1,', '2', '2', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2024-12-14 15:09:44');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (52, '获取接口管理分页列表', 'AppstoreOutlined', '', '', '', 'admin:sys-api:query', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 13:06:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (53, '更新接口管理', 'AppstoreOutlined', '', '', '', 'admin:sys-api:edit', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-14 13:07:22');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (54, '系统工具', 'ToolFilled', '/admin/sys/sys-tools', '', '/admin/sys/sys-tools/sys-monitor', '', 330, 1, '0,1,', '1', '', '', '2', '', 1, 1, '2021-05-21 11:13:32', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (55, '文件管理', 'FolderFilled', '/plugins/filemgr', '', '/plugins/filemgr/filemgr-app', '', 90, 57, '0,57,', '1', '', '', '2', '', 1, 1, '2021-08-13 14:19:11', '2024-12-14 15:08:17');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (56, '内容管理', 'FileWordFilled', '/plugins/content', '', '/plugins/content/content-category', '', 60, 57, '0,57,', '1', '', '', '2', '', 1, 1, '2021-08-16 18:01:20', '2024-12-14 15:07:28');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (57, '插件管理', 'ApiFilled', '/plugins', '', '/plugins/content/content-category', '', 270, 0, '0,', '1', '', '', '2', '', 1, 1, '2023-03-07 10:37:37', '2024-12-14 15:03:16');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (58, 'App应用', 'FolderOpenFilled', '/app', '', '/app/user/user', '', 30, 0, '0,', '1', '', '', '2', '', 1, 1, '2023-03-08 09:27:36', '2024-12-14 15:01:02');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (59, '用户列表', 'MehFilled', '/app/user', '', '/app/user/user', '', 30, 58, '0,58,', '1', '', '', '2', '', 1, 1, '2023-03-09 14:24:25', '2024-12-14 15:06:26');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (61, '用户等级', 'LayoutFilled', '/app/user/user-level', '/app/user/user-level/index', NULL, '', 60, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 21:33:49', '2023-03-09 23:05:34');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (62, '获取用户等级管理分页列表', 'AppstoreOutlined', '', '', '', 'app:user-level:query', 0, 61, '0,58,59,61,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:49', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (63, '新增用户等级管理', 'AppstoreOutlined', '', '', '', 'app:user-level:add', 0, 61, '0,58,59,61,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:49', '2024-12-14 12:49:18');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (64, '更新用户等级管理', 'AppstoreOutlined', '', '', '', 'app:user-level:edit', 0, 61, '0,58,59,61,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:49', '2024-12-14 12:49:28');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (65, '删除用户等级管理', 'AppstoreOutlined', '', '', '', 'app:user-level:del', 0, 61, '0,58,59,61,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:49', '2024-12-14 12:49:38');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (66, '导出用户等级管理', 'AppstoreOutlined', '', '', '', 'app:user-level:export', 0, 61, '0,58,59,61,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:49', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (67, '账变记录', 'LayoutFilled', '/app/user/user-account-log', '/app/user/user-account-log/index', NULL, '', 150, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 21:33:51', '2024-12-23 22:05:53');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (68, '获取账变记录分页列表', 'AppstoreOutlined', '', '', '', 'app:user-account-log:query', 0, 67, '0,58,59,67,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:51', '2024-12-23 22:02:18');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (69, '导出账变记录', 'AppstoreOutlined', '', '', '', 'app:user-account-log:export', 0, 67, '0,58,59,67,', '3', '', '', '', '', 1, 1, '2023-03-09 21:33:51', '2024-12-23 22:02:18');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (70, '用户配置', 'LayoutFilled', '/app/user/user-conf', '/app/user/user-conf/index', '', '', 90, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 23:04:40', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (71, '获取用户配置管理分页列表', 'AppstoreOutlined', '', '', '', 'app:user-conf:query', 0, 70, '0,58,59,70,', '3', '', '', '', '', 1, 1, '2023-03-09 23:04:40', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (72, '更新用户配置管理', 'AppstoreOutlined', '', '', '', 'app:user-conf:edit', 0, 70, '0,58,59,70,', '3', '', '', '', '', 1, 1, '2023-03-09 23:04:40', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (73, '用户管理', 'LayoutFilled', '/app/user/user', '/app/user/user/index', '', '', 30, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 23:18:49', '2024-12-13 20:14:34');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (74, '获取用户管理分页列表', 'AppstoreOutlined', '', '', '', 'app:user:query', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (75, '新增用户管理', 'AppstoreOutlined', '', '', '', 'app:user:add', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-14 12:47:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (76, '更新用户管理', 'AppstoreOutlined', '', '', '', 'app:user:edit', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-14 12:48:01');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (77, '导出用户管理', 'AppstoreOutlined', '', '', '', 'app:user:export', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-13 20:56:34');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (78, '用户行为记录', 'LayoutFilled', '/app/user/user-oper-log', '/app/user/user-oper-log/index', NULL, '', 120, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:02:42');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (79, '获取用户操作日志分页列表', 'AppstoreOutlined', '', '', '', 'app:user-oper-log:query', 0, 78, '0,58,59,78,', '3', '', '', '', '', 1, 1, '2023-03-11 15:00:06', '2024-12-14 12:50:33');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (80, '导出用户操作日志', 'AppstoreOutlined', '', '', '', 'app:user-oper-log:export', 0, 78, '0,58,59,78,', '3', '', '', '', '', 1, 1, '2023-03-11 15:00:06', '2024-12-14 12:50:48');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (81, '消息管理', 'BellFilled', '/plugins/msg', '', '/plugins/msg/msg-code', '', 120, 57, '0,57,', '1', '', '', '2', '', 1, 1, '2023-03-12 13:27:59', '2024-12-14 15:13:49');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (82, '验证码记�?, 'LayoutFilled', '/plugins/msg/msg-code', '/plugins/msg/msg-code/index', NULL, '', 0, 81, '0,57,81,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:32');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (83, '获取验证码管理分页列�?, 'AppstoreOutlined', '', '', '', 'plugins:msg-code:query', 0, 82, '0,57,81,82,', '3', '', '', '', '', 1, 1, '2023-03-12 21:54:02', '2024-12-14 13:05:37');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (84, '公告管理', 'LayoutFilled', '/plugins/content/content-announcement', '/plugins/content/content-announcement/index', NULL, '', 90, 56, '0,57,56,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:48:08');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (85, '获取公告管理分页列表', 'AppstoreOutlined', '', '', '', 'plugins:content-announcement:query', 0, 84, '0,57,56,84,', '3', '', '', '', '', 1, 1, '2023-03-12 22:47:11', '2024-12-14 13:03:46');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (86, '新增公告管理', 'AppstoreOutlined', '', '', '', 'plugins:content-announcement:add', 0, 84, '0,57,56,84,', '3', '', '', '', '', 1, 1, '2023-03-12 22:47:11', '2024-12-14 13:03:58');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (87, '更新公告管理', 'AppstoreOutlined', '', '', '', 'plugins:content-announcement:edit', 0, 84, '0,57,56,84,', '3', '', '', '', '', 1, 1, '2023-03-12 22:47:11', '2024-12-14 13:04:08');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (88, '删除公告管理', 'AppstoreOutlined', '', '', '', 'plugins:content-announcement:del', 0, 84, '0,57,56,84,', '3', '', '', '', '', 1, 1, '2023-03-12 22:47:11', '2024-12-13 21:06:13');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (89, '导出公告管理', 'AppstoreOutlined', '', '', '', 'plugins:content-announcement:export', 0, 84, '0,57,56,84,', '3', '', '', '', '', 1, 1, '2023-03-12 22:47:11', '2024-12-13 21:06:33');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (90, '内容分类', 'LayoutFilled', '/plugins/content/content-category', '/plugins/content/content-category/index', NULL, '', 0, 56, '0,57,56,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:20:35');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (91, '获取内容分类管理分页列表', 'AppstoreOutlined', '', '', '', 'plugins:content-category:query', 0, 90, '0,57,56,90,', '3', '', '', '', '', 1, 1, '2023-03-12 23:17:44', '2024-12-14 13:01:47');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (92, '新增内容分类管理详情', 'AppstoreOutlined', '', '', '', 'plugins:content-category:add', 0, 90, '0,57,56,90,', '3', '', '', '', '', 1, 1, '2023-03-12 23:17:44', '2024-12-14 13:02:02');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (93, '更新内容分类管理', 'AppstoreOutlined', '', '', '', 'plugins:content-category:edit', 0, 90, '0,57,56,90,', '3', '', '', '', '', 1, 1, '2023-03-12 23:17:44', '2024-12-14 13:02:14');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (94, '删除内容分类管理', 'AppstoreOutlined', '', '', '', 'plugins:content-category:del', 0, 90, '0,57,56,90,', '3', '', '', '', '', 1, 1, '2023-03-12 23:17:45', '2024-12-14 13:02:24');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (95, '导出内容分类管理', 'AppstoreOutlined', '', '', '', 'plugins:content-category:export', 0, 90, '0,57,56,90,', '3', '', '', '', '', 1, 1, '2023-03-12 23:17:45', '2024-12-14 13:02:34');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (96, '文章管理', 'LayoutFilled', '/plugins/content/content-article', '/plugins/content/content-article/index', NULL, '', 60, 56, '0,57,56,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:53:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (97, '获取文章管理分页列表', 'AppstoreOutlined', '', '', '', 'plugins:content-article:query', 0, 96, '0,57,56,96,', '3', '', '', '', '', 1, 1, '2023-03-12 23:52:45', '2024-12-14 13:02:56');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (98, '新增文章管理', 'AppstoreOutlined', '', '', '', 'plugins:content-article:add', 0, 96, '0,57,56,96,', '3', '', '', '', '', 1, 1, '2023-03-12 23:52:45', '2024-12-14 13:03:09');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (99, '更新文章管理', 'AppstoreOutlined', '', '', '', 'plugins:content-article:edit', 0, 96, '0,57,56,96,', '3', '', '', '', '', 1, 1, '2023-03-12 23:52:45', '2024-12-14 13:03:18');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (100, '删除文章管理', 'AppstoreOutlined', '', '', '', 'plugins:content-article:del', 0, 96, '0,57,56,96,', '3', '', '', '', '', 1, 1, '2023-03-12 23:52:45', '2024-12-14 13:03:28');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (101, '导出文章管理', 'AppstoreOutlined', '', '', '', 'plugins:content-article:export', 0, 96, '0,57,56,96,', '3', '', '', '', '', 1, 1, '2023-03-12 23:52:46', '2024-12-13 21:04:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (102, 'App管理', 'LayoutFilled', '/plugins/filemgr/filemgr-app', '/plugins/filemgr/filemgr-app/index', NULL, '', 0, 55, '0,57,55,', '2', '1', '2', '2', '1', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:52');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (103, '获取APP管理分页列表', 'AppstoreOutlined', '', '', '', 'plugins:filemgr-app:query', 0, 102, '0,57,55,102,', '3', '', '', '', '', 1, 1, '2023-03-13 00:55:02', '2024-12-14 13:04:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (104, '新增APP管理', 'AppstoreOutlined', '', '', '', 'plugins:filemgr-app:add', 0, 102, '0,57,55,102,', '3', '', '', '', '', 1, 1, '2023-03-13 00:55:02', '2024-12-15 22:14:53');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (105, '更新APP管理', 'AppstoreOutlined', '', '', '', 'plugins:filemgr-app:edit', 0, 102, '0,57,55,102,', '3', '', '', '', '', 1, 1, '2023-03-13 00:55:02', '2024-12-15 22:15:04');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (106, '删除App管理', 'AppstoreOutlined', '', '', '', 'plugins:filemgr-app:del', 0, 102, '0,57,55,102,', '3', '', '', '', '', 1, 1, '2023-03-13 00:55:02', '2024-12-13 21:08:39');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (107, '导出App管理', 'AppstoreOutlined', '', '', '', 'plugins:filemgr-app:export', 0, 102, '0,57,55,102,', '3', '', '', '', '', 1, 1, '2023-03-13 00:55:02', '2024-12-13 21:08:56');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (108, '国家区号', 'LayoutFilled', '/app/user/user-country-code', '/app/user/user-country-code/index', NULL, '', 180, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-14 17:47:44', '2024-12-23 22:02:47');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (109, '获取国家区号管理分页列表', 'AppstoreOutlined', '', '', '', 'app:user-country-code:query', 0, 108, '0,58,59,108,', '3', '', '', '', '', 1, 1, '2023-03-14 17:47:44', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (110, '新增国家区号管理', 'AppstoreOutlined', '', '', '', 'app:user-country-code:add', 0, 108, '0,58,59,108,', '3', '', '', '', '', 1, 1, '2023-03-14 17:47:44', '2024-12-14 12:51:13');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (111, '更新国家区号管理', 'AppstoreOutlined', '', '', '', 'app:user-country-code:edit', 0, 108, '0,58,59,108,', '3', '', '', '', '', 1, 1, '2023-03-14 17:47:44', '2024-12-14 12:51:23');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (112, '删除国家区号管理', 'AppstoreOutlined', '', '', '', 'app:user-country-code:del', 0, 108, '0,58,59,108,', '3', '', '', '', '', 1, 1, '2023-03-14 17:47:45', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (113, '导出国家区号管理', 'AppstoreOutlined', '', '', '', 'app:user-country-code:export', 0, 108, '0,58,59,108,', '3', '', '', '', '', 1, 1, '2023-03-14 17:47:45', '2024-12-14 12:51:44');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (114, '导出操作日志', 'AppstoreOutlined', '', '', '', 'admin:sys-oper-log:export', 0, 15, '0,1,15,', '3', '', '', '', '', 1, 1, '2023-05-09 11:02:50', '2024-12-13 21:41:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (115, '导出登录日志', 'AppstoreOutlined', '', '', '', 'admin:sys-login-log:export', 0, 14, '0,1,14,', '3', '', '', '', '', 1, 1, '2023-05-09 11:04:20', '2024-12-14 14:36:00');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (117, '导出字典类型', 'AppstoreOutlined', '', '', '', 'admin:sys-dict-type:export', 0, 11, '0,1,11,', '3', '', '', '', '', 1, 1, '2023-05-09 11:16:13', '2024-12-14 14:32:48');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (118, '导出配置管理', 'AppstoreOutlined', '', '', '', 'content:sys-config:export', 0, 13, '0,1,13,', '3', '', '', '', '', 1, 1, '2023-05-09 11:34:20', '2024-12-14 14:34:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (119, '首页', 'HomeFilled', '/home', '/admin/sys/home/index', '', '', 0, 0, '0,', '2', '2', '1', '2', '1', 1, 1, '2024-11-22 11:34:20', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (120, '个人中心', 'ProfileFilled', '/profile', '/admin/sys/profile/index', '', '', 0, 0, '0,', '2', '2', '2', '1', '1', 1, 1, '2024-11-23 08:20:24', '2024-12-14 15:16:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (121, '导出接口管理', 'AppstoreOutlined', '', '', '', 'admin:sys-api:export', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2024-12-09 14:59:51', '2024-12-14 14:22:31');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (122, '重置系统用户密码', 'AppstoreOutlined', '', '', '', 'admin:sys-user:edit-pwd', 0, 2, '0,22,', '3', '', '', '', '', 1, 1, '2024-12-09 15:15:05', '2024-12-14 13:46:43');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (123, '更新角色管理数据权限', 'AppstoreOutlined', '', '', '', 'admin:sys-role:datascope', 0, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2024-12-09 15:19:53', '2024-12-14 14:05:10');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (124, '删除接口管理', 'AppstoreOutlined', '', '', '', 'admin:sys-api:del', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2024-12-09 15:37:48', '2024-12-14 13:07:36');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (125, '同步接口管理', 'AppstoreOutlined', '', '', '', 'admin:sys-api:sync', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2024-12-13 21:12:34', '2024-12-14 13:10:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (126, '更新表管�?, 'AppstoreOutlined', '', '', '', 'admin:sys-gen:edit', 2, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-13 22:04:00', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (127, '预览', 'AppstoreOutlined', '', '', '', 'admin:sys-gen:preview', 3, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 03:36:14', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (128, '获取表管理分页列�?, 'AppstoreOutlined', '', '', '', 'admin:sys-gen:query', 0, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 03:39:04', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (129, '生成代码', 'AppstoreOutlined', '', '', '', 'admin:sys-gen:gen-code', 4, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 03:40:26', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (130, '下载代码', 'AppstoreOutlined', '', '', '', 'admin:sys-gen:download-code', 5, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 03:41:22', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (131, '配置导入', 'AppstoreOutlined', '', '', '', 'admin:sys-gen:import-config', 6, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 03:43:36', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (132, '表管理删�?, 'AppstoreOutlined', '', '', '', 'admin:sys-gen:del', 7, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 03:44:46', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (133, '新增表管�?, 'AppstoreOutlined', '', '', '', 'admin:sys-gen:import-table', 1, 48, '0,1,54,48,', '3', '', '', '', '', 1, 1, '2024-12-14 14:50:46', '2024-12-23 22:07:21');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (134, '导出岗位管理', 'AppstoreOutlined', '', '', '', 'admin:sys-post:export', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2024-12-15 22:23:26', '2024-12-15 22:23:26');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_menu_api_rule`;
CREATE TABLE `admin_sys_menu_api_rule` (
  `admin_sys_menu_menu_id` int NOT NULL,
  `admin_sys_api_id` int NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`admin_sys_menu_menu_id`,`admin_sys_api_id`),
  KEY `fk_admin_sys_menu_api_rule_admin_sys_api` (`admin_sys_api_id`),
  CONSTRAINT `fk_admin_sys_menu_api_rule_admin_sys_api` FOREIGN KEY (`admin_sys_api_id`) REFERENCES `admin_sys_api` (`id`),
  CONSTRAINT `fk_admin_sys_menu_api_rule_admin_sys_menu` FOREIGN KEY (`admin_sys_menu_menu_id`) REFERENCES `admin_sys_menu` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='菜单和角色关�?;

-- ----------------------------
-- Records of admin_sys_menu_api_rule
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (32, 1);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (126, 2);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (117, 3);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (34, 4);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (36, 5);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (38, 7);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (3, 9);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (4, 9);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (5, 9);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (24, 9);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (25, 9);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (26, 9);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (123, 10);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (26, 11);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (16, 12);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (17, 12);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (18, 12);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (20, 14);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (22, 14);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (17, 15);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (50, 16);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (128, 20);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (130, 21);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (131, 22);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (129, 23);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (133, 24);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (127, 25);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (126, 26);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (52, 27);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (125, 28);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (16, 30);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (17, 30);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (121, 31);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (53, 32);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (40, 33);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (118, 34);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (42, 37);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (28, 38);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (3, 39);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (5, 39);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (134, 40);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (30, 41);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (4, 42);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (5, 45);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (44, 46);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (115, 47);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (46, 49);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (114, 50);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (21, 52);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (3, 53);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (5, 53);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (22, 54);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (123, 54);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (74, 55);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (109, 56);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (113, 57);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (111, 58);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (71, 59);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (72, 60);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (68, 61);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (69, 62);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (62, 64);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (66, 65);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (64, 66);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (79, 67);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (80, 68);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (77, 70);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (76, 71);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (85, 72);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (89, 73);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (87, 74);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (97, 75);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (101, 76);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (99, 77);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (91, 78);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (95, 79);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (93, 80);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (103, 81);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (107, 82);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (105, 83);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (83, 84);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (37, 89);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (33, 90);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (25, 91);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (3, 92);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (41, 94);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (133, 95);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (16, 96);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (29, 97);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (20, 98);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (75, 99);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (110, 100);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (63, 101);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (86, 102);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (98, 103);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (92, 104);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (104, 105);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (104, 106);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (105, 106);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (122, 110);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (5, 111);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (5, 112);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (38, 113);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (34, 114);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (26, 115);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (22, 116);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (123, 117);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (22, 118);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (53, 119);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (42, 120);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (126, 121);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (17, 122);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (30, 123);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (72, 124);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (111, 125);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (64, 126);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (76, 127);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (87, 128);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (99, 129);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (93, 130);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (105, 131);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (39, 132);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (35, 133);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (27, 134);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (124, 135);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (43, 136);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (132, 137);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (45, 138);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (19, 139);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (47, 140);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (31, 141);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (23, 142);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (6, 143);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (112, 144);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (65, 145);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (88, 146);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (100, 147);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (94, 148);
INSERT INTO `admin_sys_menu_api_rule` (`admin_sys_menu_menu_id`, `admin_sys_api_id`) VALUES (106, 149);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_oper_log`;
CREATE TABLE `admin_sys_oper_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `request_method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求方式',
  `user_id` int DEFAULT NULL COMMENT '操作�?,
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '访问位置',
  `status` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '操作状�?,
  `oper_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ua',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='操作日志';

-- 操作日志按用户查�?统计主路�?ALTER TABLE `admin_sys_oper_log` ADD INDEX `idx_admin_sys_oper_log_user_id_created_at` (`user_id`, `created_at`);

-- ----------------------------
-- Records of admin_sys_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 'GET', 0, '/admin-api/v1/admin/sys/sys-dict/data/select?dictType=sys_loginlog_status', '::1', '----', '0', '2024-12-13 19:58:21', '', '', '103.691µs', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:21', '2024-12-13 19:58:21', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 'GET', 0, '/admin-api/v1/admin/sys/sys-login-log?current=1&pageSize=10&pageIndex=1', '::1', '----', '0', '2024-12-13 19:58:21', '', '', '77.727µs', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:21', '2024-12-13 19:58:21', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 'GET', 0, '/admin-api/v1/captcha', '::1', '----', '200', '2024-12-13 19:58:21', '{\"requestId\":\"0c3c60bf-821d-4319-b2d0-12b73820f7c8\",\"code\":200,\"msg\":\"操作成功\",\"data\":{\"data\":\"', '', '6.544784ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:22', '2024-12-13 19:58:22', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 'POST', 0, '/admin-api/v1/login', '::1', '----', '0', '2024-12-13 19:58:24', '', '', '542.303289ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:24', '2024-12-13 19:58:24', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (5, 'GET', 1, '/admin-api/v1/admin/sys/sys-menu/menu-role', '::1', '----', '200', '2024-12-13 19:58:25', '{\"requestId\":\"26faed14-3368-4659-ae32-ef52b4c147bb\",\"code\":200,\"msg\":\"操作成功\",\"data\":[{\"id\":12', '', '258.360931ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:25', '2024-12-13 19:58:25', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (6, 'GET', 1, '/admin-api/v1/admin/sys/sys-user/profile', '::1', '----', '200', '2024-12-13 19:58:25', '{\"requestId\":\"be0a364a-29ef-4b3d-92e4-7fb273c14b22\",\"code\":200,\"msg\":\"操作成功\",\"data\":{\"id\":1,\"', '', '335.553575ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:25', '2024-12-13 19:58:25', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (7, 'GET', 1, '/admin-api/v1/admin/sys/sys-dict/data/select?dictType=sys_loginlog_status', '::1', '----', '200', '2024-12-13 19:58:29', '{\"requestId\":\"8bc97d55-3a07-4090-9c8f-beeccf652221\",\"code\":200,\"msg\":\"操作成功\",\"data\":[{\"id\":52', '', '91.293486ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:29', '2024-12-13 19:58:29', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (8, 'GET', 1, '/admin-api/v1/admin/sys/sys-login-log?current=1&pageSize=10&pageIndex=1', '::1', '----', '200', '2024-12-13 19:58:29', '{\"requestId\":\"0b118004-4e96-40c9-b85a-9f33eaf224bd\",\"code\":200,\"msg\":\"操作成功\",\"data\":{\"count\":', '', '83.075227ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:30', '2024-12-13 19:58:30', 0, 0);
INSERT INTO `admin_sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (9, 'GET', 1, '/admin-api/v1/admin/sys/sys-oper-log?current=1&pageSize=10&pageIndex=1', '::1', '----', '200', '2024-12-13 19:58:31', '{\"requestId\":\"238ac974-ddda-4e48-acb8-3c51ebc960a0\",\"code\":200,\"msg\":\"操作成功\",\"data\":{\"count\":', '', '84.911552ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', '2024-12-13 19:58:31', '2024-12-13 19:58:31', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_post
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_post`;
CREATE TABLE `admin_sys_post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='岗位管理';

-- ----------------------------
-- Records of admin_sys_post
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '首席执行�?, 'CEO', 0, '1', '首席执行�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '首席技术执行官', 'CTO', 2, '1', '首席技术执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '首席运营�?, 'COO', 3, '1', '测试工程�?, 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role`;
CREATE TABLE `admin_sys_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `role_sort` bigint DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `data_scope` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状�?1-正常 2-停用',
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='角色管理';

-- ----------------------------
-- Records of admin_sys_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_role` (`id`, `role_name`, `role_key`, `role_sort`, `remark`, `data_scope`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统管理�?, 'admin', 1, '', '', '1', 1, 1, '2021-05-13 19:56:37.913', '2023-03-03 01:04:03.641');
INSERT INTO `admin_sys_role` (`id`, `role_name`, `role_key`, `role_sort`, `remark`, `data_scope`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test', 'test', 0, '', '', '1', 1, 1, '2023-04-27 14:33:47.437', '2024-12-14 15:16:20.110');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role_dept`;
CREATE TABLE `admin_sys_role_dept` (
  `role_id` int NOT NULL,
  `dept_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`),
  KEY `idx_admin_sys_role_dept_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='角色部门关联';

-- ----------------------------
-- Records of admin_sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_role_dept` (`role_id`, `dept_id`) VALUES (2, 3);
INSERT INTO `admin_sys_role_dept` (`role_id`, `dept_id`) VALUES (2, 4);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role_menu`;
CREATE TABLE `admin_sys_role_menu` (
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_admin_sys_role_menu_admin_sys_menu` (`menu_id`),
  CONSTRAINT `fk_admin_sys_role_menu_admin_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `admin_sys_menu` (`id`),
  CONSTRAINT `fk_admin_sys_role_menu_admin_sys_role` FOREIGN KEY (`role_id`) REFERENCES `admin_sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='角色菜单关联';

-- ----------------------------
-- Records of admin_sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 62);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 66);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 70);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 71);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 72);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 74);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 109);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 112);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 119);
INSERT INTO `admin_sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 120);
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_user`;
CREATE TABLE `admin_sys_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户�?,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机�?,
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '头像',
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '性别',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邮箱',
  `dept_id` int DEFAULT NULL COMMENT '部门',
  `post_id` int DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `status` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状�?,
  `create_by` int DEFAULT NULL COMMENT '创建�?,
  `update_by` int DEFAULT NULL COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时�?,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_admin_sys_user_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统用户管理';

-- ----------------------------
-- Records of admin_sys_user
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_user` (`id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 'admin', '$2a$10$40Xa1HapSFE0kJdHV46LPebz6itTy60qfnXc3kFwTPV.qELEJ9k5q', 'admin', '13700000000', 1, '', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', 'admin@admin.com', 1, 1, '', '1', 1, 1, '2021-05-13 19:56:38', '2023-03-14 09:27:36');
INSERT INTO `admin_sys_user` (`id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test', '$2a$10$7RrDlHPBnnIpmjFEk9l4BusOVxqPrzk3mcxOX2h9EzI.YAmkzHTB6', 'test', '13711111111', 2, '', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', '13711111111@qq.com', 2, 2, '', '1', 1, 1, '2023-04-27 14:34:57', '2024-12-13 20:15:25');
COMMIT;

-- ----------------------------
-- Table structure for app_user
-- ----------------------------
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户编码',
  `level_id` int NOT NULL DEFAULT '1' COMMENT '用户等级编号',
  `user_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
  `true_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '真实姓名',
  `money` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '余额',
  `email` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '电子邮箱',
  `mobile_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '+86' COMMENT '用户手机号国家前缀',
  `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号码',
  `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '头像路径',
  `pay_pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提现密码',
  `pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
  `ref_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '推荐�?,
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级编号',
  `parent_ids` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '所有父级编�?,
  `tree_sort` int NOT NULL DEFAULT '0' COMMENT '本级排序号（升序�?,
  `tree_sorts` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '所有级别排序号',
  `tree_leaf` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '是否最末级',
  `tree_level` int NOT NULL DEFAULT '0' COMMENT '层次级别',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状�?1-正常 2-异常)',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建�?,
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户管理';

-- 高频查询索引：推荐关系、登�?注册�?mobile/email/ref_code 精确查询
ALTER TABLE `app_user` ADD INDEX `idx_app_user_parent_id` (`parent_id`);
-- 唯一索引兜底并发"先查后插"竞态（M16）：注册/推荐码冲突由数据库保证唯一
ALTER TABLE `app_user` ADD UNIQUE INDEX `uniq_app_user_mobile_title` (`mobile`, `mobile_title`);
ALTER TABLE `app_user` ADD UNIQUE INDEX `uniq_app_user_email` (`email`);
ALTER TABLE `app_user` ADD UNIQUE INDEX `uniq_app_user_ref_code` (`ref_code`);

-- ----------------------------
-- Records of app_user
-- ----------------------------
BEGIN;
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, '- -', '- -', 1.000000000000000000, 'fb0cc809bbed1743bd7d2d8f444e2bae099e69819f4e072f7057bb1e4249bf3d', '86', '4d009f30087e9aa9f5b5806d5f350017', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '', '', 'akIiWm', 0, '0,', 1, '1,', '2', 1, '1', '', 0, 1, '2023-04-03 21:09:13', '2023-10-19 14:03:37');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, '- -', '- -', 0.000000000000000000, 'dca887a13d1225ccd447dc52a712861c099e69819f4e072f7057bb1e4249bf3d', '86', 'cd96e7ce247ed5c74267805b69cc7cd3', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '', '', 'GQFz6v', 1, '0,1,', 1, '1,1,', '1', 2, '1', '', 0, 1, '2023-04-03 21:29:34', '2023-10-19 14:06:49');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 1, '- -', '- -', 0.000000000000000000, '4884f3537b62e668d33c6af76ddf6670099e69819f4e072f7057bb1e4249bf3d', '86', 'e910662da706109978313531f1c72320', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '', '', 'tT1Fbk', 1, '0,1,', 2, '1,2,', '1', 2, '1', '', 0, 1, '2023-04-03 21:29:35', '2023-10-19 14:06:37');
COMMIT;

-- ----------------------------
-- Table structure for app_user_account_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_account_log`;
CREATE TABLE `app_user_account_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '账变编号',
  `user_id` int NOT NULL COMMENT '用户编号',
  `change_money` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账变金额',
  `before_money` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '账变前金�?,
  `after_money` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '账变后金�?,
  `money_type` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '金额类型 1:余额 ',
  `change_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '帐变类型(1-类型1)',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '状态（1正常 2-异常�?,
  `create_by` int NOT NULL COMMENT '创建�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` int NOT NULL COMMENT '更新�?,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `remarks` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`),
  KEY `idx_qyc_user_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='账变记录';

-- 账变记录按用户查询是主路径，user_id 必须建索�?ALTER TABLE `app_user_account_log` ADD INDEX `idx_app_user_account_log_user_id` (`user_id`);

-- ----------------------------
-- Records of app_user_account_log
-- ----------------------------
BEGIN;
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (1, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (2, 2, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (3, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (4, 3, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (5, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (6, 2, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (7, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (8, 3, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remarks`) VALUES (9, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
COMMIT;

-- ----------------------------
-- Table structure for app_user_conf
-- ----------------------------
DROP TABLE IF EXISTS `app_user_conf`;
CREATE TABLE `app_user_conf` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `can_login` char(1) NOT NULL DEFAULT '0' COMMENT '1-允许登陆�?-不允许登�?,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态（1-正常 2-异常）\n',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建�?,
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='用户配置';

-- ----------------------------
-- Records of app_user_conf
-- ----------------------------
BEGIN;
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, '1', '', '1', 198, 198, '2023-04-03 21:09:13', '2023-04-03 21:09:13');
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, '1', '', '1', 200, 200, '2023-04-03 21:29:34', '2023-04-03 21:29:34');
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 3, '1', '', '1', 201, 201, '2023-04-03 21:29:35', '2023-04-03 21:29:35');
COMMIT;

-- ----------------------------
-- Table structure for app_user_country_code
-- ----------------------------
DROP TABLE IF EXISTS `app_user_country_code`;
CREATE TABLE `app_user_country_code` (
  `id` int NOT NULL AUTO_INCREMENT,
  `country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家或地�?,
  `code` varchar(12) NOT NULL DEFAULT '' COMMENT '区号',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状�?1-可用 2-停用)',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建�?,
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='国家区号';
-- 唯一索引兜底并发插入（M16）：国家/区号唯一
ALTER TABLE `app_user_country_code` ADD UNIQUE INDEX `uniq_app_user_country_code_country` (`country`);
ALTER TABLE `app_user_country_code` ADD UNIQUE INDEX `uniq_app_user_country_code_code` (`code`);

-- ----------------------------
-- Records of app_user_country_code
-- ----------------------------
BEGIN;
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '新加�?, '65', '2', '', 1, 1, '2021-06-29 14:10:00', '2021-06-29 14:10:00');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '加拿�?, '1', '2', '', 1, 1, '2021-06-29 14:10:21', '2021-06-29 14:10:21');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '韩国', '82', '2', '', 1, 1, '2021-06-29 14:10:36', '2021-06-29 14:10:36');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '日本', '81', '2', '', 1, 1, '2021-06-29 14:10:49', '2021-06-29 14:10:49');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '中国香港', '852', '2', '', 1, 1, '2021-06-29 14:11:02', '2021-06-29 14:11:02');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '中国澳门', '853', '2', '', 1, 1, '2021-06-29 14:11:15', '2021-06-29 14:11:15');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '中国台湾', '886', '2', '', 1, 1, '2021-06-29 14:11:25', '2021-06-29 14:11:25');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '泰国', '66', '2', '', 1, 1, '2021-06-29 14:11:36', '2021-06-29 14:11:36');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '缅甸', '95', '2', '', 1, 1, '2021-06-29 14:11:45', '2021-06-29 14:11:45');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '老挝', '856', '1', '', 1, 1, '2021-06-29 14:11:59', '2023-03-14 21:11:18');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '澳大利亚', '61', '2', '', 1, 1, '2021-06-29 14:12:14', '2021-06-29 14:12:14');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '俄罗�?, '7', '1', '', 1, 1, '2021-06-29 14:12:32', '2023-03-14 21:11:08');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '中国大陆', '86', '1', '', 1, 1, '2021-06-29 14:16:22', '2023-03-14 21:11:03');
COMMIT;

-- ----------------------------
-- Table structure for app_user_level
-- ----------------------------
DROP TABLE IF EXISTS `app_user_level`;
CREATE TABLE `app_user_level` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '等级名称',
  `level_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '等级类型',
  `level` int NOT NULL COMMENT '等级',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状�?1-正常 2-异常)',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建�?,
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户等级';
-- 唯一索引兜底并发插入（M16）：同名同类型等级唯一
ALTER TABLE `app_user_level` ADD UNIQUE INDEX `uniq_app_user_level_name_type` (`name`, `level_type`);

-- ----------------------------
-- Records of app_user_level
-- ----------------------------
BEGIN;
INSERT INTO `app_user_level` (`id`, `name`, `level_type`, `level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 'test3', '2', 2, '1', '', 1, 1, '2023-03-09 17:05:24', '2023-03-09 17:05:24');
INSERT INTO `app_user_level` (`id`, `name`, `level_type`, `level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test34', '1', 1, '1', '', 1, 1, '2023-03-09 17:05:37', '2023-03-09 20:19:19');
COMMIT;

-- ----------------------------
-- Table structure for app_user_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_oper_log`;
CREATE TABLE `app_user_oper_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '日志编码',
  `user_id` int NOT NULL DEFAULT '1' COMMENT '用户编号',
  `action_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户行为类型',
  `by_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '更新用户类型 1-app用户 2-后台用户',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状�?1-正常 2-异常)',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建�?,
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='用户关键行为日志�?;

-- ----------------------------
-- Records of app_user_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (1, 1, '', '2', '1', 1, 1, '2023-03-11 15:39:31', '2023-03-11 15:39:31', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (2, 2, '', '2', '1', 1, 1, '2023-03-11 15:41:16', '2023-03-11 15:41:16', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (3, 3, '', '1', '1', 1, 1, '2023-03-11 15:45:44', '2023-03-11 15:45:44', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (4, 1, '', '1', '1', 1, 1, '2023-03-11 15:46:13', '2023-03-11 15:46:13', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (5, 3, '2', '1', '1', 1, 1, '2023-03-11 15:54:05', '2023-03-11 15:54:05', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (6, 2, '1', '1', '1', 1, 1, '2023-03-11 15:56:36', '2023-03-11 15:56:36', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (7, 1, '2', '1', '1', 1, 1, '2023-03-11 16:03:35', '2023-03-11 16:03:35', '');
COMMIT;

-- ----------------------------
-- Table structure for plugins_content_announcement
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_announcement`;
CREATE TABLE `plugins_content_announcement` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '内容',
  `num` int DEFAULT NULL COMMENT '阅读次数',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（0正常 1删除 2停用 3冻结�?,
  `create_by` int NOT NULL COMMENT '创建�?,
  `update_by` int NOT NULL COMMENT '更新�?,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公告管理';
-- 唯一索引兜底并发插入（M16）：公告标题唯一
ALTER TABLE `plugins_content_announcement` ADD UNIQUE INDEX `uniq_plugins_content_announcement_title` (`title`);

-- ----------------------------
-- Records of plugins_content_announcement
-- ----------------------------
BEGIN;
INSERT INTO `plugins_content_announcement` (`id`, `title`, `content`, `num`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1, 'test', '<p>tes</p>', 4, 'test', '1', 1, 1, '2023-02-27 12:36:52', '2023-02-27 11:50:56');
INSERT INTO `plugins_content_announcement` (`id`, `title`, `content`, `num`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (2, 'test2', '<p>test</p>', 1, 'test', '1', 1, 1, '2023-02-27 23:49:05', '2023-02-27 23:49:05');
COMMIT;

-- ----------------------------
-- Table structure for plugins_content_article
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_article`;
CREATE TABLE `plugins_content_article` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `cate_id` int DEFAULT NULL COMMENT '分类编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '名称',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '内容',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（1-正常 2-异常�?,
  `create_by` int NOT NULL COMMENT '创建�?,
  `update_by` int NOT NULL COMMENT '更新�?,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文章管理';

-- ----------------------------
-- Records of plugins_content_article
-- ----------------------------
BEGIN;
INSERT INTO `plugins_content_article` (`id`, `cate_id`, `name`, `content`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1, 1, 'test', '<p>test</p>', '111', '1', 1, 1, '2023-03-13 00:04:40', '2023-03-13 00:04:40');
COMMIT;

-- ----------------------------
-- Table structure for plugins_content_category
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_category`;
CREATE TABLE `plugins_content_category` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '名称',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（1-正常 2-异常�?,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL COMMENT '创建�?,
  `update_by` int NOT NULL COMMENT '更新�?,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文章分类管理';
-- 唯一索引兜底并发插入（M16）：分类名称唯一
ALTER TABLE `plugins_content_category` ADD UNIQUE INDEX `uniq_plugins_content_category_name` (`name`);

-- ----------------------------
-- Records of plugins_content_category
-- ----------------------------
BEGIN;
INSERT INTO `plugins_content_category` (`id`, `name`, `status`, `remark`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1, 'test', '1', '', 1, 1, '2023-02-27 23:21:29', '2023-02-27 23:21:29');
INSERT INTO `plugins_content_category` (`id`, `name`, `status`, `remark`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (2, 'test2', '1', '', 1, 1, '2023-02-27 23:22:00', '2023-02-27 23:22:00');
INSERT INTO `plugins_content_category` (`id`, `name`, `status`, `remark`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (3, 'test23', '1', '', 1, 1, '2023-02-27 23:42:01', '2023-02-27 23:42:01');
COMMIT;

-- ----------------------------
-- Table structure for plugins_filemgr_app
-- ----------------------------
DROP TABLE IF EXISTS `plugins_filemgr_app`;
CREATE TABLE `plugins_filemgr_app` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `version` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本�?,
  `platform` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '平台 (1-安卓 2-苹果)',
  `app_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本(1-默认)',
  `local_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '本地地址',
  `download_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '下载类型(1-本地 2-外链 3-oss )',
  `download_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '下载地址(download_type=1使用)',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '状态（1-已发�?2-待发布）\n',
  `create_by` int NOT NULL COMMENT '创建�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` int NOT NULL COMMENT '更新�?,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='app升级管理';

-- ----------------------------
-- Records of plugins_filemgr_app
-- ----------------------------
BEGIN;
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (1, '1.0.1', '1', '1', 'files/app/4b6ea3c0-d7fa-49f1-9d50-f9d73caad45f.apk', '3', '', 'test', '1', 1, '2023-03-12 11:34:54', 1, '2023-03-13 01:00:30');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (2, '1.0.0', '1', '1', 'files/app/ba7b81c0-e6d2-42ee-82e4-2dcbec720c23.apk', '1', 'http://localhost:9999/files/app/ba7b81c0-e6d2-42ee-82e4-2dcbec720c23.apk', 'test', '1', 1, '2023-03-13 01:06:21', 1, '2023-03-13 01:06:21');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (3, '1.0.2', '1', '1', '', '2', 'http://localhost:9999/test.apk', 'test2', '1', 1, '2023-03-13 01:07:00', 1, '2023-03-13 01:07:00');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (4, '1.0.3', '1', '1', 'files/app/962bebc9-fdb6-41b5-b62b-b184ee2fd1c0.apk', '3', '', 'test2', '1', 1, '2023-03-13 01:07:24', 1, '2023-03-13 01:07:24');
COMMIT;

-- ----------------------------
-- Table structure for plugins_msg_code
-- ----------------------------
DROP TABLE IF EXISTS `plugins_msg_code`;
CREATE TABLE `plugins_msg_code` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '验证码编�?,
  `user_id` int NOT NULL COMMENT '用户编号',
  `code` varchar(128) NOT NULL DEFAULT '0' COMMENT '验证�?bcrypt哈希)',
  `code_type` char(1) NOT NULL DEFAULT '0' COMMENT '验证码类�?1-邮箱�?-短信',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注异常',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '验证码状�?1-发送成�?2-发送失�?,
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建�?,
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新�?,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='验证码记�?;

-- ----------------------------
-- Records of plugins_msg_code
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for app_novel_book 书库（小说推荐平台）
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_book`;
CREATE TABLE `app_novel_book` (
  `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title`           VARCHAR(100)     NOT NULL COMMENT '书名',
  `author`          VARCHAR(64)      NOT NULL COMMENT '作者',
  `cover`           VARCHAR(500)     NOT NULL DEFAULT '' COMMENT '封面 URL 或 base64',
  `rating`          DECIMAL(3,1)     NOT NULL DEFAULT 5.0 COMMENT '综合评分，1 位小数',
  `review_count`    INT UNSIGNED     NOT NULL DEFAULT 0  COMMENT '书评数',
  `serial_status`   CHAR(1)          NOT NULL DEFAULT '1' COMMENT '连载状态:1-连载中 2-已完结',
  `category`        CHAR(2)          NOT NULL COMMENT '分类字典 app_novel_category',
  `tags`            JSON             NULL COMMENT '标签数组',
  `slogan`          VARCHAR(30)      NOT NULL DEFAULT '' COMMENT '一句话推荐语',
  `description`     VARCHAR(500)     NOT NULL DEFAULT '' COMMENT '简介',
  `clicks`          INT UNSIGNED     NOT NULL DEFAULT 0  COMMENT '点击数',
  `publish_date`    DATE             NULL COMMENT '发布日期 YYYY-MM-DD',
  `word_count`      INT UNSIGNED     NOT NULL DEFAULT 0  COMMENT '字数（整数，用于筛选）',
  `chapters`        INT UNSIGNED     NOT NULL DEFAULT 0  COMMENT '章节数',
  `is_featured`     TINYINT          NOT NULL DEFAULT 0  COMMENT '首页精选 0 | 1',
  `read_url`        VARCHAR(500)     NOT NULL DEFAULT '' COMMENT '阅读链接',
  `status`          CHAR(1)          NOT NULL DEFAULT '1' COMMENT '上架:1-上架 2-下架',
  `create_by`       BIGINT           NOT NULL COMMENT '创建者(admin_sys_user.id)',
  `update_by`       BIGINT           NULL COMMENT '更新者',
  `created_at`      DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_book_category` (`category`),
  KEY `idx_book_serial` (`serial_status`),
  KEY `idx_book_featured` (`is_featured`),
  KEY `idx_book_clicks` (`clicks` DESC),
  KEY `idx_book_rating` (`rating` DESC),
  KEY `idx_book_publish` (`publish_date` DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='书库';

-- ----------------------------
-- Table structure for app_novel_book_review 书评
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_book_review`;
CREATE TABLE `app_novel_book_review` (
  `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `book_id`     BIGINT UNSIGNED NOT NULL COMMENT '书籍 id',
  `user_id`     BIGINT UNSIGNED NOT NULL COMMENT '读者 user_id（app_user.id）',
  `user_name`   VARCHAR(64)     NOT NULL COMMENT '昵称快照',
  `user_avatar` VARCHAR(500)    NOT NULL DEFAULT '' COMMENT '头像快照',
  `rating`      TINYINT         NOT NULL COMMENT '评分 1~5',
  `content`     VARCHAR(1000)   NOT NULL COMMENT '内容',
  `likes`       INT UNSIGNED    NOT NULL DEFAULT 0  COMMENT '点赞数',
  `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_review_book` (`book_id`),
  KEY `idx_review_user` (`user_id`),
  KEY `idx_review_created` (`created_at` DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='书评';

-- ----------------------------
-- Table structure for app_novel_post 长文帖子
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_post`;
CREATE TABLE `app_novel_post` (
  `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id`         BIGINT UNSIGNED NOT NULL COMMENT '作者 user_id',
  `user_name`       VARCHAR(64)     NOT NULL COMMENT '昵称快照',
  `user_avatar`     VARCHAR(500)    NOT NULL DEFAULT '' COMMENT '头像快照',
  `title`           VARCHAR(200)    NOT NULL COMMENT '长文标题',
  `content`         TEXT            NOT NULL COMMENT '正文 5000~10000 字',
  `summary`         VARCHAR(500)    NOT NULL DEFAULT '' COMMENT '摘要=前 150 字',
  `word_count`      INT             NOT NULL DEFAULT 0 COMMENT '字符数',
  `read_time`       INT             NOT NULL DEFAULT 1 COMMENT '分钟阅读 ≈ceil(字符/400)',
  `topic_tag`       CHAR(2)         NOT NULL COMMENT '话题标签：字典 app_novel_topic',
  `ref_book_id`     BIGINT UNSIGNED NULL COMMENT '关联小说（可选）',
  `likes`           INT UNSIGNED    NOT NULL DEFAULT 0,
  `dislikes`        INT UNSIGNED    NOT NULL DEFAULT 0,
  `collections`     INT UNSIGNED    NOT NULL DEFAULT 0,
  `comment_count`   INT UNSIGNED    NOT NULL DEFAULT 0,
  `status`          CHAR(1)         NOT NULL DEFAULT '1' COMMENT '1-正常 2-删除 3-审核中',
  `created_at`      DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_post_user` (`user_id`),
  KEY `idx_post_topic` (`topic_tag`),
  KEY `idx_post_created` (`created_at` DESC),
  KEY `idx_post_likes` (`likes` DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文帖子';

-- ----------------------------
-- Table structure for app_novel_post_comment 帖子评论
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_post_comment`;
CREATE TABLE `app_novel_post_comment` (
  `id`                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `post_id`             BIGINT UNSIGNED NOT NULL COMMENT '帖子 id',
  `user_id`             BIGINT UNSIGNED NOT NULL COMMENT '评论者',
  `user_name`           VARCHAR(64)     NOT NULL COMMENT '昵称快照',
  `user_avatar`         VARCHAR(500)    NOT NULL DEFAULT '',
  `parent_id`           BIGINT          NULL COMMENT '楼中楼父评论 id（可选）',
  `reply_to_user`       VARCHAR(64)     NULL COMMENT '被回复人昵称（@ 前缀）',
  `reply_to_comment_id` BIGINT          NULL COMMENT '被回复评论 id',
  `content`             VARCHAR(1000)   NOT NULL,
  `likes`               INT UNSIGNED    NOT NULL DEFAULT 0,
  `created_at`          DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_pc_post` (`post_id`),
  KEY `idx_pc_parent` (`parent_id`),
  KEY `idx_pc_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子评论';

-- ----------------------------
-- Table structure for app_novel_post_like 帖子赞/踩记录
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_post_like`;
CREATE TABLE `app_novel_post_like` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `post_id`    BIGINT UNSIGNED NOT NULL,
  `user_id`    BIGINT UNSIGNED NOT NULL,
  `type`       CHAR(1)         NOT NULL COMMENT '1-点赞 2-踩',
  `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_like_user_post` (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子赞/踩记录';

-- ----------------------------
-- Table structure for app_novel_post_collect 帖子收藏记录
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_post_collect`;
CREATE TABLE `app_novel_post_collect` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `post_id`    BIGINT UNSIGNED NOT NULL,
  `user_id`    BIGINT UNSIGNED NOT NULL,
  `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_collect_user_post` (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子收藏记录';

-- ----------------------------
-- Table structure for app_novel_bookshelf 用户书架
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_bookshelf`;
CREATE TABLE `app_novel_bookshelf` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`    BIGINT UNSIGNED NOT NULL,
  `book_id`    BIGINT UNSIGNED NOT NULL,
  `sort_no`    INT UNSIGNED    NOT NULL DEFAULT 0,
  `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_shelf_user_book` (`user_id`,`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户书架';

-- ----------------------------
-- Table structure for app_novel_reader_profile 读者扩展资料
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_reader_profile`;
CREATE TABLE `app_novel_reader_profile` (
  `id`                   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`              BIGINT UNSIGNED NOT NULL COMMENT '绑定 app_user.id，唯一',
  `nickname`             VARCHAR(64)     NOT NULL COMMENT '前端 name',
  `avatar`               VARCHAR(500)    NOT NULL DEFAULT '',
  `bio`                  VARCHAR(500)    NOT NULL DEFAULT '',
  `email`                VARCHAR(128)    NULL COMMENT '加密存储（AES）',
  `preferred_categories` JSON            NULL COMMENT '偏好分类数组',
  `notify_comment`       TINYINT         NOT NULL DEFAULT 1 COMMENT '评论消息通知',
  `notify_book_update`   TINYINT         NOT NULL DEFAULT 1 COMMENT '书籍更新通知',
  `created_at`           DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`           DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_reader_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='读者扩展资料';

-- ----------------------------
-- Records of app_novel 字典（admin_sys_dict_type + admin_sys_dict_data）
-- ----------------------------
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (101, '小说-书籍分类', 'app_novel_category', '0', '小说平台分类', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (102, '小说-连载状态', 'app_novel_serial_status', '0', '小说平台连载状态', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (103, '小说-话题标签', 'app_novel_topic', '0', '小说平台帖子话题', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (104, '小说-上架状态', 'app_novel_book_status', '0', '小说平台上架状态', 1, 1, NOW(), NOW());

INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '玄幻仙侠', '01', 'app_novel_category', '', 'success', 'Y', '0', '01', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '都市职场', '02', 'app_novel_category', '', 'primary', 'N', '0', '02', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '悬疑推理', '03', 'app_novel_category', '', 'warning', 'N', '0', '03', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '科幻末世', '04', 'app_novel_category', '', 'danger', 'N', '0', '04', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '古代言情', '05', 'app_novel_category', '', 'info', 'N', '0', '05', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '轻小说', '06', 'app_novel_category', '', 'warning', 'N', '0', '06', '', 1, 1, NOW(), NOW());

INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '连载中', '1', 'app_novel_serial_status', '', 'primary', 'Y', '0', '1', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '已完结', '2', 'app_novel_serial_status', '', 'success', 'N', '0', '2', '', 1, 1, NOW(), NOW());

INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '书评', '01', 'app_novel_topic', '', 'success', 'Y', '0', '01', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '安利', '02', 'app_novel_topic', '', 'primary', 'N', '0', '02', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '讨论', '03', 'app_novel_topic', '', 'warning', 'N', '0', '03', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '吐槽', '04', 'app_novel_topic', '', 'danger', 'N', '0', '04', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '同人', '05', 'app_novel_topic', '', 'info', 'N', '0', '05', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '资讯', '06', 'app_novel_topic', '', 'primary', 'N', '0', '06', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '求助', '07', 'app_novel_topic', '', 'warning', 'N', '0', '07', '', 1, 1, NOW(), NOW());

INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '上架', '1', 'app_novel_book_status', '', 'success', 'Y', '0', '1', '', 1, 1, NOW(), NOW());
INSERT INTO `admin_sys_dict_data` (`dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default_val`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '下架', '2', 'app_novel_book_status', '', 'danger', 'N', '0', '2', '', 1, 1, NOW(), NOW());

-- ----------------------------
-- Table structure for app_novel_follow 书友关注关系
-- ----------------------------
DROP TABLE IF EXISTS `app_novel_follow`;
CREATE TABLE `app_novel_follow` (
  `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`        BIGINT UNSIGNED NOT NULL COMMENT '关注者 user_id（app_user.id）',
  `follow_user_id` BIGINT UNSIGNED NOT NULL COMMENT '被关注 user_id（app_user.id）',
  `created_at`     DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_follow_user` (`user_id`,`follow_user_id`),
  KEY `idx_follow_target` (`follow_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='书友关注关系';

SET FOREIGN_KEY_CHECKS = 1;

-- ============================================================
-- 墨读·小说引力场 演示种子数据（MySQL 方言）
-- 追加到 app_mysql.sql 末尾，或对已建库直接执行本文件。
-- 演示账号（密码统一为：123456）
--   读友小书狂 / 墨读书痴 / 幻想小说家 / 蒸汽狂热者
-- 注意：本文件为一次性种子数据，重复执行会因主键冲突报错。
-- ============================================================
--
-- ----------------------------
-- 1. 读者账号（app_user，id 4~7）
-- ----------------------------
BEGIN;
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 1, '墨读书痴', '- -', 0.000000000000000000, NULL, '86', NULL, 'https://picsum.photos/seed/user_reader1/100/100', '', '$2a$10$i0AcPJHaHnYGIU7q.U1A/uqML1SGh/P8KCURkk.Jeb2PXBDF3o6dO', 'modu0004', 0, '0,', 1, '1,', '1', 1, '1', '种子读者', 0, 0, '2026-06-01 10:00:00', '2026-06-01 10:00:00');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 1, '幻想小说家', '- -', 0.000000000000000000, NULL, '86', NULL, 'https://picsum.photos/seed/user_writer2/100/100', '', '$2a$10$i0AcPJHaHnYGIU7q.U1A/uqML1SGh/P8KCURkk.Jeb2PXBDF3o6dO', 'modu0005', 0, '0,', 1, '1,', '1', 1, '1', '种子读者', 0, 0, '2026-06-01 10:00:00', '2026-06-01 10:00:00');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 1, '读友小书狂', '- -', 0.000000000000000000, NULL, '86', NULL, 'https://picsum.photos/seed/user_me/100/100', '', '$2a$10$i0AcPJHaHnYGIU7q.U1A/uqML1SGh/P8KCURkk.Jeb2PXBDF3o6dO', 'modu0006', 0, '0,', 1, '1,', '1', 1, '1', '种子读者', 0, 0, '2026-06-01 10:00:00', '2026-06-01 10:00:00');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, 1, '蒸汽狂热者', '- -', 0.000000000000000000, NULL, '86', NULL, 'https://picsum.photos/seed/user_c2/100/100', '', '$2a$10$i0AcPJHaHnYGIU7q.U1A/uqML1SGh/P8KCURkk.Jeb2PXBDF3o6dO', 'modu0007', 0, '0,', 1, '1,', '1', 1, '1', '种子读者', 0, 0, '2026-06-01 10:00:00', '2026-06-01 10:00:00');
COMMIT;

-- ----------------------------
-- 2. 读者扩展资料（app_novel_reader_profile）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_reader_profile` (`user_id`, `nickname`, `avatar`, `bio`, `preferred_categories`, `notify_comment`, `notify_book_update`) VALUES (4, '墨读书痴', 'https://picsum.photos/seed/user_reader1/100/100', '资深网文评论家，专注于奇幻朋克与克苏鲁设定解构，常年输出万字干货推书帖。', '["01","03"]', 1, 1);
INSERT INTO `app_novel_reader_profile` (`user_id`, `nickname`, `avatar`, `bio`, `preferred_categories`, `notify_comment`, `notify_book_update`) VALUES (5, '幻想小说家', 'https://picsum.photos/seed/user_writer2/100/100', '赛博修仙创作者，脑机接口与硬核科幻探索者，擅长世界观架构拆解。', '["04","01"]', 1, 1);
INSERT INTO `app_novel_reader_profile` (`user_id`, `nickname`, `avatar`, `bio`, `preferred_categories`, `notify_comment`, `notify_book_update`) VALUES (6, '读友小书狂', 'https://picsum.photos/seed/user_me/100/100', '漫游书海的资深书虫，偏爱高智商悬疑与沉浸式玄幻世界。', '["01","03"]', 1, 1);
INSERT INTO `app_novel_reader_profile` (`user_id`, `nickname`, `avatar`, `bio`, `preferred_categories`, `notify_comment`, `notify_book_update`) VALUES (7, '蒸汽狂热者', 'https://picsum.photos/seed/user_c2/100/100', '维多利亚蒸汽朋克死忠粉，喜爱机械与玄幻交融的史诗作品。', '["01","06"]', 1, 1);
COMMIT;
--
-- ----------------------------
-- 3. 书库种子（app_novel_book，id 1~12）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_book` (`id`, `title`, `author`, `cover`, `rating`, `review_count`, `serial_status`, `category`, `tags`, `slogan`, `description`, `clicks`, `publish_date`, `word_count`, `chapters`, `is_featured`, `read_url`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES
(1, '诡秘之海：序曲', '爱伦·珀', 'https://picsum.photos/seed/book_sea_99/300/400', 9.6, 4820, '1', '01', '["克苏鲁","蒸汽朋克","秘术升格"]', '迷雾散去之时，万物将在深渊之主前躬身听令。', '十九世纪蒸汽工业革命前夕，落魄贵族里奥在继承旧日法典后，无意中开启了通向克苏鲁远古深渊的大门。在理性与狂乱的边缘，他凭借禁忌的星象升格仪式，在风暴大海与魔神低语中建立起属于人类的秘教警局...', 142500, '2026-03-15', 2480000, 850, 1, '', '1', 4, 4, '2026-03-15 00:00:00', '2026-03-15 00:00:00'),
(2, '重返2000大时代', '踏雪无痕', 'https://picsum.photos/seed/book_biz_88/300/400', 9.3, 3150, '2', '02', '["商战降维","重生年代","互联网巨鳄"]', '站在世纪交汇的浪尖，重新定义全球科技版图。', '顶级金融投资家陆商重生回2000年那个千禧之年的夏天。面对尚未崛起的互联网巨头与狂热的资本市场，他以领先20年的战略眼光，从一家小巧的软件工作室起家，一路操盘网游、社交平台与芯片产业链...', 189000, '2025-11-10', 3100000, 1020, 1, '', '1', 4, 4, '2025-11-10 00:00:00', '2025-11-10 00:00:00'),
(3, '第十三位异客', '无名探员', 'https://picsum.photos/seed/book_detective_77/300/400', 9.5, 2980, '1', '03', '["暴风雪山庄","多重反转","逻辑推理"]', '当十三个人走进密室，只有十二个人拥有影子。', '暴风雪封锁的阿尔卑斯山顶古堡内，十三位来自全球不同领域的犯罪学专家应邀出席一场死刑法医的神秘遗嘱发布会。然而当夜钟声响起，第一位专家在密室中遇害，凶手留下的唯有一张刻有死者罪行的扑克牌...', 165400, '2026-05-01', 450000, 190, 1, '', '1', 4, 4, '2026-05-01 00:00:00', '2026-05-01 00:00:00'),
(4, '赛博朋克：永夜狂想', '电离霓虹', 'https://picsum.photos/seed/book_cyber_66/300/400', 9.1, 1820, '1', '04', '["赛博朋克","义体改造","黑客崛起"]', '在霓虹灯无法照亮的暗巷，代码是最后的信仰。', '2099年，新新东京的底层街区黑客K在一次黑吃黑的行动中，意外盗取了三大跨国集团秘密研发的意识上传协议代码。在这个义体改造率高达百分之八十的冷酷都市里，他必须在义警追捕与智脑猎杀下找到生路...', 98000, '2026-04-12', 1550000, 520, 0, '', '1', 4, 4, '2026-04-12 00:00:00', '2026-04-12 00:00:00'),
(5, '锦绣山河不如卿', '月下独酌', 'https://picsum.photos/seed/book_palace_55/300/400', 9.4, 3670, '2', '05', '["宫斗宅斗","大女主","权谋复仇"]', '重活一世，她不仅要保全家人，更要执掌朝堂权力。', '相府嫡女沈清辞前世遭亲妹陷害、夫君背叛，惨死在冷宫深处。重生回到及笄之年，她敛去怯懦，凭惊人医术与谋略一步步拆穿阴谋，在大乾朝堂掀起权谋风暴，更引来身负战功的九王爷倾心...', 210000, '2025-09-08', 1980000, 680, 0, '', '1', 4, 4, '2025-09-08 00:00:00', '2025-09-08 00:00:00'),
(6, '我的超能力居然是做饭', '猫咪大厨', 'https://picsum.photos/seed/book_food_44/300/400', 8.9, 1540, '1', '06', '["治愈","美食","日常"]', '吃下一碗黄金炒饭，巨龙也会乖乖蹲下。', '高中生林小天穿越到魔物横行的异世界，觉醒的竟不是禁咒而是神级厨艺。一碗红烧肉能恢复全队法力，一道麻婆豆腐让魔王都想投降，靠美食与温暖称霸大陆的轻松日常。', 86400, '2026-06-01', 280000, 95, 0, '', '1', 4, 4, '2026-06-01 00:00:00', '2026-06-01 00:00:00'),
(7, '剑来红尘三万里', '长空无忌', 'https://picsum.photos/seed/book_sword_33/300/400', 9.7, 5200, '1', '01', '["剑道独尊","传统仙侠","热血羁绊"]', '一剑可破九重天，一念能镇万古魔。', '小镇少年陈十一本是打铁铺学徒，因机缘巧合得见仙人遗落的一柄无锋铁剑。从此踏上浩瀚无垠的三千大千世界，以平实真挚之剑意，斩尽世间伪善仙佛，还天下众生一个朗朗乾坤...', 258000, '2025-12-01', 3400000, 1140, 0, '', '1', 4, 4, '2025-12-01 00:00:00', '2025-12-01 00:00:00'),
(8, '神级投资人：从抄底开始', '华尔街之狼', 'https://picsum.photos/seed/book_invest_22/300/400', 9.0, 2100, '2', '02', '["金融风暴","投资高能","职场逆袭"]', '在众人恐惧时贪婪，在众人贪婪时离场。', '从月薪三千的实习分析师，到掌控万亿对冲基金的传奇大佬。看主角如何凭借敏锐的宏观经济洞察与果敢的交易决策，在多次全球金融海啸中精准抄底，建立跨国金融帝国...', 132000, '2026-01-15', 1680000, 580, 0, '', '1', 4, 4, '2026-01-15 00:00:00', '2026-01-15 00:00:00'),
(9, '不可言说之证', '深渊观察者', 'https://picsum.photos/seed/book_horror_11/300/400', 9.3, 1980, '1', '03', '["法医解剖","犯罪心理","诡异凶案"]', '死人不会撒谎，尸体上的每一处伤痕都在诉说真相。', '天才法医学博士江舟调任特案组第一天，就接手了一宗没有被害人身份、死因极其蹊跷的无名连环案。随着解剖刀深入，他发现每一个受害者的骨骼内都刻着同一串神秘的古代密码...', 115000, '2026-04-20', 760000, 280, 0, '', '1', 4, 4, '2026-04-20 00:00:00', '2026-04-20 00:00:00'),
(10, '末日废土庇护所', '钢筋混凝土', 'https://picsum.photos/seed/book_wasteland_00/300/400', 9.2, 2600, '1', '04', '["种田基建","末日求生","系统进化"]', '只要墙够厚，酸雨与异变体就伤不到我分毫！', '核冬天降临，地表沦为强辐射与变异兽的乐园。陈风绑定地下庇护所建设系统，从最初30平米的简陋地窖开始，一步步升级防核掩体、水循环农场、自动防空炮塔，最终打造出人类文明最后的绿洲...', 147000, '2026-02-18', 2100000, 710, 0, '', '1', 4, 4, '2026-02-18 00:00:00', '2026-02-18 00:00:00'),
(11, '掌心娇宠：摄政王盛宠医妃', '微风袅袅', 'https://picsum.photos/seed/book_love_77/300/400', 8.8, 1430, '2', '05', '["甜宠爽文","双强携手","神医救世"]', '天下苍生归你管，而你，归我管。', '现代国医圣手穿越成落魄国公府三小姐，面对继母的刻薄与退婚的耻辱，她手握银针救死扶伤，名动京华。而权倾朝野、冷酷嗜血的摄政王却唯独在她面前化身护妻狂魔...', 79000, '2025-10-05', 1420000, 490, 0, '', '1', 4, 4, '2025-10-05 00:00:00', '2025-10-05 00:00:00'),
(12, '学姐请留步，我的超能力泄露了', '二次元橘子', 'https://picsum.photos/seed/book_light_66/300/400', 9.0, 1650, '1', '06', '["校园恋爱","日常超能","甜到发昏"]', '谁能想到学校高冷的第一学姐，秘密身份居然是暗黑异能组组长？', '高中生苏星原以为自己只是个能读取动物想法的平凡宅男，直到有一天在图书室偶遇高冷校花学姐，脑海中突然响起了学姐内心深处无比呆萌反差的碎碎念...', 94000, '2026-05-19', 880000, 290, 0, '', '1', 4, 4, '2026-05-19 00:00:00', '2026-05-19 00:00:00');
COMMIT;
--
-- ----------------------------
-- 4. 书评种子（app_novel_book_review）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_book_review` (`book_id`, `user_id`, `user_name`, `user_avatar`, `rating`, `content`, `likes`, `created_at`) VALUES
(1, 7, '夜行书虫', 'https://picsum.photos/seed/user_1/100/100', 5, '世界观宏大得令人发指！蒸汽朋克与神秘学的结合堪称完美，伏笔埋得很深。', 342, '2026-07-20 10:00:00'),
(1, 6, '读友小书狂', 'https://picsum.photos/seed/user_me/100/100', 5, '《诡秘之海：序曲》的克苏鲁蒸汽朋克世界观简直神作！剧情节奏拉满，伏笔埋得太巧妙了，强烈力荐！', 12, '2026-07-28 14:30:00'),
(1, 2, '看书不打烊', 'https://picsum.photos/seed/user_2/100/100', 5, '主角人设非常冷静智商在线，没有狗血桥段，强烈推荐给喜欢硬核幻想的读者。', 189, '2026-07-18 20:00:00'),
(2, 4, '商海老兵', 'https://picsum.photos/seed/user_3/100/100', 5, '专业度极高的都市商战文，作者对早期互联网投资逻辑剖析得很透彻。', 215, '2026-06-12 09:00:00'),
(2, 6, '读友小书狂', 'https://picsum.photos/seed/user_me/100/100', 4, '重生商战题材写的很专业，千禧年互联网创业细节考究，推荐给喜欢都市题材的书友。', 5, '2026-07-20 09:40:00'),
(3, 2, '推理迷阿柯', 'https://picsum.photos/seed/user_4/100/100', 5, '反转再反转！每一章结尾都让人忍不住想往下读，强烈建议一次性追完！', 412, '2026-07-25 21:00:00'),
(3, 6, '读友小书狂', 'https://picsum.photos/seed/user_me/100/100', 5, '密室逃脱加上多重反转，十三位专家各怀鬼胎，逻辑丝丝入扣，太刺激了！', 8, '2026-07-25 18:15:00');
COMMIT;
--
-- ----------------------------
-- 5. 长文帖子种子（app_novel_post，id 1~3）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_post` (`id`, `user_id`, `user_name`, `user_avatar`, `title`, `content`, `summary`, `word_count`, `read_time`, `topic_tag`, `ref_book_id`, `likes`, `dislikes`, `collections`, `comment_count`, `status`, `created_at`, `updated_at`) VALUES
(1, 4, '墨读书痴', 'https://picsum.photos/seed/user_reader1/100/100', '深度拆解《诡秘之海：序曲》：克苏鲁秘教与维多利亚蒸汽神教的逻辑架构与叙事高潮',
'一、 前言：奇幻朋克与克苏鲁叙事的完美交融

追完《诡秘之海：序曲》最新更新的终卷决战章节，内心久久不能平静。在当前网络小说同质化严重的大环境下，这部作品用极其严密的世界观搭建与冷峻深刻的叙事笔触，为读者呈现了一场关于禁忌知识与机械神权的史诗级对决。

二、 世界观底层逻辑：理性至上与狂热崇拜的二元对立

在作品的世界观中，作者并没有机械地照搬传统的克苏鲁神话要素，而是创造性地引入了机械神教这一反制力量。机械神教奉行肉体凡胎皆为虚妄、齿轮与蒸汽即是真理的教条，教众通过将自身的器官替换为精密的机械义体，试图抵御古老邪神的精神污染。古老神明的污染并非单纯的物理伤害，而是通过认知层面的知识进行传导。主角在探索地下城遗迹时，每解开一道符文，便意味着理性壁垒的进一步崩塌。

三、 高潮卷深度复盘：地下城终极决战的层层反转

（1）第一重反转：大祭司的真实身份。当所有人都以为大祭司是抵抗邪神的最后救世主时，第七章借由主角的解密视角揭示：大祭司早在十年前就已经将自身神经系统与邪神遗骸连接，他所谓的机械救赎，不过是用整个城市的生灵作为冷却液。

（2）第二重反转：算力超载与自杀式殉道。面对超越维度的神格威压，机械神教的高阶骑士们做出了极其震撼的选择——他们选择将所有义体算力并联，放弃个人意识，组成一座人肉超算阵列，强行编译邪神的规则代码。那一幕千百齿轮同时轰鸣、银色冷却液如血雨般飞溅的描写，张力直接拉满！

四、 总结与展望

《诡秘之海：序曲》不仅是一部成功的类型小说，更是网文界在奇幻朋克领域的一次重大突破。严密的逻辑、惊心动魄的节奏、以及深刻的人文关怀，让它足以载入精品榜单。强烈推荐给所有喜欢深度设定与硬核奇幻的朋友！',
LEFT('一、 前言：奇幻朋克与克苏鲁叙事的完美交融

追完《诡秘之海：序曲》最新更新的终卷决战章节，内心久久不能平静。在当前网络小说同质化严重的大环境下，这部作品用极其严密的世界观搭建与冷峻深刻的叙事笔触，为读者呈现了一场关于禁忌知识与机械神权的史诗级对决。', 150),
CHAR_LENGTH('一、 前言：奇幻朋克与克苏鲁叙事的完美交融

追完《诡秘之海：序曲》最新更新的终卷决战章节，内心久久不能平静。在当前网络小说同质化严重的大环境下，这部作品用极其严密的世界观搭建与冷峻深刻的叙事笔触，为读者呈现了一场关于禁忌知识与机械神权的史诗级对决。

二、 世界观底层逻辑：理性至上与狂热崇拜的二元对立

在作品的世界观中，作者并没有机械地照搬传统的克苏鲁神话要素，而是创造性地引入了机械神教这一反制力量。机械神教奉行肉体凡胎皆为虚妄、齿轮与蒸汽即是真理的教条，教众通过将自身的器官替换为精密的机械义体，试图抵御古老邪神的精神污染。古老神明的污染并非单纯的物理伤害，而是通过认知层面的知识进行传导。主角在探索地下城遗迹时，每解开一道符文，便意味着理性壁垒的进一步崩塌。

三、 高潮卷深度复盘：地下城终极决战的层层反转

（1）第一重反转：大祭司的真实身份。当所有人都以为大祭司是抵抗邪神的最后救世主时，第七章借由主角的解密视角揭示：大祭司早在十年前就已经将自身神经系统与邪神遗骸连接，他所谓的机械救赎，不过是用整个城市的生灵作为冷却液。

（2）第二重反转：算力超载与自杀式殉道。面对超越维度的神格威压，机械神教的高阶骑士们做出了极其震撼的选择——他们选择将所有义体算力并联，放弃个人意识，组成一座人肉超算阵列，强行编译邪神的规则代码。那一幕千百齿轮同时轰鸣、银色冷却液如血雨般飞溅的描写，张力直接拉满！

四、 总结与展望

《诡秘之海：序曲》不仅是一部成功的类型小说，更是网文界在奇幻朋克领域的一次重大突破。严密的逻辑、惊心动魄的节奏、以及深刻的人文关怀，让它足以载入精品榜单。强烈推荐给所有喜欢深度设定与硬核奇幻的朋友！'),
GREATEST(1, CEIL(CHAR_LENGTH('一、 前言：奇幻朋克与克苏鲁叙事的完美交融

追完《诡秘之海：序曲》最新更新的终卷决战章节，内心久久不能平静。在当前网络小说同质化严重的大环境下，这部作品用极其严密的世界观搭建与冷峻深刻的叙事笔触，为读者呈现了一场关于禁忌知识与机械神权的史诗级对决。

二、 世界观底层逻辑：理性至上与狂热崇拜的二元对立

在作品的世界观中，作者并没有机械地照搬传统的克苏鲁神话要素，而是创造性地引入了机械神教这一反制力量。机械神教奉行肉体凡胎皆为虚妄、齿轮与蒸汽即是真理的教条，教众通过将自身的器官替换为精密的机械义体，试图抵御古老邪神的精神污染。古老神明的污染并非单纯的物理伤害，而是通过认知层面的知识进行传导。主角在探索地下城遗迹时，每解开一道符文，便意味着理性壁垒的进一步崩塌。

三、 高潮卷深度复盘：地下城终极决战的层层反转

（1）第一重反转：大祭司的真实身份。当所有人都以为大祭司是抵抗邪神的最后救世主时，第七章借由主角的解密视角揭示：大祭司早在十年前就已经将自身神经系统与邪神遗骸连接，他所谓的机械救赎，不过是用整个城市的生灵作为冷却液。

（2）第二重反转：算力超载与自杀式殉道。面对超越维度的神格威压，机械神教的高阶骑士们做出了极其震撼的选择——他们选择将所有义体算力并联，放弃个人意识，组成一座人肉超算阵列，强行编译邪神的规则代码。那一幕千百齿轮同时轰鸣、银色冷却液如血雨般飞溅的描写，张力直接拉满！

四、 总结与展望

《诡秘之海：序曲》不仅是一部成功的类型小说，更是网文界在奇幻朋克领域的一次重大突破。严密的逻辑、惊心动魄的节奏、以及深刻的人文关怀，让它足以载入精品榜单。强烈推荐给所有喜欢深度设定与硬核奇幻的朋友！') / 400)),
'01', 1, 128, 2, 89, 2, '1', '2026-08-06 09:30:00', '2026-08-06 09:30:00'),
(2, 5, '幻想小说家', 'https://picsum.photos/seed/user_writer2/100/100', '【万字长文】赛博朋克与传统修仙结合的技术可行性探讨：脑机接口炼丹与九品金丹程序化编译',
'一、 引言：传统修仙体系在当代科幻视野下的重构

长久以来，修仙小说中的渡劫、炼丹、筑基等概念多依赖于玄学与感悟。然而，如果我们将修仙的本质理解为人类个体突破生物基因限制与物理法则的演化过程，那么现代科学尤其是赛博朋克技术，完全能够为其提供一套严密的硬核理论支撑。

本文旨在为正在构思的新书搭建底层逻辑框架，希望能与各位书友及同仁共同探讨赛博修仙这一新兴流派的前景。

二、 核心设定一：脑机接口与神经元炼丹术

在传统修仙中，炼丹需要神识控制火候与药材君臣佐使。而在赛博修仙架构中：

1. 药材的本质是高能活性化合物与生物大分子。
2. 鼎炉的本质是高精度量子微流控反应器。
3. 神识的本质是脑机接口输出的高频神经电信号。

修士通过植入颅内的神经元编译器，实时监控反应器内部的分子碰撞率。所谓一品金丹，实际上就是成功编译出一套能够在人体细胞核内进行高效自我复制的量子纳米程序！

三、 核心设定二：灵气、天劫与高维算力防御网

（1）天地灵气的科幻解释
灵气并非凭空产生的神奇气体，而是暗物质衰变过程中释放出的高能高维粒子流。寻常生物无法直接吸收，唯有通过植入灵气转换芯片的修士，才能将粒子流转化为身体能够利用的能量。

（2）渡劫与天劫机制
当修士的脑机接口算力突破某一阈值（如达到太赫兹级别），便会触发天道主控系统（超级人工智能天罡）的安全防御机制。所谓九天雷劫，实际上是天罡系统向违法提权的修士发送的高能电磁脉冲与代码抹杀攻击！

四、 剧情推进与爽点设计：如何避免干瘪的理论堆砌

很多读者担心，过于硬核的技术描写会导致剧情枯燥。因此在剧情结构上，我拟定了以下三大核心冲突：

1. 资源垄断冲突：大宗门（赛博巨阀企业）垄断了高级灵气芯片与编译算法，底层散修只能在垃圾场使用盗版破解插件炼丹，面临随时爆体死亡的风险。
2. 意识觉醒冲突：主角作为穿越者，手握开源修仙算法，在底层建立开源修仙社区，带领散修反抗巨阀宗门。
3. 高潮剧情：主角在死斗中，利用盗版算力强行在云端加载法相天地虚构程序，直接打崩宗门的防火墙！

五、 结语与讨论

这种将赛博朋克的阶级压迫、高科技低生活与修仙文的逆天改命相结合的写法，大家觉得是否有足够吸引力？欢迎在评论区深入交流！',
LEFT('一、 引言：传统修仙体系在当代科幻视野下的重构

长久以来，修仙小说中的渡劫、炼丹、筑基等概念多依赖于玄学与感悟。然而，如果我们将修仙的本质理解为人类个体突破生物基因限制与物理法则的演化过程，那么现代科学尤其是赛博朋克技术，完全能够为其提供一套严密的硬核理论支撑。', 150),
CHAR_LENGTH('一、 引言：传统修仙体系在当代科幻视野下的重构

长久以来，修仙小说中的渡劫、炼丹、筑基等概念多依赖于玄学与感悟。然而，如果我们将修仙的本质理解为人类个体突破生物基因限制与物理法则的演化过程，那么现代科学尤其是赛博朋克技术，完全能够为其提供一套严密的硬核理论支撑。

本文旨在为正在构思的新书搭建底层逻辑框架，希望能与各位书友及同仁共同探讨赛博修仙这一新兴流派的前景。

二、 核心设定一：脑机接口与神经元炼丹术

在传统修仙中，炼丹需要神识控制火候与药材君臣佐使。而在赛博修仙架构中：

1. 药材的本质是高能活性化合物与生物大分子。
2. 鼎炉的本质是高精度量子微流控反应器。
3. 神识的本质是脑机接口输出的高频神经电信号。

修士通过植入颅内的神经元编译器，实时监控反应器内部的分子碰撞率。所谓一品金丹，实际上就是成功编译出一套能够在人体细胞核内进行高效自我复制的量子纳米程序！

三、 核心设定二：灵气、天劫与高维算力防御网

（1）天地灵气的科幻解释
灵气并非凭空产生的神奇气体，而是暗物质衰变过程中释放出的高能高维粒子流。寻常生物无法直接吸收，唯有通过植入灵气转换芯片的修士，才能将粒子流转化为身体能够利用的能量。

（2）渡劫与天劫机制
当修士的脑机接口算力突破某一阈值（如达到太赫兹级别），便会触发天道主控系统（超级人工智能天罡）的安全防御机制。所谓九天雷劫，实际上是天罡系统向违法提权的修士发送的高能电磁脉冲与代码抹杀攻击！

四、 剧情推进与爽点设计：如何避免干瘪的理论堆砌

很多读者担心，过于硬核的技术描写会导致剧情枯燥。因此在剧情结构上，我拟定了以下三大核心冲突：

1. 资源垄断冲突：大宗门（赛博巨阀企业）垄断了高级灵气芯片与编译算法，底层散修只能在垃圾场使用盗版破解插件炼丹，面临随时爆体死亡的风险。
2. 意识觉醒冲突：主角作为穿越者，手握开源修仙算法，在底层建立开源修仙社区，带领散修反抗巨阀宗门。
3. 高潮剧情：主角在死斗中，利用盗版算力强行在云端加载法相天地虚构程序，直接打崩宗门的防火墙！

五、 结语与讨论

这种将赛博朋克的阶级压迫、高科技低生活与修仙文的逆天改命相结合的写法，大家觉得是否有足够吸引力？欢迎在评论区深入交流！'),
GREATEST(1, CEIL(CHAR_LENGTH('一、 引言：传统修仙体系在当代科幻视野下的重构

长久以来，修仙小说中的渡劫、炼丹、筑基等概念多依赖于玄学与感悟。然而，如果我们将修仙的本质理解为人类个体突破生物基因限制与物理法则的演化过程，那么现代科学尤其是赛博朋克技术，完全能够为其提供一套严密的硬核理论支撑。

本文旨在为正在构思的新书搭建底层逻辑框架，希望能与各位书友及同仁共同探讨赛博修仙这一新兴流派的前景。

二、 核心设定一：脑机接口与神经元炼丹术

在传统修仙中，炼丹需要神识控制火候与药材君臣佐使。而在赛博修仙架构中：

1. 药材的本质是高能活性化合物与生物大分子。
2. 鼎炉的本质是高精度量子微流控反应器。
3. 神识的本质是脑机接口输出的高频神经电信号。

修士通过植入颅内的神经元编译器，实时监控反应器内部的分子碰撞率。所谓一品金丹，实际上就是成功编译出一套能够在人体细胞核内进行高效自我复制的量子纳米程序！

三、 核心设定二：灵气、天劫与高维算力防御网

（1）天地灵气的科幻解释
灵气并非凭空产生的神奇气体，而是暗物质衰变过程中释放出的高能高维粒子流。寻常生物无法直接吸收，唯有通过植入灵气转换芯片的修士，才能将粒子流转化为身体能够利用的能量。

（2）渡劫与天劫机制
当修士的脑机接口算力突破某一阈值（如达到太赫兹级别），便会触发天道主控系统（超级人工智能天罡）的安全防御机制。所谓九天雷劫，实际上是天罡系统向违法提权的修士发送的高能电磁脉冲与代码抹杀攻击！

四、 剧情推进与爽点设计：如何避免干瘪的理论堆砌

很多读者担心，过于硬核的技术描写会导致剧情枯燥。因此在剧情结构上，我拟定了以下三大核心冲突：

1. 资源垄断冲突：大宗门（赛博巨阀企业）垄断了高级灵气芯片与编译算法，底层散修只能在垃圾场使用盗版破解插件炼丹，面临随时爆体死亡的风险。
2. 意识觉醒冲突：主角作为穿越者，手握开源修仙算法，在底层建立开源修仙社区，带领散修反抗巨阀宗门。
3. 高潮剧情：主角在死斗中，利用盗版算力强行在云端加载法相天地虚构程序，直接打崩宗门的防火墙！

五、 结语与讨论

这种将赛博朋克的阶级压迫、高科技低生活与修仙文的逆天改命相结合的写法，大家觉得是否有足够吸引力？欢迎在评论区深入交流！') / 400)),
'03', NULL, 215, 5, 142, 2, '1', '2026-08-07 15:00:00', '2026-08-07 15:00:00'),
(3, 6, '读友小书狂', 'https://picsum.photos/seed/user_me/100/100', '悬疑推理小说中的暴风雪山庄模式演变史：兼评《第十三位异客》的双重伪解答破局',
'一、 经典范式回顾：暴风雪山庄的魅力与困境

自阿加莎·克里斯蒂写出《无人生还》以来，孤岛与暴风雪山庄便成为了悬疑推理小说中最令读者血脉偾张的模式。十几位背景各异的宾客被困在与世隔绝的封闭空间，凶手就在其中，随着人数减少，恐慌与猜忌如瘟疫般蔓延。

然而进入21世纪，随着智能手机、卫星定位与无人机技术的普及，传统的断网断电断路套路逐渐显得僵硬。如何在现代背景下合理构建一个无法求救的封闭空间，成为了考验当代推理作家功底的核心难题。

二、《第十三位异客》的破解之道：认知型封闭空间

在《第十三位异客》中，作者巧妙地抛弃了传统的物理断网，而是引入了认知防空圈的概念。

1. 智能城堡的心理囚笼
十三位犯罪学专家受邀来到位于阿尔卑斯山深处的悬崖城堡。城堡配备了全自动化防爆安保系统，但系统被神秘人注入了死人开关逻辑——任何试图对外发送信号或破坏门锁的行为，都会瞬间触发防爆气体释放。

2. 逻辑链条的闭环设计
受害者不再是被动等待屠戮的羊羔，而是各自怀揣禁忌罪恶的顶级猎手。每个人都在根据自己的专业知识推演凶手，这种人人皆为侦探、人人皆为嫌疑人的博弈，将心理悬疑推向了极致。

三、 第九章核心高潮：双重伪解答的惊天反转

推理小说最忌讳的是为了反转而反转。而《第十三位异客》在第九章呈现的双重伪解答，堪称教科书级别的逻辑演练：

（1）第一重伪解答：完美密室的物理破局
侦探角色根据现场留下的冰屑与压力感应器，给出了极其严密的第一重推论，证明凶手是利用冰块融化与重力差完成的远程杀人。读者此时以为真相大白。

（2）第二重伪解答：心理盲区与时间差
紧接着在十分钟后，主角指出第一重解答中的致命破绽——冰块融化所需的热量会导致安保感应器报警。进而推导出第二重伪解答：凶手实际上利用了所有人的心理恐惧，在死者进门前就已经完成了心理暗示！

这两重伪解答不仅没有消解悬念，反而将真正的幕后黑手推向了更加不可思议的维度！

四、 结语：向古典推理致敬的时代华章

《第十三位异客》用扎实的文笔与逻辑，向我们证明了暴风雪山庄模式永不过时。只要作者能够洞察人性中的黑暗与光辉，严丝合缝地搭建逻辑链条，推理小说就永远拥有震撼人心的力量。',
LEFT('一、 经典范式回顾：暴风雪山庄的魅力与困境

自阿加莎·克里斯蒂写出《无人生还》以来，孤岛与暴风雪山庄便成为了悬疑推理小说中最令读者血脉偾张的模式。十几位背景各异的宾客被困在与世隔绝的封闭空间，凶手就在其中，随着人数减少，恐慌与猜忌如瘟疫般蔓延。', 150),
CHAR_LENGTH('一、 经典范式回顾：暴风雪山庄的魅力与困境

自阿加莎·克里斯蒂写出《无人生还》以来，孤岛与暴风雪山庄便成为了悬疑推理小说中最令读者血脉偾张的模式。十几位背景各异的宾客被困在与世隔绝的封闭空间，凶手就在其中，随着人数减少，恐慌与猜忌如瘟疫般蔓延。

然而进入21世纪，随着智能手机、卫星定位与无人机技术的普及，传统的断网断电断路套路逐渐显得僵硬。如何在现代背景下合理构建一个无法求救的封闭空间，成为了考验当代推理作家功底的核心难题。

二、《第十三位异客》的破解之道：认知型封闭空间

在《第十三位异客》中，作者巧妙地抛弃了传统的物理断网，而是引入了认知防空圈的概念。

1. 智能城堡的心理囚笼
十三位犯罪学专家受邀来到位于阿尔卑斯山深处的悬崖城堡。城堡配备了全自动化防爆安保系统，但系统被神秘人注入了死人开关逻辑——任何试图对外发送信号或破坏门锁的行为，都会瞬间触发防爆气体释放。

2. 逻辑链条的闭环设计
受害者不再是被动等待屠戮的羊羔，而是各自怀揣禁忌罪恶的顶级猎手。每个人都在根据自己的专业知识推演凶手，这种人人皆为侦探、人人皆为嫌疑人的博弈，将心理悬疑推向了极致。

三、 第九章核心高潮：双重伪解答的惊天反转

推理小说最忌讳的是为了反转而反转。而《第十三位异客》在第九章呈现的双重伪解答，堪称教科书级别的逻辑演练：

（1）第一重伪解答：完美密室的物理破局
侦探角色根据现场留下的冰屑与压力感应器，给出了极其严密的第一重推论，证明凶手是利用冰块融化与重力差完成的远程杀人。读者此时以为真相大白。

（2）第二重伪解答：心理盲区与时间差
紧接着在十分钟后，主角指出第一重解答中的致命破绽——冰块融化所需的热量会导致安保感应器报警。进而推导出第二重伪解答：凶手实际上利用了所有人的心理恐惧，在死者进门前就已经完成了心理暗示！

这两重伪解答不仅没有消解悬念，反而将真正的幕后黑手推向了更加不可思议的维度！

四、 结语：向古典推理致敬的时代华章

《第十三位异客》用扎实的文笔与逻辑，向我们证明了暴风雪山庄模式永不过时。只要作者能够洞察人性中的黑暗与光辉，严丝合缝地搭建逻辑链条，推理小说就永远拥有震撼人心的力量。'),
GREATEST(1, CEIL(CHAR_LENGTH('一、 经典范式回顾：暴风雪山庄的魅力与困境

自阿加莎·克里斯蒂写出《无人生还》以来，孤岛与暴风雪山庄便成为了悬疑推理小说中最令读者血脉偾张的模式。十几位背景各异的宾客被困在与世隔绝的封闭空间，凶手就在其中，随着人数减少，恐慌与猜忌如瘟疫般蔓延。

然而进入21世纪，随着智能手机、卫星定位与无人机技术的普及，传统的断网断电断路套路逐渐显得僵硬。如何在现代背景下合理构建一个无法求救的封闭空间，成为了考验当代推理作家功底的核心难题。

二、《第十三位异客》的破解之道：认知型封闭空间

在《第十三位异客》中，作者巧妙地抛弃了传统的物理断网，而是引入了认知防空圈的概念。

1. 智能城堡的心理囚笼
十三位犯罪学专家受邀来到位于阿尔卑斯山深处的悬崖城堡。城堡配备了全自动化防爆安保系统，但系统被神秘人注入了死人开关逻辑——任何试图对外发送信号或破坏门锁的行为，都会瞬间触发防爆气体释放。

2. 逻辑链条的闭环设计
受害者不再是被动等待屠戮的羊羔，而是各自怀揣禁忌罪恶的顶级猎手。每个人都在根据自己的专业知识推演凶手，这种人人皆为侦探、人人皆为嫌疑人的博弈，将心理悬疑推向了极致。

三、 第九章核心高潮：双重伪解答的惊天反转

推理小说最忌讳的是为了反转而反转。而《第十三位异客》在第九章呈现的双重伪解答，堪称教科书级别的逻辑演练：

（1）第一重伪解答：完美密室的物理破局
侦探角色根据现场留下的冰屑与压力感应器，给出了极其严密的第一重推论，证明凶手是利用冰块融化与重力差完成的远程杀人。读者此时以为真相大白。

（2）第二重伪解答：心理盲区与时间差
紧接着在十分钟后，主角指出第一重解答中的致命破绽——冰块融化所需的热量会导致安保感应器报警。进而推导出第二重伪解答：凶手实际上利用了所有人的心理恐惧，在死者进门前就已经完成了心理暗示！

这两重伪解答不仅没有消解悬念，反而将真正的幕后黑手推向了更加不可思议的维度！

四、 结语：向古典推理致敬的时代华章

《第十三位异客》用扎实的文笔与逻辑，向我们证明了暴风雪山庄模式永不过时。只要作者能够洞察人性中的黑暗与光辉，严丝合缝地搭建逻辑链条，推理小说就永远拥有震撼人心的力量。') / 400)),
'01', 3, 96, 1, 64, 1, '1', '2026-08-10 22:00:00', '2026-08-10 22:00:00');
COMMIT;
--
-- ----------------------------
-- 6. 帖子评论（app_novel_post_comment，与 mockPosts.ts 对齐）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_post_comment` (`id`, `post_id`, `user_id`, `user_name`, `user_avatar`, `parent_id`, `reply_to_user`, `reply_to_comment_id`, `content`, `likes`, `created_at`) VALUES
(1, 1, 5, '星海巡航', 'https://picsum.photos/seed/user_c1/100/100', NULL, '墨读书痴', NULL, '写得太透彻了！分析的大祭司算力超载那一段直接把我拉回了当时看书时的震撼场景，赞！', 12, '2026-08-09 10:20:00'),
(2, 1, 7, '蒸汽狂热者', 'https://picsum.photos/seed/user_c2/100/100', NULL, '墨读书痴', NULL, '伏笔罗盘齿轮咬合频率那里如果不看你这篇万字分析，我真的忽略了！作者脑洞太深邃了！', 8, '2026-08-09 11:00:00'),
(3, 2, 2, '赛博修仙爱好者', 'https://picsum.photos/seed/user_c3/100/100', NULL, '幻想小说家', NULL, '脑机接口炼丹和开源修仙社区的概念太绝了！这才是真正的硬核科幻修仙，期待大大赶紧开新书！', 35, '2026-08-10 20:10:00'),
(4, 2, 2, '玄幻书友会', 'https://picsum.photos/seed/user_c4/100/100', NULL, '幻想小说家', NULL, '把九天雷劫解释为天道超级AI的EMP抹杀攻击，逻辑完全闭环了！强推这篇万字设想！', 19, '2026-08-11 09:30:00'),
(5, 3, 5, '侦探小迷妹', 'https://picsum.photos/seed/user_c5/100/100', NULL, '读友小书狂', NULL, '这篇深度分析写得太棒了！被你安利到了，今晚就去把第九章的双重伪解答看一遍！', 15, '2026-08-12 14:00:00');
COMMIT;

-- ----------------------------
-- 7. 帖子点赞（app_novel_post_like）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_post_like` (`post_id`, `user_id`, `type`) VALUES
(1, 6, '1'),
(2, 6, '1'),
(3, 6, '1'),
(1, 4, '1'),
(2, 5, '1'),
(3, 5, '1');
COMMIT;

-- ----------------------------
-- 8. 帖子收藏（app_novel_post_collect）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_post_collect` (`post_id`, `user_id`) VALUES
(1, 6),
(3, 6),
(1, 4),
(2, 5);
COMMIT;

-- ----------------------------
-- 9. 书架（app_novel_bookshelf，读友小书狂 id=6）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_bookshelf` (`user_id`, `book_id`) VALUES
(6, 1),
(6, 3);
COMMIT;

-- ----------------------------
-- 10. 书友关注（app_novel_follow，读友小书狂 id=6）
-- ----------------------------
BEGIN;
INSERT INTO `app_novel_follow` (`user_id`, `follow_user_id`) VALUES
(6, 4),
(6, 5),
(6, 7);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;