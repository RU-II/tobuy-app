'use strict';

var utils = require('../utils/writer.js');
var DeleteItem = require('../service/DeleteItemService');

module.exports.itemsIdDeletePOST = function itemsIdDeletePOST (req, res, next, id) {
  DeleteItem.itemsIdDeletePOST(id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
