-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 28, 2019 at 08:24 AM
-- Server version: 10.1.34-MariaDB
-- PHP Version: 7.2.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `class_attendance`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendance`
--

CREATE TABLE `attendance` (
  `attendance_id` varchar(24) NOT NULL,
  `class_id` varchar(24) NOT NULL,
  `student_id` varchar(24) DEFAULT NULL,
  `attend_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `attendance`
--

INSERT INTO `attendance` (`attendance_id`, `class_id`, `student_id`, `attend_time`) VALUES
('1a9415f0-510f-11e9-9624-', '5891693d-5107-11e9-9624-', '2fa5830d-5106-11e9-9624-', '2019-03-28 04:08:12'),
('4f7f4ff1-5128-11e9-9624-', '7a4b31ce-5127-11e9-9624-', '43bb63e2-5106-11e9-9624-', '2019-03-28 07:08:38'),
('7a4c9b36-5127-11e9-9624-', '7a4b31ce-5127-11e9-9624-', NULL, '2019-03-28 07:02:40'),
('cbaaece2-5107-11e9-9624-', '5891693d-5107-11e9-9624-', NULL, '2019-03-28 03:15:53'),
('d79f1ffc-5121-11e9-9624-', '5891693d-5107-11e9-9624-', '43bb63e2-5106-11e9-9624-', '2019-03-28 06:22:20');

-- --------------------------------------------------------

--
-- Table structure for table `class`
--

CREATE TABLE `class` (
  `id` varchar(24) NOT NULL,
  `class_name` varchar(30) NOT NULL,
  `class_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `room` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `class`
--

INSERT INTO `class` (`id`, `class_name`, `class_time`, `room`, `created_at`, `updated_at`) VALUES
('5891693d-5107-11e9-9624-', 'Software Testing', '2019-02-10 07:00:00', 'E4', '2019-03-28 03:12:40', '2019-03-28 03:12:40'),
('7a4b31ce-5127-11e9-9624-', 'Computer Network', '2019-02-09 05:00:00', 'E9', '2019-03-28 07:02:40', '2019-03-28 07:02:40'),
('d516c283-5106-11e9-9624-', 'Basic Programming', '2019-02-02 04:30:00', 'E2', '2019-03-28 03:08:59', '2019-03-28 03:08:59'),
('d516dea8-5106-11e9-9624-', 'Machine Learning', '2019-03-28 03:13:21', 'E6', '2019-03-28 03:08:59', '2019-03-28 03:08:59');

-- --------------------------------------------------------

--
-- Table structure for table `student`
--

CREATE TABLE `student` (
  `id` varchar(24) NOT NULL,
  `name` varchar(30) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `student`
--

INSERT INTO `student` (`id`, `name`, `phone`, `created_at`, `updated_at`) VALUES
('2fa5830d-5106-11e9-9624-', 'James', '089475827364', '2019-03-28 03:04:22', '2019-03-26 17:18:21'),
('3774a81a-5106-11e9-9624-', 'Laza', '084756321739', '2019-03-28 03:04:35', '2019-03-26 17:19:07'),
('3ddcc619-5106-11e9-9624-', 'Kram', '082367894630', '2019-03-28 03:04:46', '2019-03-26 17:19:48'),
('43bb63e2-5106-11e9-9624-', 'Keroro', '084763298323', '2019-03-28 03:04:55', '2019-03-26 17:21:21'),
('48ea161e-5106-11e9-9624-', 'Jambu', '083673849000', '2019-03-28 03:05:04', '2019-03-26 17:22:07');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`attendance_id`),
  ADD KEY `fk_class_id` (`class_id`),
  ADD KEY `fk_student_id` (`student_id`);

--
-- Indexes for table `class`
--
ALTER TABLE `class`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `attendance`
--
ALTER TABLE `attendance`
  ADD CONSTRAINT `fk_class_id` FOREIGN KEY (`class_id`) REFERENCES `class` (`id`),
  ADD CONSTRAINT `fk_student_id` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
