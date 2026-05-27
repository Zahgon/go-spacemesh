package config

func MainnetConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// July 15, 2024, 10:00:00 AM UTC

// starting from epoch 13 activesets below 12 will be pruned

// https://github.com/spacemeshos/go-spacemesh/issues/4559
// 3000 of spends

// NOTE(dshulyak) this is intentional. we increased committee size with hare3 upgrade
// but certifier continues to use 200 committee size.
// this will be upgraded in future with scheduled upgrade.

// RequestTimeout = RequestRetryDelay * 2 * MaxRequestRetries*(MaxRequestRetries+1)/2

// 3h
