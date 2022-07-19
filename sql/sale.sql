/*
 Navicat MySQL Data Transfer

 Source Server         : 52.11.26.186_3306
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 52.11.26.186:3306
 Source Schema         : sale

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 19/07/2022 10:49:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for spu_activity
-- ----------------------------
DROP TABLE IF EXISTS `spu_activity`;
CREATE TABLE `spu_activity` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `act_name` varchar(100) DEFAULT NULL,
  `spu_id` bigint DEFAULT NULL,
  `sale_price` decimal(10,2) DEFAULT NULL,
  `act_status` tinyint DEFAULT NULL,
  `prime` tinyint DEFAULT NULL,
  `special` tinyint DEFAULT NULL,
  `start_time` timestamp NULL DEFAULT NULL,
  `end_time` timestamp NULL DEFAULT NULL,
  `total_stock` int DEFAULT NULL,
  `avail_stock` int DEFAULT NULL,
  `lock_stock` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for spu_item
-- ----------------------------
DROP TABLE IF EXISTS `spu_item`;
CREATE TABLE `spu_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `spu_name` varchar(255) NOT NULL,
  `spu_price` decimal(10,2) NOT NULL,
  `spu_desc` varchar(255) DEFAULT NULL,
  `spu_image` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for spu_order
-- ----------------------------
DROP TABLE IF EXISTS `spu_order`;
CREATE TABLE `spu_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` bigint DEFAULT NULL,
  `activity_id` bigint DEFAULT NULL,
  `state` tinyint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=162 DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;
