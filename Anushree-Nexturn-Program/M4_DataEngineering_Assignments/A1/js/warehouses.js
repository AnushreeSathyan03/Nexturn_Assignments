const mongoose = require('mongoose');

const warehouseSchema = new mongoose.Schema({
  warehouseId: { type: String, required: true },
  location: {
    type: { type: String, enum: ['Point'], required: true },
    coordinates: { type: [Number], required: true },
  },
  products: [String], // Array of product IDs available in the warehouse
});

warehouseSchema.index({ location: '2dsphere' });

module.exports = mongoose.model('Warehouse', warehouseSchema);
