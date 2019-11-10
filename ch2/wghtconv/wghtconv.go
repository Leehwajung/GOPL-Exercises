//!+

package wghtconv

import "fmt"

type Pound float64
type Kilogram float64

func (lbs Pound) String() string    { return fmt.Sprintf("%g lbs", lbs) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%g kg", kg) }

//!-
