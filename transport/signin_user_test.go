package transport

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndSigninUser(t *testing.T) {
	client := &http.Client{}
	payload := map[string]string{
		"username": gofakeit.Username(),
		"password": gofakeit.Password(true, true, true, false, false, 15),
	}

	path := fmt.Sprintf("%s/1.0/user", server.URL)
	b, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(b))
	assert.NoError(t, err)

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusCreated)

	// signin
	pathSignin := fmt.Sprintf("%s/1.0/user/signin", server.URL)
	payloadSignin, err := json.Marshal(payload)
	assert.NoError(t, err)

	reqSignin, err := http.NewRequest(http.MethodPost, pathSignin, bytes.NewBuffer(payloadSignin))
	assert.NoError(t, err)

	resp, err = client.Do(reqSignin)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
