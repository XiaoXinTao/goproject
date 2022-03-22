CREATE TABLE `item_groups`
(
    `unique_id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `item_groupid` bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品的group',
    `passport_uid` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户账号',
    `create_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
    `update_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_time',
    PRIMARY KEY (`unique_id`),
    UNIQUE KEY `uniq_passport_item_group` (`passport_uid`, `item_groupid`),
    KEY            `idx_passport_item_group` (`passport_uid`, `item_groupid`)
) ENGINE=InnoDB AUTO_INCREMENT=244 DEFAULT CHARSET=utf8mb4 COMMENT='商家创建的店铺';