-- +goose Up
-- +goose StatementBegin

create table web_banner
(
    `id`         int unsigned     not null auto_increment,
    `layout`     varchar(10)      not null default '' comment '布局：mobile=移动端；ipad=iPad；pc=电脑',
    `picture`    varchar(255)     not null default '' comment '图片链接',
    `name`       varchar(32)      not null default '' comment '名称',
    `target`     varchar(5)       not null default '' comment '链接打开方式：blank=新窗口；self=本窗口',
    `url`        varchar(255)     not null default '' comment '链接',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `order`      tinyint unsigned not null default 0 comment '序号',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment '官网轮播表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists web_banner;

-- +goose StatementEnd
