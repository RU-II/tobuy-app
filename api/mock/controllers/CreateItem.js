'use strict';

var utils = require('../utils/writer.js');
var CreateItem = require('../service/CreateItemService');

module.exports.itemsPOST = function itemsPOST (req, res, next, body) {
  CreateItem.itemsPOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
