package menu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/menu"
)

func TestBuildMenuTree(t *testing.T) {
	// Setup test data
	menus := []*SysMenuResponse{
		{ID: 1, ParentID: 0, MenuType: menu.TypeMenu, Title: "System", Auths: ""},
		{ID: 2, ParentID: 1, MenuType: menu.TypeMenu, Title: "User", Auths: ""},
		{ID: 3, ParentID: 2, MenuType: menu.TypeButton, Title: "Add User", Auths: "sys:user:add"},
		{ID: 4, ParentID: 2, MenuType: menu.TypeButton, Title: "Edit User", Auths: "sys:user:edit"},
		{ID: 5, ParentID: 1, MenuType: menu.TypeMenu, Title: "Role", Auths: ""},
		{ID: 6, ParentID: 5, MenuType: menu.TypeButton, Title: "Delete Role", Auths: "sys:role:delete"},
	}

	// Execute
	tree := buildMenuTree(menus)

	// Verify
	assert.Equal(t, 1, len(tree))
	systemMenu := tree[0]
	assert.Equal(t, "System", systemMenu.Title)
	assert.Equal(t, 2, len(systemMenu.Children))

	// Check User Menu
	userMenu := systemMenu.Children[0]
	assert.Equal(t, "User", userMenu.Title)
	assert.Equal(t, 0, len(userMenu.Children), "Buttons should not be in Children")
	assert.Contains(t, userMenu.Auths, "sys:user:add")
	assert.Contains(t, userMenu.Auths, "sys:user:edit")
	assert.Contains(t, userMenu.Auths, ",")

	// Check Role Menu
	roleMenu := systemMenu.Children[1]
	assert.Equal(t, "Role", roleMenu.Title)
	assert.Equal(t, 0, len(roleMenu.Children), "Buttons should not be in Children")
	assert.Equal(t, "sys:role:delete", roleMenu.Auths)
}
