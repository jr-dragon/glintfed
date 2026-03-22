package fositestore

import (
	"encoding/json"
	"time"

	"github.com/ory/fosite"
)

func marshalSession(session fosite.Session) ([]byte, error) {
	return json.Marshal(session)
}

func unmarshalSession(data []byte, session fosite.Session) error {
	return json.Unmarshal(data, session)
}

func requesterFromRow(
	requestID, _ /*clientID*/, subject string,
	scopes []string,
	sessionData []byte,
	requestedAt, _ /*expiresAt*/ time.Time,
	client fosite.Client,
	session fosite.Session,
) (fosite.Requester, error) {
	if err := unmarshalSession(sessionData, session); err != nil {
		return nil, err
	}
	req := fosite.NewRequest()
	req.ID = requestID
	req.Client = client
	req.RequestedAt = requestedAt
	req.Session = session
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, s := range scopes {
		req.GrantScope(s)
	}
	_ = subject // subject is carried in the session
	return req, nil
}
