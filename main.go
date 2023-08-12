package main

import (
	"flag"
	"inmem-db/store"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var initialSize int64
var compress bool
var hashtable store.HashTable

func init() {
	sizeArg := flag.Int("size", 1000, "How many Entries Memory DB should allow, Default is 1000")
	compressArg := flag.Bool("compress", false, "Compress Values sent")

	flag.Parse()

	initialSize = int64(*sizeArg)
	compress = *compressArg

	hashtable = *store.NewHashTable(initialSize, compress)

}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:           true,
		ReduceMemoryUsage: true,
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		v := new(Data)
		err := c.BodyParser(v)
		if err != nil {
			return err
		}
		if v.Key == "" {
			return c.Status(http.StatusFailedDependency).JSON(map[string]any{
				"message": "Failed to Read Request",
				"status":  false,
			})
		}
		if v.Value == nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]any{
				"message": "Value not Provided",
				"status":  false,
			})
		}
		err2 := hashtable.Add(v.Key, v.Value)
		if err2 != nil {
			return c.Status(http.StatusExpectationFailed).JSON(map[string]any{
				"message": "Failed to Add Data",
				"status":  false,
			})
		}
		return c.Status(http.StatusCreated).JSON(map[string]any{
			"message": "Successfully Added Data",
			"status":  true,
		})

	})

	app.Post("/get", func(c *fiber.Ctx) error {
		v := new(Data)
		err := c.BodyParser(v)
		if err != nil {
			return c.Status(http.StatusFailedDependency).JSON(map[string]any{
				"message": "Failed to Read Request",
				"status":  false,
			})
		}
		if v.Key == "" {
			return c.Status(http.StatusBadRequest).JSON(map[string]any{
				"message": "Key not Provided",
				"status":  false,
			})
		}
		val, state := hashtable.Get(v.Key)

		if !state {
			return c.Status(http.StatusExpectationFailed).JSON(map[string]any{
				"message": "No Such Key in DB",
				"status":  true,
			})
		}
		return c.Status(http.StatusAccepted).JSON(map[string]any{
			"data":   val,
			"status": true,
		})

	})

	app.Delete("/del", func(c *fiber.Ctx) error {
		v := new(Data)
		err := c.BodyParser(v)
		if err != nil {
			return c.Status(http.StatusFailedDependency).JSON(map[string]any{
				"message": "Failed to Read Request",
				"status":  false,
			})
		}
		if v.Key == "" {
			return c.Status(http.StatusBadRequest).JSON(map[string]any{
				"message": "Key not Provided",
				"status":  false,
			})
		}
		hashtable.Del(v.Key)
		return c.Status(http.StatusAccepted).JSON(map[string]any{
			"message": "Successfully Deleted Data",
			"status":  true,
		})
	})

	app.Listen("localhost:6000")
}

type Data struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}
