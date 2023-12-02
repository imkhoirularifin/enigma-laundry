CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(100),
  email VARCHAR(50),
  username VARCHAR(50),
  password VARCHAR(100),
  role VARCHAR(40),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(100),
  price BIGINT,
  type VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE customers (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(100),
  phone_number VARCHAR(16),
  address TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE bills (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  bill_date DATE,
  customer_id UUID,
  user_id UUID,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE bill_details (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  bill_id UUID,
  product_id UUID,
  qty int,
  price BIGINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (bill_id) REFERENCES bills(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Users
INSERT INTO public.users
(id, "name", email, username, "password", "role", created_at, updated_at)
VALUES('df29b765-7235-40e3-abe7-43ebcfda6e2a'::uuid, 'Destry Faradila', 'destry.faradila@gmail.com', 'destry.faradila', '$2a$10$AuwgzIrYaZDKQHfo4WA4x.BtNMUxU/gCa9qdfozDHPdqOavqfV5Vm', 'employee', '2023-11-07 13:17:31.791', '2023-11-07 13:17:31.776');
INSERT INTO public.users
(id, "name", email, username, "password", "role", created_at, updated_at)
VALUES('fc5812b9-60e0-4d1d-9130-6f9d7e8cf99c'::uuid, 'Dinda Aditiya', 'dinda.aditiya@gmail.com', 'dinda.aditiya', '$2a$10$4StptFVnCmbgoPrclxN7ROQSHgZdLBVgt8Q.L3L0Gw4sdNHvSoigq', 'employee', '2023-11-07 13:23:23.175', '2023-11-07 13:23:23.161');
INSERT INTO public.users
(id, "name", email, username, "password", "role", created_at, updated_at)
VALUES('33d765b2-7170-45f2-9f3e-8794a6360531'::uuid, 'Dinda Aditiya', 'dinda.aditiya@gmail.com', 'dinda.aditiya', '$2a$10$CSycnwsGb8VdBRg.BS.84OJ8.5GeCXWtfXELXESK6Q70Ur4xx7HMS', 'employee', '2023-11-07 14:43:25.155', '2023-11-07 14:43:25.137');
INSERT INTO public.users
(id, "name", email, username, "password", "role", created_at, updated_at)
VALUES('adfef570-8c30-43e0-bf5d-3a34ab908fcf'::uuid, 'Jution Candra Kirana', 'jutionck@gmail.com', 'jutionck', '$2a$10$4StptFVnCmbgoPrclxN7ROQSHgZdLBVgt8Q.L3L0Gw4sdNHvSoigq', 'admin', '2023-11-06 13:53:36.324', '2023-11-06 13:53:36.324');
INSERT INTO public.users
(id, "name", email, username, "password", "role", created_at, updated_at)
VALUES('7e8097d8-fef1-49ed-9f7b-cecc9520cca8'::uuid, 'Tiara Agisti', 'tiara.agisti@gmail.com', 'tiara.agisti', '$2a$10$4StptFVnCmbgoPrclxN7ROQSHgZdLBVgt8Q.L3L0Gw4sdNHvSoigq', 'employee', '2023-11-06 13:53:36.324', '2023-11-06 13:53:36.324');


-- Products
INSERT INTO public.products
(id, "name", price, "type", created_at, updated_at)
VALUES('6afb0c8f-a1eb-4107-983b-2a100d8e8d8b'::uuid, 'Cuci', 5000, 'Kg', '2023-11-06 14:10:54.920', '2023-11-06 14:10:54.920');
INSERT INTO public.products
(id, "name", price, "type", created_at, updated_at)
VALUES('7891cc2e-4839-4d71-a9a3-4ca682fff990'::uuid, 'Setrika', 4000, 'Kg', '2023-11-06 14:10:54.920', '2023-11-06 14:10:54.920');
INSERT INTO public.products
(id, "name", price, "type", created_at, updated_at)
VALUES('23c18bdc-f4dd-4af3-ba72-f971f8af521d'::uuid, 'Cuci + Setrika', 8000, 'Kg', '2023-11-06 14:10:54.920', '2023-11-06 14:10:54.920');


-- Customers
INSERT INTO public.customers
(id, "name", phone_number, address, created_at, updated_at)
VALUES('7ed12603-cf60-494c-91e6-d38ac8022da6'::uuid, 'John Lbf', '018393939', 'Jakarta Selatan', '2023-11-06 14:09:46.545', '2023-11-06 14:09:46.545');
INSERT INTO public.customers
(id, "name", phone_number, address, created_at, updated_at)
VALUES('1ca4bd3a-9b19-45ca-ba14-021ff41c08c2'::uuid, 'Amin Rais', '28282910', 'Jakarta Selatan', '2023-11-06 14:09:46.545', '2023-11-06 14:09:46.545');
INSERT INTO public.customers
(id, "name", phone_number, address, created_at, updated_at)
VALUES('b8c8afff-fa55-4eec-8f6e-4f4039873118'::uuid, 'Anisa R', '283910111', 'Jakarta Selatan', '2023-11-06 14:09:46.545', '2023-11-06 14:09:46.545');