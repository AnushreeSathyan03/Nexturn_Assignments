const Order = require('../models/orderModel');
const Product = require('../models/productModel');

// Find orders within a specific time range
const findOrdersInTimeRange = async (startDate, endDate) => {
  const result = await Order.aggregate([
    { $match: { date: { $gte: new Date(startDate), $lte: new Date(endDate) } } },
    { $lookup: {
      from: "users",
      localField: "userId",
      foreignField: "_id",
      as: "userDetails"
    }},
    { $unwind: "$userDetails" },
    { $project: { "userDetails.name": 1, "productDetails": 1, "date": 1 } }
  ]);
  return result;
};

// Update stock after order completion
const updateStockAfterOrder = async (order) => {
  for (const item of order.items) {
    await Product.updateOne(
      { _id: item.productId },
      { $inc: { stock: -item.quantity } }
    );
  }
};

module.exports = {
  findOrdersInTimeRange,
  updateStockAfterOrder
};
