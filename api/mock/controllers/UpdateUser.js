'use strict';

var utils = require('../utils/writer.js');
var UpdateUser = require('../service/UpdateUserService');

module.exports.meUpdatePOST = function meUpdatePOST (req, res, next, body) {
  UpdateUser.meUpdatePOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
