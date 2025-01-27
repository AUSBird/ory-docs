package session

import (
	"context"
)

func DisableSession(ctx context.Context, sessionId string) (err error) {
	// highlight-start
	_, err = ory.IdentityApi.DisableSession(ContextWithToken(ctx), sessionId).
		Execute()
	// highlight-end
	return err
}
