'use strict';

var utils = require('../utils/writer.js');
var FetchAllItems = require('../service/FetchAllItemsService');

module.exports.itemsGET = function itemsGET (req, res, next) {
  FetchAllItems.itemsGET()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
