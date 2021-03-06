'use strict';

var utils = require('../utils/writer.js');
var UpdateItem = require('../service/UpdateItemService');

module.exports.itemsIdUpdatePOST = function itemsIdUpdatePOST (req, res, next, body, id) {
  UpdateItem.itemsIdUpdatePOST(body, id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
