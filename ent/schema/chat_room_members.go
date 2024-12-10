// ent/schema/chatroommember.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatRoomMember holds the schema definition for the ChatRoomMember entity.
type ChatRoomMember struct {
	ent.Schema
}

// Fields of the ChatRoomMember.
func (ChatRoomMember) Fields() []ent.Field {
	return []ent.Field{
		field.Time("joined_at").
			Default(func() time.Time { return time.Now() }),
		field.Int("user_id"),
		field.Int("room_id"),
	}
}

// Edges of the ChatRoomMember.

func (ChatRoomMember) Edges() []ent.Edge {
	return []ent.Edge{
		// 1:N - ユーザーが中間テーブルに参加
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),

		// 1:N - チャットルームが中間テーブルに参加
		edge.To("room", ChatRoom.Type).
			Unique().
			Required().
			Field("room_id"),
	}
}

/*
func (ChatRoomMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", ChatRoom.Type).
			Ref("memberships").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("memberships").
			Unique().
			Required(),
	}
}
*/
