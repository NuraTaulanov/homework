package main

type RoomInterface interface {
	GetAllExits() string
	GetAllItems() string
	CheckItem(item string) bool
	DeleteItem(item string)
}

//	func NewRoom(name, description, status string, items map[string][]string, exits []string) *Room {
//		room := &Room{
//			name:        name,
//			description: description,
//			status:      status,
//			items:       items,
//			exits:       exits,
//		}
//		return room
//	}
func (r Room) GetAllExits() string {
	allExits := ""
	if len(r.exits) > 0 {
		allExits = "можно пройти - "
		for i, exit := range r.exits {
			if i > 0 {
				allExits += ", "
			}
			allExits += exit
		}
	}
	return allExits
}

func (r Room) GetAllItems() string {
	var allItems string
	if len(r.items) > 0 {
		var a int
		for place, items := range r.items {
			if a++; a > 1 {
				allItems += ", "
			}
			allItems = allItems + place + ": "
			for index, item := range items {
				if index > 0 {
					allItems += ", "
				}
				allItems += item
			}
		}
	} else {
		switch r.name {
		case "кухня", "комната", "улица":
			allItems += "пустая " + r.name
		case "коридор":
			allItems += "пустой " + r.name
		}
	}
	return allItems
}

func (r Room) CheckItem(item string) bool {
	for _, i := range r.items {
		for _, j := range i {
			if j == item {
				return true
			}
		}
	}
	return false
}

func (r Room) DeleteItem(item string) {
	for key, value := range r.items {
		for index, i := range value {
			if item == i {
				r.items[key] = append(value[:index], value[index+1:]...)
				if len(r.items[key]) == 0 {
					delete(r.items, key)
				}
			}
		}
	}
}
