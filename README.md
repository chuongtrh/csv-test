Parse csv file with Node.js & Go

## go-csv

- Fiber
- CSV encoding


``` golang
package main

import (
	"encoding/csv"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		temp, err := file.Open()
		if err != nil {
			return err
		}
		reader := csv.NewReader(temp)

		if _, err := reader.Read(); err != nil {
			panic(err)
		}

		records, err := reader.ReadAll()
		if err != nil {
			return err
		}
		return c.SendString(fmt.Sprintf("len: %d", len(records)))
	})
	log.Fatal(app.Listen("localhost:3000"))
}
```
## node-csv

- Express
- Fast-csv
- Multer

``` javascript
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
```