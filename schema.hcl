table "chat_room_members" {
  schema = schema.test
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "joined_at" {
    null = true
    type = timestamp
  }
  column "user_id" {
    null = false
    type = bigint
  }
  column "room_id" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "chat_room_members_chat_rooms_room" {
    columns     = [column.room_id]
    ref_columns = [table.chat_rooms.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "chat_room_members_users_user" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "chat_room_members_chat_rooms_room" {
    columns = [column.room_id]
  }
  index "chatroommember_user_id_room_id" {
    unique  = true
    columns = [column.user_id, column.room_id]
  }
}
table "chat_rooms" {
  schema = schema.test
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "user_chat_rooms" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "chat_rooms_users_chat_rooms" {
    columns     = [column.user_chat_rooms]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "user_chat_rooms" {
    unique  = true
    columns = [column.user_chat_rooms]
  }
}
table "messages" {
  schema = schema.test
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "content" {
    null = false
    type = longtext
  }
  column "created_at" {
    null = true
    type = timestamp
  }
  column "chat_room_messages" {
    null = false
    type = bigint
  }
  column "user_messages" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "messages_chat_rooms_messages" {
    columns     = [column.chat_room_messages]
    ref_columns = [table.chat_rooms.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "messages_users_messages" {
    columns     = [column.user_messages]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "messages_chat_rooms_messages" {
    columns = [column.chat_room_messages]
  }
  index "messages_users_messages" {
    columns = [column.user_messages]
  }
}
table "users" {
  schema = schema.test
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "username" {
    null = false
    type = varchar(255)
  }
  column "email" {
    null = false
    type = varchar(255)
  }
  column "password_hash" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "email" {
    unique  = true
    columns = [column.email]
  }
  index "username" {
    unique  = true
    columns = [column.username]
  }
}
schema "test" {
  charset = "utf8mb4"
  collate = "utf8mb4_bin"
}
