CREATE TABLE user_resource (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    system_id INT UNSIGNED NOT NULL COMMENT '[fk]user_system.id，用户id',
    resource_id INT UNSIGNED NOT NULL COMMENT '[fk]product.id，资源id',
    amount BIGINT UNSIGNED DEFAULT 0 NOT NULL COMMENT '数量',
    update_time DATETIME DEFAULT NOW() NOT NULL COMMENT '更新时间',
    create_time DATETIME DEFAULT NOW() NOT NULL COMMENT '创建时间'
);

ALTER TABLE user_resource ADD INDEX(system_id, resource_id);
