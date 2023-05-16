package time_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeZero(t *testing.T) {
	nt := time.Time{}

	fmt.Println(nt.IsZero())
}
