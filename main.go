package main

import "strings"

var player PlayerInterface

func initGame() {
	player = &Player{
		currentRoom: "кухня",
		Exits: map[string]Room{
			"кухня": Room{
				name:        "кухня",
				description: "ты находишься на кухне",
				status:      "кухня, ничего интересного",
				exits:       []string{"коридор"},
				items: map[string][]string{
					"на столе": []string{"чай"},
				},
			},
			"коридор": Room{
				name:        "коридор",
				description: "",
				status:      "ничего интересного",
				exits:       []string{"кухня", "комната", "улица"},
				items:       map[string][]string{},
			},
			"комната": Room{
				name:        "комната",
				description: "",
				status:      "ты в своей комнате",
				exits:       []string{"коридор"},
				items: map[string][]string{
					"на столе": []string{"ключи", "конспекты"},
					"на стуле": []string{"рюкзак"},
				},
			},
			"улица": Room{
				name:        "улица",
				description: "",
				status:      "на улице весна",
				exits:       []string{"домой"},
				items:       map[string][]string{},
			},
		},
		DoorIsOpen:  false,
		HasBackPack: false,
		inventory:   []string{},
		Task:        "надо собрать рюкзак и идти в универ",
	}
}

func handleCommand(command string) string {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "осмотреться":
		return player.LookAround()
	case "идти":
		if len(parts) < 2 {
			return "неверная команда"
		}
		return player.Move(parts[1])
	case "взять":
		if len(parts) < 2 {
			return "неверная команда"
		}
		return player.TakeItem(parts[1])
	case "надеть":
		if len(parts) < 2 || parts[1] != "рюкзак" {
			return "нет такого"
		}
		return player.WearBackpack(parts[1])
	case "применить":
		if len(parts) < 3 {
			return "неверная команда"
		}
		return player.UseItem([]string{parts[1], parts[2]})
	default:
		return "неизвестная команда"
	}
}
