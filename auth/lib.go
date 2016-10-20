package auth

import api "github.com/bep/kittn/api"

func Authorize(secret string) *api.KittnStore {
	return new(api.KittnStore)
}
