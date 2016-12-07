CREATE TABLE user_feature (
    id INT UNSIGNED AUTO_INCREMET PRIMARY KEY NOT NULL,
    system_id INT UNSIGNED NOT NULL COMMENT '[fk]user_system.id，用户id',
    feature_id INT UNSIGNED NOT NULL COMMENT '[fk]product.id，特权id',
    start_time DATETIME DEFAULT NULL COMMENT '开始时间',
    end_time DATETIME DEFAULT NULL COMMENT '结束时间',
    update_time DATETIME DEFAULT NOW() NOT NULL COMMENT '更新时间',
    create_time DATETIME DEFAULT NOW() NOT NULL COMMENT '创建时间'
);

ALTER TABLE user_feature ADD INDEX(system_id, feature_id);
