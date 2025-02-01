const mongoose = require('mongoose');
const Warehouse = require('./models/warehouseModel');

async function createGeoIndex() {
  await mongoose.connect('mongodb://localhost:27017/onlineShopping');
  await Warehouse.createIndexes();
  console.log('Geospatial index created on location field');
}

createGeoIndex().catch(err => console.error(err));
