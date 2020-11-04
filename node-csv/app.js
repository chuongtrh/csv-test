'use strict';

const http = require('http');
const fs = require('fs');

const express = require('express');
const multer = require('multer');
const csv = require('fast-csv');

const Router = express.Router;
const upload = multer({ dest: 'tmp/csv/' });
const app = express();
const router = new Router();
const server = http.createServer(app);
const port = 9000

router.post('/', upload.single('document'), function (req, res) {
    const fileRows = [];
    // open uploaded file
    fs.createReadStream(req.file.path)
        .pipe(csv.parse({ headers: true }))
        .on("data", function (data) {
            fileRows.push(data); // push each row
        })
        .on("end", (rowCount) => {
            // remove temp file
            fs.unlinkSync(req.file.path);
            res.send(`len: ${rowCount}`)
        })
});

app.use('/upload', router);

// Start server
function startServer() {
    server.listen(port, function () {
        console.log('Express server listening on ', port);
    });
}

setImmediate(startServer);