CREATE DATABASE  `grpc-myboot` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
create user 'myboot'@'%' identified by '123456';
grant all privileges on `grpc-myboot`.*  to 'myboot'@'%';
flush privileges;

use grpc-myboot;

CREATE TABLE `sys_users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` datetime                                DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime                                DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime                                DEFAULT NULL,
    `username`   varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
    `password`   varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
    `mobile`     varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
    `nickname`   varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户昵称',
    PRIMARY KEY (`id`),
    KEY          `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `sys_users`
VALUES (1, '2021-10-23 12:00:44', '2021-10-23 12:00:48', NULL, 'admin', NULL, '13888888888', 'admin');
INSERT INTO `sys_users`
VALUES (2, '2021-10-23 12:00:44', '2021-10-23 12:00:44', NULL, 'admin2', NULL, '13888888889', 'admin2');
