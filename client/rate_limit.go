package client

import "context"

import "golang.org/x/time/rate"

//rateLimit uses the routeKey of the method to check if a request is allowed according to the clients limiter.
//if no rate limit exists for the key one is created
func (c *client) rateLimit(ctx context.Context, key routeKey) error {
	limiters, ok := c.limiters[key]
	if !ok { // if limiters have not been set create them
		limiters = newLimiters(c, key)
	}
	err := limiters[0].Wait(ctx)
	if err != nil {
		return err
	}
	err = limiters[1].Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}

//newLimiters creates and sets the client's limiters for the given routeKey
func newLimiters(c *client, key routeKey) []*rate.Limiter {
	switch c.Variant() {
	case devClient:
		lim1 := rate.NewLimiter(20, 20)
		lim2 := rate.NewLimiter(5/6, 100)
		c.limiters[key] = []*rate.Limiter{lim1, lim2}
	case personalClient:
		lim1 := rate.NewLimiter(20, 20)
		lim2 := rate.NewLimiter(5/6, 100)
		c.limiters[key] = []*rate.Limiter{lim1, lim2}
	case productionClient:
		lim1 := rate.NewLimiter(300, 3000)
		lim2 := rate.NewLimiter(300, 180000)
		c.limiters[key] = []*rate.Limiter{lim1, lim2}
	}
	return c.limiters[key]
}
