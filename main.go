package main

import "ecommerce/cmd"

func main() {
	cmd.Serve()

	// jwt, err := util.CreateJwt("my-secret", util.PayLoad{
	// 	Sub:         45,
	// 	FirstName:   "Shompa",
	// 	LastName:    "Akter",
	// 	Email:       "shompa@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)
}
