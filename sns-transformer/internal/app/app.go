package app

import "context"

type Application struct {
}

func (p *Application) Init() error {
	return nil
}

func (p *Application) Fini() error {
	return nil
}

func (p *Application) Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
