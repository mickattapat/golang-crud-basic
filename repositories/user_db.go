package repositories

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) Login(user User) (*User, error) {
	tx := r.db.Where("username=?", user.Username).Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r userRepositoryDB) Create(user User) (*User, error) {

	tx := r.db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Update(id int, user User) error {

	tx := r.db.Model(&User{}).Where("user_id=?", id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r userRepositoryDB) GetAll() ([]User, error) {

	users := []User{}
	tx := r.db.Find(&users)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r userRepositoryDB) GetByID(id int) (*User, error) {
	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Delete(id int) error {
	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return tx.Error
	}

	tx = r.db.Unscoped().Delete(&user, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
