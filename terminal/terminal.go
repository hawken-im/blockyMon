package terminal

import (
	"fmt"
)

func Input() {
	for {
		var action string
		fmt.Print("Enter your action: ")
		fmt.Scanf("%s", &action)

		switch action {
		case "pet":
			fmt.Println(action)
		case "clean":
			fmt.Println(action)
		case "feed":
			fmt.Println(action)
		case "heal":
			fmt.Println(action)
		default:
			fmt.Println("unknown action")
		}
	}
}
