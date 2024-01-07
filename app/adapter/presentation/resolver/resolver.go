package resolver

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/yuorei/video-server/app/application"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	application *application.Application
	usecase     *application.UseCase
}

func NewResolver(app *application.Application) *Resolver {
	return &Resolver{
		application: app,
		usecase:     application.NewUseCase(app),
	}
}
