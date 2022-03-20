import { MongoClient } from "mongodb";
import { Author, Book, Publisher } from "./types";

const connectionString = "mongodb://localhost:27017";
const dbName = "digital_outlet";

export const client = new MongoClient(connectionString);

export async function connector() {
  await client.connect();
  const db = client.db(dbName);
  const collections = {
    books: db.collection<Book>("books"),
    authors: db.collection<Author>("authors"),
    publishers: db.collection<Publisher>("publishers"),
  };
  return collections;
}
