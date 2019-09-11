package main

import "github.com/eddieowens/simaia/app"

func main() {
	_ = app.CreateInjector().GetStructPtr(app.Key).(app.App).Start()
}
