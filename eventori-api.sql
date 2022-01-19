/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `model_catalogues` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `Model_name` varchar(255) DEFAULT NULL,
  `Image_url` text,
  `Description` text,
  `Created_at` timestamp NULL DEFAULT NULL,
  `Updated_at` timestamp NULL DEFAULT NULL,
  `Deleted_at` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `model_schedules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `Model_id` int DEFAULT NULL,
  `Schedule_date` timestamp NULL DEFAULT NULL,
  `Created_at` timestamp NULL DEFAULT NULL,
  `Updated_at` timestamp NULL DEFAULT NULL,
  `Deleted_at` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `model_catalogues` (`id`, `Model_name`, `Image_url`, `Description`, `Created_at`, `Updated_at`, `Deleted_at`) VALUES
(2, 'model 2', 'https://images.unsplash.com/photo-1453728013993-6d66e9c9123a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80', 'model 2 description', '2021-12-12 00:00:00', '2021-12-12 00:00:00', '2021-12-12 00:00:00');
INSERT INTO `model_catalogues` (`id`, `Model_name`, `Image_url`, `Description`, `Created_at`, `Updated_at`, `Deleted_at`) VALUES
(3, 'Model 5+6+7+1', 'a', 'Model 5+6+7+1', '2021-12-12 00:00:00', '2021-12-12 00:00:00', '2021-12-12 00:00:00');
INSERT INTO `model_catalogues` (`id`, `Model_name`, `Image_url`, `Description`, `Created_at`, `Updated_at`, `Deleted_at`) VALUES
(4, 'Model 4', 'https://images.unsplash.com/photo-1516035069371-29a1b244cc32?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=764&q=80', 'Description 4', NULL, NULL, NULL);
INSERT INTO `model_catalogues` (`id`, `Model_name`, `Image_url`, `Description`, `Created_at`, `Updated_at`, `Deleted_at`) VALUES
(5, 'Model 5+6+7', 'a', 'Model 5+6+7', NULL, NULL, NULL),
(6, 'Model 5 6', 'awdawd', 'Model 5 6', NULL, NULL, NULL);

INSERT INTO `model_schedules` (`id`, `Model_id`, `Schedule_date`, `Created_at`, `Updated_at`, `Deleted_at`) VALUES
(1, 1, '2021-12-12 00:00:00', '2021-12-12 00:00:00', '2021-12-12 00:00:00', '2021-12-12 00:00:00');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;