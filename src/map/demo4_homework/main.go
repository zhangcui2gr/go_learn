package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	map1, ok := users[name]
	if ok {
		map1["pwd"] = "88888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = ""
		users[name]["nickname"] = ""

		// map1 = make(map[string]string)
		// map1["pwd"] = ""
		// map1["nickname"] = ""
	}

}

func main() {

	users := make(map[string]map[string]string)
	users["xiaoguan"] = make(map[string]string)
	users["xiaoguan"]["pwd"] = "77777777"

	modifyUser(users, "xiaoguan")
	fmt.Println(users)

	modifyUser(users, "xiaozhang")
	fmt.Println(users)

}
