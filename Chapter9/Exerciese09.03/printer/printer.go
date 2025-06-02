package printer

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintNewUUID() string {
	uuid := uuid.New()
	return fmt.Sprintf("Generated UUID: %s\n", uuid)
}
