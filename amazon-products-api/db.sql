CREATE TABLE IF NOT EXISTS products (
    product_url VARCHAR(255) PRIMARY KEY,
    product_name VARCHAR(255) UNIQUE KEY,
    price INT,
    details VARCHAR(255),
    image_url VARCHAR(255),
    review_count INT,
    date_updated datetime
);


Insert into products(product_url,product_name,price,details,image_url,review_count) Values("?","?","?",?","?","?")