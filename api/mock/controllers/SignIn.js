'use strict';

var utils = require('../utils/writer.js');
var SignIn = require('../service/SignInService');

module.exports.signinPOST = function signinPOST (req, res, next, body) {
  SignIn.signinPOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
