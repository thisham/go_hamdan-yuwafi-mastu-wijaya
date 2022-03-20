export type Book = {
  _id: number;
  _author_id: number;
  _publisher_id: number;
  title: string;
  price: number;
  stats: {
    page: number;
    weight: number;
  };
  publishedAt: Date;
  category: string[];
};

export type Author = {
  _id: number;
  firstName: string;
  lastName: string;
};

export type Publisher = {
  _id: number;
  publisherName: string;
};
