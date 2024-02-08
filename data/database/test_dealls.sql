-- -------------------------------------------------------------
-- TablePlus 5.6.2(516)
--
-- https://tableplus.com/
--
-- Database: test_dealls
-- Generation Time: 2024-02-08 17:16:46.6050
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `swipes`;
CREATE TABLE `swipes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `swiping_user_id` bigint unsigned DEFAULT NULL,
  `swiped_user_id` bigint unsigned DEFAULT NULL,
  `direction` varchar(255) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_swipes_swiping_user_id` (`swiping_user_id`),
  KEY `idx_swipes_swiped_user_id` (`swiped_user_id`),
  KEY `idx_swipes_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_swipes_swiped_user` FOREIGN KEY (`swiped_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_swipes_swiping_user` FOREIGN KEY (`swiping_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_users_swiped` FOREIGN KEY (`swiped_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_users_swiped_user` FOREIGN KEY (`swiped_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_users_swiping` FOREIGN KEY (`swiping_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_users_swiping_user` FOREIGN KEY (`swiping_user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `premium` tinyint(1) DEFAULT '0',
  `last_swipe` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;;

INSERT INTO `swipes` (`id`, `swiping_user_id`, `swiped_user_id`, `direction`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 2, 'left', '2024-02-07 05:38:07.000', '2024-02-07 05:38:07.000', NULL),
(2, 1, 3, 'right', '2024-02-07 05:38:07.000', '2024-02-07 05:38:07.000', NULL),
(3, 1, 4, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(4, 1, 5, 'left', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(5, 4, 1, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(6, 4, 5, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(7, 4, 2, 'left', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(8, 1, 6, 'left', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(9, 1, 7, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(10, 1, 8, 'left', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(11, 1, 9, 'left', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(12, 1, 10, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(13, 1, 11, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(14, 1, 12, 'left', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(15, 1, 13, 'right', '2024-02-08 05:38:07.000', '2024-02-08 05:38:07.000', NULL),
(17, 1, 28, 'right', '2024-02-06 15:58:02.390', '2024-02-08 15:58:02.390', NULL),
(21, 1, 28, 'right', '2024-02-07 16:10:23.232', '2024-02-08 16:10:23.232', NULL),
(22, 1, 28, 'right', '2024-02-05 16:19:14.274', '2024-02-08 16:19:14.274', NULL),
(28, 1, 28, 'right', '2024-02-08 16:47:31.860', '2024-02-08 16:47:31.860', NULL),
(29, 1, 29, 'right', '2024-02-08 16:56:21.823', '2024-02-08 16:56:21.823', NULL),
(30, 1, 30, 'right', '2024-02-08 17:08:20.513', '2024-02-08 17:08:20.513', NULL),
(31, 1, 27, 'right', '2024-02-08 17:13:30.312', '2024-02-08 17:13:30.312', NULL);

INSERT INTO `users` (`id`, `name`, `email`, `password`, `premium`, `last_swipe`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Zulkifli Raihan', 'zuran2907@gmail.com', '$2a$10$wy59G/8mOkyLqw7YAjzZAuahjpRGZAJ4WEv8kJl8ncZfhAMSWqRyO', 1, '2024-02-08 17:13:30.323', '2024-02-08 16:44:00.000', '2024-02-08 17:13:30.323', NULL),
(2, 'Zulkifli Raihan', 'zuran290701@gmail.com', '$2a$10$5on5COX2ri0st67KYe3Thex6J3D68HMmn17koc1eQEbfZh6TUyNPG', 0, NULL, '2024-02-08 05:36:09.941', '2024-02-08 05:36:09.941', NULL),
(3, 'Zulkifli Raihan', 'zuran290702@gmail.com', '$2a$10$OFjpOSJEsWC/HFKTWInCFuybmWZKiWYIUdmI0B8pnE9Xb4IBA6ba6', 0, NULL, '2024-02-08 05:36:11.344', '2024-02-08 05:36:11.344', NULL),
(4, 'Zulkifli Raihan', 'zuran290703@gmail.com', '$2a$10$njqTIw2rljB7dYav2YUfculHfHPM1WHDiwY6bs.1xJIQ0Q/sqjERW', 0, NULL, '2024-02-08 05:36:12.863', '2024-02-08 05:36:12.863', NULL),
(5, 'Zulkifli Raihan', 'zuran290704@gmail.com', '$2a$10$Hca3fn/CCB.t4NcZhkbCTex6/GC7xhe..Xz5m/3jZv0YBFvYbU/Zy', 0, NULL, '2024-02-08 05:36:14.415', '2024-02-08 05:36:14.415', NULL),
(6, 'Zulkifli Raihan', 'zuran290705@gmail.com', '$2a$10$aI3wqvoyPKCeHDugRh.TROh06DxinbviGLnHrkbnofDr8sRGvRini', 0, NULL, '2024-02-08 05:36:16.125', '2024-02-08 05:36:16.125', NULL),
(7, 'Zulkifli Raihan', 'zuran290706@gmail.com', '$2a$10$C07fWO10SjmGOBXoahDv.O2Pne50IdPjiDwHIcCIr5vOz/zU6RQPe', 0, NULL, '2024-02-08 05:36:17.738', '2024-02-08 05:36:17.738', NULL),
(8, 'Zulkifli Raihan', 'zuran290707@gmail.com', '$2a$10$O27zwVLY6djogI6LRUUTgOyLpcbUsfsMBOEifcCQVH92Pbcb8Oi2e', 0, NULL, '2024-02-08 05:36:19.623', '2024-02-08 05:36:19.623', NULL),
(9, 'Zulkifli Raihan', 'zuran290708@gmail.com', '$2a$10$dZwJW9w6Fz9CoPyLrYTNGOo7TMNVPt2SCNA93XUtuXHsUyZOvKIIS', 0, NULL, '2024-02-08 05:36:21.755', '2024-02-08 05:36:21.755', NULL),
(10, 'Zulkifli Raihan', 'zuran290709@gmail.com', '$2a$10$2/wJBhC1Z5sSHOkE5UB1GeJ3MtteshO0jLDQNI81I18486H.QOZXy', 0, NULL, '2024-02-08 05:36:23.528', '2024-02-08 05:36:23.528', NULL),
(11, 'Zulkifli Raihan', 'zuran290710@gmail.com', '$2a$10$Q5GSoNPn/6JHD6.h.ObQsOxviODc9Vod17C8eTPWiGCmXMVSCrhhS', 0, NULL, '2024-02-08 05:36:25.309', '2024-02-08 05:36:25.309', NULL),
(12, 'Zulkifli Raihan', 'zuran290712@gmail.com', '$2a$10$7c5rH8kJbC0uO9JcRgWD2.f6Tihg0JbTT3B1nbtNqOgqLJ0zTaSI6', 0, NULL, '2024-02-08 14:37:54.852', '2024-02-08 14:37:54.852', NULL),
(13, 'Zulkifli Raihan', 'zuran290713@gmail.com', '$2a$10$krL0wWff04Az9HKd/vjRsuiuFHDq2/RVQJAIBYoyxZbZbDTiRMweK', 0, NULL, '2024-02-08 14:37:56.354', '2024-02-08 14:37:56.354', NULL),
(14, 'Zulkifli Raihan', 'zuran290714@gmail.com', '$2a$10$SCBC2Z0TRINTC4ZvpNpAEeBXlY.gGK3xKf4OkMBuRfQb0LW2CoHkq', 0, NULL, '2024-02-08 14:37:57.637', '2024-02-08 14:37:57.637', NULL),
(15, 'Zulkifli Raihan', 'zuran290715@gmail.com', '$2a$10$yYw2FPZuYALMsjAT26Npd.XLiKebn27qaFgR4/8i.pbVH7cyx2oIW', 0, NULL, '2024-02-08 14:37:58.841', '2024-02-08 14:37:58.841', NULL),
(16, 'Zulkifli Raihan', 'zuran290716@gmail.com', '$2a$10$TA0YPcHO7YWlgOMe9wr0bOOAW7yK9phYBep2jUK282iw0omhpY516', 0, NULL, '2024-02-08 14:38:00.893', '2024-02-08 14:38:00.893', NULL),
(17, 'Zulkifli Raihan', 'zuran290717@gmail.com', '$2a$10$1qlBELHCGhzOhGrrNBPTiOXwsv1WX7aOCjlnXhLPieXgB.Q3fZpyu', 0, NULL, '2024-02-08 14:38:02.684', '2024-02-08 14:38:02.684', NULL),
(18, 'Zulkifli Raihan', 'zuran290718@gmail.com', '$2a$10$z3fb.4NqIZtcZ2qdG9T7/OiA5VuCQrDAq//Df7uygQZzZaK6Rl4lK', 0, NULL, '2024-02-08 14:38:04.448', '2024-02-08 14:38:04.448', NULL),
(19, 'Zulkifli Raihan', 'zuran290719@gmail.com', '$2a$10$IElwUGa/4xcvDucZjNRHR.5BpIEXmDK45GgaHRWD6y4BIkZvLuzGu', 0, NULL, '2024-02-08 14:38:06.343', '2024-02-08 14:38:06.343', NULL),
(20, 'Zulkifli Raihan', 'zuran290720@gmail.com', '$2a$10$BCX4ImEYq0Ey25EeFketA.X4xo/pp3An9FTGMpRvBjN5PaZjnWUKK', 0, NULL, '2024-02-08 14:38:07.931', '2024-02-08 14:38:07.931', NULL),
(21, 'Zulkifli Raihan', 'zuran290721@gmail.com', '$2a$10$4kbTexQaOMegSOb0L3uKZ.RteGorKqR7RdjmfrTYtpt6gUWZg5OCO', 0, NULL, '2024-02-08 14:38:09.541', '2024-02-08 14:38:09.541', NULL),
(22, 'Zulkifli Raihan', 'zuran290722@gmail.com', '$2a$10$1liqroPniWu7N9N.WWJ5Qe4bmcTXonKc6noQwOlXoCRwJyyhzZwGa', 0, NULL, '2024-02-08 14:38:10.817', '2024-02-08 14:38:10.817', NULL),
(23, 'Zulkifli Raihan', 'zuran290723@gmail.com', '$2a$10$zfe6qtr0fF0iP70XrO4ZtecqXz9zb6b5kPF3zEsLqLAjBXVVrdO0G', 0, NULL, '2024-02-08 14:38:12.241', '2024-02-08 14:38:12.241', NULL),
(24, 'Zulkifli Raihan', 'zuran290724@gmail.com', '$2a$10$3Q6d41..q9s.klZndAL57OoKThhPjLLHJMfaB6mEuO50Hxa8FFYRe', 0, NULL, '2024-02-08 14:38:14.610', '2024-02-08 14:38:14.610', NULL),
(25, 'Zulkifli Raihan', 'zuran290725@gmail.com', '$2a$10$S/DtK27x/sXGFypvVf2zdefbp2FVui7392G3qZ387gfluP.2dcD6a', 0, NULL, '2024-02-08 14:38:16.094', '2024-02-08 14:38:16.094', NULL),
(26, 'Zulkifli Raihan', 'zuran290726@gmail.com', '$2a$10$gf/Nh/vQtkiLY95eWmwnveBwOqDiGW8frOB/vpZCTGjf0/YPth3zC', 0, NULL, '2024-02-08 14:38:18.185', '2024-02-08 14:38:18.185', NULL),
(27, 'Zulkifli Raihan', 'zuran290727@gmail.com', '$2a$10$d2zBAxx7FHqyOgcERPQUkOgyPAJNvnRjirpccHVVXBsA2MIA0JM4C', 0, NULL, '2024-02-08 14:38:20.572', '2024-02-08 14:38:20.572', NULL),
(28, 'Zulkifli Raihan', 'zuran290728@gmail.com', '$2a$10$6r4pNzDHDd2oTI9NrXvjsuAwITxM/1euh6fwDCWxCoaT7Lrpr5bau', 0, NULL, '2024-02-08 14:38:22.356', '2024-02-08 14:38:22.356', NULL),
(29, 'Zulkifli Raihan', 'zuran290729@gmail.com', '$2a$10$HOEPDKvb8hOvPPc.I046v.0w1P2Yjuipd6ccT77xmVKEr3/CWonum', 0, NULL, '2024-02-08 14:38:23.971', '2024-02-08 14:38:23.971', NULL),
(30, 'Zulkifli Raihan', 'zuran290730@gmail.com', '$2a$10$1ABuxc/hEBk5PqZiSH/MzO5ezlpxefb8sAXgrWpFmPca4H6A6eaAe', 0, NULL, '2024-02-08 14:38:30.073', '2024-02-08 14:38:30.073', NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;