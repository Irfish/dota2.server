DROP DATABASE IF EXISTS dota2_db;
CREATE DATABASE dota2_db;
USE dota2_db;

/*用户表*/
create table `user`(
  `id` bigint(20) AUTO_INCREMENT not NULL,
  `steam_id` varchar(200) NOT NULL,
  `steam_name` varchar(200) NOT NULL,
  `steam_gold` bigint(10) NOT NULL,
  `steam_silver` bigint(10) NOT NULL,
  `steam_vip_exp` bigint(10) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*卡密表*/
create table `card_key`(
  `id` bigint(20) AUTO_INCREMENT not NULL,
  `key_code` varchar(200) NOT NULL,
  `key_state` int NOT NULL,
  `key_type` bigint(10) NOT NULL,
  `key_cost` bigint(10) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;








