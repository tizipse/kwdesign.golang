-- +goose Up
-- +goose StatementBegin

CREATE TABLE `web_contact`
(
    `id`         int unsigned     NOT NULL AUTO_INCREMENT,
    `city`       varchar(32)      NOT NULL DEFAULT '' COMMENT '城市',
    `address`    varchar(255)     NOT NULL DEFAULT '' COMMENT '地址',
    `telephone`  varchar(32)      NOT NULL DEFAULT '' COMMENT '电话',
    `order`      tinyint unsigned not null default 0 comment '序号',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    PRIMARY KEY (`id`)
) AUTO_INCREMENT = 1000
  DEFAULT COLLATE = utf8mb4_unicode_ci COMMENT ='官网联系表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists web_contact;

-- +goose StatementEnd
