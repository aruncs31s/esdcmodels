# ESDC Backend Shared Models v0.2.0

## User Model
```go
package model

import (
	"time"
)

type User struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	Name          string        `gorm:"column:name;not null" json:"name"`
	Username      string        `gorm:"column:username;unique;not null" json:"username"`
	Email         string        `gorm:"column:email;unique;not null" json:"email"`
	Password      string        `gorm:"column:password;not null" json:"-"`
	Image         *string       `gorm:"column:image" json:"image"`
	Bio           *string       `gorm:"column:bio" json:"bio"`
	Role          string        `gorm:"column:role;not null;default:user" json:"role"`
	Verified      *bool         `gorm:"column:verified;not null;default:false" json:"verified"`
	Status        string        `gorm:"column:status;not null;default:active" json:"status"`
	CreatedAt     time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Teams         *[]Teams      `gorm:"many2many:team_members;" json:"teams"`
	Github        *Github       `gorm:"foreignKey:UserID;onDelete:cascade" json:"github"`
	Submissions   *[]Submission `gorm:"foreignKey:UserID;onDelete:cascade" json:"submissions"`
	Details       *UserDetails  `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"details"`
	Projects      *[]Project    `gorm:"many2many:project_contributors;" json:"projects"`
	LikedProjects *[]Project    `gorm:"many2many:project_likes;" json:"liked_projects"`
	Notifications []Notification `gorm:"foreignKey:UserID;onDelete:cascade" json:"notifications"`
	Location      *Location     `gorm:"foreignKey:UserID;onDelete:cascade" json:"location"`
}

type UserDetails struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	UserID        uint    `gorm:"column:user_id;not null;unique" json:"user_id"`
	Twitter       *string `gorm:"column:twitter" json:"twitter"`
	Location      *string `gorm:"column:location" json:"location"`
	LinkedIn      *string `gorm:"column:linkedin" json:"linkedin"`
	Facebook      *string `gorm:"column:facebook" json:"facebook"`
	Instagram     *string `gorm:"column:instagram" json:"instagram"`
	StackOverflow *string `gorm:"column:stackoverflow" json:"stackoverflow"`
	Website       *string `gorm:"column:website" json:"website"`
}

func (UserDetails) TableName() string {
	return "user_details"
}

type Location struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	UserID  uint    `gorm:"column:user_id;not null;unique" json:"user_id"`
	City    *string `gorm:"column:city" json:"city"`
	Country *string `gorm:"column:country" json:"country"`
}

func (Location) TableName() string {
	return "locations"
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

### User Relations Documentation

| Relation | Type | Description | Cascade |
|----------|------|-------------|---------|
| **Teams** | many2many | User can be a member of multiple teams | No |
| **Github** | 1-to-1 | User's GitHub profile information | ✅ Yes - When user is deleted, Github record is deleted |
| **Submissions** | 1-to-many | User's problem submissions | ✅ Yes - When user is deleted, all submissions are deleted |
| **Details** | 1-to-1 | Extended user profile details (social links, website) | ✅ Yes - When user is deleted, details are deleted |
| **Projects** | many2many | User as a contributor to projects | No |
| **LikedProjects** | many2many | Projects liked by the user | No |
| **Notifications** | 1-to-many | User's notifications | ✅ Yes - When user is deleted, all notifications are deleted |
| **Location** | 1-to-1 | User's city and country information | ✅ Yes - When user is deleted, location is deleted |


## Team Model
```go
package model

type Teams struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:name;uniqueIndex"`
	Image     *string   `gorm:"column:image"`
	GithubUrl *string   `gorm:"column:github_url"`
	Projects  []Project `gorm:"many2many:project_teams;" json:"projects"`
}

func (Teams) TableName() string {
	return "teams"
}
```

### Team Relations Documentation

| Relation | Type | Description | Cascade |
|----------|------|-------------|---------|
| **Projects** | many2many | Projects associated with the team | No |



## Submission Model
```go
package model

import "time"

type Submission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"column:user_id;not null"`
	ProblemID   uint      `gorm:"column:problem_id;not null"`
	Attempts    int       `gorm:"column:attempts"`
	SubmittedAt time.Time `gorm:"column:submitted_at"`
	Status      bool      `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Submission) TableName() string {
	return "submissions"
}
```

### Submission Relations Documentation

| Relation | Type | Description | Cascade |
|----------|------|-------------|---------|
| **UserID** | foreign key | Reference to the user who submitted | ✅ Yes - When user is deleted, all submissions are deleted |


## Tag Model
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

### Tag Relations Documentation

| Relation | Type | Description |
|----------|------|-------------|
| **Projects** | many2many | Projects tagged with this tag (via project_tags join table) |



## Technologies Model
```go
package model

type Technologies struct {
	ID   uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;unique;not null"`
}

func (Technologies) TableName() string {
	return "technologies"
}
```

### Technologies Relations Documentation

| Relation | Type | Description |
|----------|------|-------------|
| **Projects** | many2many | Projects using this technology (via project_technologies join table) |


## Project Model
```go
package model

import "time"

