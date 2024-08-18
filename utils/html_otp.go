package utils

import "fmt"

func CreateHTMLOTP(name, province, city, otp string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            padding: 20px;
        }
        .container {
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
        }
        p {
            color: #666;
        }
        .otp {
            font-size: 24px;
            font-weight: bold;
            color: #d9534f;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>OTP Verification</h1>
        <p>Jangan berikan kode ini kepada siapapun, ini merupakan kredensial. <strong>Pihak Developer tidak pernah menanyakan OTP.</strong> Gunakan ini untuk aplikasi Islamind</p>
        <p>Name: %s</p>
        <p>Province: %s</p>
        <p>City: %s</p>
        <p>Your OTP is: <span class="otp">%s</span></p>
    </div>
</body>
</html>
`, name, province, city, otp)
}
