package github

type MockClient struct {
	LatestTag string
	Err       error
}

func (m *MockClient) FetchLatestTag(owner, repo string) (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.LatestTag, nil
}
