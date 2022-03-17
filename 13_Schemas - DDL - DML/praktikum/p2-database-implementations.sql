-- ----------------------------------
-- initiating a database (clause 2:1)
-- create database alta_online_shop;

-- will be had: 
--  - many payment methods (clause 2:7c)
--  - many addresses (clause 2:7b)
create table users (
  `id` varchar(36) primary key not null,
  `name` varchar(255) not null,
  `birthdate` date,
  `gender` enum('male', 'female'),
  `status` enum('active', 'inactive'),
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

-- will be had an user (clause 2:7b)
create table addresses (
  `id` varchar(36) primary key not null,
  `user_id` varchar(36) not null,
  `line_one` varchar(255),
  `line_two` varchar(255),
  `city` varchar(63),
  `province` varchar(63),
  `country` varchar(31)
);

-- will be had a product
create table product_descriptions (
  `id` int primary key not null,
  `name` varchar(255) not null,
  `price` float not null,
  `description` text
);

-- will be had many product
create table product_types (
  `id` int not null primary key,
  `name` varchar(127) not null
);

-- will be had:
--  - a product description
--  - a product type
create table products (
  `id` int primary key not null,
  `product_description_id` int not null unique,
  `product_type_id` int not null
);

-- will be had many transactions
create table operators (
  `id` varchar(36) not null primary key,
  `name` varchar(255) not null
);

-- will be had:
--  - many users
--  - an payment method description (clause 2:7a)
create table payment_methods (
  `id` int not null primary key,
  `payment_method_description_id` int unique not null
);

-- will be had a payment method (clause 2:7a)
create table payment_method_descriptions (
  `id` int not null primary key,
  `name` varchar(63) not null,
  `description` text
);

-- relation table between users and payment methods (clause 2:7c)
create table user_payment_method_details (
  `user_id` varchar(36) not null,
  `payment_method_id` int not null
);

-- will be had a transaction detail
create table transactions (
  `id` varchar(36) not null primary key,
  `product_id` int not null,
  `user_id` varchar(36) not null,
  `operator_id` varchar(36) not null,
  `transaction_detail_id` varchar(36) not null unique,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

-- will be had a transaction
create table transaction_details (
  `id` varchar(36) not null primary key,
  `amount` int not null,
  `costs` float not null
);

-- clause 2:3
create table couriers (
  `id` varchar(36) primary key not null,
  `name` varchar(127) not null,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

-- ------------
-- modify table

-- clause 2:4
alter table couriers
  add column `basic_cost` float;

-- clause 2:5
alter table couriers
  rename shippings;

-- ----------------------
-- create index creations
alter table addresses
  add key `address_with_user` (`user_id`);

alter table products
  add key `product_with_product_type`(`product_type_id`),
  add key `product_meets_product_description`(`product_description_id`);

alter table payment_methods
  add key `payment_method_meets_payment_description_method`(`payment_method_description_id`);

alter table transactions
  add key `transaction_with_product`(`product_id`),
  add key `transaction_with_operator`(`operator_id`),
  add key `transaction_with_user`(`user_id`),
  add key `transaction_meets_transaction_detail`(`transaction_detail_id`);

alter table user_payment_method_details
  add key `payment_method_with_user`(`payment_method_id`),
  add key `user_with_payment_method`(`user_id`);

-- -----------
-- constraints
alter table addresses
  add constraint `address_with_user` foreign key (`user_id`) references users(`id`);

alter table products
  add constraint `product_with_product_type` foreign key (`product_type_id`) references product_types(`id`),
  add constraint `product_meets_product_description` foreign key (`product_description_id`) references product_descriptions(`id`);

alter table payment_methods
  add constraint `payment_method_meets_payment_method_description` foreign key (`payment_method_description_id`) references payment_methods(`id`);

alter table transactions
  add constraint `transaction_with_product` foreign key (`product_id`) references products(`id`),
  add constraint `transaction_with_operator` foreign key (`operator_id`) references operators(`id`),
  add constraint `transaction_with_user` foreign key (`user_id`) references users(`id`),
  add constraint `transaction_meets_transaction_detail` foreign key (`transaction_detail_id`) references transaction_details(`id`);

alter table user_payment_method_details
  add constraint `user_with_payment_method` foreign key (`user_id`) references users(`id`),
  add constraint `payment_method_with_user` foreign key (`payment_method_id`) references payment_methods(`id`);

-- ---------------
-- table deletions
-- clause 2:6
drop table shippings;
