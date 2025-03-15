package cipher

import (
	"github.com/Archetarcher/gophkeeper/internal/vault"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidCipherCardData = errors.New("cipher has to have  valid fields")
)

// CipherCardData is an aggregate
type CipherCardData struct {
	cipher *vault.Cipher

	cardHolderName []byte
	brand          []byte
	number         []byte
	expMonth       []byte
	expYear        []byte
	code           []byte
}

// NewCipherCardData is a Factory to create a new CipherCardData aggregate
// It will validate that the cardHolderName, brand, number, expMonth, expYear, code are not empty
func NewCipherCardData(cardHolderName, brand, number, expMonth, expYear, code, meta []byte, userId uuid.UUID) (*CipherCardData, error) {
	if len(cardHolderName) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "card holder name does not provided")
	}
	if len(brand) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "brand does not provided")
	}
	if len(number) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "number does not provided")
	}
	if len(expMonth) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "expiration month does not provided")
	}
	if len(expYear) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "expiration year does not provided")
	}
	if len(code) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "code does not provided")
	}

	if userId == uuid.Nil {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "incorrect userId")
	}

	return &CipherCardData{
		cipher: &vault.Cipher{
			Id:        uuid.New(),
			UserId:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			MetaData:  meta,
		},
		cardHolderName: cardHolderName,
		brand:          brand,
		number:         number,
		expMonth:       expMonth,
		expYear:        expYear,
		code:           code,
	}, nil
}

// UnmarshalCipherCardDataFromDatabase marshals db model to domain aggregate
// It's not constructor, use only for db unmarshalling
func UnmarshalCipherCardDataFromDatabase(id uuid.UUID, cardHolderName, brand, number, expMonth, expYear, code, meta []byte, userId uuid.UUID, createdAt, updatedAt, deletedAt time.Time) (*CipherCardData, error) {
	c, err := NewCipherCardData(cardHolderName, brand, number, expMonth, expYear, code, meta, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal db model")
	}
	c.cipher.Id = id
	c.cipher.CreatedAt = createdAt
	c.cipher.UpdatedAt = updatedAt
	c.cipher.DeletedAt = deletedAt
	return c, nil
}

func (u *CipherCardData) GetId() uuid.UUID {
	return u.cipher.Id
}
func (u *CipherCardData) GetCardHolderName() []byte {
	return u.cardHolderName
}
func (u *CipherCardData) GetBrand() []byte {
	return u.brand
}
func (u *CipherCardData) GetNumber() []byte {
	return u.number
}
func (u *CipherCardData) GetExpMonth() []byte {
	return u.expMonth
}
func (u *CipherCardData) GetExpYear() []byte {
	return u.expYear
}
func (u *CipherCardData) GetCode() []byte {
	return u.code
}
func (u *CipherCardData) GetUserId() uuid.UUID {
	return u.cipher.UserId
}
func (u *CipherCardData) GetCreatedAt() time.Time {
	return u.cipher.CreatedAt
}
func (u *CipherCardData) GetUpdatedAt() time.Time {
	return u.cipher.UpdatedAt
}
