-- +goose Up
-- +goose StatementBegin

CREATE TABLE `web_project_picture`
(
    `id`                int unsigned NOT NULL auto_increment,
    `classification_id` varchar(64)  not null comment '分类ID',
    `project_id`        varchar(64)  not null comment '项目ID',
    `url`               varchar(255) NOT NULL DEFAULT '' COMMENT '链接',
    `created_at`        timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`        timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`        timestamp             default null,
    PRIMARY KEY (`id`),
    key (`classification_id`),
    key (`project_id`)
) DEFAULT COLLATE = utf8mb4_unicode_ci COMMENT ='官网项目图片表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists `web_project_picture`;

-- +goose StatementEnd
