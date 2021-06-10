-- 富媒体彩信发送日志
CREATE table fmt_sms_send_log (
    log_id serial PRIMARY KEY,
    task_id VARCHAR(100) not null DEFAULT '',
    template_id VARCHAR(100) not null DEFAULT '',
    phone VARCHAR(20) not null DEFAULT '',
    status SMALLINT not null DEFAULT 0,
    data  JSON not null default '{}',
    create_yw INTEGER not null DEFAULT 0,
    send_time VARCHAR(100) not null DEFAULT '',
    create_time INTEGER NOT NULL DEFAULT extract(EPOCH FROM now()), -- 创建时间
    create_date date NOT NULL DEFAULT CURRENT_DATE             -- 创建日期
);