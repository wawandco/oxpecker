package buffalo

import "context"

type Initializer struct{}

func (i Initializer) Name() string {
	return "buffalo/initializer"
}

// Creates dirs:
// name/app
// name/app/actions
// name/app/assets
// name/app/middleware
// name/app/models
// name/app/render
// name/app/render/helpers
// name/app/tasks
// name/app/templates
// public
// db
// db/migrations
func (i *Initializer) Initialize(ctx context.Context, root string) error {

	return nil
}
