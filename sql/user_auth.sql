CREATE TABLE `user_auth` (
    `id` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `app_id` VARCHAR(64) NOT NULL,
    `token` VARCHAR(64) DEFAULT NULL,
    `username` VARCHAR(320) NOT NULL,
    `ip` VARCHAR(64) NOT NULL,
    `expired_at` DATETIME NOT NULL,
    `last_login` DATETIME NOT NULL,
    `create_at` DATETIME NOT NULL DEFAULT NOW()
) DEFAULT CHARSET=utf8;
ALTER TABLE `user_auth` ADD UNIQUE(`username`, `app_id`);
ALTER TABLE `user_auth` ADD UNIQUE(`token`, `app_id`);
