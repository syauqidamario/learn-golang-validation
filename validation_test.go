package learn_golang_validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationField(t *testing.T){
	validate := validator.New()
	var user string = ""

	err := validate.Var(user, "required")

	if err != nil{
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariables(t *testing.T){
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "salah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil{
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T){
	validate := validator.New()
	var user string = "syauqi12345"

	err := validate.Var(user, "required,number")

	if err != nil{
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T){
	validate := validator.New()
	user := "11"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil{
		fmt.Println(err.Error())
	}
}

