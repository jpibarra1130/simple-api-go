
-- +goose Up
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL DEFAULT '',
  `hashed_password` varchar(255) NOT NULL DEFAULT '',
  `created_at` BIGINT DEFAULT NULL,
  `updated_at` BIGINT DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_users_on_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

-- +goose Down
DROP TABLE `users`;
