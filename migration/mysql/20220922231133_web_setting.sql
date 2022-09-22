-- +goose Up
-- +goose StatementBegin

CREATE TABLE `web_setting`
(
    `id`         int unsigned     NOT NULL AUTO_INCREMENT,
    `type`       varchar(10)      NOT NULL DEFAULT '' COMMENT '字段类型：input/textarea/enable/url/email',
    `label`      varchar(10)      NOT NULL DEFAULT '' COMMENT '名称',
    `key`        varchar(20)      NOT NULL DEFAULT '' COMMENT '键',
    `val`        text                      default null COMMENT '值',
    `required`   tinyint unsigned NOT NULL DEFAULT 0 COMMENT '必填：1=是；2=否',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    PRIMARY KEY (`id`)
) AUTO_INCREMENT = 1000
  DEFAULT COLLATE = utf8mb4_unicode_ci COMMENT ='官网设置表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists web_setting;

-- +goose StatementEnd
