package main

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/vasiliy-t/gotrain/iface/mocking/mock_redis_client"
	"gopkg.in/redis.v4"
)

func TestGetValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mock_redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get("qwerty").Return(redis.NewStringResult([]byte("qwerty"), nil))

	res, err := GetValue(redisMock, "qwerty")

	if res != "qwerty" {
		t.Errorf("failed to assert that expected %s equals %s", "qwerty", res)
	}

	if err != nil {
		t.Errorf("failed to assert that err is nil, got %s", err)
	}
}

func TestSetValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mock_redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Set("qwerty", "qwerty", time.Duration(0)).Return(redis.NewStatusResult("OK", nil))

	res, err := SetValue(redisMock, "qwerty", "qwerty")

	if res != "OK" {
		t.Errorf("failed to assert that expected %s equals %s", "OK", res)
	}

	if err != nil {
		t.Errorf("failed to assert that eerr is nil, got %s", err)
	}
}
