CREATE TABLE user_system (
    id INT UNSIGNED AUTO_INCREMET PRIMARY KEY NOT NULL,
    username VARCHAR(11) NOT NULL COMMENT '登录名',
    password VARCHAR(65) NOT NULL COMMENT '密码：user_token|rand_salt',
    inviter_id INT UNSIGNED DEFAULT NULL COMMENT '[fk]user_system.id，邀请人id',
    channel VARCHAR(11) DEFAULT NULL COMMENT '渠道',
    status CHAR(1) NOT NULL COMMENT '账号状态，USER_SYSTEM_STATUS_*',
    setting TEXT DEFAULT NULL COMMENT '系统设置，JSON',
    device_code VARCHAR(64) DEFAULT NULL COMMENT '设备码',
    version VARCHAR(11) DEFAULT NULL COMMENT '客户端版本',
    last_login_time DATETIME DEFAULT NOW() NOT NULL COMMENT '最后登录时间',
    update_time DATETIME DEFAULT NOW() NOT NULL COMMENT '更新时间',
    create_time DATETIME DEFAULT NOW() NOT NULL COMMENT '创建时间'
);

ALTER TABLE user_system ADD INDEX (username), ADD INDEX(status), ADD INDEX(inviter_id);
