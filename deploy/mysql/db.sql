/*
 Navicat Premium Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : sim_db

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 26/08/2024 10:31:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat_models
-- ----------------------------
DROP TABLE IF EXISTS `chat_models`;
CREATE TABLE `chat_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `send_user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `rev_user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `msg_type` tinyint NULL DEFAULT NULL,
  `msg_preview` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `msg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `system_msg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_chat_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat_models
-- ----------------------------
INSERT INTO `chat_models` VALUES (1, '2024-07-21 10:51:04.000', '2024-07-21 10:51:07.000', NULL, 1, 2, 1, '哈哈哈', '{\r\n    \"type\": 1,\r\n    \"textMsg\": \"1发送的第一条消息\"\r\n}', NULL);
INSERT INTO `chat_models` VALUES (2, '2024-07-21 10:51:37.000', '2024-07-21 10:51:40.000', NULL, 1, 2, 1, '啦啦啦', '{\r\n    \"type\": 1,\r\n    \"textMsg\": \"1发送的第二条消息\"\r\n}', '{\"type\":5,\"content\":\"对方已经同意了你的好友请求\"}');
INSERT INTO `chat_models` VALUES (3, '2024-07-21 10:52:10.000', '2024-07-21 10:52:13.000', NULL, 2, 1, 1, '嘻嘻嘻', '{\r\n    \"type\": 1,\r\n    \"textMsg\": \"1接收的第一条消息\"\r\n}', NULL);
INSERT INTO `chat_models` VALUES (4, '2024-07-21 10:52:35.000', '2024-07-21 10:52:37.000', NULL, 2, 3, 1, '啦啦啦', '{\r\n    \"type\": 1,\r\n    \"textMsg\": \"2发给3的第一条消息\"\r\n}', NULL);
INSERT INTO `chat_models` VALUES (5, '2024-07-21 10:53:01.000', '2024-07-21 10:53:03.000', NULL, 3, 2, 1, '嘻嘻嘻', '{\r\n    \"type\": 1,\r\n    \"textMsg\": \"3发给2的第一条消息\"\r\n}', NULL);
INSERT INTO `chat_models` VALUES (6, '2024-08-03 17:07:41.000', '2024-08-03 17:07:44.000', NULL, 1, 3, 1, '8.3 号发送的第一条消息', '{\r\n    \"type\": 1,\r\n    \"textMsg\": \"8.3 森发送给 张亨睿 的第一条消息\"\r\n}', NULL);
INSERT INTO `chat_models` VALUES (8, '2024-08-04 15:56:50.649', '2024-08-04 15:56:50.649', NULL, 7, 1, 1, '[系统消息]对方已经同意了你的好友请求', '{\"type\":1,\"textMsg\":\"我们已经是好友了,开始聊天吧！\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null}', '{\"type\":5,\"content\":\"对方已经同意了你的好友请求\"}');
INSERT INTO `chat_models` VALUES (9, '2024-08-09 16:32:44.312', '2024-08-09 16:32:44.312', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (10, '2024-08-09 16:34:25.023', '2024-08-09 16:34:25.023', NULL, 1, 3, 13, 'sen上线了', '{\"type\":13,\"textMsg\":\"sen上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (11, '2024-08-09 16:35:54.172', '2024-08-09 16:35:54.172', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (12, '2024-08-09 16:40:56.249', '2024-08-09 16:40:56.249', NULL, 1, 3, 13, 'sen上线了', '{\"type\":13,\"textMsg\":\"sen上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (13, '2024-08-09 16:41:40.662', '2024-08-09 16:41:40.662', NULL, 1, 3, 13, 'sen上线了', '{\"type\":13,\"textMsg\":\"sen上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (14, '2024-08-09 16:43:34.735', '2024-08-09 16:43:34.735', NULL, 1, 3, 13, 'sen上线了', '{\"type\":13,\"textMsg\":\"sen上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (15, '2024-08-09 16:45:13.591', '2024-08-09 16:45:13.591', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (16, '2024-08-09 16:48:09.014', '2024-08-09 16:48:09.014', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (17, '2024-08-09 16:50:17.767', '2024-08-09 16:50:17.767', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (18, '2024-08-09 20:27:23.701', '2024-08-09 20:27:23.701', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (19, '2024-08-09 20:31:59.171', '2024-08-09 20:31:59.171', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (20, '2024-08-09 20:35:09.620', '2024-08-09 20:35:09.620', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (21, '2024-08-09 20:35:32.131', '2024-08-09 20:35:32.131', NULL, 1, 3, 13, 'sen上线了', '{\"type\":13,\"textMsg\":\"sen上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (22, '2024-08-09 20:36:40.115', '2024-08-09 20:36:40.115', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (23, '2024-08-09 20:58:13.719', '2024-08-09 20:58:13.719', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (24, '2024-08-09 20:58:20.054', '2024-08-09 20:58:20.054', NULL, 1, 3, 1, '你好呀！', '{\"type\":1,\"textMsg\":\"你好呀！\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (25, '2024-08-09 21:19:01.109', '2024-08-09 21:19:01.109', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (26, '2024-08-09 21:19:06.236', '2024-08-09 21:27:01.002', NULL, 1, 3, 8, '[撤回消息]-撤回了一条消息猜猜我撤回了什么？', '{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"msgId\":26},\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (27, '2024-08-09 21:19:15.588', '2024-08-09 21:19:15.588', NULL, 3, 1, 1, '森哥你也好！', '{\"type\":1,\"textMsg\":\"森哥你也好！\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (28, '2024-08-09 21:24:17.373', '2024-08-09 21:24:17.373', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (29, '2024-08-09 21:27:01.015', '2024-08-09 21:27:01.015', NULL, 1, 3, 8, '[撤回消息] - 你好呀！', '{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"msgId\":26},\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (30, '2024-08-09 22:10:17.589', '2024-08-09 22:10:17.589', NULL, 3, 1, 13, '张亨睿上线了', '{\"type\":13,\"textMsg\":\"张亨睿上线了\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (31, '2024-08-09 22:12:11.760', '2024-08-09 22:15:44.375', NULL, 1, 3, 8, '[撤回消息]-撤回了一条消息猜猜我撤回了什么？', '{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"msgId\":31},\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);
INSERT INTO `chat_models` VALUES (32, '2024-08-09 22:15:44.383', '2024-08-09 22:15:44.383', NULL, 1, 3, 8, '[撤回消息] - 哈哈哈哈哈哈哈哈哈哈！', '{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"msgId\":31},\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL);

-- ----------------------------
-- Table structure for file_models
-- ----------------------------
DROP TABLE IF EXISTS `file_models`;
CREATE TABLE `file_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uid` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `file_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `size` bigint NULL DEFAULT NULL,
  `path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `hash` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_file_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of file_models
-- ----------------------------
INSERT INTO `file_models` VALUES (3, '2024-07-09 15:44:34.877', '2024-07-09 15:44:34.877', NULL, 'af7dde69-afc1-43c2-9155-030ae962fec9', 1, '110.jpg', 3046146, 'file/api/uploads/avatar/af7dde69-afc1-43c2-9155-030ae962fec9.jpg', 'bad15a09fedae249b3022a1b77b1887f');
INSERT INTO `file_models` VALUES (10, '2024-08-08 16:07:51.199', '2024-08-08 16:07:51.199', NULL, 'de2e641f-4432-48e8-ab1f-0261da29e0a7', 1, '最长括号长度.txt', 1347, 'file/api/uploads/file/user_id_1/de2e641f-4432-48e8-ab1f-0261da29e0a7.txt', '6d8e3af87af7d587c6c1c5e9b10e16f2');
INSERT INTO `file_models` VALUES (11, '2024-08-08 17:51:40.447', '2024-08-08 17:51:40.447', NULL, 'f1ba55ae-f03e-42c8-b171-b38461de50de', 1, 'Go语言开发-刘宏森简历.docx', 11502442, 'file/api/uploads/file/user_id_1/f1ba55ae-f03e-42c8-b171-b38461de50de.docx', 'ae30a3636a4ccbfda6bc15fc95772734');

-- ----------------------------
-- Table structure for friend_verify_models
-- ----------------------------
DROP TABLE IF EXISTS `friend_verify_models`;
CREATE TABLE `friend_verify_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `send_user_id` bigint UNSIGNED NOT NULL,
  `rev_user_id` bigint UNSIGNED NOT NULL,
  `status` tinyint NULL DEFAULT 0,
  `send_status` tinyint NULL DEFAULT 0,
  `rev_status` tinyint NULL DEFAULT 0,
  `additional_messages` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `verification_question` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `send_time` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_friend_verify_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_friend_verify_models_send_user`(`send_user_id` ASC) USING BTREE,
  INDEX `fk_friend_verify_models_rev_user`(`rev_user_id` ASC) USING BTREE,
  CONSTRAINT `fk_friend_verify_models_rev_user` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_friend_verify_models_send_user` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friend_verify_models
-- ----------------------------
INSERT INTO `friend_verify_models` VALUES (6, '2024-08-04 15:56:50.643', '2024-08-04 15:56:50.643', NULL, 1, 7, 0, 0, 1, '你好，我想添加你为好友！', NULL, 1722758210);

-- ----------------------------
-- Table structure for friendship_models
-- ----------------------------
DROP TABLE IF EXISTS `friendship_models`;
CREATE TABLE `friendship_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `friend_id` bigint UNSIGNED NOT NULL,
  `accepted` tinyint(1) NULL DEFAULT 0,
  `alias` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_friendship_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_user_models_friends`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_user_models_friends` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friendship_models
-- ----------------------------
INSERT INTO `friendship_models` VALUES (1, '2024-07-07 22:50:13.000', '2024-07-10 11:48:14.703', NULL, 1, 2, 1, '森哥');
INSERT INTO `friendship_models` VALUES (2, '2024-07-10 10:23:47.000', '2024-07-10 10:23:49.000', NULL, 1, 3, 1, '亨睿');
INSERT INTO `friendship_models` VALUES (3, '2024-07-10 10:24:13.000', '2024-07-10 10:24:15.000', NULL, 4, 1, 1, '正清');
INSERT INTO `friendship_models` VALUES (4, '2024-07-10 10:24:35.000', '2024-07-10 10:24:37.000', NULL, 3, 4, 1, '正清');
INSERT INTO `friendship_models` VALUES (5, '2024-07-10 10:25:13.000', '2024-07-10 10:25:15.000', NULL, 4, 2, 0, NULL);
INSERT INTO `friendship_models` VALUES (6, '2024-07-10 10:25:13.000', '2024-07-10 10:25:15.000', NULL, 1, 5, 1, NULL);
INSERT INTO `friendship_models` VALUES (7, '2024-07-10 10:25:13.000', '2024-07-10 10:25:15.000', NULL, 1, 6, 1, NULL);
INSERT INTO `friendship_models` VALUES (14, '2024-08-04 15:56:50.636', '2024-08-04 15:56:50.636', NULL, 1, 7, 1, '');

-- ----------------------------
-- Table structure for group_member_models
-- ----------------------------
DROP TABLE IF EXISTS `group_member_models`;
CREATE TABLE `group_member_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `group_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_forbid` tinyint(1) NULL DEFAULT NULL,
  `forbid_time` bigint NULL DEFAULT NULL,
  `join_time` bigint NULL DEFAULT NULL,
  `role` tinyint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_member_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_group_models_member_list`(`group_id` ASC) USING BTREE,
  CONSTRAINT `fk_group_models_member_list` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_member_models
-- ----------------------------
INSERT INTO `group_member_models` VALUES (1, '2024-08-17 21:17:33.000', NULL, NULL, 1, 1, '森(群聊版)', NULL, NULL, NULL, 2);
INSERT INTO `group_member_models` VALUES (2, '2024-08-17 21:18:12.000', NULL, NULL, 1, 2, '花生(群聊版)', NULL, NULL, NULL, 1);
INSERT INTO `group_member_models` VALUES (3, '2024-08-17 21:19:03.000', NULL, NULL, 1, 3, '张亨睿(群聊版)', NULL, NULL, NULL, 0);
INSERT INTO `group_member_models` VALUES (4, '2024-08-17 21:19:38.000', NULL, NULL, 1, 4, '雷正清(群聊版)', NULL, NULL, NULL, 0);

-- ----------------------------
-- Table structure for group_member_roles
-- ----------------------------
DROP TABLE IF EXISTS `group_member_roles`;
CREATE TABLE `group_member_roles`  (
  `group_member_model_id` bigint UNSIGNED NOT NULL,
  `group_user_role_model_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`group_member_model_id`, `group_user_role_model_id`) USING BTREE,
  INDEX `fk_group_member_roles_group_user_role_model`(`group_user_role_model_id` ASC) USING BTREE,
  CONSTRAINT `fk_group_member_roles_group_member_model` FOREIGN KEY (`group_member_model_id`) REFERENCES `group_member_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_group_member_roles_group_user_role_model` FOREIGN KEY (`group_user_role_model_id`) REFERENCES `group_user_role_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_member_roles
-- ----------------------------

-- ----------------------------
-- Table structure for group_models
-- ----------------------------
DROP TABLE IF EXISTS `group_models`;
CREATE TABLE `group_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `group_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `group_desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `group_avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `creator_id` bigint UNSIGNED NULL DEFAULT NULL,
  `is_search` tinyint(1) NULL DEFAULT NULL,
  `verify_type` tinyint NULL DEFAULT NULL,
  `verify_question` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_invite` tinyint(1) NULL DEFAULT NULL,
  `is_temp_session` tinyint(1) NULL DEFAULT NULL,
  `is_forbidden` tinyint(1) NULL DEFAULT NULL,
  `size` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_models
-- ----------------------------
INSERT INTO `group_models` VALUES (1, '2024-08-17 20:57:49.000', NULL, NULL, '测试群聊', '测试', 'https://profile-avatar.csdnimg.cn/010e6e0b3ee6405cb7e3b500c9b5dbf5_qq_21879995.jpg', 1, 1, 1, NULL, 1, NULL, NULL, 50);

-- ----------------------------
-- Table structure for group_msg_models
-- ----------------------------
DROP TABLE IF EXISTS `group_msg_models`;
CREATE TABLE `group_msg_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `group_id` bigint UNSIGNED NULL DEFAULT NULL,
  `sender_id` bigint UNSIGNED NULL DEFAULT NULL,
  `group_member_id` bigint UNSIGNED NULL DEFAULT NULL,
  `msg_type` tinyint NULL DEFAULT NULL,
  `msg_preview` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `msg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `system_msg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `timestamp` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_msg_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_group_member_models_msg_list`(`group_member_id` ASC) USING BTREE,
  INDEX `fk_group_models_messages`(`group_id` ASC) USING BTREE,
  CONSTRAINT `fk_group_member_models_msg_list` FOREIGN KEY (`group_member_id`) REFERENCES `group_member_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_group_models_messages` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_msg_models
-- ----------------------------
INSERT INTO `group_msg_models` VALUES (23, '2024-08-21 21:55:07.253', '2024-08-21 21:55:07.253', NULL, 1, 1, 1, 1, '群主发送第一条消息', '{\"type\":1,\"textMsg\":\"群主发送第一条消息\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL, '2024-08-21 21:55:07.252');
INSERT INTO `group_msg_models` VALUES (24, '2024-08-21 21:56:21.118', '2024-08-21 21:56:21.118', NULL, 1, 1, 1, 1, '群主发送第一条消息', '{\"type\":1,\"textMsg\":\"群主发送第一条消息\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL, '2024-08-21 21:56:21.118');
INSERT INTO `group_msg_models` VALUES (25, '2024-08-21 21:56:41.679', '2024-08-21 21:56:41.679', NULL, 1, 4, 4, 1, '雷正清发送一条消息', '{\"type\":1,\"textMsg\":\"雷正清发送一条消息\",\"imageMsg\":null,\"fileMsg\":null,\"audioMsg\":null,\"videoMsg\":null,\"voiceMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"forwardMsg\":null,\"replyMsg\":null,\"atMsg\":null,\"tipMsg\":null,\"friendOnlineMsg\":null,\"imageTextMsg\":null}', NULL, '2024-08-21 21:56:41.679');

-- ----------------------------
-- Table structure for group_user_role_models
-- ----------------------------
DROP TABLE IF EXISTS `group_user_role_models`;
CREATE TABLE `group_user_role_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `role_desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `permissions` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_user_role_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_user_role_models
-- ----------------------------

-- ----------------------------
-- Table structure for group_verify_models
-- ----------------------------
DROP TABLE IF EXISTS `group_verify_models`;
CREATE TABLE `group_verify_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `group_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT NULL,
  `additional_messages` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `verification_question` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `type` tinyint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_verify_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_group_verify_models_group_model`(`group_id` ASC) USING BTREE,
  CONSTRAINT `fk_group_verify_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_verify_models
-- ----------------------------

-- ----------------------------
-- Table structure for log_models
-- ----------------------------
DROP TABLE IF EXISTS `log_models`;
CREATE TABLE `log_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `log_type` tinyint NULL DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `level` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `service` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_log_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_models
-- ----------------------------

-- ----------------------------
-- Table structure for msg_models
-- ----------------------------
DROP TABLE IF EXISTS `msg_models`;
CREATE TABLE `msg_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `sender_id` bigint UNSIGNED NULL DEFAULT NULL,
  `recver_id` bigint UNSIGNED NULL DEFAULT NULL,
  `msg_type` tinyint NULL DEFAULT NULL,
  `msg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `system_msg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `timestamp` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_msg_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_user_models_message_sent`(`sender_id` ASC) USING BTREE,
  INDEX `fk_user_models_message_recv`(`recver_id` ASC) USING BTREE,
  CONSTRAINT `fk_user_models_message_recv` FOREIGN KEY (`recver_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_models_message_sent` FOREIGN KEY (`sender_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of msg_models
-- ----------------------------

-- ----------------------------
-- Table structure for top_user_models
-- ----------------------------
DROP TABLE IF EXISTS `top_user_models`;
CREATE TABLE `top_user_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `top_user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_top_user_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of top_user_models
-- ----------------------------

-- ----------------------------
-- Table structure for user_chat_delete_models
-- ----------------------------
DROP TABLE IF EXISTS `user_chat_delete_models`;
CREATE TABLE `user_chat_delete_models`  (
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `chat_id` bigint UNSIGNED NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_chat_delete_models
-- ----------------------------
INSERT INTO `user_chat_delete_models` VALUES (1, 1);

-- ----------------------------
-- Table structure for user_conf_models
-- ----------------------------
DROP TABLE IF EXISTS `user_conf_models`;
CREATE TABLE `user_conf_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `recall_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `friend_online` tinyint(1) NULL DEFAULT NULL,
  `sound` tinyint(1) NULL DEFAULT NULL,
  `secure_link` tinyint(1) NULL DEFAULT NULL,
  `save_pwd` tinyint(1) NULL DEFAULT NULL,
  `search_user` tinyint NULL DEFAULT NULL,
  `verification` tinyint NULL DEFAULT NULL,
  `verification_question` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `online_status` tinyint(1) NULL DEFAULT NULL,
  `forbid_chat` tinyint(1) NULL DEFAULT NULL,
  `forbid_add_user` tinyint(1) NULL DEFAULT NULL,
  `forbid_create_group` tinyint(1) NULL DEFAULT NULL,
  `forbid_in_group_chat` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_conf_models_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_user_models_user_conf`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_user_models_user_conf` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_conf_models
-- ----------------------------
INSERT INTO `user_conf_models` VALUES (1, '2024-07-03 10:57:16.000', '2024-07-03 11:00:27.657', NULL, 1, '猜猜我撤回了什么？', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `user_conf_models` VALUES (2, '2024-07-07 22:35:28.000', '2024-07-07 22:35:30.000', NULL, 2, '卑微地撤回了一条消息!', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `user_conf_models` VALUES (3, '2024-07-07 22:35:28.000', '2024-07-07 22:35:30.000', NULL, 3, '卑微地撤回了一条消息!', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `user_conf_models` VALUES (4, '2024-07-07 22:35:28.000', '2024-07-07 22:35:30.000', NULL, 4, '卑微地撤回了一条消息!', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `user_conf_models` VALUES (5, '2024-07-07 22:35:28.000', '2024-07-07 22:35:30.000', NULL, 5, '卑微地撤回了一条消息!', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `user_conf_models` VALUES (6, '2024-07-07 22:35:28.000', '2024-07-07 22:35:30.000', NULL, 6, '卑微地撤回了一条消息!', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `user_conf_models` VALUES (7, '2024-07-07 22:35:28.000', '2024-07-07 22:35:30.000', NULL, 7, '卑微地撤回了一条消息!', 1, 0, 1, 1, 2, 1, NULL, 1, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for user_models
-- ----------------------------
DROP TABLE IF EXISTS `user_models`;
CREATE TABLE `user_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `pwd` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `abstract` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `open_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `register_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `role` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_models
-- ----------------------------
INSERT INTO `user_models` VALUES (1, '2024-07-02 15:48:57.000', '2024-07-03 11:00:27.650', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', 'sen', '啦啦啦啦啦啦啦啦啦', 'https://thirdwx.qlogo.cn/mmopen/vi_32/icZJCCzianSibe7kHIwCSCXbjbbnpHMKDaePZQlrhFTIOraziap9Pxp9vhWEJhvTZTAg4sbRbDvzVeKbZxzicJILMzqOs1XRrqpLH9HbvySjUIWo/132', '179621078@qq.com', NULL, '江西省赣州市', NULL, '系统注册', 2);
INSERT INTO `user_models` VALUES (2, '2024-07-07 22:31:08.000', '2024-07-07 22:31:11.000', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', '花生', '哈哈哈哈哈哈', 'https://thirdwx.qlogo.cn/mmopen/vi_32/icZJCCzianSibe7kHIwCSCXbjbbnpHMKDaePZQlrhFTIOraziap9Pxp9vhWEJhvTZTAg4sbRbDvzVeKbZxzicJILMzqOs1XRrqpLH9HbvySjUIWo/132', '2323831454@qq.com', NULL, '江西省赣州市', NULL, '系统注册', 2);
INSERT INTO `user_models` VALUES (3, '2024-07-10 10:21:34.000', '2024-07-10 10:21:35.000', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', '张亨睿', '嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻', NULL, '1111111111@qq.com', NULL, '四川省', NULL, '系统注册', 1);
INSERT INTO `user_models` VALUES (4, '2024-07-10 10:22:56.000', '2024-07-10 10:22:58.000', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', '雷正清', '水水水水水水水水', NULL, '2222222222@qq.com', NULL, '江西省抚州市', NULL, '系统注册', 1);
INSERT INTO `user_models` VALUES (5, '2024-07-11 11:08:22.000', '2024-07-11 11:08:25.000', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', '小小森', '强强强强强强强强', 'https://thirdwx.qlogo.cn/mmopen/vi_32/icZJCCzianSibe7kHIwCSCXbjbbnpHMKDaePZQlrhFTIOraziap9Pxp9vhWEJhvTZTAg4sbRbDvzVeKbZxzicJILMzqOs1XRrqpLH9HbvySjUIWo/132', 'Hsen1015@gmail.com', NULL, '江西省赣州市', NULL, '系统注册', 1);
INSERT INTO `user_models` VALUES (6, '2024-07-11 11:09:16.000', '2024-07-11 11:09:18.000', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', '大大森', '哇哇哇哇哇哇哇哇哇哇', 'https://thirdwx.qlogo.cn/mmopen/vi_32/icZJCCzianSibe7kHIwCSCXbjbbnpHMKDaePZQlrhFTIOraziap9Pxp9vhWEJhvTZTAg4sbRbDvzVeKbZxzicJILMzqOs1XRrqpLH9HbvySjUIWo/132', 'Hsen1015@gmail.com', NULL, '江西省赣州市', NULL, '系统注册', 1);
INSERT INTO `user_models` VALUES (7, '2024-07-11 11:09:16.000', '2024-07-11 11:09:18.000', NULL, '$2a$04$8q7IjvD/wJUzxkVxMgNEPuGiJ/5JmafKci94MpQIBOJLbUCeUPPdS', '刘宏森', 'AAAAAAAAAAAAAAA', 'https://thirdwx.qlogo.cn/mmopen/vi_32/icZJCCzianSibe7kHIwCSCXbjbbnpHMKDaePZQlrhFTIOraziap9Pxp9vhWEJhvTZTAg4sbRbDvzVeKbZxzicJILMzqOs1XRrqpLH9HbvySjUIWo/132', 'Hsen1015@gmail.com', NULL, '江西省赣州市', NULL, '系统注册', 1);

-- ----------------------------
-- Table structure for user_role_models
-- ----------------------------
DROP TABLE IF EXISTS `user_role_models`;
CREATE TABLE `user_role_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `role_desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `permissions` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_role_models_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role_models
-- ----------------------------
INSERT INTO `user_role_models` VALUES (0, '2024-07-02 15:52:50.000', '2024-07-02 15:52:52.000', NULL, '普通用户', '普通用户', '1,2');
INSERT INTO `user_role_models` VALUES (1, '2024-07-02 15:51:53.000', '2024-07-02 15:51:56.000', NULL, '管理员', '管理员', '1,2,3,4');
INSERT INTO `user_role_models` VALUES (2, '2024-07-07 22:33:52.000', '2024-07-07 22:33:54.000', NULL, '卑微研究生', '卑微研究生', '1,2');
INSERT INTO `user_role_models` VALUES (3, '2024-07-07 22:34:19.000', '2024-07-07 22:34:21.000', NULL, '卑微打工人', '卑微打工人', '1,2,3,4');
INSERT INTO `user_role_models` VALUES (4, '2024-07-07 22:34:44.000', '2024-07-07 22:34:46.000', NULL, 'Go开发工程师', 'Goer', '1,2,3,4,5');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`  (
  `user_model_id` bigint UNSIGNED NOT NULL,
  `user_role_model_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`user_model_id`, `user_role_model_id`) USING BTREE,
  INDEX `fk_user_roles_user_role_model`(`user_role_model_id` ASC) USING BTREE,
  CONSTRAINT `fk_user_roles_user_model` FOREIGN KEY (`user_model_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_roles_user_role_model` FOREIGN KEY (`user_role_model_id`) REFERENCES `user_role_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES (1, 1);
INSERT INTO `user_roles` VALUES (2, 2);

-- ----------------------------
-- Table structure for users_groups
-- ----------------------------
DROP TABLE IF EXISTS `users_groups`;
CREATE TABLE `users_groups`  (
  `user_model_id` bigint UNSIGNED NOT NULL,
  `group_model_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`user_model_id`, `group_model_id`) USING BTREE,
  INDEX `fk_users_groups_group_model`(`group_model_id` ASC) USING BTREE,
  CONSTRAINT `fk_users_groups_group_model` FOREIGN KEY (`group_model_id`) REFERENCES `group_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_users_groups_user_model` FOREIGN KEY (`user_model_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users_groups
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
