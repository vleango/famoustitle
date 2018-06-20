// get all items
var params = {
    TableName: 'famoustitle_users'
};
dynamodb.scan(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});


// update item
var params = {
    TableName: 'famoustitle_users',
    Key: {
        email: "tha@test.com"
    },
    UpdateExpression: 'SET is_writer = :is_writer',
    ExpressionAttributeValues: {
        ':is_writer': true
    },
    ReturnValues: 'ALL_NEW'
};
docClient.update(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});


// delete item
var params = {
    TableName: 'famoustitle_users',
    Key: {
        email: "tha@test.com"
    },
    ReturnValues: 'ALL_OLD'
};
docClient.delete(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});
