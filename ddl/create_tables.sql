DROP DATABASE messagedb;
CREATE DATABASE messagedb DEFAULT CHAR SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE messagedb;

CREATE TABLE `message` (
    `id` VARCHAR(36) NOT NULL,
    `title` JSON NOT NULL,
    `text` JSON NOT NULL,
    PRIMARY KEY (`id`)
)
