package user

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gofrs/flock"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type CsvRepository struct{}

func NewCsvRepository() *CsvRepository {
	return &CsvRepository{}
}

// domain/model/user/user_repository.go にある IUserRepositoryを実装する構造体です
// 実際にアプリケーションを作るときはここの部分は最も最後に書くように、設計もなんなら後回しで大丈夫です
// 理由としては例えばDBの選択はビジネスの開始フェイズ(または開発開始フェイズ)時に必ずしも最適なモノを選べるとは限らないからです
// とにかくはじめにinterfaceを書きましょう、そしてその後実態を書いていきます。そうするとDBの細かい特性などもどうでもよくなってくるはずですｗ
func (r *CsvRepository) Store(user *domain.User) error {
	id := user.Id()
	firstName := user.Name().FirstName()
	lastName := user.Name().LastName()
	file, err := os.OpenFile("/tmp/users.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	err = writer.Write([]string{id, firstName, lastName})
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func (r *CsvRepository) FindById(id *domain.Id) (*domain.User, error) {
	file, err := os.OpenFile("/tmp/users.csv", os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if id.Value() == record[0] {
			user, err := domain.NewUser(
				domain.NewId(record[0]),
				record[1],
				record[2],
			)
			if err != nil {
				return nil, err
			}
			return user, nil
		}
	}
	return nil, errors.New("Couldn't find specified user")
}

func (r *CsvRepository) Update(user *domain.User) error {

	var newRecords [][]string
	var err error

	// 簡易的な lock の機構が必要かな
	fileLock := flock.New("/tmp/users.csv")
	locked, err := fileLock.TryLock()
	if err != nil {
		return err
	}

	// lock failed
	if !locked {
		return fmt.Errorf("csv file is locked. Please try again.")
	}

	defer fileLock.Unlock()

	rfile, err := os.OpenFile("/tmp/users.csv", os.O_RDONLY, 0600)
	if err != nil {
		return err
	}

	reader := csv.NewReader(rfile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if user.Id() == record[0] {
			newRecords = append(newRecords, []string{
				record[0],
				user.Name().FirstName(),
				user.Name().LastName(),
			})
		} else {
			// no need to change the record
			newRecords = append(newRecords, record)
		}
	}

	// 書き込みの為に再度 Open
	wfile, err := os.OpenFile("/tmp/users.csv", os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer wfile.Close()
	out := csv.NewWriter(wfile)

	// WriteAll が error を返すので、そのまま return に渡す
	return out.WriteAll(newRecords)
}
