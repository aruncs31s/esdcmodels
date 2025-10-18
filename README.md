# ESDC Backend Shared Models v0.1.0

## User Model
```go
package model

type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"column:name;not null" json:"name"`
	Username    string        `gorm:"column:username;unique;not null" json:"username"`
	Email       string        `gorm:"column:email;unique;not null" json:"email"`
	Password    string        `gorm:"column:password;not null" json:"password"`
	Role        string        `gorm:"column:role;not null;default:user" json:"role"`
	Verified    *bool         `gorm:"column:verified;not null;default:false" json:"verified"`
	Status      string        `gorm:"column:status;not null;default:active" json:"status"`
	CreatedAt   int64         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   int64         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Teams       *[]Teams      `gorm:"many2many:team_members;" json:"teams"`
	Github      *Github       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Submissions *[]Submission `gorm:"foreignKey:UserID"`
	Details     *UserDetails  `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`

	// fixed here:
	Projects      *[]Project     `gorm:"many2many:project_contributors;" json:"projects"`
	LikedProjects *[]Project     `gorm:"many2many:project_likes;" json:"liked_projects"`
	Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications"`
}

type Github struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	UserID   uint   `gorm:"column:user_id"`
	Username string `gorm:"column:username;unique;not null"`
}

func (Github) TableName() string {
	return "github"
}

func (User) TableName() string {
	return "users"
}

```
## Team Model
```go
package model

type Teams struct {
	ID        int64     `gorm:"primaryKey;"`
	Name      string    `gorm:"column:name;uniqueIndex"`
	Image     *string   `gorm:"column:image;"`
	GithubUrl *string   `gorm:"column:github_url;"`
	Projects  []Project `gorm:"many2many:project_teams;" json:"projects"`
}

func (Teams) TableName() string {
	return "teams"
}
```

## Submission Model
```go
package model

import "time"

type Submission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"column:user_id"`
	ProblemID   uint      `gorm:"column:problem_id"`
	Attempts    int       `gorm:"column:attempts"`
	SubmittedAt time.Time `gorm:"column:submitted_at"`
	Status      bool      `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Submission) TableName() string {
	return "submissions"
}
```
## Tags Model
```go
package model

type Tag struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"column:name;unique;not null"`
}

func (Tag) TableName() string {
	return "tags"
}

```

## Technologies Model
```go
package model

type Technologies struct {
	ID   int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;unique;not null"`
}

func (Technologies) TableName() string {
	return "technologies"
}

```
## Project Model
```go
package model

import "time"

type Project struct {
	ID           int             `gorm:"primaryKey"`
	Title        string          `gorm:"column:title;unique"`
	Image        *string         `gorm:"column:image"`
	Description  string          `gorm:"column:description"`
	GithubLink   string          `gorm:"column:github_link"`
	LiveUrl      *string         `gorm:"column:live_url"`
	Likes        int             `gorm:"column:likes;default:0"`
	Category     string          `gorm:"column:category;default:'General'"`
	Cost         int             `gorm:"column:cost;default:0"`
	Status       string          `gorm:"column:status;default:'active'"` // active, inactive, archived
	Visibility   int             `gorm:"column:visibility;default:0"`    // 0: public, 1: private
	CreatedBy    uint            `gorm:"column:created_by;not null"`
	ModifiedBy   *uint           `gorm:"column:modified_by"`
	CreatedAt    *time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	Creator      User            `gorm:"foreignKey:CreatedBy;references:ID" json:"creator"`
	Contributors *[]User         `gorm:"many2many:project_contributors;" json:"contributors"`
	Tags         *[]Tag          `gorm:"many2many:project_tags;" json:"tags"`
	Technologies *[]Technologies `gorm:"many2many:project_technologies;" json:"technologies"`
	LikedBy      []User          `gorm:"many2many:project_likes;" json:"liked_by"`
}

func (Project) TableName() string {
	return "projects"
}

```

## Post Model
```go
package model

type Post struct {
	ID       int   `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	AuthorID int   `gorm:"not null"`
}

func (Post) TableName() string {
	return "posts"
}

```
## Notification Model
```go
package model

type Notification struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"column:user_id;not null" json:"user_id"`
	Title     string `gorm:"column:title;not null" json:"title"`
	Message   string `gorm:"column:message;not null" json:"message"`
	Achieved  bool   `gorm:"column:achieved;default:false" json:"achieved"`
	Read      bool   `gorm:"column:read;default:false" json:"read"`
	ReadAt    *int64 `gorm:"column:read_at" json:"read_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}

```

## Chatbot Model
```go
package model

import "time"

type ChatBotMessage struct {
	ID        int       `gorm:"primaryKey"`
	AskedBy   *uint     `gorm:"column:asked_by"`
	Role      string    `gorm:"column:role"`
	Category  *string   `gorm:"column:category"` // project product etc
	Content   string    `gorm:"column:content"`  // the message.
	Response  *string   `gorm:"column:response"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ChatBotMessage) TableName() string {
	return "chat_bot_messages"
}
```