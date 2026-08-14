package utility

type PasswordHasher interface {
	Hash(plaintext string) (string, error)
}