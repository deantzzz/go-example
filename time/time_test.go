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

func TestName(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(time.DateTime))
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())

	fmt.Println(time.Unix(now.Unix(), 0).Format(time.DateTime))
}
