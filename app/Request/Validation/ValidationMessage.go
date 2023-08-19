package Validation

func ValidationMessage(tag string, msg string, parms string) string {
	switch tag {
	case "required":
		return msg + " Wajib Diisi"
	case "max":
		return msg + " Maksimal " + parms + " Karakter"
	case "email":
		return "Invalid email"
	}
	return ""
}
