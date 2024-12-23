// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"chatSystem/ent/migrate"

	"chatSystem/ent/chatroom"
	"chatSystem/ent/chatroommember"
	"chatSystem/ent/message"
	"chatSystem/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ChatRoom is the client for interacting with the ChatRoom builders.
	ChatRoom *ChatRoomClient
	// ChatRoomMember is the client for interacting with the ChatRoomMember builders.
	ChatRoomMember *ChatRoomMemberClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ChatRoom = NewChatRoomClient(c.config)
	c.ChatRoomMember = NewChatRoomMemberClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		ChatRoom:       NewChatRoomClient(cfg),
		ChatRoomMember: NewChatRoomMemberClient(cfg),
		Message:        NewMessageClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		ChatRoom:       NewChatRoomClient(cfg),
		ChatRoomMember: NewChatRoomMemberClient(cfg),
		Message:        NewMessageClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ChatRoom.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ChatRoom.Use(hooks...)
	c.ChatRoomMember.Use(hooks...)
	c.Message.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ChatRoom.Intercept(interceptors...)
	c.ChatRoomMember.Intercept(interceptors...)
	c.Message.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChatRoomMutation:
		return c.ChatRoom.mutate(ctx, m)
	case *ChatRoomMemberMutation:
		return c.ChatRoomMember.mutate(ctx, m)
	case *MessageMutation:
		return c.Message.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ChatRoomClient is a client for the ChatRoom schema.
type ChatRoomClient struct {
	config
}

