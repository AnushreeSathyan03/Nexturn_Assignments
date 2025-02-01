const express = require('express');
const Order = require('../models/orderModel'); // Assuming you have an Order model

const router = express.Router();

// Get all orders
router.get('/', async (req, res) => {
  try {
    const orders = await Order.find();
    res.json(orders);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Create a new order
router.post('/', async (req, res) => {
  const { orderId, userId, orderDate, items, totalAmount, status } = req.body;

  const order = new Order({
    orderId,   // No need to convert to ObjectId if it's a string
    userId,    // Same here for userId
    orderDate,
    items,
    totalAmount,
    status
  });

  try {
    await order.save();
    res.status(201).json(order);
  } catch (err) {
    res.status(400).send(err);
  }
});


// Get a specific order
router.get('/:orderId', async (req, res) => {
  try {
    const order = await Order.findOne({ orderId: req.params.orderId });
    if (!order) return res.status(404).send('Order not found');
    res.json(order);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Update an order
router.put('/:orderId', async (req, res) => {
  try {
    const updatedOrder = await Order.findOneAndUpdate(
      { orderId: req.params.orderId },
      req.body,
      { new: true }
    );
    if (!updatedOrder) return res.status(404).send('Order not found');
    res.json(updatedOrder);
  } catch (err) {
    res.status(400).send(err);
  }
});

// Delete an order
router.delete('/:orderId', async (req, res) => {
  try {
    const deletedOrder = await Order.findOneAndDelete({ orderId: req.params.orderId });
    if (!deletedOrder) return res.status(404).send('Order not found');
    res.json({ message: 'Order deleted' });
  } catch (err) {
    res.status(500).send(err);
  }
});

module.exports = router;
