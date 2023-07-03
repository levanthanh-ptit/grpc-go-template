package cmd

import "context"

func main() {
	ctx := context.Background()
	app, err := NewApp(ctx)
	if err != nil {
		panic(err)
	}
	app.Start(ctx)
}
