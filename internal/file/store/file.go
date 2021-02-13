package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"gorm.io/gorm"
)

type FileStore interface {
	GetUserFolders(userID int) ([]*model.Folder, error)
	ListFolders(folders []int, withSubFolders, withFiles bool) ([]*model.Folder, error)

	FindFolder(folderID string) (*model.Folder, error)
	CreateFolder(folder *model.Folder) error
	AppendFolder(parentID string, folder *model.Folder) error
	UpdateFolder(folder *model.Folder) error
	DeleteFolder(folderID int) error
	DeleteFolders(folderIDs []int) error
}

type fileStore struct {
	db *gorm.DB
}

func NewFolderStore(conn *gorm.DB) *fileStore {
	return &fileStore{
		db: conn,
	}
}

func (s fileStore) GetUserFolders(uid int) ([]*model.Folder, error) {
	var folders []*model.Folder
	query := s.db.Model(&model.Folder{})
	if err := query.Where("owner_id  = ?", uid).Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

func (s fileStore) ListFolders(folders []int, withSubFolders, withFiles bool) ([]*model.Folder, error) {
	return nil, nil
}
func (s fileStore) FindFolder(folderID string) (*model.Folder, error) {
	var folder model.Folder
	query := s.db.Model(&model.Folder{})
	if err := query.Where("id = ?", folderID).Take(&folder).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (s fileStore) CreateFolder(folder *model.Folder) error {
	return s.db.Create(folder).Error
}

func (s fileStore) AppendFolder(parentID string, folder *model.Folder) error {
	query := s.db.Model(&model.Folder{ID: parentID}).Association("Subs")
	if err := query.Append(folder); err != nil {
		return err
	}
	return nil
}
func (s fileStore) UpdateFolder(folder *model.Folder) error {
	return nil
}
func (s fileStore) DeleteFolder(folderID int) error {
	return nil
}
func (s fileStore) DeleteFolders(folderIDs []int) error {
	return nil
}
