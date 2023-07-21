CREATE TABLE `grants` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `uid` varchar(100) NOT NULL,
    `role_id` varchar(100) NOT NULL,
    `resources_id` varchar(100) NOT NULL,
    `action_name` varchar(255) NOT NULL,
    `is_deleted` boolean NOT NULL DEFAULT 0,
    `created_by` varchar(100) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT now(),
    `updated_by` varchar(100) NOT NULL,
    `updated_at` datetime NOT NULL DEFAULT now(),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`, `email`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;