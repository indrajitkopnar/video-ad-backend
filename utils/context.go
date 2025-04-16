// utils/context.go
package utils

import "context"

// Ctx is a global background context used by services like Kafka
var Ctx = context.Background()
