const Product = require('../models/productModel');

// Get popular products by average rating
const listPopularProductsByRating = async () => {
  const result = await Product.aggregate([
    { $unwind: "$ratings" },
    { $group: { _id: "$_id", avgRating: { $avg: "$ratings" } } },
    { $match: { avgRating: { $gte: 4 } } },
  ]);
  return result;
};

module.exports = {
  listPopularProductsByRating
};
