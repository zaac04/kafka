package redis

import (
	"anya/structs"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func (c *Cache) SetJSON(data *structs.SetCache) (err error) {
	err = c.client.JSONSet(data.Ctx, data.Key, "$", data.Val).Err()
	if err != nil {
		return err
	}
	err = c.client.Expire(data.Ctx, data.Key, data.Expiry).Err()
	return err
}

func (c *Cache) GetJSON(param *structs.GetCache) (res []byte, err error) {
	data, err := c.client.JSONGet(param.Ctx, param.Key, "$").Result()
	if err != nil {
		return nil, err
	}
	str := strings.Trim(data, "[]")
	return []byte(str), nil
}

func (c *Cache) SetHMap(dat *structs.SetCache) error {
	err := c.client.HSet(dat.Ctx, dat.Key, dat.Val).Err()
	if err != nil {
		return err
	}
	err = c.client.Expire(dat.Ctx, dat.Key, dat.Expiry).Err()
	return err
}

func (c *Cache) GetHMap(dat *structs.GetCache) (map[string]string, error) {
	val, err := c.client.HGetAll(dat.Ctx, dat.Key).Result()
	return val, err
}
func (c *Cache) SetList(dat *structs.SetCache) error {
	err := c.client.RPush(dat.Ctx, dat.Key, dat.Val).Err()
	if err != nil {
		return err
	}
	err = c.client.Expire(dat.Ctx, dat.Key, dat.Expiry).Err()
	return err
}
func (c *Cache) GetList(dat *structs.GetCache) ([]string, error) {
	val, err := c.client.LRange(dat.Ctx, dat.Key, 0, -1).Result()
	return val, err
}
func (c *Cache) RemoveListElement(dat *structs.SetCache) error {
	err := c.client.LRem(dat.Ctx, dat.Key, 1, dat.Val).Err()
	return err
}
