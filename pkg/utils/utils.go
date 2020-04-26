package utils

import (
	"io"
	"fmt"
	"errors"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/sy-anwar/grader-ranker-server/pkg/models"
)

type MalformedRequest struct {
    Status int
    Msg	   string
}

func (mr *MalformedRequest) Error() string {
    return mr.Msg
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
    
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    err := dec.Decode(&dst)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError

        switch {
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

        default:
            return err
        }
    }

    if dec.More() {
        msg := "Request body must only contain a single JSON object"
        return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
    }

    return nil
}

func CheckNilFieldsExamsheets(examsheets models.Examsheets) bool {
	if examsheets.Data == nil {
		return true
	}
	for i, _ := range examsheets.Data {
		if CheckNilFieldsExamsheet(examsheets.Data[i]) {
			return true
		}
	}
	return false
}

func CheckNilFieldsExamsheet(examsheet models.Examsheet) bool {
	if examsheet.IDExam == nil {
		return true
	}
	if examsheet.IDUser == nil {
		return true
	}
	if examsheet.Answersheets == nil {
		return true
	}
	for i, _ := range examsheet.Answersheets {
		if CheckNilFieldsAnswersheets(examsheet.Answersheets[i]) {
			return true
		}
	}
	return false
}

func CheckNilFieldsAnswersheets(answersheets models.Answersheet) bool {
	if answersheets.IDSession == nil {
		return true
	}
	var answers models.Answers
	if answersheets.Answers == answers {
		return true
	}
	return CheckNilFieldsAnswers(answersheets.Answers)
}

func CheckNilFieldsAnswers(answers models.Answers) bool {
	if answers.Num1 == nil {
		return true
	}
	if answers.Num2 == nil {
		return true
	}
	if answers.Num3 == nil {
		return true
	}
	if answers.Num4 == nil {
		return true
	}
	if answers.Num5 == nil {
		return true
	}
	if answers.Num6 == nil {
		return true
	}
	if answers.Num7 == nil {
		return true
	}
	if answers.Num8 == nil {
		return true
	}
	if answers.Num9 == nil {
		return true
	}
	if answers.Num10 == nil {
		return true
	}
	return false
}
