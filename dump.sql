CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext,
  `password` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO users (created_at,updated_at,deleted_at,username,password) VALUES
	 ('2024-08-31 13:18:40.480','2024-08-31 13:18:40.480',NULL,'guest','$2a$10$4Sp6mmuz9AJwS2xnF4BurOUZaDrFEhIw7T1fH8OfUcNkFKLhDxMya');
