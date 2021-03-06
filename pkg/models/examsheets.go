package models

type Examsheets struct {
	Data []Examsheet `json:"data"`
}

type Examsheet struct {
	IDExam       *uint32     	`json:"id_exam"`
	IDUser       *uint32       	`json:"id_user"`
	Answersheets []Answersheet 	`json:"answersheets"`
}

type Answersheet struct {
	IDSession *uint32   `json:"id_session"`
	Answers   Answers 	`json:"answers"`
}

type Answers struct {
	Num1  *bool `json:"1"`
	Num2  *bool `json:"2"`
	Num3  *bool `json:"3"`
	Num4  *bool `json:"4"`
	Num5  *bool `json:"5"`
	Num6  *bool `json:"6"`
	Num7  *bool `json:"7"`
	Num8  *bool `json:"8"`
	Num9  *bool `json:"9"`
	Num10 *bool `json:"10"`
}
