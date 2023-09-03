
CREATE TABLE `courses` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `uid` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
    `subject` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
    `created_at` TIMESTAMP NULL DEFAULT current_timestamp(),
    `updated_at` TIMESTAMP NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`)
) COLLATE='utf8_general_ci' ENGINE=InnoDB;