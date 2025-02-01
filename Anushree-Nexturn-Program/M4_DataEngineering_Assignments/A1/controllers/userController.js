const express = require('express');
const User = require('../models/userModel'); // Assuming you have a User model

const router = express.Router();

// Get all users
router.get('/', async (req, res) => {
  try {
    const users = await User.find();
    res.json(users);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Create a new user
router.post('/', async (req, res) => {
  const { userId, name, email, age, address, createdAt } = req.body;

  const user = new User({
    userId,
    name,
    email,
    age,
    address,
    createdAt
  });

  try {
    await user.save();
    res.status(201).json(user);
  } catch (err) {
    res.status(400).send(err);
  }
});

// Get a specific user
router.get('/:userId', async (req, res) => {
  try {
    const user = await User.findOne({ userId: req.params.userId });
    if (!user) return res.status(404).send('User not found');
    res.json(user);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Update a user
router.put('/:userId', async (req, res) => {
  try {
    const updatedUser = await User.findOneAndUpdate(
      { userId: req.params.userId },
      req.body,
      { new: true }
    );
    if (!updatedUser) return res.status(404).send('User not found');
    res.json(updatedUser);
  } catch (err) {
    res.status(400).send(err);
  }
});

// Delete a user
router.delete('/:userId', async (req, res) => {
  try {
    const deletedUser = await User.findOneAndDelete({ userId: req.params.userId });
    if (!deletedUser) return res.status(404).send('User not found');
    res.json({ message: 'User deleted' });
  } catch (err) {
    res.status(500).send(err);
  }
});

module.exports = router;
