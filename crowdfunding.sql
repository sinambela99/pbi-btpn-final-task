-- phpMyAdmin SQL Dump
-- version 5.1.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 22, 2023 at 02:40 PM
-- Server version: 5.7.33
-- PHP Version: 8.2.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crowdfunding`
--

-- --------------------------------------------------------

--
-- Table structure for table `photos`
--

CREATE TABLE `photos` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `caption` text,
  `photo_url` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `photos`
--

INSERT INTO `photos` (`id`, `user_id`, `title`, `caption`, `photo_url`, `created_at`, `updated_at`) VALUES
(26, 21, 'koosng', 'kossong', '', '2023-09-21 03:03:30', '2023-09-21 03:03:30'),
(27, 21, 'koosng', 'kossong', 'images/%d-%s2179072313.jfif', '2023-09-21 03:04:01', '2023-09-21 03:04:01'),
(28, 21, 'koosng', 'kossong', 'images/%d-%s2179072313.jfif', '2023-09-21 03:04:17', '2023-09-21 03:04:17'),
(29, 21, 'koosng', 'kossong', 'images/%d-%s2179072313.jfif', '2023-09-21 03:04:27', '2023-09-21 03:04:27'),
(30, 21, 'koosng', 'kossong', 'images/%d-%s2120b96b02a65f21197a0ca29b28d401b9.jpg', '2023-09-21 03:04:43', '2023-09-21 03:04:43'),
(31, 24, 'update', 'update', 'images/%d-%s2420b96b02a65f21197a0ca29b28d401b9-removebg-preview.png', '2023-09-21 23:04:22', '2023-09-22 00:26:54'),
(32, 24, 'koosng', 'kossong', 'images/%d-%s2420b96b02a65f21197a0ca29b28d401b9.jpg', '2023-09-21 23:04:24', '2023-09-21 23:04:24'),
(33, 24, 'koosng', 'kossong', 'images/%d-%s2420b96b02a65f21197a0ca29b28d401b9.jpg', '2023-09-21 23:04:25', '2023-09-21 23:04:25'),
(34, 24, 'koosng', 'kossong', 'images/%d-%s2420b96b02a65f21197a0ca29b28d401b9.jpg', '2023-09-21 23:04:26', '2023-09-21 23:04:26'),
(35, 25, 'koosng', 'kossong', '', '2023-09-22 20:29:18', '2023-09-22 20:29:18');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password_hash`, `created_at`, `updated_at`) VALUES
(25, 'Yudistira update', 'test@gmail.com', '$2a$04$k5.1aQYM2ZY9dRAuwdxH8u7s64sNhTCZKGPc8IAoGkCR80ixgOG4G', '2023-09-22 20:28:59', '2023-09-22 20:28:59');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `photos`
--
ALTER TABLE `photos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `photos`
--
ALTER TABLE `photos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
