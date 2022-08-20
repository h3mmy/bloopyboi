package providers

import (
	"fmt"
	"time"
	"context"
	"github.com/alexliesenfeld/health"
)


func onComponentStatusChanged(_ context.Context, name string, state health.CheckState) {
	logger.Println(fmt.Sprintf("component %s changed status to %s", name, state.Status))
}

func onSystemStatusChanged(_ context.Context, state health.CheckerState) {
	logger.Println(fmt.Sprintf("system status changed to %s", state.Status))
}

func onReadinessStatusChanged(_ context.Context, state health.CheckerState) {
	logger.Println(fmt.Sprintf("readiness status changed to %s", state.Status))
}

func NewReadinessChecker() health.Checker {
	return health.NewChecker(
		health.WithTimeout(10*time.Second),
		// The following check will be executed periodically every 15 seconds
		// started with an initial delay of 3 seconds. The check function will NOT
		// be executed for each HTTP request.
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name: "CustomReady",
			// If the check function returns an error, this component will be considered unavailable ("down").
			// The context contains a deadline according to the configuration of the Checker.
			Check: func(ctx context.Context) error {
				return nil
			},
		}),
		// Set a status listener that will be invoked when the health status changes.
		// More powerful hooks are also available (see docs). For guidance, please refer to the links
		// listed in the main function documentation.
		health.WithStatusListener(onReadinessStatusChanged),
	)
}
