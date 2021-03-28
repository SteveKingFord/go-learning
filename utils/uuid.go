package utils

import(
	"fmt"
	"github.com/gofrs/uuid"
)

func GenUUID() (string,error){
	v4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return "", err
	}
	return v4.String(), nil
}