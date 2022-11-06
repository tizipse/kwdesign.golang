-- +goose Up
-- +goose StatementBegin

CREATE TABLE `web_classification`
(
    `id`          varchar(64)  NOT NULL,
    `name`        varchar(32)  NOT NULL DEFAULT '' COMMENT '名称',
    `alias`       varchar(32)  NOT NULL DEFAULT '' COMMENT '别名',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT 'SEO 标题',
    `keyword`     varchar(255) NOT NULL DEFAULT '' COMMENT 'SEO 关键词',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT 'SEO 描述',
    `order`       tinyint unsigned not null default 0 comment '序号',
    `is_enable`   tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `created_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp             default null,
    PRIMARY KEY (`id`)
) DEFAULT COLLATE = utf8mb4_unicode_ci COMMENT ='官网分类表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists `web_classification`;

-- +goose StatementEnd