type Project struct {
	ID               uint            `gorm:"primaryKey"`
	Title            string          `gorm:"column:title;unique"`
	Image            *string         `gorm:"column:image"`
	Description      string          `gorm:"column:description"`
	GithubLink       string          `gorm:"column:github_link"`
	LiveURL          *string         `gorm:"column:live_url"`
	Likes            int             `gorm:"column:likes;default:0"`
	Views            int             `gorm:"column:views;default:0"`
	Category         string          `gorm:"column:category;default:'General'"`
	Cost             int             `gorm:"column:cost;default:0"`
	Status           string          `gorm:"column:status;default:'active'"` // active, inactive, archived
	Visibility       int             `gorm:"column:visibility;default:0"`    // 0: public, 1: private
	CreatedBy        uint            `gorm:"column:created_by;not null"`
	ModifiedBy       *uint           `gorm:"column:modified_by"`
	CreatedAt        time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;autoUpdateTime"`
	Creator          User            `gorm:"foreignKey:CreatedBy;references:ID" json:"creator"`
	Contributors     *[]User         `gorm:"many2many:project_contributors;" json:"contributors"`
	Tags             *[]Tag          `gorm:"many2many:project_tags;" json:"tags"`
	Technologies     *[]Technologies `gorm:"many2many:project_technologies;" json:"technologies"`
	LikedBy          []User          `gorm:"many2many:project_likes;" json:"liked_by"`
	ViewedBy         []User          `gorm:"many2many:project_views;" json:"viewed_by"`
	ForkedFrom       *uint           `gorm:"column:forked_from"`
	ForkedByProjects *[]Project      `gorm:"foreignKey:ForkedFrom;references:ID;onDelete:cascade"`
	Comments         []Comments      `gorm:"foreignKey:ProjectID;references:ID;onDelete:cascade" json:"comments"`
	Reviews          []Reviews       `gorm:"foreignKey:ProjectID;references:ID;onDelete:cascade" json:"reviews"`
}

type Comments struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"column:project_id;not null"`
	UserID    uint      `gorm:"column:user_id;not null"`
	Content   string    `gorm:"column:content;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type Reviews struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"column:project_id;not null"`
	UserID    uint      `gorm:"column:user_id;not null"`
	Rating    int       `gorm:"column:rating;not null"`
	Comment   string    `gorm:"column:comment"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Project) TableName() string {
	return "projects"
}

func (Comments) TableName() string {
	return "comments"
}

func (Reviews) TableName() string {
	return "reviews"
}
```

### Project Relations Documentation

| Relation | Type | Description | Cascade |
|----------|------|-------------|---------|
| **Creator** | many-to-1 | User who created the project | No |
| **Contributors** | many2many | Users contributing to the project | No |
| **Tags** | many2many | Tags associated with the project | No |
| **Technologies** | many2many | Technologies used in the project | No |
| **LikedBy** | many2many | Users who liked the project | No |
| **ViewedBy** | many2many | Users who viewed the project | No |
| **ForkedFrom** | self-referencing | Original project if this is a fork | ✅ Yes - Forked projects deleted when original is deleted |
| **ForkedByProjects** | self-referencing | Projects forked from this project | ✅ Yes - When project is deleted, all forks are deleted |
| **Comments** | 1-to-many | Comments on the project | ✅ Yes - When project is deleted, all comments are deleted |
| **Reviews** | 1-to-many | Reviews/ratings on the project | ✅ Yes - When project is deleted, all reviews are deleted |



## Overall Relations Summary

### Cascade Delete Behavior
When the following records are deleted, related records are automatically deleted:

**User Deletion Cascade:**
- ✅ Github profile
- ✅ All Submissions
- ✅ User Details
- ✅ Location information
- ✅ All Notifications

**Project Deletion Cascade:**
- ✅ All Comments on the project
- ✅ All Reviews on the project
- ✅ All Forked projects (projects that are forks of this project)

### Key Design Principles

1. **User Identity**: All user-related data (Github, Details, Location) is automatically cleaned up when a user is deleted
2. **Project Integrity**: Comments and reviews are dependent on projects and cascade delete for data consistency
3. **Fork Tracking**: Self-referencing foreign key allows tracking project forks with proper cleanup
4. **Many-to-Many Flexibility**: Contributor, like, view relationships use join tables for flexibility without cascade (data preservation)
5. **Audit Trail**: Timestamps on all models track creation and modification times

### Database Constraints
- All foreign keys are configured with `onDelete:cascade` where appropriate
- Unique constraints on email, username, and project titles ensure data uniqueness
- Not-null constraints on essential fields ensure data completeness
## Notification Model
```go
package model

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	Title     string    `gorm:"column:title;not null" json:"title"`
	Message   string    `gorm:"column:message;not null" json:"message"`
	Achieved  bool      `gorm:"column:achieved;default:false" json:"achieved"`
	Read      bool      `gorm:"column:read;default:false" json:"read"`
	ReadAt    *time.Time `gorm:"column:read_at" json:"read_at"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
```

### Notification Relations Documentation

| Relation | Type | Description | Cascade |
|----------|------|-------------|---------|
| **UserID** | foreign key | Reference to the user receiving the notification | ✅ Yes - When user is deleted, all notifications are deleted |



## ChatBot Model
```go
package model

import "time"

type ChatBotMessage struct {
	ID        uint      `gorm:"primaryKey"`
	AskedBy   *uint     `gorm:"column:asked_by"`
	Role      string    `gorm:"column:role"`
	Category  *string   `gorm:"column:category"` // project, product, etc
	Content   string    `gorm:"column:content"`  // the message
	Response  *string   `gorm:"column:response"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ChatBotMessage) TableName() string {
	return "chat_bot_messages"
}
```

### ChatBot Relations Documentation

| Relation | Type | Description |
|----------|------|-------------|
| **AskedBy** | foreign key (optional) | Reference to the user who asked the question. Can be NULL for anonymous queries |

