-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `user_id` int(0) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                          `wechat_avatar` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信头像',
                          `wechat_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信名称',
                          `type` int(0) NULL DEFAULT NULL COMMENT '微信头像类型(1微信头像地址2上传头像地址)',
                          `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'openid(全局唯一标识)',
                          `unionid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'unionid',
                          `is_admin` tinyint(1) NULL DEFAULT NULL COMMENT '是否管理员',
                          `is_enabled` int(0) NULL DEFAULT NULL COMMENT '是否启用',
                          `create_at_utc` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                          PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000095 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

INSERT INTO `users` VALUES (1000001, '165ce2b9-cb69-4f33-8caa-1d72f9991d1', '张峰', 2, '11111', '', 1, 1, '2022-12-28 18:30:59');