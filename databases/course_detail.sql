
CREATE TABLE `course_detail` (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `uid` VARCHAR(100) NULL DEFAULT NULL,
     `courses_uid` VARCHAR(255) NULL DEFAULT NULL,
     `sub_courses` VARCHAR(255) NULL DEFAULT NULL,
     `created_at` TIMESTAMP NULL DEFAULT current_timestamp(),
     `updated_at` TIMESTAMP NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
     PRIMARY KEY (`id`),
     UNIQUE KEY (`uid`),
     FOREIGN KEY (`courses_uid`) REFERENCES courses(`uid`)
) COLLATE='utf8_general_ci' ENGINE=InnoDB;