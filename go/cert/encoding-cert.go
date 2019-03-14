package main

import (
	"encoding/pem"
	"os"
)

func main() {
	keyStr := "MIIEpAIBAAKCAQEAjX88TOJrzsyJtif0nfIb80Q0N6KKPMbgS/HCV8Jv4NnFksnwUU3k16tbVcPh\njtIEzkeZe/DEa+0dPaQN7xYLJyZEX8L2jtLGJaxLiCMSt7vhf1oxxVxIz4ckVQzAzMOpf7DoJmnM\nFO4jHT1WibDpQ2ZAHOUua+v/VCvTMWDlHC7mb2KkkRq+FvLz/9C/MUTi0ViZIWzpxsjOFeIWZqfh\nASZvMfYtuEODI+3zBOjITtDYWoIlq/fqKxhxf6QR+xiSFEoOr21ND+w1GWkA+pzVz9/JloGPNw4L\nsm7C8aaD6LJ9rqEzR51fEejGNPdEa1nKZkSOoG3xNuRpGU3Wk66EJwIDAQABAoIBAQCAK1SPwDkv\nZY8IdUDbIUe4H3u9OXpGe08I+8CyhemZnymcQVdyDkDQKuPc2EYkfJDl2W1qGzmRGYN3JneIrksD\n92HP7GYyETJzRZJVio+V5pXoKecTgvhhI2UBtT+fO0z54WEoCoVWIC+meL0+lAV3i9NfWSmnkGoD\n8H6pJ7/gui+r0QHE3s8ZQ/QAuCOQKxvaIIxeoBjWcduIsSnx9xhpItDMmxT2cfHYeu+PPeK5KiuJ\nxA23GARXt0r3LUnkQIX3iGsc6EEFhHlsi+HaZX8NtSUW7zKqaY4eHsD6iToujG8kpb/tOj2nyz1r\nhv+Nv+/7lLpeNHrlEBOvfvBZLgmBAoGBANavXxdujFDPDMQ172rcUsf3aEW+m0ZHk9oGwFQfJvzL\naFOMTYIanCCvS0Bc4rJ0XKV0yslkiz8xz24WoIsf250Iu/9sJweAR9AJlerr02srgAqE0f1xbjHA\n/YLclKoyljd0XFERmNZQ58X6AojcpLxbDsogZ4EvoxK/7rH2T1sRAoGBAKi6Mqcjn5Kx2YmNWd4n\ncn4vmR5hv1bYqx8IrXOJ/vDNE27YD/z5pnOtauVBzAPZVC5B6qmp+xCkvzZw3Hgiy8GfxQomtkgf\npuIdyoC4CrMgwmbJJn92mnwFXJ+w1y6HNwnMQ90KYmTDPhOA2WVpAIOVg3fdeID+xdywrWecf7u3\nAoGBAIy/UOJ10XwxIjeI/LvcUlFTPINtITOEPcSU5IbCHbdiS7VqqG4BsBYllnpN6zn/27WFymUa\ni8gUid+s6fgMHNjpHweoCjZ6vZoAjs8DaQo77Kf4yAsW0J3B3vJqVvT2UVFt8hAcL01ioXFEJWFA\n9qeeMyYE9JbXelwShRUbXI9hAoGAGpMI+04fCem9C28Z4mVwIPO6OBDJJBKIESE6wGHRp3eFQgYB\nrqRbVQjRNz5/tXN7ULXNEd/VzS9FdnucHcfNM7fIPusN9poVnLN4e4GBHelg9PNlftndRv9JMVWW\n7psZTg+/+IZ0LM6bblLy6of5LsBTroNRQW/fFf01soXxkE0CgYBGzTSkQ0ZTycJxMb9bP7EmErhr\nMCfhU7+Aa0E13a5p89c2CGG7GD6Zw5tRKcG7ztoh/3JW+8PQjDCDZUhdSK93kYH7n0syAQ7My3af\n+iIRHedKpofKSFAshrXwSS/vzQMSsodiGqNvewGM34qgeQIYHwI7zBKA62z82zXD/LrCAw=="
	savePEMKey("./key.pem", []byte(keyStr))
}

func savePEMKey(fileName string, key []byte) {
	outFile, _ := os.Create(fileName)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: key,
	}

	_ = pem.Encode(outFile, privateKey)
}
