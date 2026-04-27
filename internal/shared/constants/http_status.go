package constants

import "errors"

const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFail    = "fail"
)

const (
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeUnauthorized = "UNAUTHORIZED"
)

var (
	ErrDataNotFound              = errors.New("data tidak ditemukan")
	ErrAccountNotFound           = errors.New("akun tidak ditemukan")
	ErrAccountNotFoundOrInactive = errors.New("akun tidak ditemukan atau tidak aktif")
	ErrAccountInactive           = errors.New("akun tidak aktif")
	ErrInvalidPassword           = errors.New("password yang anda masukan salah")
	ErrInvalidRequest            = errors.New("request tidak valid")
	ErrDataAlreadyExists         = errors.New("data anda sudah ada")
	ErrNimAlreadyExists          = errors.New("nim anda sudah terdaftar")
	ErrEmailAlreadyExists        = errors.New("email anda sudah terdaftar")
	ErrInternalServer            = errors.New("terjadi kesalahan sistem")
)
