CREATE TABLE cities (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    province_id BIGINT UNSIGNED NOT NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_cities_province
        FOREIGN KEY (province_id) REFERENCES provinces(id)
        ON DELETE CASCADE
);