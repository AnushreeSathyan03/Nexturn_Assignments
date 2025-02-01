const mongoose = require('mongoose');

const addressSchema = new mongoose.Schema({
  city: String,
  state: String,
  zip: String
}, { _id: false });

const userSchema = new mongoose.Schema({
  userId: { type: String, required: true, unique: true },
  name: { type: String, required: true },
  email: { type: String, required: true, unique: true },
  age: { type: Number, required: true },
  address: addressSchema,
  createdAt: { type: Date, default: Date.now }
});

const User = mongoose.model('User', userSchema);

// Sample User Data
const sampleUsers = [
  {
    userId: "U001",
    name: "John Doe",
    email: "john.doe@example.com",
    age: 28,
    address: { city: "New York", state: "NY", zip: "10001" },
    createdAt: "2024-01-01T10:00:00Z"
  },
  {
    userId: "U002",
    name: "Jane Smith",
    email: "jane.smith@example.com",
    age: 32,
    address: { city: "Los Angeles", state: "CA", zip: "90001" },
    createdAt: "2023-06-15T08:30:00Z"
  },
  {
    userId: "U003",
    name: "Michael Johnson",
    email: "michael.johnson@example.com",
    age: 24,
    address: { city: "Chicago", state: "IL", zip: "60601" },
    createdAt: "2024-03-10T11:45:00Z"
  }
];

// Insert sample users
User.insertMany(sampleUsers)
  .then(() => console.log('Sample Users added'))
  .catch(err => console.error('Error inserting users:', err));

module.exports = User;
