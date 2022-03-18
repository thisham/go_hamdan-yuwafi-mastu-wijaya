-- clause 1:1
insert into operators (`id`, `name`) values 
  (1, 'operator satu'), 
  (2, 'operator dua'), 
  (3, 'operator tiga'), 
  (4, 'operator empat'),
  (5, 'operator lima');

insert into product_types (`id`, `name`) values
  (1, 'laptop'),
  (2, 'handheld'),
  (3, 'perangkat tambahan');

insert into products (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`) values
  (1, 1, 3, 'LAP389', 'SUSA-X402', 1),
  (2, 1, 3, 'LAP380', 'ASHER-ASD1', 1),

  (3, 2, 1, 'HAN131', 'OPO-3193', 1),
  (4, 2, 1, 'HAN492', 'SIOMAY-X3', 1),
  (5, 2, 1, 'HAN213', 'VIVA-V32', 1),
  
  (6, 3, 4, 'DEV302', 'Headset 30an', 1),
  (7, 3, 4, 'DEV381', 'Kabel data', 1),
  (8, 3, 4, 'DEV323', 'SDCard 32GB', 1);

insert into product_descriptions (`product_id`, `description`) values
  (1, 'Laptop apik SUSA berkualiti'),
  (2, 'Laptop apik ASHER berkualiti'),

  (3, 'HP apik OPO kamera mantap'),
  (4, 'HP apik SIOMAY nikmat lezat'),
  (5, 'HP apik VIVA tergahar'),

  (6, 'Devkit paling ngebass'),
  (7, 'Kabel data kuat awet'),
  (8, 'Kapasitas lega buat simpen banyak di kandang');

insert into payment_methods (`id`, `name`, `status`) values
  (1, 'tukar beras', 1),
  (2, 'linkwae', 1),
  (3, 'bayar pake tanah', 1);

insert into users (`id`, `status`, `birthdate`, `gender`) values 
  (1, 1, '2002-02-02', 'L'),
  (2, 1, '2000-02-03', 'P'),
  (3, 1, '1989-04-02', 'L'),
  (4, 1, '1995-02-05', 'P'),
  (5, 1, '1999-12-25', 'L');

delimiter #

create procedure insertTransactions()
begin
  declare currentTransaction int default 0;
  declare currentUser int default 0;
  declare currentUserTransaction int default 0;
  declare currentTransactionProduct int default 1;
  declare flatPrice int default 100000;
  
  while currentUser < 5 do
    set currentUserTransaction = 0;
    set currentUser = currentUser + 1;
    
    while currentUserTransaction < 3 do
      set currentTransaction = currentTransaction + 1;
      set currentUserTransaction = currentUserTransaction + 1;
      insert into transactions (`id`, `user_id`, `payment_method_id`, `status`, `total_quantity`, `total_price`) values
        (currentTransaction, currentUser, currentUserTransaction, 'paid', 3, 3 * flatPrice);
      insert into transaction_details (`transaction_id`, `product_id`, `status`, `quantity`, `price`) values 
        (currentTransaction, 1, 'paid', 1, flatPrice),
        (currentTransaction, 2, 'paid', 1, flatPrice),
        (currentTransaction, 3, 'paid', 1, flatPrice);
    end while;
  end while;
  commit;
end #

delimiter ;

call insertTransactions();

-- clause 1:2
select * from users where gender = 'L';
select * from products where id = 3;
select * from users where (day(now) - day(created_at)) <= 7 and `name` like '%a%';
select * from users where gender = 'P';
select * from users order by `name` asc;
select * from products limit 5;

-- clause 1:3
update products set `name` = 'product dummy' where id = 1;
update transaction_details set quantity = 3 where product_id = 1;

-- clause 1:4
delete from products where id = 1;
delete from products where product_type_id = 1;
