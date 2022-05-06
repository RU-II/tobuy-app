'use strict';

var utils = require('../utils/writer.js');
var SignUp = require('../service/SignUpService');

module.exports.signupPOST = function signupPOST (req, res, next, body) {
  SignUp.signupPOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
