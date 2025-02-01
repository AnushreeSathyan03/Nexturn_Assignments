// orderModel.js

const mongoose = require('mongoose');

const orderSchema = new mongoose.Schema({
  orderId: { 
    type: String,  // Change this to String if orderId is not an ObjectId
    required: true 
  },
  userId: { 
    type: String,  // Same here for userId
    required: true 
  },
  orderDate: { type: Date },
  items: [
    { 
      productId: String,  // Assuming productId should also be a string, update if necessary
      quantity: { type: Number, required: true },
      price: { type: Number, required: true }
    }
  ],
  totalAmount: { type: Number, required: true },
  status: { type: String, default: 'Pending' }
});

module.exports = mongoose.model('Order', orderSchema);
