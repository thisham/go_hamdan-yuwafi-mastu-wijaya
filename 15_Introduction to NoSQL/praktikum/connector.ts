import { MongoClient } from "mongodb";
import {
  Operator,
  PaymentMethod,
  Product,
  ProductDescription,
  ProductType,
  Transaction,
  TransactionDetail,
  User,
} from "./types";

const connectionString = "mongodb://localhost:27017";
const dbName = "digital_outlet";

export const client = new MongoClient(connectionString);

export async function connector() {
  await client.connect();
  const db = client.db(dbName);
  const collections = {
    users: db.collection<User>("users"),
    products: db.collection<Product>("products"),
    payment_methods: db.collection<PaymentMethod>("payment_methods"),
    transactions: db.collection<Transaction>("transactions"),
    transaction_details: db.collection<TransactionDetail>(
      "transaction_details"
    ),
    product_descriptions: db.collection<ProductDescription>(
      "product_descriptions"
    ),
    product_types: db.collection<ProductType>("product_types"),
    operators: db.collection<Operator>("operators"),
  };
  return collections;
}
