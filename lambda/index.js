var mysql = require('mysql');

var con = mysql.createConnection({
    host: "mariadb",
    user: "root",
    password: "password",
    database: "database"
});

exports.handler =  async function(event, context) {    

    con.connect(function(err) {
        if (err) throw err;
        console.log("Connected!");
    });

    return {
        statusCode: 200,
        body: `Received: ${event.message}`
    };
}
