'use strict';

var utils = require('../utils/writer.js');
var FetchItemBuyId = require('../service/FetchItemBuyIdService');

module.exports.itemsIdGET = function itemsIdGET (req, res, next, id) {
  FetchItemBuyId.itemsIdGET(id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