// NewChatRoomClient returns a client for the ChatRoom from the given config.
func NewChatRoomClient(c config) *ChatRoomClient {
	return &ChatRoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chatroom.Hooks(f(g(h())))`.
func (c *ChatRoomClient) Use(hooks ...Hook) {
	c.hooks.ChatRoom = append(c.hooks.ChatRoom, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chatroom.Intercept(f(g(h())))`.
func (c *ChatRoomClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChatRoom = append(c.inters.ChatRoom, interceptors...)
}

// Create returns a builder for creating a ChatRoom entity.
func (c *ChatRoomClient) Create() *ChatRoomCreate {
	mutation := newChatRoomMutation(c.config, OpCreate)
	return &ChatRoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatRoom entities.
func (c *ChatRoomClient) CreateBulk(builders ...*ChatRoomCreate) *ChatRoomCreateBulk {
	return &ChatRoomCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChatRoomClient) MapCreateBulk(slice any, setFunc func(*ChatRoomCreate, int)) *ChatRoomCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChatRoomCreateBulk{err: fmt.Errorf("calling to ChatRoomClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChatRoomCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChatRoomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatRoom.
func (c *ChatRoomClient) Update() *ChatRoomUpdate {
	mutation := newChatRoomMutation(c.config, OpUpdate)
	return &ChatRoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatRoomClient) UpdateOne(cr *ChatRoom) *ChatRoomUpdateOne {
	mutation := newChatRoomMutation(c.config, OpUpdateOne, withChatRoom(cr))
	return &ChatRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatRoomClient) UpdateOneID(id int) *ChatRoomUpdateOne {
	mutation := newChatRoomMutation(c.config, OpUpdateOne, withChatRoomID(id))
	return &ChatRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatRoom.
func (c *ChatRoomClient) Delete() *ChatRoomDelete {
	mutation := newChatRoomMutation(c.config, OpDelete)
	return &ChatRoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatRoomClient) DeleteOne(cr *ChatRoom) *ChatRoomDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatRoomClient) DeleteOneID(id int) *ChatRoomDeleteOne {
	builder := c.Delete().Where(chatroom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatRoomDeleteOne{builder}
}

// Query returns a query builder for ChatRoom.
func (c *ChatRoomClient) Query() *ChatRoomQuery {
	return &ChatRoomQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChatRoom},
		inters: c.Interceptors(),
	}
}

// Get returns a ChatRoom entity by its id.
func (c *ChatRoomClient) Get(ctx context.Context, id int) (*ChatRoom, error) {
	return c.Query().Where(chatroom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatRoomClient) GetX(ctx context.Context, id int) *ChatRoom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCreator queries the creator edge of a ChatRoom.
func (c *ChatRoomClient) QueryCreator(cr *ChatRoom) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroom.Table, chatroom.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, chatroom.CreatorTable, chatroom.CreatorColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a ChatRoom.
func (c *ChatRoomClient) QueryMessages(cr *ChatRoom) *MessageQuery {
	query := (&MessageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroom.Table, chatroom.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chatroom.MessagesTable, chatroom.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParticipants queries the participants edge of a ChatRoom.
func (c *ChatRoomClient) QueryParticipants(cr *ChatRoom) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroom.Table, chatroom.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, chatroom.ParticipantsTable, chatroom.ParticipantsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMemberships queries the memberships edge of a ChatRoom.
func (c *ChatRoomClient) QueryMemberships(cr *ChatRoom) *ChatRoomMemberQuery {
	query := (&ChatRoomMemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroom.Table, chatroom.FieldID, id),
			sqlgraph.To(chatroommember.Table, chatroommember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, chatroom.MembershipsTable, chatroom.MembershipsColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChatRoomClient) Hooks() []Hook {
	return c.hooks.ChatRoom
}

// Interceptors returns the client interceptors.
func (c *ChatRoomClient) Interceptors() []Interceptor {
	return c.inters.ChatRoom
}

func (c *ChatRoomClient) mutate(ctx context.Context, m *ChatRoomMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatRoomCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatRoomUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatRoomDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChatRoom mutation op: %q", m.Op())
	}
}

// ChatRoomMemberClient is a client for the ChatRoomMember schema.
type ChatRoomMemberClient struct {
	config
}

// NewChatRoomMemberClient returns a client for the ChatRoomMember from the given config.
func NewChatRoomMemberClient(c config) *ChatRoomMemberClient {
	return &ChatRoomMemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chatroommember.Hooks(f(g(h())))`.
func (c *ChatRoomMemberClient) Use(hooks ...Hook) {
	c.hooks.ChatRoomMember = append(c.hooks.ChatRoomMember, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chatroommember.Intercept(f(g(h())))`.
func (c *ChatRoomMemberClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChatRoomMember = append(c.inters.ChatRoomMember, interceptors...)
}

// Create returns a builder for creating a ChatRoomMember entity.
func (c *ChatRoomMemberClient) Create() *ChatRoomMemberCreate {
	mutation := newChatRoomMemberMutation(c.config, OpCreate)
	return &ChatRoomMemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatRoomMember entities.
func (c *ChatRoomMemberClient) CreateBulk(builders ...*ChatRoomMemberCreate) *ChatRoomMemberCreateBulk {
	return &ChatRoomMemberCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChatRoomMemberClient) MapCreateBulk(slice any, setFunc func(*ChatRoomMemberCreate, int)) *ChatRoomMemberCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChatRoomMemberCreateBulk{err: fmt.Errorf("calling to ChatRoomMemberClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChatRoomMemberCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChatRoomMemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatRoomMember.
func (c *ChatRoomMemberClient) Update() *ChatRoomMemberUpdate {
	mutation := newChatRoomMemberMutation(c.config, OpUpdate)
	return &ChatRoomMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatRoomMemberClient) UpdateOne(crm *ChatRoomMember) *ChatRoomMemberUpdateOne {
	mutation := newChatRoomMemberMutation(c.config, OpUpdateOne, withChatRoomMember(crm))
	return &ChatRoomMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatRoomMemberClient) UpdateOneID(id int) *ChatRoomMemberUpdateOne {
	mutation := newChatRoomMemberMutation(c.config, OpUpdateOne, withChatRoomMemberID(id))
	return &ChatRoomMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatRoomMember.
func (c *ChatRoomMemberClient) Delete() *ChatRoomMemberDelete {
	mutation := newChatRoomMemberMutation(c.config, OpDelete)
	return &ChatRoomMemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChatRoomMemberClient) DeleteOne(crm *ChatRoomMember) *ChatRoomMemberDeleteOne {
	return c.DeleteOneID(crm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChatRoomMemberClient) DeleteOneID(id int) *ChatRoomMemberDeleteOne {
	builder := c.Delete().Where(chatroommember.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatRoomMemberDeleteOne{builder}
}

// Query returns a query builder for ChatRoomMember.
func (c *ChatRoomMemberClient) Query() *ChatRoomMemberQuery {
	return &ChatRoomMemberQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChatRoomMember},
		inters: c.Interceptors(),
	}
}

// Get returns a ChatRoomMember entity by its id.
func (c *ChatRoomMemberClient) Get(ctx context.Context, id int) (*ChatRoomMember, error) {
	return c.Query().Where(chatroommember.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatRoomMemberClient) GetX(ctx context.Context, id int) *ChatRoomMember {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a ChatRoomMember.
func (c *ChatRoomMemberClient) QueryUser(crm *ChatRoomMember) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := crm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroommember.Table, chatroommember.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, chatroommember.UserTable, chatroommember.UserColumn),
		)
		fromV = sqlgraph.Neighbors(crm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoom queries the room edge of a ChatRoomMember.
func (c *ChatRoomMemberClient) QueryRoom(crm *ChatRoomMember) *ChatRoomQuery {
	query := (&ChatRoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := crm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroommember.Table, chatroommember.FieldID, id),
			sqlgraph.To(chatroom.Table, chatroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, chatroommember.RoomTable, chatroommember.RoomColumn),
		)
		fromV = sqlgraph.Neighbors(crm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChatRoomMemberClient) Hooks() []Hook {
	return c.hooks.ChatRoomMember
}

// Interceptors returns the client interceptors.
func (c *ChatRoomMemberClient) Interceptors() []Interceptor {
	return c.inters.ChatRoomMember
}

func (c *ChatRoomMemberClient) mutate(ctx context.Context, m *ChatRoomMemberMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChatRoomMemberCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChatRoomMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChatRoomMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChatRoomMemberDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChatRoomMember mutation op: %q", m.Op())
	}
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `message.Intercept(f(g(h())))`.
func (c *MessageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Message = append(c.inters.Message, interceptors...)
}

// Create returns a builder for creating a Message entity.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MessageClient) MapCreateBulk(slice any, setFunc func(*MessageCreate, int)) *MessageCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MessageCreateBulk{err: fmt.Errorf("calling to MessageClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MessageCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id int) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MessageClient) DeleteOneID(id int) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMessage},
		inters: c.Interceptors(),
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id int) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id int) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoom queries the room edge of a Message.
func (c *MessageClient) QueryRoom(m *Message) *ChatRoomQuery {
	query := (&ChatRoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(chatroom.Table, chatroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.RoomTable, message.RoomColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Message.
func (c *MessageClient) QueryAuthor(m *Message) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.AuthorTable, message.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// Interceptors returns the client interceptors.
func (c *MessageClient) Interceptors() []Interceptor {
	return c.inters.Message
}

func (c *MessageClient) mutate(ctx context.Context, m *MessageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MessageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MessageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Message mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChatRooms queries the chat_rooms edge of a User.
func (c *UserClient) QueryChatRooms(u *User) *ChatRoomQuery {
	query := (&ChatRoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(chatroom.Table, chatroom.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.ChatRoomsTable, user.ChatRoomsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a User.
func (c *UserClient) QueryMessages(u *User) *MessageQuery {
	query := (&MessageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.MessagesTable, user.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParticipatisRooms queries the participatis_rooms edge of a User.
func (c *UserClient) QueryParticipatisRooms(u *User) *ChatRoomQuery {
	query := (&ChatRoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(chatroom.Table, chatroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ParticipatisRoomsTable, user.ParticipatisRoomsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMemberships queries the memberships edge of a User.
func (c *UserClient) QueryMemberships(u *User) *ChatRoomMemberQuery {
	query := (&ChatRoomMemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(chatroommember.Table, chatroommember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.MembershipsTable, user.MembershipsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ChatRoom, ChatRoomMember, Message, User []ent.Hook
	}
	inters struct {
		ChatRoom, ChatRoomMember, Message, User []ent.Interceptor
	}
)
