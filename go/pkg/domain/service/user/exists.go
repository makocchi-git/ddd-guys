package user

import (
	"github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

// これは次回のrepositoryパターンでも取り上げる予定ですが
// repositoryにある isDuplicatedUser 的な関数は "ユーザ名の重複は許可しない" というルールがドメイン的であるため、
// 存在可否のようなチェックと、実際の保存は責務を分けられるべきであります
// (関数名が重複しているユーザかどうかという点で技術要素よりかはビジネスルールっぽいですよね??)
// そしてapplication service内で、
// 1. ユーザオブジェクトを作成
// 2. 存在可否チェック: いたらerrorを返す
// 3. いなければ保存
// という処理のフローを書きます
// また
// - ユーザの重複を許可しない、なのか
// - ユーザの氏名の重複を許可しない、なのか
// これによって渡す引数がUserを渡すべきか、User.Name()を渡すべきか変わってくるはずです
// ただ通常はなんとなく重複を許可しないではなく、氏名による重複排除、メアドによる重複排除というのが、ビジネスルールとして決められるはずです

// UserServiceにしてもいいですし、ケースごとにわけてもいいかなとは思います
type UserExistsService struct {
	repository user.IUserRepository
}

// 実際のところ、repositoryの実装パターンとしては2通りあります(domain/user/repository.go にて詳細を残しておきます)
func NewUserExistsService(repository user.IUserRepository) *UserExistsService {
	return &UserExistsService{repository}
}

// パターン1: repositoryにExistsメソッドを生やす場合
// 例としてユーザの氏名で重複チェックを行わせるという名目でやってます(実際はメアドとかが自然だとは思いますがｗ)
func (s *UserExistsService) ExistsV1(n user.Name) bool {
	return s.repository.Exists(n) 
}
// パターン2: repositoryにFindByIdメソッドを生やす場合
func (s *UserExistsService) ExistsV2(n user.Name) bool {
	u := s.repository.FindByName(n) // 実際はbool, errorと返す感じになるかもですね。
	if u == nil {
		return false
	}
	return true
}

// for just passing assembly check
func (s *UserExistsService) Exists(u user.User) bool {
	return s.ExistsV1(u)
}
