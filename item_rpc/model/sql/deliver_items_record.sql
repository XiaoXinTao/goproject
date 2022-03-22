CREATE TABLE `deliever_items_record`
(
    `unique_id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `item_groupid`       bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品的group',
    `item_id`            bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品的id',
    `deliver_item_count` int       NOT NULL DEFAULT 0 COMMENT '购买的商品',
    `passport_uid`       bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户账号',
    `create_time`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
    `update_time`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_time',
    PRIMARY KEY (`unique_id`)
) ENGINE=InnoDB AUTO_INCREMENT=244 DEFAULT CHARSET=utf8mb4 COMMENT='购买记录表格';