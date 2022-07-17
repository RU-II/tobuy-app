'use strict';

var utils = require('../utils/writer.js');
var UpdatePassword = require('../service/UpdatePasswordService');

module.exports.meUpdatePasswordPOST = function meUpdatePasswordPOST (req, res, next, body) {
  UpdatePassword.meUpdatePasswordPOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
