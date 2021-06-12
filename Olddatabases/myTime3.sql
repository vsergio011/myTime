-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: May 20, 2021 at 08:29 PM
-- Server version: 5.7.33
-- PHP Version: 7.4.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `myTime`
--

-- --------------------------------------------------------

--
-- Table structure for table `Company`
--

CREATE TABLE `Company` (
  `owner` varchar(50) NOT NULL,
  `id` int(20) NOT NULL,
  `name` varchar(50) NOT NULL,
  `address` varchar(80) NOT NULL,
  `country` varchar(3) NOT NULL,
  `phone` varchar(14) NOT NULL,
  `cif` varchar(40) NOT NULL,
  `region` varchar(40) NOT NULL,
  `type` int(11) DEFAULT NULL,
  `subtype` int(11) DEFAULT NULL,
  `admin` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `Company`
--

INSERT INTO `Company` (`owner`, `id`, `name`, `address`, `country`, `phone`, `cif`, `region`, `type`, `subtype`, `admin`) VALUES
('UIDOO', 1, 'penguins', 'penguin street', 'esp', '685658', '23234324', 'eu', 23, 22, 0);

-- --------------------------------------------------------

--
-- Table structure for table `place`
--

CREATE TABLE `place` (
  `id` int(11) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `company` int(11) NOT NULL,
  `address` varchar(80) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `place`
--

INSERT INTO `place` (`id`, `phone`, `company`, `address`) VALUES
(1, '685658', 1, 'penguins');

-- --------------------------------------------------------

--
-- Table structure for table `task`
--

CREATE TABLE `task` (
  `id` int(20) NOT NULL,
  `title` varchar(40) NOT NULL,
  `date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `uid_user` varchar(50) NOT NULL,
  `place` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `task`
--

INSERT INTO `task` (`id`, `title`, `date`, `uid_user`, `place`) VALUES
(1, 'Example task', '2021-02-28 15:00:17', 'nnaGlkOw9uPJVrqf1pwppxOSXOc2', 1),
(2, 'Example task2 update', '2021-02-01 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(3, 'Example task3', '2021-02-28 15:00:17', 'nnaGlkOw9uPJVrqf1pwppxOSXOc2', 1),
(4, 'Example task4', '2021-02-28 15:00:17', 'nnaGlkOw9uPJVrqf1pwppxOSXOc2', 1),
(5, 'Example task2', '2021-02-28 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(6, 'Example task6', '2021-02-28 15:00:17', 'nnaGlkOw9uPJVrqf1pwppxOSXOc2', 1),
(7, 'Example task2 update', '2021-02-28 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(8, 'Example taskpostman', '2021-05-20 18:31:04', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(9, 'Example taddskpostman', '2021-05-20 18:31:43', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(10, 'Example taddskpostman', '2021-05-20 18:31:56', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(11, 'Example taddskpostman', '2021-05-20 18:39:36', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(12, 'Example new', '2021-05-20 18:31:56', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(13, 'Example new', '2025-05-20 18:31:56', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(14, 'Example new', '2029-05-20 18:31:56', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(15, 'Example new', '2029-05-20 18:31:56', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(16, 'Example new', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(17, 'Example new', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(18, 'Example new', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(19, 'Example new', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(20, 'pruebaapp', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(21, 'pruebaapp', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(22, 'pruebaapp', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(23, 'Example task2sdd', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(24, 'Example task5', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(25, 'Example task5', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(26, 'Example task5', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(27, 'Example task5', '2021-02-28 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(28, 'Example task223234234', '2021-02-26 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(29, 'Example new', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1),
(30, 'Example new', '2029-05-20 00:00:00', 'M1VXKsoMHqO9vMFZaREVSVKJxQX2', 1);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `uid` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `surname` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`uid`, `name`, `surname`, `email`) VALUES
('frKvD1aaITa9aZkuod4SDcjsjQj2', 'displayname2', 'ss', 'miemailnuevo@vera.com'),
('M1VXKsoMHqO9vMFZaREVSVKJxQX2', 'sergio', 'surnanme', 'sergiov0011@gmail.com'),
('nnaGlkOw9uPJVrqf1pwppxOSXOc2', 'miname', 'ss', 'miemail@gmal.com'),
('UIDOO', 'bruce', 'surnameprueba', 'prueba@gmail');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Company`
--
ALTER TABLE `Company`
  ADD PRIMARY KEY (`id`),
  ADD KEY `company_owner` (`owner`);

--
-- Indexes for table `place`
--
ALTER TABLE `place`
  ADD PRIMARY KEY (`id`),
  ADD KEY `place_company` (`company`);

--
-- Indexes for table `task`
--
ALTER TABLE `task`
  ADD PRIMARY KEY (`id`),
  ADD KEY `User_task` (`uid_user`),
  ADD KEY `task_place` (`place`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`uid`),
  ADD UNIQUE KEY `uid_prim` (`uid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Company`
--
ALTER TABLE `Company`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `place`
--
ALTER TABLE `place`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `task`
--
ALTER TABLE `task`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Company`
--
ALTER TABLE `Company`
  ADD CONSTRAINT `company_owner` FOREIGN KEY (`owner`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `place`
--
ALTER TABLE `place`
  ADD CONSTRAINT `place_company` FOREIGN KEY (`company`) REFERENCES `Company` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION;

--
-- Constraints for table `task`
--
ALTER TABLE `task`
  ADD CONSTRAINT `User_task` FOREIGN KEY (`uid_user`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE NO ACTION,
  ADD CONSTRAINT `task_place` FOREIGN KEY (`place`) REFERENCES `place` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
