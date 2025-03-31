package main

type PlayerInterface interface {
	LookAround() string
	Move(direction string) string
	TakeItem(item string) string
	WearBackpack(item string) string
	UseItem(data []string) string
}

//func NewPlayer(rooms []*Room) PlayerInterface {
//	return &Player{
//		Exits: map[string]Room{},
//	}
//}

func (p *Player) LookAround() string {
	var result string
	room := p.Exits[p.currentRoom]
	result += room.GetAllItems()
	if room.description != "" {
		result = room.description + ", " + result + ", " + p.Task
	}
	result += ". " + room.GetAllExits()
	return result
}

func (p *Player) Move(direction string) string {
	result := ""
	room := p.Exits[p.currentRoom]
	if _, ok := p.Exits[direction]; ok {
		if contains(room.exits, direction) {
			if direction == "улица" && p.DoorIsOpen == false {
				return "дверь закрыта"
			}
			p.currentRoom = direction
			room := p.Exits[p.currentRoom]
			result = room.status + ". " + room.GetAllExits()
		} else {
			result = "нет пути в " + direction
		}
	}
	return result
}

func (p *Player) TakeItem(item string) string {
	result := ""
	room := p.Exits[p.currentRoom]
	if room.CheckItem(item) {
		if p.HasBackPack {
			p.inventory = append(p.inventory, item)
			room.DeleteItem(item)
			result = "предмет добавлен в инвентарь: " + item
		} else {
			result = "некуда класть"
		}
	} else {
		result = "нет такого"
	}
	return result
}

func (p *Player) WearBackpack(item string) string {
	result := "нет такого"
	if item == "рюкзак" && p.HasBackPack == false {
		room := p.Exits[p.currentRoom]
		if room.CheckItem(item) {
			p.HasBackPack = true
			p.Task = "надо идти в универ"
			room.DeleteItem(item)
			result = "вы надели: " + item
		}
	}
	return result
}

func (p *Player) UseItem(data []string) string {
	result := ""
	key := data[0]
	to := data[1]
	if contains(p.inventory, key) {
		if key == "ключи" && to == "дверь" {
			p.DoorIsOpen = true
			result = "дверь открыта"
		} else {
			result = "не к чему применить"
		}
	} else {
		result = "нет предмета в инвентаре - " + key
	}
	return result
}
