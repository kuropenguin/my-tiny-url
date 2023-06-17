DROP TABLE IF EXISTS `urls`;

CREATE TABLE
    `urls` (
        `original_url` varchar(512) NOT NULL,
        `tiny_url` varchar(255) NOT NULL,
        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`original_url`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
