//!+

package lengconv

func FtToM(ft Feet) Meter { return Meter(ft * 0.3048) }

func MToFt(m Meter) Feet { return Feet(m / 0.3048) }

//!-
