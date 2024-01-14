DROP DATABASE IF EXISTS todo_db;
CREATE DATABASE todo_db;
USE todo_db;

CREATE TABLE IF NOT EXISTS `user_t` (
    `user_id` varchar(128) NOT NULL,
    `user_name` varchar(128) NOT NULL,
    `password` varchar(128) NOT NULL,
    `deleted_time` datetime,
    `register_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSEt=utf8mb4;

CREATE TABLE IF NOT EXISTS `todo_t` (
    `todo_id` varchar(128) NOT NULL,
    `user_id` varchar(128) NOT NULL,
    `todo` TEXT NOT NULL,
    `deleted_time` datetime,
    `register_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`todo_id`, `user_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user_t`(`user_id`)
) ENGINE=InnoDB DEFAULT CHARSEt=utf8mb4;