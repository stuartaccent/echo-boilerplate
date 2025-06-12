package lib

type DropdownMenuItemType string

const (
	DropdownMenuItemTypeLabel     DropdownMenuItemType = "label"
	DropdownMenuItemTypeSeparator DropdownMenuItemType = "separator"
	DropdownMenuItemTypeItem      DropdownMenuItemType = "item"
)

type DropdownMenuItem struct {
	Type DropdownMenuItemType `json:"type"`
	Text string               `json:"text,omitempty"`
	Href string               `json:"href,omitempty"`
}
