// ent/schema/user.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("password_hash").
			NotEmpty(),
	}
}

// Edges of the User.

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// 1:N - ユーザーが作成したチャットルーム
		edge.To("chat_rooms", ChatRoom.Type).
			Unique(),

		// 1:N - ユーザーが送信したメッセージ
		edge.To("messages", Message.Type),

		// N:N - ユーザーが参加しているチャットルーム

		edge.To("participatis_rooms", ChatRoom.Type).
			Through("memberships", ChatRoomMember.Type),

		/*
			edge.From("participating_rooms", ChatRoom.Type).
				Ref("participants").
				Through("memberships", ChatRoomMember.Type),
		*/
		/*
			edge.From("participating_roomsRoom", ChatRoom.Type).
				Ref("participantsRoom").
				Through("memberships", ChatRoomMember.Type),
		*/

		// チャットルームメンバーシップへのエッジを追加
		//edge.To("memberships", ChatRoomMember.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
		index.Fields("username").Unique(),
	}
}

/*
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat_rooms", ChatRoom.Type).
			From("creator").
			Unique(),
		edge.To("messages", Message.Type),
		edge.To("participating_rooms", ChatRoom.Type).
			Through("memberships", ChatRoomMember.Type),
	}
}
*/
