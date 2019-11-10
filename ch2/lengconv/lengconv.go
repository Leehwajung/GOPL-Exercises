//!+

package lengconv

import "fmt"

type Feet float64
type Meter float64

func (ft Feet) String() string  { return fmt.Sprintf("%g ft", ft) }
func (m Meter) String() string  { return fmt.Sprintf("%g m", m) }

//!-
