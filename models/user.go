package models

import (
    "github.com/go-ozzo/ozzo-validation/v4"
    "golang.org/x/crypto/bcrypt"
    "time"
)

type User struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Username  string    `gorm:"unique;not null" json:"username"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Bookings  []Booking `json:"bookings"`
}

func (user *User) HashPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return nil
}

func (user *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}

// проверяет корректность полей модели User
func (user *User) Validate() error {
    return validation.ValidateStruct(user,
        validation.Field(&user.Username, validation.Required, validation.Length(5, 20)),
        validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
    )
}
