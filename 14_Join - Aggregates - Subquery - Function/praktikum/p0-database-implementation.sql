create table operators (
  `id` int primary key not null auto_increment,
  `name` varchar(255),
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

create table product_types (
  `id` int primary key not null auto_increment,
  `name` varchar(255),
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

create table products (
  `id` int primary key not null auto_increment,
  `product_type_id` int not null,
  `operator_id` int not null,
  `code` varchar(50),
  `name` varchar(100),
  `status` smallint,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

alter table products
  add constraint `product_with_product_type` foreign key (`product_type_id`) references product_types(`id`),
  add constraint `product_with_operator` foreign key (`operator_id`) references operators(`id`);

create table product_descriptions (
  `product_id` int primary key not null,
  `description` text,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

alter table product_descriptions
  add constraint `product_detail_meets_product` foreign key (`product_id`) references products(`id`);

create table payment_methods (
  `id` int not null primary key auto_increment,
  `name` varchar(255) not null,
  `status` smallint,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

create table users (
  `id` int not null primary key,
  `status` smallint,
  `birthdate` date,
  `gender` char(1),
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

create table transactions (
  `id` int not null primary key,
  `user_id` int not null,
  `payment_method_id` int not null,
  `status` varchar(10),
  `total_quantity` int not null,
  `total_price` int not null,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

alter table transactions 
  add constraint `transaction_with_user` foreign key (`user_id`) references users(`id`),
  add constraint `transaction_with_payment_method` foreign key (`payment_method_id`) references payment_methods(`id`);

create table transaction_details (
  `transaction_id` int not null,
  `product_id` int not null,
  `status` varchar(10),
  `quantity` int not null,
  `price` float,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

alter table transaction_details 
  add constraint `transaction_detail_primary_key` primary key (`transaction_id`, `product_id`),
  add constraint `transaction_detail_with_transaction` foreign key (`transaction_id`) references transactions(`id`),
  add constraint `transaction_detail_with_product` foreign key (`product_id`) references products(`id`);
