package user

import (
	"encoding/csv"
	"os"

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
	file, err := os.OpenFile("/tmp/users.csv", os.O_WRONLY|os.O_APPEND, 0600)
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
