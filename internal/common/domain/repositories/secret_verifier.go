package common_repositories

type SecretVerifier interface {
	Verify(hashed string, secret string) bool
	Hash(content string) (string, error)
}
