package roman

type Roman struct {
  Number int
}

func (roman Roman) GetI() string {
  var I string
  breaking := roman.Number % 5

  for i:= 0; i < breaking; i++ {
	I += "I"	
  }
return I
}

func (roman Roman) GetRoman() string {
  var romanMap = map[int]string{
    1 : "V",
    2 : "X",
  }
  var romanNumber string
    var mapIndex int
    mapIndex = roman.Number / 5
    if (roman.Number % 5) == 4 {
	romanNumber = "I"+romanMap[mapIndex+1]
    }else{
	romanNumber = romanMap[mapIndex] 
	romanNumber+=roman.GetI()
    }
  return romanNumber 
}

