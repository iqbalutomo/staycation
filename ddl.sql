CREATE TYPE user_role_enum AS ENUM('customer', 'hotel_owner');
CREATE TYPE room_status_enum AS ENUM('available', 'booked', 'maintenance');
CREATE TYPE booking_status_enum AS ENUM('booked', 'cancelled', 'completed');
CREATE TYPE invoice_status_enum AS ENUM('PENDING', 'PAID', 'EXPIRED');

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    phone VARCHAR(14) UNIQUE NOT NULL,
    role user_role_enum NOT NULL
);

CREATE TABLE balances(
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    balance DECIMAL(10, 2) DEFAULT 0
);

CREATE TABLE hotels(
    id SERIAL PRIMARY KEY,
    owner_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    zipcode VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    phone VARCHAR(14) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    star INT DEFAULT 1
);

CREATE TABLE room_types(
    id SERIAL PRIMARY KEY,
    hotel_id INT REFERENCES hotels(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) CHECK(price >= 0) NOT NULL,
    room_size DECIMAL(4, 1) NOT NULL,
    guest INT CHECK(guest >= 1) NOT NULL
);

CREATE TABLE room_bed_types(
    room_type_id INT PRIMARY KEY REFERENCES room_types(id) ON DELETE CASCADE,
    double_bed INT DEFAULT 0,
    single_bed INT DEFAULT 0,
    king_bed INT DEFAULT 0
);

CREATE TABLE room_facilities(
    room_type_id INT PRIMARY KEY REFERENCES room_types(id) ON DELETE CASCADE,
    has_shower BOOLEAN DEFAULT FALSE,
    has_refrigerator BOOLEAN DEFAULT FALSE,
    seating_area BOOLEAN DEFAULT FALSE,
    air_conditioning BOOLEAN DEFAULT FALSE,
    has_breakfast BOOLEAN DEFAULT FALSE,
    has_wifi BOOLEAN DEFAULT FALSE,
    smoking_allowed BOOLEAN DEFAULT FALSE
);

CREATE TABLE rooms(
    id SERIAL PRIMARY KEY,
    room_type_id INT REFERENCES room_types(id) ON DELETE CASCADE,
    room_number INT NOT NULL,
    status room_status_enum DEFAULT 'available'
);

CREATE TABLE bookings(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    hotel_id INT REFERENCES hotels(id) ON DELETE CASCADE,
    room_id INT REFERENCES rooms(id) ON DELETE CASCADE,
    check_in_date TIMESTAMP NOT NULL,
    check_out_date TIMESTAMP NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status booking_status_enum DEFAULT 'booked'
);

CREATE TABLE invoices(
    booking_id INT PRIMARY KEY REFERENCES bookings(id) ON DELETE CASCADE,
    xendit_invoice_id TEXT NOT NULL,
    invoice_url TEXT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status invoice_status_enum DEFAULT 'PENDING'
);

CREATE TABLE payments(
    id INT PRIMARY KEY,
    invoice_id INT REFERENCES invoices(booking_id) ON DELETE CASCADE,
    payment_method TEXT NOT NULL,
    paid_amount DECIMAL(10, 2) NOT NULL,
    paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

