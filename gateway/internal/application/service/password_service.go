package service

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// PasswordService defines methods for hashing and verifying passwords
type PasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hashed string) (bool, error)
}

// Config holds Argon2 parameters for password hashing
type Config struct {
	TimeCost    uint32 // Number of iterations
	MemoryCost  uint32 // Memory usage in kibibytes
	Parallelism uint8  // Number of threads
	SaltLength  uint32 // Salt length in bytes
	HashLength  uint32 // Hash length in bytes
}

// passwordServiceImpl is the implementation of PasswordService
type passwordServiceImpl struct {
	config *Config
}

// NewPasswordService creates a new PasswordService with the provided configuration
func NewPasswordService(config *Config) PasswordService {
	return &passwordServiceImpl{config: config}
}

// GenerateSalt creates a random salt
func (p *passwordServiceImpl) generateSalt() ([]byte, error) {
	salt := make([]byte, p.config.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}

// HashPassword hashes the password using Argon2 with the configured parameters
func (p *passwordServiceImpl) HashPassword(password string) (string, error) {
	salt, err := p.generateSalt()
	if err != nil {
		return "", err
	}

	// Hash the password
	hash := argon2.IDKey([]byte(password), salt, p.config.TimeCost, p.config.MemoryCost, p.config.Parallelism, p.config.HashLength)

	// Encode the salt and hash for storage
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s.%s", saltEncoded, hashEncoded), nil
}

// VerifyPassword compares a plaintext password with a stored hash
func (p *passwordServiceImpl) VerifyPassword(password, hashed string) (bool, error) {
	// Split the stored hash into salt and hash components
	parts := strings.Split(hashed, ".")
	if len(parts) != 2 {
		return false, errors.New("invalid hash format")
	}

	// Decode the salt and original hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	originalHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	// Hash the input password with the extracted salt
	hash := argon2.IDKey([]byte(password), salt, p.config.TimeCost, p.config.MemoryCost, p.config.Parallelism, p.config.HashLength)

	// Compare the computed hash to the original hash
	return bytes.Equal(hash, originalHash), nil
}

func ProvidePasswordConfig() *Config {
	return &Config{
		TimeCost:    1,
		MemoryCost:  64 * 1024, // 64 MB
		Parallelism: 2,
		SaltLength:  16,
		HashLength:  32,
	}
}
