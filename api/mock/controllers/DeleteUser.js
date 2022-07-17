'use strict';

var utils = require('../utils/writer.js');
var DeleteUser = require('../service/DeleteUserService');

module.exports.meDeletePOST = function meDeletePOST (req, res, next, body) {
  DeleteUser.meDeletePOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
