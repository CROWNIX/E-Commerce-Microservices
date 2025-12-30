CREATE TABLE IF NOT EXISTS products (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    category_id BIGINT UNSIGNED DEFAULT NULL,
    name VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    images JSON NOT NULL,
    description TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
    price BIGINT UNSIGNED NOT NULL,
    stock INT UNSIGNED NOT NULL DEFAULT 0,
    final_price BIGINT UNSIGNED NOT NULL,
    discount_percent TINYINT UNSIGNED NOT NULL DEFAULT 0,
    minimum_purchase TINYINT UNSIGNED NOT NULL DEFAULT 1,
    maximum_purchase TINYINT UNSIGNED DEFAULT NULL,
    created_at TIMESTAMP NULL DEFAULT NULL,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    PRIMARY KEY (id),
    KEY products_category_id_foreign (category_id),
    
    CONSTRAINT products_category_id_foreign 
        FOREIGN KEY (category_id) REFERENCES categories(id) 
        ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
