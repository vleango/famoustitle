var params = {
    TableName: 'tech_writer_users',
    Item: { // a map of attribute name to AttributeValue
        // attribute_value (string | number | boolean | null | Binary | DynamoDBSet | Array | Object)
        // more attributes...
        id: "5f7d8353-f976-4737-a6cc-65bd0df6867d",
        first_name: "Tha",
        last_name: "Leang",
        email: "tha@test.com",
        password_digest: '$2a$10$sLJXlfwvwJgYEx102y.dPO3OzkR3K..LFt2Tl/lHXTIff6xrB8oOO',
        admin: true
    }
};
docClient.put(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});
