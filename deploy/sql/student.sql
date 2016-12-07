CREATE TABLE student (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    system_id INT UNSIGNED NOT NULL COMMENT '[fk]user_system.id，用户id',
    name VARCHAR(20) DEFAULT NULL COMMENT '姓名',
    mobile VARCHAR(11) DEFAULT NULL COMMENT '手机号',
    sex TINYINT(1) DEFAULT NULL COMMENT 'SEX_*，性别',
    head_img VARCHAR(200) DEFAULT NULL COMMENT '用户头像地址',
    update_time DATETIME DEFAULT NOW() NOT NULL COMMENT '更新时间',
    create_time DATETIME DEFAULT NOW() NOT NULL COMMENT '创建时间'
);

ALTER TABLE student ADD INDEX(system_id);
