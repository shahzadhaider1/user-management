-- Create Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(50) UNIQUE NOT NULL,
    otp VARCHAR(10),
    otp_expiration_time TIMESTAMP
);

CREATE TABLE facts (
    id BIGSERIAL PRIMARY KEY,
    bedroom TEXT[],
    bathroom BIGINT[],
    plot_area FLOAT,
    built_up_area FLOAT,
    view BIGINT[],
    furnished BIGINT,
    ownership BIGINT,
    sc_currency_id VARCHAR,
    unit_of_measure VARCHAR
);
