-- init.sql

-- Create Addresses table
CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    street_address VARCHAR(255),
    city VARCHAR(255),
    state VARCHAR(255),
    postal_code VARCHAR(20),
    country VARCHAR(255)
);

-- Create Profile table
CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    addresses_id BIGINT REFERENCES addresses(id),
    profile_image_url VARCHAR(255),
    phone_number VARCHAR(255),
    company_number VARCHAR(255),
    whatsapp_number VARCHAR(255),
    botim VARCHAR(255),
    tawasal VARCHAR(255),
    gender BIGINT,
    all_languages_id BIGINT[],
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    ref_no VARCHAR(50) UNIQUE NOT NULL,
    cover_image_url VARCHAR(255),
    nic_no VARCHAR(50),
    nic_image_url VARCHAR(255),
    passport_no VARCHAR(50),
    passport_image_url VARCHAR(255)
);

-- Create Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    profile_id BIGINT REFERENCES profile(id) ON DELETE SET NULL,
    phone_number VARCHAR(50) UNIQUE NOT NULL,
    otp VARCHAR(10),
    otp_expiration_time TIMESTAMP
);
