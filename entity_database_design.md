Berikut adalah contoh yang telah disesuaikan dengan DDL yang kamu berikan:

---

## A. Entities

### a. Entity: users
##### Attributes:
- `id` SERIAL PRIMARY KEY
- `name` VARCHAR(100) NOT NULL
- `email` VARCHAR(100) UNIQUE NOT NULL
- `password` VARCHAR(60) NOT NULL
- `phone` VARCHAR(14) UNIQUE NOT NULL
- `role` ENUM("customer", "hotel_owner") NOT NULL

### b. Entity: hotels
##### Attributes:
- `id` SERIAL PRIMARY KEY
- `name` VARCHAR(100) NOT NULL
- `description` TEXT
- `address` VARCHAR(255) NOT NULL
- `city` VARCHAR(100) NOT NULL
- `zipcode` VARCHAR(20) NOT NULL
- `country` VARCHAR(100) NOT NULL
- `phone` VARCHAR(14) UNIQUE NOT NULL
- `email` VARCHAR(100) UNIQUE NOT NULL
- `star` INT DEFAULT 1

### c. Entity: room_types
##### Attributes:
- `id` SERIAL PRIMARY KEY
- `hotel_id` INT REFERENCES hotels(id) ON DELETE CASCADE
- `name` VARCHAR(100) NOT NULL
- `description` TEXT
- `price` DECIMAL(10, 2) CHECK(price >= 0) NOT NULL
- `room_size` DECIMAL(4, 1) NOT NULL
- `guest` INT CHECK(guest >= 0) NOT NULL

### d. Entity: room_bed_types
##### Attributes:
- `room_type_id` INT PRIMARY KEY REFERENCES room_types(id) ON DELETE CASCADE
- `double_bed` INT DEFAULT 0
- `single_bed` INT DEFAULT 0
- `king_bed` INT DEFAULT 0

### e. Entity: room_facilities
##### Attributes:
- `room_type_id` INT PRIMARY KEY REFERENCES room_types(id) ON DELETE CASCADE
- `has_shower` BOOLEAN DEFAULT FALSE
- `has_refrigerator` BOOLEAN DEFAULT FALSE
- `seating_area` BOOLEAN DEFAULT FALSE
- `air_conditioning` BOOLEAN DEFAULT FALSE
- `has_breakfast` BOOLEAN DEFAULT FALSE
- `has_wifi` BOOLEAN DEFAULT FALSE
- `smoking_allowed` BOOLEAN DEFAULT FALSE

### f. Entity: rooms
##### Attributes:
- `id` SERIAL PRIMARY KEY
- `room_type_id` INT REFERENCES room_types(id) ON DELETE CASCADE
- `room_number` INT NOT NULL
- `status` ENUM("available", "booked", "maintenance") DEFAULT "available"

### g. Entity: bookings
##### Attributes:
- `id` SERIAL PRIMARY KEY
- `user_id` INT REFERENCES users(id) ON DELETE CASCADE
- `hotel_id` INT REFERENCES hotels(id) ON DELETE CASCADE
- `room_id` INT REFERENCES rooms(id) ON DELETE CASCADE
- `check_in_date` TIMESTAMP NOT NULL
- `check_out_date` TIMESTAMP NOT NULL
- `total_price` DECIMAL(10, 2) NOT NULL
- `status` ENUM("booked", "cancelled", "completed") DEFAULT "booked"

### h. Entity: invoices
##### Attributes:
- `booking_id` INT PRIMARY KEY REFERENCES bookings(id) ON DELETE CASCADE
- `xendit_invoice_id` TEXT NOT NULL
- `amount` DECIMAL(10, 2) NOT NULL
- `status` ENUM("PENDING", "PAID", "EXPIRED") DEFAULT "PENDING"

### i. Entity: payments
##### Attributes:
- `invoice_id` INT PRIMARY KEY REFERENCES invoices(booking_id) ON DELETE CASCADE
- `payment_method` TEXT NOT NULL
- `paid_amount` DECIMAL(10, 2) NOT NULL
- `paid_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP

### j. Entity: balances
##### Attributes:
- `user_id` INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE
- `balance` DECIMAL(10, 2) DEFAULT 0

---

## B. Relationships:

### 1. users to bookings
- **Type:** One to Many
- **Description:** One user can make many bookings, but each booking is associated with only one user.

### 2. hotels to room_types
- **Type:** One to Many
- **Description:** One hotel can have many room types, but each room type is associated with only one hotel.

### 3. room_types to rooms
- **Type:** One to Many
- **Description:** One room type can be associated with many rooms, but each room is associated with only one room type.

### 4. room_types to room_bed_types
- **Type:** One to One
- **Description:** Each room type can have only one set of bed types, and each set of bed types is associated with only one room type.

### 5. room_types to room_facilities
- **Type:** One to One
- **Description:** Each room type can have only one set of facilities, and each set of facilities is associated with only one room type.

### 6. bookings to invoices
- **Type:** One to One
- **Description:** Each booking has only one invoice, and each invoice is associated with only one booking.

### 7. invoices to payments
- **Type:** One to Many 
- **Description:** One invoice can have many payments, but each payment is associated with only one invoice.

### 8. users to balances
- **Type:** One to One 
- **Description:** One user has only one balance, and each balance is associated with only one user.

---

## C. Adjusted Integrity Constraints:
- The `email` in the `users` entity must be unique.
- All attributes must be non-null (`NOT NULL`).
- Foreign keys (FK) in related entities such as `hotel_id`, `room_type_id`, `room_id`, `user_id`, `booking_id`, and `invoice_id` must reference valid primary keys.

---

## D. Adjusted Additional Notes:
- The system allows hotel owners to define different types of rooms with specific bed configurations and facilities.
- Customers can book rooms, and their bookings generate invoices that can be paid through various methods.
- The system also allows tracking of room availability based on bookings and other statuses like maintenance.

---

## 1NF (First Normal Form):
- Each column must contain atomic values, and each entity must have a primary key.
- This is fulfilled since all attributes are atomic, and each entity has a primary key.

## 2NF (Second Normal Form):
- The entity must be in 1NF, and all non-prime attributes must fully depend on the primary key.
- All entities are in 1NF.
- All non-prime attributes in each entity fully depend on the primary key.

## 3NF (Third Normal Form):
- The entity must be in 2NF, and there must be no transitive dependency of non-prime attributes on the primary key.
- All entities are in 2NF.
- There are no transitive dependencies in these entities, as all non-prime attributes directly depend on the primary key.