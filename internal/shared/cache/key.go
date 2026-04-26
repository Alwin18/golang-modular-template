package cache

import "fmt"

// Key constants & helpers for consistent Redis key naming.

func UserKey(id uint) string {
	return fmt.Sprintf("user:%d", id)
}

func SessionKey(token string) string {
	return fmt.Sprintf("session:%s", token)
}

func OTPKey(identifier string) string {
	return fmt.Sprintf("otp:%s", identifier)
}

func RateLimitKey(identifier string) string {
	return fmt.Sprintf("ratelimit:%s", identifier)
}

func LockKey(resource string) string {
	return fmt.Sprintf("lock:%s", resource)
}

func BlacklistTokenKey(token string) string {
	return fmt.Sprintf("blacklist:%s", token)
}
