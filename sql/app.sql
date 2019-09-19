CREATE TABLE `app` (
    `id` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `app_id` VARCHAR(64) NOT NULL,
    `secret` VARCHAR(64) NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `name` VARCHAR(128) NOT NULL,
    `description` TEXT,
    `oidc_redirect_uri` VARCHAR(256),
    `update_at` DATETIME NOT NULL DEFAULT NOW(),
    `create_at` DATETIME NOT NULL DEFAULT NOW()
) DEFAULT CHARSET=utf8mb4;
ALTER TABLE `app` ADD UNIQUE(`app_id`);


