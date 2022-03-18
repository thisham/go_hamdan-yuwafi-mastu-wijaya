-- clause 2:1
select * from transactions where user_id = 1 
  union select * from transactions where user_id = 2;

-- clause 2:2
select sum(price) as sum_prices 
  from transactions 
  join transaction_details on transactions.id = transaction_details.transaction_id 
  where user_id = 1;

-- clause 2:3
select count(*) as transactions, sum(price) as sum_prices, sum(quantity) as sold_products
  from transactions
  join transaction_details on transactions.id = transaction_details.transaction_id
  join products on transaction_details.product_id = products.id
  where products.product_type_id = 2;

-- clause 2:4
select `products`.`id` as `id`, `product_types`.`name` as `product_type`, 
  `operator_id`, `code`, `products`.`name` as `name`, `status`, `products`.`created_at`, 
  `products`.`updated_at`
  from products
  join product_types on products.product_type_id = product_types.id;

-- clause 2:5
select *
  from transactions 
  join users on transactions.user_id = users.id
  join transaction_details on transactions.id = transaction_details.transaction_id
  join products on transaction_details.product_id = products.id;

-- clause 2:6
delimiter #
create trigger deleteTransactionDetail
  before delete on transactions 
  for each row
begin
  delete from transaction_details where transaction_id = old.id;
end #

-- delimiter ;

-- clause 2:7
delimiter #
create trigger decreaseQuantityWhenDetailDeleted
  after delete on transaction_details
  for each row
begin
  update transactions set 
    total_price = total_price - old.price, 
    total_quantity = total_quantity - old.quantity
    where id = old.transaction_id;
end #

delimiter ;

-- clause 2:8
select * from products where id not in
  (select product_id from transaction_details);
