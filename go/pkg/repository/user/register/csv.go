package register

import (
	"encoding/csv"
	"fmt"
	"os"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

type CSVRepository struct{}

func NewCSVRepository() *CSVRepository {
	return &CSVRepository{}
}

func (r *CSVRepository) Store(user *domain.User) error {
	id := user.Id()
	firstName := user.Name().FirstName()
	lastName := user.Name().LastName()
	timeStamp := user.RegisteredAt().String()

	// TODO parameterize?
	// そうですね、ここはmainから渡して上げてもいいかなとは思います
	fileName := "/tmp/users.csv"

	// ここから重複チェックと書き込みの処理
	// それぞれ function を分けた
	// - function の引数は primitive なもの(string)にしているんだけど、
	//   そのまま User の構造体を渡してしまう方法でもいいのかどうかが気になるところ
	// ここは、重複してはいけないかどうかって技術的要素ではなく、ドメインのルールだと思うので、このルール自体は、domain層のコードにアル方が自然ですね
	// domain層にユーザ名チェック用のサービスのinterfaceを切って、こっちでその実装をしてあげるのがよいかと
	// - open した file(os.File) を使い回すこともできるんだけど
	// 非機能要件としてパフォーマンスが問題になったら考えるというのがDDD的ではありますが、
	// 公開されている関数なら、スレッドセーフにするために使い回さない方がよさげです
	// privateな関数であれば影響があるのはそのモジュール内部だけなので、使いまわしちゃってもいいかもです。
	// ただいずれにせよ非機能要件として出てきたら考える、なのでそのときにはCSV以外のものにしようという話になりそうです
	//   reader と writer で open する時の flag を分けたほうがいいかなぁ
	//   open x close が 2 回になるので多少 over head はあるものの、2 回程度ならそんなに気にしないでもいいか
	//   なのでそれぞれの function には fileName を渡している

	// check whether name was already registered
	dup, err := isDupulicatedName(fileName, firstName, lastName)
	if err != nil {
		return err
	}
	if dup {
		return fmt.Errorf("%s %s was already registerd", firstName, lastName)
	}

	// write user data
	// if err := writeIntoCSV(...); err != nil { return err }
	// と丁寧に書きたくなるけどもここはスマートに
	return writeIntoCSV(fileName, timeStamp, id, firstName, lastName)
}

func isDupulicatedName(fileName, firstName, lastName string) (bool, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0400)
	defer file.Close()
	if err != nil {
		return true, err
	}
	reader := csv.NewReader(file)
	entries, rerr := reader.ReadAll()
	if rerr != nil {
		return true, rerr
	}

	// colmun
	//   1: timestamp
	//   2: id
	//   3: firstName
	//   4: lastName
	for _, entry := range entries {
		if entry[2] == firstName && entry[3] == lastName {
			// found entry
			return true, nil
		}
	}

	// not found
	return false, nil
}

func writeIntoCSV(fileName, timeStamp, id, firstName, lastName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{timeStamp, id, firstName, lastName})
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
