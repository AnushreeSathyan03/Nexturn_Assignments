const express = require('express');
const Product = require('../models/productModel');
 // Assuming you have a Product model

const router = express.Router();

// Get all products
router.get('/', async (req, res) => {
  try {
    const products = await Product.find();
    res.json(products);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Create a new product
router.post('/', async (req, res) => {
  const { productId, name, category, price, stock, ratings } = req.body;

  const product = new Product({
    productId,
    name,
    category,
    price,
    stock,
    ratings
  });

  try {
    await product.save();
    res.status(201).json(product);
  } catch (err) {
    res.status(400).send(err);
  }
});

// Get a specific product
router.get('/:productId', async (req, res) => {
  try {
    const product = await Product.findOne({ productId: req.params.productId });
    if (!product) return res.status(404).send('Product not found');
    res.json(product);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Update a product
router.put('/:productId', async (req, res) => {
  try {
    const updatedProduct = await Product.findOneAndUpdate(
      { productId: req.params.productId },
      req.body,
      { new: true }
    );
    if (!updatedProduct) return res.status(404).send('Product not found');
    res.json(updatedProduct);
  } catch (err) {
    res.status(400).send(err);
  }
});

// Delete a product
router.delete('/:productId', async (req, res) => {
  try {
    const deletedProduct = await Product.findOneAndDelete({ productId: req.params.productId });
    if (!deletedProduct) return res.status(404).send('Product not found');
    res.json({ message: 'Product deleted' });
  } catch (err) {
    res.status(500).send(err);
  }
});

module.exports = router;
