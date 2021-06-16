-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: mysql-server
-- Generation Time: Jun 16, 2021 at 01:49 PM
-- Server version: 8.0.19
-- PHP Version: 7.4.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `nozom_reports`
--

-- --------------------------------------------------------

--
-- Table structure for table `servers`
--

CREATE TABLE `servers` (
  `id` int UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `db_user` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `db_password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `db_ip` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `db_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `server_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `servers`
--

INSERT INTO `servers` (`id`, `created_at`, `updated_at`, `deleted_at`, `db_user`, `db_password`, `db_ip`, `db_name`, `server_name`) VALUES
(1, '2021-06-16 14:07:28', '2021-06-16 14:07:28', NULL, 'mcs', '123', '41.38.87.59', 'stock_main', 'halal'),
(2, '2021-06-16 14:07:28', '2021-06-16 14:07:28', NULL, 'web', '123', 'dental.myfirewall.co', 'STOCK2021', 'dental'),
(3, '2021-06-16 14:07:28', '2021-06-16 14:07:28', NULL, 'mcs', '123', '192.168.1.167', 'mcstree', 'local');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `server_id` int UNSIGNED DEFAULT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `email`, `server_id`, `password`) VALUES
(1, '2021-06-16 14:07:28', '2021-06-16 14:07:28', NULL, 'halal@elnozom.com', 1, '$2a$10$5BvuVFKOGOqOItuGgPCAkOoy1HWFsyhG0Fc2obboWa5uyx3chQ5xq'),
(2, '2021-06-16 14:07:28', '2021-06-16 14:07:28', NULL, 'dental@elnozom.com', 2, '$2a$10$5BvuVFKOGOqOItuGgPCAkOoy1HWFsyhG0Fc2obboWa5uyx3chQ5xq'),
(3, '2021-06-16 14:07:28', '2021-06-16 14:07:28', NULL, 'local@elnozom.com', 3, '$2a$10$5BvuVFKOGOqOItuGgPCAkOoy1HWFsyhG0Fc2obboWa5uyx3chQ5xq');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `servers`
--
ALTER TABLE `servers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `db_ip` (`db_ip`),
  ADD UNIQUE KEY `db_name` (`db_name`),
  ADD UNIQUE KEY `server_name` (`server_name`),
  ADD KEY `idx_servers_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uix_users_email` (`email`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`),
  ADD KEY `users_server_id_servers_id_foreign` (`server_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `servers`
--
ALTER TABLE `servers`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_server_id_servers_id_foreign` FOREIGN KEY (`server_id`) REFERENCES `servers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
