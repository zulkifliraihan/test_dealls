package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type TypeValidationError struct {
    Field 		string
    Message   	string
}

func ValidationCustomMessage(err validator.ValidationErrors) []TypeValidationError {

	var ve validator.ValidationErrors

	if errors.As(err, &ve) {

		formatErrors := make([]TypeValidationError, len(ve))
		fmt.Println("Error ve : ", ve)
		for i, fe := range ve {

			// fmt.Print(
			// 	`Error fe `, i , ` ActualTag: `, fe.ActualTag(), "\n",
			// 	`Error fe `, i , ` Type: `, fe.Type(), "\n",
			// 	`Error fe `, i , ` Value: `, fe.Value(), "\n",
			// 	`Error fe `, i , ` Error: `, fe.Error(), "\n",
			// 	`Error fe `, i , ` Kind: `, fe.Kind(), "\n",
			// 	`Error fe `, i , ` Namespace: `, fe.Namespace(), "\n",
			// 	`Error fe `, i , ` Param: `, fe.Param(), "\n",
			// 	`Error fe `, i , ` StructField: `, fe.StructField(), "\n",
			// 	`Error fe `, i , ` StructNamespace: `, fe.StructNamespace(), "\n",
			// 	`Error fe `, i , ` Kind().String(): `, fe.Kind().String(), "\n",
			// 	`Error fe `, i , ` Type().Align(): `, fe.Type().Align(), "\n",
			// )
			
			formatErrors[i] = TypeValidationError{
				fe.Field(), 
				msgForTag(fe.Tag(), fe.Param()),
			}
		}

		return formatErrors
	}
	
	return nil
}

func msgForTag(tag string, param string) string {
    switch tag {
    	case "required":
        	return "This field is required."
    	case "number":
        	return "Please enter a valid numeric value."
    	case "email":
        	return "Please enter a valid email address."
    	case "min":
        	return "Please enter at least " + param + " characters.."
    	case "max":
        	return "Please enter no more than " + param + " characters."
    }

    return ""
}
