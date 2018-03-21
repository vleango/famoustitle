var params = {
    TableName: 'articles',
    Item: { // a map of attribute name to AttributeValue
        // attribute_value (string | number | boolean | null | Binary | DynamoDBSet | Array | Object)
        // more attributes...
        id: "123",
        author: "tha",
        title: "this is my title",
        body: "this is my body"
    }
};
docClient.put(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});

var params = {
    TableName: 'articles',
    Item: { // a map of attribute name to AttributeValue
        // attribute_value (string | number | boolean | null | Binary | DynamoDBSet | Array | Object)
        // more attributes...
        id: "456",
        author: "leang",
        title: "this is my title 2",
        body: "this is my body 2"
    }
};
docClient.put(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});
