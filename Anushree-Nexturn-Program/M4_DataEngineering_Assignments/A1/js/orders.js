const mongoose = require('mongoose');

// Define the schema for an Order
const orderSchema = new mongoose.Schema({
  userId: { type: mongoose.Schema.Types.ObjectId, ref: 'User', required: true },
  productIds: [{ type: mongoose.Schema.Types.ObjectId, ref: 'Product', required: true }],
  totalAmount: { type: Number, required: true },
  status: { type: String, default: 'pending' },
  createdAt: { type: Date, default: Date.now }
});

// Create the Order model based on the schema
const Order = mongoose.model('Order', orderSchema);

module.exports = Order;
