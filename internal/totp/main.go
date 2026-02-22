package totp

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func Startup() {
	// key, err := totp.Generate(totp.GenerateOpts{
	// 	Issuer:      "Example Application",
	// 	AccountName: "user@example.com",
	// 	SecretSize:  20, // Default size
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// // 2. Display key information to the console
	// fmt.Printf("Issuer: %s\n", key.Issuer())
	// fmt.Printf("Account Name: %s\n", key.AccountName())
	// fmt.Printf("Secret: %s\n", key.Secret())
	// fmt.Printf("URL: %s\n", key.URL()) // The otpauth:// URL

	// // 3. Generate and save the QR code image for the user to scan
	// img, err := key.Image(200, 200)
	// if err != nil {
	// 	panic(err)
	// }
	// var buf bytes.Buffer
	// png.Encode(&buf, img)
	// os.WriteFile("qr-code.png", buf.Bytes(), 0644)
	// fmt.Println("QR code written to qr-code.png. Please scan with your authenticator app.")
	secret := "JBSWY3DPEHPK3PXP"

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Important to stop the ticker when done

	for range ticker.C {
		passcode, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			panic(err)
		}

		key, _ := totp.Generate(totp.GenerateOpts{
			Issuer:      "Example Application",
			AccountName: "user@example.com",
			Secret:      []byte("JBSWY3DPEHPK3PXP")})

		img, err := key.Image(200, 200)

		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		png.Encode(&buf, img)
		err = os.Remove("qr-code.png")

		if err != nil {
			panic(err)
		}

		os.WriteFile("qr-code.png", buf.Bytes(), 0644)

		fmt.Printf("Current TOTP Passcode: %s\n", passcode)
	}
}
