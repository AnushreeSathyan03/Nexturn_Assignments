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

module.exports = Product;
