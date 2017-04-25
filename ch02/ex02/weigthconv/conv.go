package weigthconv

// KToP converts a Kilogram temperature to Pound.
func KToP(k Kilogram) Pound { return Pound(k * 2.2046) }

// PToK converts a Pound temperature to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
