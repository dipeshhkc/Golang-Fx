CREATE TABLE IF NOT EXISTS `users` (
  `id` BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
  `email` VARCHAR(100) NOT NULL,
  `first_name` VARCHAR(20) NOT NULL,
  `last_name` VARCHAR(20) NOT NULL,
  `user_name` VARCHAR(20) NOT NULL,
  `uid` VARCHAR(30) NOT NULL,
  `stripe_account_id` VARCHAR(100) NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;