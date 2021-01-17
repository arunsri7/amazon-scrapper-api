CREATE TABLE IF NOT EXISTS products (
    product_url VARCHAR(255) PRIMARY KEY,
    product_name VARCHAR(255) UNIQUE KEY,
    price VARCHAR(255),
    details VARCHAR(255),
    image_url VARCHAR(255),
    review_count VARCHAR(255),
    date_updated datetime
);


Insert into products(product_url,product_name,price,details,image_url,review_count) Values("?","?","?",?","?","?")