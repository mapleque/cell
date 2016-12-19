CREATE TABLE resource (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id INT UNSIGNED NOT NULL COMMENT '[fk]user.id，用户id',
    product_id INT UNSIGNED NOT NULL COMMENT '[fk]product.id，资源id',
    amount BIGINT UNSIGNED DEFAULT 0 NOT NULL COMMENT '数量',
    update_time DATETIME DEFAULT NOW() NOT NULL COMMENT '更新时间',
    create_time DATETIME DEFAULT NOW() NOT NULL COMMENT '创建时间'
);

ALTER TABLE resource ADD INDEX(user_id, product_id);
