package service

import (
	"fmt"
	"main/model"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	Clear()
}

type UserService struct {
	cache Cache
}

func NewUserService(cache Cache) *UserService {
	return &UserService{
		cache: cache,
	}
}

func (u UserService) CompareAndClassify(users []model.User) []model.CompareResult {
	res := []model.CompareResult{}
	for i, originalUser := range users {
		for j, comparingUser := range users {
			// remove same user comparision
			if i == j {
				continue
			}
			//check if comparission was already made
			if val, ok := u.cache.Get(u.crateUserCacheKey(originalUser.ID, comparingUser.ID)); ok {
				res = append(res, val.(model.CompareResult))
				continue
			}
			score := u.classifyScore(originalUser.GetScore(comparingUser))
			result := model.CompareResult{
				ContactIDSource: originalUser.ID,
				ContactIDMatch:  comparingUser.ID,
				Accuracy:        score,
			}
			res = append(res, result)
			//storing in cache users so we avoid comparing them again
			u.cache.Set(u.crateUserCacheKey(originalUser.ID, comparingUser.ID), result, time.Minute*5)
		}
	}
	return res
}

func (u UserService) crateUserCacheKey(i1, i2 int) string {
	// Ensure the smallest number is first
	if i1 > i2 {
		i1, i2 = i2, i1
	}
	return fmt.Sprintf("%d-%d", i1, i2)
}

func (u UserService) classifyScore(score int) string {
	switch {
	case score > 8:
		return "HIGH"
	case score > 4:
		return "MID"
	default:
		return "LOW"
	}
}
