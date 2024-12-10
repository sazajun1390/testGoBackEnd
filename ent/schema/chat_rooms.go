// ent/schema/chatroom.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatRoom holds the schema definition for the ChatRoom entity.
type ChatRoom struct {
	ent.Schema
}

// Fields of the ChatRoom.
func (ChatRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the ChatRoom.

func (ChatRoom) Edges() []ent.Edge {
	return []ent.Edge{
		// 1:N - チャットルームを作成したユーザー
		edge.From("creator", User.Type).
			Ref("chat_rooms").
			Unique().
			Required(),

		// 1:N - チャットルーム内のメッセージ
		edge.To("messages", Message.Type),

		// N:N - チャットルームに参加しているユーザー
		edge.From("participants", User.Type).
			Ref("participatis_rooms").
			Through("memberships", ChatRoomMember.Type),
		/*
			edge.To("participantsRoom", User.Type).
				Through("memberships", ChatRoomMember.Type),
		*/

		// チャットルームメンバーシップへのエッジを追加
		//edge.To("memberships", ChatRoomMember.Type),
	}
}

/*
func (ChatRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("chat_rooms").
			Unique().
			Required(),
		edge.To("messages", Message.Type),
		edge.To("participants", User.Type).
			Through("memberships", ChatRoomMember.Type),
	}
}
*/
