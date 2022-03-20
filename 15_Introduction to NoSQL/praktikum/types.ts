export type User = {
  _id: number;
  name: string;
  status: boolean;
  birthdate: Date;
  gender: "M" | "F";
  created_at?: Date;
  updated_at?: Date;
};

export type Operator = {
  _id: number;
  name: string;
  created_at?: Date;
  updated_at?: Date;
};

export type ProductType = {
  _id: number;
  name: string;
  created_at?: Date;
  updated_at?: Date;
};

export type ProductDescription = {
  _id: number;
  description: string;
  created_at?: Date;
  updated_at?: Date;
};

export type Product = {
  _id: number;
  _product_type_id: number;
  _operator_id: number;
  code: string;
  name: string;
  status: boolean;
  created_at?: Date;
  updated_at?: Date;
};

export type TransactionDetail = {
  _transaction_id: number;
  _product_id: number;
  status: string;
  quantity: number;
  price: number;
  created_at?: Date;
  updated_at?: Date;
};

export type PaymentMethod = {
  _id: number;
  name: string;
  status: boolean;
  created_at?: Date;
  updated_at?: Date;
};

export type Transaction = {
  _id: number;
  _user_id: number;
  _payment_method_id: number;
  status: string;
  total_quantity: number;
  total_price: number;
  created_at?: Date;
  updated_at?: Date;
};
