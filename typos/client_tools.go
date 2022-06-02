package typos

import (
	"encoding/base64"
	"strings"
)

func (ct *ClientData) UnmarshalJSON(d []byte) error {
	data := string(d)
	*ct = ClientData(data)
	return nil
}

func (ct *ClientData) String() string {
	return string(*ct)
}

func (ct *ClientScope) UnmarshalJSON(d []byte) error {
	scope := strings.Replace(string(d), "\"", "", -1)
	scopes := strings.Split(scope, " ")
	*ct = ClientScope(scopes)
	return nil
}

func (ct *ClientSecretBase64) UnmarshalJSON(d []byte) error {
	data := base64.StdEncoding
	secret := data.EncodeToString(d)
	*ct = ClientSecretBase64(secret)
	return nil
}

func (ct *ClientSecretBase64) String() string {
	return string(*ct)
}
