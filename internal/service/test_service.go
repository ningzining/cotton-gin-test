package service

import (
	"cotton-gin-test/api/v1/test"
)

func Sum(req test.SumReq) (int, error) {
	return req.A + req.B, nil
}
