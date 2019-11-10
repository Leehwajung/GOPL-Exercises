//!+

package wghtconv

func LbsToKg(lbs Pound) Kilogram { return Kilogram(lbs * 0.45359237) }

func KgToLbs(kg Kilogram) Pound { return Pound(kg / 0.45359237) }

//!-
