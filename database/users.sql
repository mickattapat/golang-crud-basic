-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 17, 2022 at 09:57 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `atplabs`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `username` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `start_date` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `exp_date` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `status` int(11) NOT NULL,
  `role` varchar(100) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `username`, `password`, `name`, `last_name`, `start_date`, `exp_date`, `status`, `role`) VALUES
(23, 'mick', '$2a$10$mRcPIIL52bjj4sBgQU7otesFmMOoZGTBk4gUAJ/E/trXByFdXcOfe', 'attapattest', 'kayasatest', '2022-08-16 04:33:24', '2023-07-29 15:55:00', 1, 'admin'),
(51, 'attapatka01', '$2a$10$KeqWhjD3IkwntmWXDElW/uWmC8.DNPaOBnsZMmpzY7/EoUIG8kBxq', 'attapattest', 'kayasatest', '2022-08-16 05:08:12', '2023-07-29 15:55:00', 1, ''),
(54, 'user01', '$2a$10$QNly22g4P3DMDqQLEEWm7e6pBQFN6bM0.LjODZ5KySp9215D801Zi', 'user01', 'user01', '0000-00-00 00:00:00', '2022-08-10 17:00:00', 1, ''),
(55, 'user02', '$2a$10$eMkI.ajnnwKtgQlbJDnhtup8WmtUQ5t4PnOmP0ICnBkVhj.POSBfa', 'user02', 'user02', '0000-00-00 00:00:00', '2022-07-31 17:00:00', 1, ''),
(59, 'user05', '$2a$10$QcvGIgeEy9qVxTVQyyNBquVv2Bo/Idw2tQJ9Ne6Fm9CC40qzzehgO', 'user05', 'user05', '0000-00-00 00:00:00', '2022-08-15 17:00:00', 1, ''),
(60, 'user06', '$2a$10$2IRUwlU09YdU21.dpGi.T.4nhPLzPpkrr6lAD8EP5DkcU//Rbe1Nq', 'user06', 'user06', '0000-00-00 00:00:00', '2022-08-15 17:00:00', 1, ''),
(61, 'user07', '$2a$10$5IFmXqVEsHMIReck832bw.HHh.rPhHduWZ9.zwpb0gGmjr.tp/8MK', 'user07', 'user07', '0000-00-00 00:00:00', '2022-08-16 17:00:00', 1, ''),
(62, 'user08', '$2a$10$GpZGAGpI0he3DYbD81JeauEvN2WM/Tn9wgX9BvBCyZKiyhMj3IqWW', 'user08', 'user08', '2022-08-16 03:38:57', '2022-08-24 17:00:00', 1, ''),
(63, 'attapatka55', '$2a$10$b7J/8Zg6oj72ihXkVFiPqurHuhC5Pqls5mDhmARrj5n7N4dR6iSA2', 'attapatka55', 'attapatka55', '0000-00-00 00:00:00', '2022-08-23 17:00:00', 1, ''),
(64, 'attapatka66', '$2a$10$/GznUcmO6Cl8K6IF.V7DQeO6xP7fWUkXS82eIeaWzT6hdUSB5lMhW', 'attapatka66', 'attapatka66', '0000-00-00 00:00:00', '2022-08-15 17:00:00', 1, ''),
(65, 'attapatka77', '$2a$10$0LJu92gKPzIn5FZdoaJEVezT12OgXBKkvIbxHk.9EJhaCep32qi4y', 'attapatka77', 'attapatka77', '0000-00-00 00:00:00', '2022-08-19 17:00:00', 1, ''),
(66, 'mick90', '$2a$10$504h33q2C2/pLOgLsni8Xew6v04NLI0/5ZqTWnlJtZ6yAFokQhpYG', 'mick', 'attapat', '0000-00-00 00:00:00', '2023-10-29 15:55:00', 1, ''),
(67, 'mick101', '$2a$10$8jCvHnG4zc2tWNu2NWa7cegsA1NZt0nAJv.JKZuwzSZ0PgU0tx3NC', 'mick', 'attapat', '0000-00-00 00:00:00', '2023-10-29 15:55:00', 1, ''),
(68, 'mick01', '$2a$10$kVLVlzm1JbuvgJhTAhbd3.iBC8V9bTk9s34aU3PhOrqN5hme0PS1a', 'mick', 'attapat', '0000-00-00 00:00:00', '2023-10-29 15:55:00', 1, ''),
(69, 'mick01', '$2a$10$SC3WwTnxeetoUi6Fajz2s.AY3QJUwZSCXs11l01TFBE/QEZ6QvGSW', 'mick', 'attapat', '0000-00-00 00:00:00', '2023-10-29 15:55:00', 1, ''),
(70, 'mick01', '$2a$10$rt1/jRIkuuZ.Ib5osnaWj.ec41P/QhBy.Egq6mZbc2B7axsgEroSG', 'mick', 'attapat', '0000-00-00 00:00:00', '2023-10-29 15:55:00', 1, ''),
(71, 'mick01', '$2a$10$ysRV0MAgcq1KVqDiNrq9EelfICNRMJMZ82Sn9CMFnUzMe6v4ZLt2K', 'mick', 'attapat', '2022-08-16 10:19:57', '2023-10-29 15:55:00', 1, '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=72;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
