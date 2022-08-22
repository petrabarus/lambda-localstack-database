const pg = require('pg');
const R = require('ramda');

const db_dsn = 'postgres://postgres:postgres@postgres:5432/postgres';

exports.handler =  async function(event, context) {    
    const client = new pg.Client(db_dsn);
    client.connect();
    
    const { rows  } = await client.query('SELECT 1 + 4');
    const result = R.head(R.values(R.head(rows)));
    client.end();

    return {
        statusCode: 200,
        body: `Received: ${event.message}. Result: ${result}`
    };
}
