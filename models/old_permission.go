package models

type OldPermission struct {
	Path     string           `json:"path,omitempty"`
	Name     string           `json:"name,omitempty"`
	Meta     Meta             `json:"meta,omitempty"`
	Children []*OldPermission `json:"children,omitempty"`
}

type Meta struct {
	Title string   `json:"title,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Auths []string `json:"auths,omitempty"`
	Icon  string   `json:"icon,omitempty"`
	Rank  int      `json:"rank,omitempty"`
}

func QueryOldPermissionList() (list []*OldPermission, total int64, err error) {
	list = make([]*OldPermission, 0)

	list = append(list, &OldPermission{
		Path: "/permission",
		Meta: Meta{
			Title: "menus.permission",
			Icon:  "lollipop",
			Rank:  10,
		},
		Children: []*OldPermission{
			{
				Path: "/permission/page/index",
				Name: "PermissionPage",
				Meta: Meta{
					Title: "menus.permissionPage",
					Roles: []string{"admin", "common"},
				},
				Children: nil,
			},
			{
				Path: "/permission/button/index",
				Name: "PermissionButton",
				Meta: Meta{
					Title: "menus.permissionButton",
					Roles: []string{"admin", "common"},
					Auths: []string{"btn_add", "btn_edit", "btn_delete"},
				},
				Children: nil,
			},
		},
	})

	total = int64(len(list))
	return
}

func SetEmptyArray(m *OldPermission) {
	if m.Meta.Roles == nil {
		m.Meta.Roles = make([]string, 0)
	}
	if m.Meta.Auths == nil {
		m.Meta.Auths = make([]string, 0)
	}
	if m.Children == nil {
		m.Children = make([]*OldPermission, 0)
	}
	for _, v := range m.Children {
		SetEmptyArray(v)
	}
}
