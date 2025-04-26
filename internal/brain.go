package methodius

import (
	"fmt"
	"math/rand"
	"net/http"
)

const (
	UsageKey   = "status"
	UsageValue = "smiling"
)

func PrintUsage(methodToAction map[string]string, port int) {
	curls := make(map[string]string)
	mapping := "Use "

	for method, action := range methodToAction {
		if method == MethodQuit {
			continue
		}

		mapping += fmt.Sprintf("%s to %s, ", method, action)

		path := "/"
		if action == ActionRead || action == ActionCreate || action == ActionUpdate || action == ActionRemove {
			path += UsageKey
		}

		data := ""
		if action == ActionCreate || action == ActionUpdate {
			data = " -d '" + UsageValue + "'"
		}

		curls[action] = fmt.Sprintf("curl -X %s 'localhost:%d%s'%s", method, port, path, data)
	}

	fmt.Println(mapping + "and QUIT to... quit?")
	for action, curl := range curls {
		fmt.Printf(curl+" # %s\n", action)
	}
}

func GetMethodMaps(getStatic bool) (methodToAction map[string]string) {
	methodToAction = map[string]string{
		http.MethodGet:    ActionRead, // if it's List will be decided inside
		http.MethodPost:   ActionCreate,
		http.MethodPut:    ActionUpdate,
		http.MethodDelete: ActionRemove,
	}

	if !getStatic {
		methodToAction = getRandomMethodMaps(methodToAction)
	}

	methodToAction[MethodQuit] = ActionQuit

	return
}

func getRandomMethodMaps(static map[string]string) (methodToAction map[string]string) {
	methodToAction = make(map[string]string)

	var methods []string
	var actions []string

	for k, v := range static {
		methods = append(methods, k)
		actions = append(actions, v)
	}

	rand.Shuffle(len(actions), func(i, j int) {
		actions[i], actions[j] = actions[j], actions[i]
	})

	var i int
	for _, method := range methods {
		methodToAction[method] = actions[i]
		i++
	}

	matches := 0
	for method, action := range static {
		if methodToAction[method] == action {
			matches++
		}
	}

	if matches > 2 {
		return getRandomMethodMaps(static)
	}

	return
}
