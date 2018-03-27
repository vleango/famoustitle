var params = {
    TableName: 'users',
    Item: { // a map of attribute name to AttributeValue
        // attribute_value (string | number | boolean | null | Binary | DynamoDBSet | Array | Object)
        // more attributes...
        id: "aaaa",
        first_name: "Tha",
        last_name: "Leang",
        email: "tha@test.com",
        password: 'hogehoge',
        admin: true,
        logins: [
          { provider: 'web', token: 'abcd', refresh_token: '1234' }
        ]
    }
};
docClient.put(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});
