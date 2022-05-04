'use strict';


/**
 * SignUp
 * SignUp
 *
 * body Models.SignUpRequest SignUp Request
 * returns models.AuthResponse
 **/
exports.signupPOST = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "user" : {
    "updated_at" : "2022-05-01T17:23:17.494039+09:00",
    "name" : "test user",
    "created_at" : "2022-05-01T17:23:17.494039+09:00",
    "id" : 1,
    "deleted_at" : "",
    "email" : "test@example.com"
  },
  "token" : "ajiji1j98a"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

