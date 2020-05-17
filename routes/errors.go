package routes

import "errors"

// ErrIncorrectPassword is for incorrect password login attempts
var ErrIncorrectPassword = errors.New("Incorrect password")

// ErrBadUser is for errors retrieving login user
var ErrBadUser = errors.New("Bad username")
