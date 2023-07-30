CREATE TABLE `users` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `uid` varchar(100) NOT NULL,
    `first_name` varchar(255) NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `birthdate` date NOT NULL,
    `username` varchar(50) NOT NULL,
    `password` char(64) NOT NULL,
    `is_deleted` boolean NOT NULL DEFAULT 0,
    `created_by` varchar(100) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT now(),
    `updated_by` varchar(100) NOT NULL,
    `updated_at` datetime NOT NULL DEFAULT now(),
    FULLTEXT KEY (`first_name`, `last_name`),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`, `email`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;