const mongoose = require('mongoose');

const ratingSchema = new mongoose.Schema({
  userId: String,
  rating: Number
}, { _id: false });

const productSchema = new mongoose.Schema({
  productId: { type: String, required: true, unique: true },
  name: { type: String, required: true },
  category: { type: String, required: true },
  price: { type: Number, required: true },
  stock: { type: Number, required: true },
  ratings: [ratingSchema]
});

const Product = mongoose.model('Product', productSchema);

// Sample Product Data
const sampleProducts = [
  {
    productId: "P001",
    name: "Wireless Mouse",
    category: "Electronics",
    price: 50,
    stock: 200,
    ratings: [
      { userId: "U002", rating: 4.5 },
      { userId: "U003", rating: 3.0 }
    ]
  },
  {
    productId: "P002",
    name: "Bluetooth Headphones",
    category: "Electronics",
    price: 100,
    stock: 150,
    ratings: [
      { userId: "U001", rating: 5.0 },
      { userId: "U003", rating: 4.0 }
    ]
  },
  {
    productId: "P003",
    name: "Smartphone",
    category: "Electronics",
    price: 500,
    stock: 50,
    ratings: [
      { userId: "U001", rating: 4.7 },
      { userId: "U002", rating: 4.0 }
    ]
  }
];

// Insert sample products
Product.insertMany(sampleProducts)
  .then(() => console.log('Sample Products added'))
  .catch(err => console.error('Error inserting products:', err));

module.exports = Product;
