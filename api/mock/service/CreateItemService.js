'use strict';


/**
 * Create item
 * Create user's item
 *
 * body Models.MutationItemRequest Create Item Request
 * returns models.ItemResponse
 **/
exports.itemsPOST = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "item" : {
    "number" : 1,
    "category_id" : 1,
    "updated_at" : "2022-05-01T17:23:17.494039+09:00",
    "name" : "test item",
    "created_at" : "2022-05-01T17:23:17.494039+09:00",
    "description" : "This is a test item",
    "counter" : "数",
    "id" : 1,
    "deleted_at" : "",
    "status" : 0
  }
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

