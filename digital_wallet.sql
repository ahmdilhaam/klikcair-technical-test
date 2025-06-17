# Host: localhost  (Version 5.5.5-10.4.24-MariaDB)
# Date: 2025-06-18 00:33:27
# Generator: MySQL-Front 6.0  (Build 2.20)


#
# Structure for table "transcations"
#

DROP TABLE IF EXISTS `transcations`;
CREATE TABLE `transcations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` enum('withdrawal','deposit') DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "transcations"
#

INSERT INTO `transcations` VALUES (1,'withdrawal',10000,1,'2025-06-17 21:51:36',NULL),(2,'withdrawal',150000,1,'2025-06-17 21:54:16',NULL);

#
# Structure for table "users"
#

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) DEFAULT NULL,
  `phone_number` varchar(30) NOT NULL DEFAULT '',
  `balance` double NOT NULL DEFAULT 0,
  `email` varchar(50) NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

#
# Data for table "users"
#

INSERT INTO `users` VALUES (1,'ilham','6282123456789',1040000,'ilham_postman@yopmail.com','2025-06-17 18:06:00','2025-06-17 21:54:16');
