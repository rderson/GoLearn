package conv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func MToFt(m Meters) Feet { return Feet(m * 3.28) }

func FtToM(f Feet) Meters { return Meters(f / 3.28) }

func KToP(k Kilos) Pounds { return Pounds(k * 2.205) }

func PToK(p Pounds) Kilos { return Kilos(p / 2.205) }
