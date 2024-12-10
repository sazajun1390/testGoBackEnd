// ent/schema/message.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").
			NotEmpty(),
		field.Time("created_at").
			Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		// 1:N - メッセージが所属するチャットルーム
		edge.From("room", ChatRoom.Type).
			Ref("messages").
			Unique().
			Required(),

		// 1:N - メッセージを送信したユーザー
		edge.From("author", User.Type).
			Ref("messages").
			Unique().
			Required(),
	}
}

/*
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", ChatRoom.Type).
			Ref("messages").
			Unique().
			Required(),
		edge.From("author", User.Type).
			Ref("messages").
			Unique().
			Required(),
	}
}
*/
