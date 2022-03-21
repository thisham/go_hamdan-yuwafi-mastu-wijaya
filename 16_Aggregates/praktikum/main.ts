import { client, connector } from "./connector";

async function main() {
  const db = await connector();

  // queries

  // init [passed]
  db.books.insertMany([
    {
      _id: 1,
      _author_id: 1,
      _publisher_id: 1,
      title: "Wawasan Pancasila",
      price: 71200,
      stats: { page: 238, weight: 300 },
      publishedAt: new Date("2018-10-01"),
      category: ["social", "politics"],
    },
    {
      _id: 2,
      _author_id: 1,
      _publisher_id: 2,
      title: "Mata Air Keteladanan",
      price: 106250,
      stats: { page: 672, weight: 650 },
      publishedAt: new Date("2017-09-01"),
      category: ["social", "politics"],
    },
    {
      _id: 3,
      _author_id: 1,
      _publisher_id: 1,
      title: "Revolusi Pancasila",
      price: 54400,
      stats: { page: 220, weight: 500 },
      publishedAt: new Date("2015-05-01"),
      category: ["social", "politics"],
    },
    {
      _id: 4,
      _author_id: 2,
      _publisher_id: 1,
      title: "Self Driving",
      price: 58650,
      stats: { page: 286, weight: 300 },
      publishedAt: new Date("2018-05-01"),
      category: ["self-development"],
    },
    {
      _id: 5,
      title: "Self Disruption",
      _author_id: 2,
      _publisher_id: 2,
      price: 83300,
      stats: { page: 400, weight: 800 },
      publishedAt: new Date("2018-05-01"),
      category: ["self-development"],
    },
  ]);

  db.authors.insertMany([
    { _id: 1, firstName: "Yudi", lastName: "Latif" },
    { _id: 2, firstName: "Rhenald", lastName: "Kasali" },
  ]);

  db.publishers.insertMany([
    { _id: 1, publisherName: "Expose" },
    { _id: 2, publisherName: "Mizan" },
  ]);

  // clause 2:1 [passed]
  // union books of author 1 + 2
  db.books.aggregate([
    {
      $project: {
        _id: 1,
        _author_id: 1,
        _publisher_id: 1,
        title: 1,
        price: 1,
        stats: { page: 1, weight: 1 },
        published_at: 1,
        category: 1,
      },
    },
  ]);

  // clause 2:2 [passed]
  // _id, buku, harga dari author 1
  db.books.aggregate([
    { $match: { _author_id: 1 } },
    { $project: { _id: 1, title: 1, price: 1 } },
  ]);

  // clause 2:3 [passed]
  // total halaman buku buku author 2
  db.books.aggregate([
    { $match: { _author_id: 2 } },
    { $group: { _id: "$_author_id", totalPages: { $sum: "$stats.page" } } },
  ]);

  // clause 2:4a [passed]
  // books + author via author
  db.authors.aggregate([
    {
      $lookup: {
        from: "books",
        localField: "_id",
        foreignField: "_author_id",
        as: "books",
      },
    },
  ]);

  // clause 2:4b [passed]
  // books + author via book
  db.books.aggregate([
    {
      $lookup: {
        from: "authors",
        localField: "_author_id",
        foreignField: "_id",
        as: "authors",
      },
    },
  ]);

  // clause 2:5 [passed]
  // book + author + publisher
  db.books.aggregate([
    {
      $lookup: {
        from: "authors",
        localField: "_author_id",
        foreignField: "_id",
        as: "authors",
      },
    },
    {
      $lookup: {
        from: "publishers",
        localField: "_publisher_id",
        foreignField: "_id",
        as: "publishers",
      },
    },
  ]);

  // clause 2:6 [failed]
  // refer
  // @error: Unrecognized pipeline stage name: '$sum'
  db.authors.aggregate([
    {
      $lookup: {
        from: "books",
        localField: "_id",
        foreignField: "_author_id",
        as: "books",
      },
    },
    {
      $lookup: {
        from: "publishers",
        localField: "_publisher_id",
        foreignField: "_id",
        as: "publishers",
      },
    },
    {
      $project: {
        _id: { $concat: ["$firstName", " ", "$lastName"] },
        publishers: "$publishers",
        list_of_books: {
          $map: {
            input: "$books",
            as: "book",
            in: {
              $concat: ["$$book.title", " - "],
            },
          },
        },
      },
    },
  ]);

  // clause 2:7 [passed]
  // diskon (refer)
  db.books.aggregate([
    {
      $project: {
        _id: 1,
        name: 1,
        price: 1,
        discount: {
          $switch: {
            branches: [
              { case: { $gt: ["$price", 90000] }, then: "3%" },
              { case: { $gt: ["$price", 60000] }, then: "2%" },
            ],
            default: "1%",
          },
        },
      },
    },
  ]);

  // clause 2:8 [passed]
  // book + price + author + publisher sort by price
  db.books.aggregate([
    {
      $lookup: {
        from: "publishers",
        localField: "_publisher_id",
        foreignField: "_id",
        as: "publishers",
      },
    },
    { $unwind: "$publishers" },
    {
      $lookup: {
        from: "authors",
        localField: "_author_id",
        foreignField: "_id",
        as: "authors",
      },
    },
    { $unwind: "$authors" },
    { $sort: { price: -1 } },
    {
      $project: {
        _id: 0,
        title: 1,
        price: 1,
        author: { $concat: ["$authors.firstName", " ", "$authors.lastName"] },
        publisher: "$publishers.publisherName",
      },
    },
  ]);

  // clause 2:9 [passed]
  db.books.aggregate([
    {
      $lookup: {
        from: "publishers",
        localField: "_publisher_id",
        foreignField: "_id",
        as: "publishers",
      },
    },
    { $unwind: "$publishers" },
    {
      $project: {
        title: 1,
        price: 1,
        publisher: "$publishers.publisherName",
      },
    },
    {
      $unionWith: {
        coll: "books",
        pipeline: [
          { $match: { $or: [{ _id: 3 }, { _id: 4 }] } },
          {
            $project: {
              title: 1,
              price: 1,
              publisher: "$publishers.publisherName",
            },
          },
        ],
      },
    },
  ]);
}

main()
  .then(() => console.log)
  .catch(() => console.error)
  .finally(() => client.close());
