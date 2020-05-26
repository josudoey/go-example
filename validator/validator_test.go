package validator_test

// ref https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func ExampleValidate_Struct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}
	// Output:
	// User.Age
	// Age
	// User.Age
	// Age
	// lte
	// lte
	// uint8
	// uint8
	// 135
	// 130

	// User.FavouriteColor
	// FavouriteColor
	// User.FavouriteColor
	// FavouriteColor
	// iscolor
	// hexcolor|rgb|rgba|hsl|hsla
	// string
	// string
	// #000-

	// User.Addresses[0].City
	// City
	// User.Addresses[0].City
	// City
	// required
	// required
	// string
	// string
}

func ExampleValidate_Variable() {
	myEmail := "josudoey.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs)
		return
	}

	// Output:
	// Key: '' Error:Field validation for '' failed on the 'email' tag
}
