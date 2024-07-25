package service

import (
	"gin-casbin-admin/internal/model"
	"testing"
)

func Test_buildTree(t *testing.T) {

	permissions := []*model.AdminPermissions{
		{
			Id:   1,
			Name: "权限1",
			PId:  0,
		},
		{
			Id:   2,
			Name: "权限2",
			PId:  0,
		},
		{
			Id:   3,
			Name: "权限3",
			PId:  1,
		},
		{
			Id:   4,
			Name: "权限4",
			PId:  1,
		},
		{
			Id:   5,
			Name: "权限5",
			PId:  2,
		},
		{
			Id:   6,
			Name: "权限6",
			PId:  2,
		},
	}

	permissionsTree := buildTree(permissions, 0)

	t.Log(permissionsTree)
}
