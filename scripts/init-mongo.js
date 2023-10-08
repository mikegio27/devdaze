// init-mongo.js
db.createUser({
    user: 'admin-bot',
    pwd: 'password',
    roles: [
        {
            role: 'readWrite',
            db: 'resume',
        },
    ],
});

db = new Mongo().getDB('resume');