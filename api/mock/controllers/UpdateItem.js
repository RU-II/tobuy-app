'use strict';

var utils = require('../utils/writer.js');
var UpdateItem = require('../service/UpdateItemService');

module.exports.itemsIdPOST = function itemsIdPOST (req, res, next, body) {
  UpdateItem.itemsIdPOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
