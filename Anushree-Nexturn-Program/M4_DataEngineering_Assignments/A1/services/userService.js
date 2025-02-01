const User = require('../models/userModel');
const Order = require('../models/orderModel');

// Find high-spending users
const findHighSpendingUsers = async () => {
  const result = await User.aggregate([
    { $lookup: {
      from: "orders",
      localField: "_id",
      foreignField: "userId",
      as: "orders"
    }},
    { $unwind: "$orders" },
    { $group: {
      _id: "$_id",
      totalSpending: { $sum: "$orders.totalAmount" }
    }},
    { $match: { totalSpending: { $gt: 500 } } }
  ]);
  return result;
};

module.exports = {
  findHighSpendingUsers
};
