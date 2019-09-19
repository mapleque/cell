CREATE TABLE `user_captcha` (
    `id` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(320) NOT NULL,
    `captcha` VARCHAR(64) NOT NULL,
    `expired_at` DATETIME NOT NULL,
    `create_at` DATETIME NOT NULL DEFAULT NOW()
) DEFAULT CHARSET=utf8;
ALTER TABLE `user_captcha` ADD UNIQUE(`username`);
