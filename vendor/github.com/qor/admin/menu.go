package admin

import (
	"path"

	"github.com/qor/qor"
	"github.com/qor/roles"
)

// Menu qor admin sidebar menus definiation
type Menu struct {
	Name       string
	Link       string
	Ancestors  []string
	Permission *roles.Permission
	subMenus   []*Menu
	rawPath    string
}

func (menu Menu) HasPermission(mode roles.PermissionMode, context *qor.Context) bool {
	if menu.Permission == nil {
		return true
	}
	return menu.Permission.HasPermission(mode, context.Roles...)
}

// GetMenus get menus for admin sidebar
func (admin Admin) GetMenus() []*Menu {
	return admin.menus
}

// GetSubMenus get submenus for a menu
func (menu *Menu) GetSubMenus() []*Menu {
	return menu.subMenus
}

// AddMenu add a menu to admin
func (admin *Admin) AddMenu(menu *Menu) {
	admin.menus = appendMenu(admin.menus, menu.Ancestors, menu)
}

// GetMenu get menu with name from admin
func (admin Admin) GetMenu(name string) *Menu {
	return getMenu(admin.menus, name)
}

func getMenu(menus []*Menu, name string) *Menu {
	for _, m := range menus {
		if m.Name == name {
			return m
		}

		if len(m.subMenus) > 0 {
			if mc := getMenu(m.subMenus, name); mc != nil {
				return mc
			}
		}
	}

	return nil
}

// Generate menu links by current route. e.g "/products" to "/admin/products"
func (admin *Admin) generateMenuLinks() {
	prefixMenuLinks(admin.menus, admin.router.Prefix)
}

func prefixMenuLinks(menus []*Menu, prefix string) {
	for _, m := range menus {
		if m.rawPath != "" {
			m.Link = path.Join(prefix, m.rawPath)
		}
		if len(m.subMenus) > 0 {
			prefixMenuLinks(m.subMenus, prefix)
		}
	}
}

func newMenu(menus []string, menu *Menu) *Menu {
	menuCount := len(menus)
	for index := range menus {
		menu = &Menu{Name: menus[menuCount-index-1], subMenus: []*Menu{menu}}
	}

	return menu
}

func appendMenu(menus []*Menu, ancestors []string, menu *Menu) []*Menu {
	if len(ancestors) > 0 {
		for _, m := range menus {
			if m.Name != ancestors[0] {
				continue
			}

			if len(ancestors) > 1 {
				m.subMenus = appendMenu(m.subMenus, ancestors[1:], menu)
			} else {
				m.subMenus = append(m.subMenus, newMenu(nil, menu))
			}
			return menus
		}
	}

	return append(menus, newMenu(ancestors, menu))
}
