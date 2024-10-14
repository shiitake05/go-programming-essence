// 「go run -mod=mod entgo.io/ent/cmd/ent init Todo」を実行して生成されたファイル
// Facebook社が開発しているORMで、高度に構造化されており、モデルとクライアントを分離しつつ、IDEと親和性の良いインターフェースを提供している
// entでデータベースを扱うときには、スキーマを定義する必要がある
// 「cd $_」でtodoディレクトリ
// 「go generate ./ent」でentディレクトリ内にTodoを操作するためのCRUD APIが生成される
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
// func (Todo) Fields() []ent.Field {
// 	return nil
// }

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Int("priority").
			Default(0),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
