package utility

type PasswordHasher interface {
	Hash(plaintext string) (string, error)
	Compare(hashedPassword, plainPassword string) (bool, error)
}
