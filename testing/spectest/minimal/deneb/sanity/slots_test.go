package sanity

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/deneb/sanity"
)

func TestMinimal_Deneb_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "minimal")
}
