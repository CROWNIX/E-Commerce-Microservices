CREATE TABLE sub_districts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    city_id BIGINT UNSIGNED NOT NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_sub_districts_city
        FOREIGN KEY (city_id) REFERENCES cities(id)
        ON DELETE CASCADE
);