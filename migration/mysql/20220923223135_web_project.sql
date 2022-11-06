-- +goose Up
-- +goose StatementBegin

CREATE TABLE `web_project`
(
    `id`                varchar(64)      NOT NULL,
    `classification_id` varchar(64)      not null comment '分类ID',
    `theme`             varchar(5)       NOT NULL DEFAULT '' COMMENT '主题：light=明亮；dark=黑暗',
    `name`              varchar(32)      NOT NULL DEFAULT '' COMMENT '名称',
    `address`           varchar(64)      NOT NULL DEFAULT '' COMMENT '地点',
    `height`            tinyint unsigned NOT NULL DEFAULT 0 COMMENT '高度（%）',
    `picture`           varchar(255)     NOT NULL DEFAULT '' COMMENT '图片',
    `title`             varchar(255)     NOT NULL DEFAULT '' COMMENT 'SEO 标题',
    `keyword`           varchar(255)     NOT NULL DEFAULT '' COMMENT 'SEO 关键词',
    `description`       varchar(255)     NOT NULL DEFAULT '' COMMENT 'SEO 描述',
    `html`              text                      default null comment '',
    `is_enable`         tinyint unsigned not null default 0 comment '是否启用：1=是；2=否；',
    `dated_at`          date                      default null,
    `created_at`        timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at`        timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at`        timestamp                 default null,
    PRIMARY KEY (`id`),
    key (`classification_id`)
) DEFAULT COLLATE = utf8mb4_unicode_ci COMMENT ='官网项目表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists `web_project`;

-- +goose StatementEnd
