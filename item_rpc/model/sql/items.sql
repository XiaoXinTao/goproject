CREATE TABLE `items`
(
    `unique_id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `item_groupid` bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品的group',
    `item_id`      bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品的id',
    `item_name`    varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
    `item_count`   int          NOT NULL DEFAULT 0 COMMENT '商品剩余量',
    `is_limited`   int          NOT NULL DEFAULT 0 COMMENT '商品是否限量',
    `passport_uid` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户账号',
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_time',
    PRIMARY KEY (`unique_id`),
    UNIQUE KEY `uniq_passport_item` (`item_groupid`, `item_id`),
    KEY            `idx_passport_item` (`passport_uid`, `item_groupid`, `item_id`),
    KEY            `idx_item_group` (`item_groupid`,`item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=244 DEFAULT CHARSET=utf8mb4 COMMENT='商品表格';