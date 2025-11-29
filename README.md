# ESDC Backend Shared Models

[![Go Version](https://img.shields.io/badge/Go-1.25.2-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A shared Go module containing GORM models for the ESDC backend services.

## Installation

```bash
go get github.com/aruncs31s/esdcmodels
```

## Usage

```go
import model "github.com/aruncs31s/esdcmodels"
```

## Table of Contents

- [Models Overview](#models-overview)
- [User Model](#user-model)
- [Project Model](#project-model)
- [Team Model](#team-model)
- [Comments Model](#comments-model)
- [Reviews Model](#reviews-model)
- [Tag Model](#tag-model)
- [Technologies Model](#technologies-model)
- [Submission Model](#submission-model)
- [Notification Model](#notification-model)
- [ChatBot Model](#chatbot-model)
- [Ollama Model](#ollama-model)
- [Order Model](#order-model)
- [Post Model](#post-model)
- [Analytics Models](#analytics-models)
- [Relations Summary](#relations-summary)

---

## Models Overview

| Model | Table Name | Description |
|-------|------------|-------------|
| `User` | `users` | User accounts with profile information |
| `UserDetails` | `user_details` | Extended user profile (social links) |
| `Location` | `locations` | User location information |
| `Github` | `github` | GitHub profile integration |
| `Project` | `projects` | Project listings with metadata |
| `Comments` | `comments` | Project comments |
| `Reviews` | `reviews` | Project reviews and ratings |
| `Teams` | `teams` | Team/organization records |
| `Tag` | `tags` | Project categorization tags |
| `Technologies` | `technologies` | Technology stack items |
| `Submission` | `submissions` | Problem submissions |
| `Notification` | `notifications` | User notifications |
| `ChatBotMessage` | `chat_bot_messages` | Chatbot conversation logs |
| `Ollama` | `ollama` | Ollama LLM query cache |
| `Order` | `orders` | Customer orders |
| `OrderItem` | `order_items` | Order line items |
| `Post` | `posts` | Blog/content posts |
| `ProjectTemplate` | `project_templates` | Reusable project templates |
| `ProjectStats` | `project_stats` | Project statistics |
| `ProjectAnalytics` | `project_analytics` | Project analytics data |
| `TrendingTech` | `trending_techs` | Trending technologies |
| `TrendingTag` | `trending_tags` | Trending tags |

---

## User Model

```go
type User struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    Name          string         `gorm:"column:name;not null" json:"name"`
    Username      string         `gorm:"column:username;unique;not null" json:"username"`
    Email         string         `gorm:"column:email;unique;not null" json:"email"`
    Password      string         `gorm:"column:password;not null" json:"-"`
    Image         *string        `gorm:"column:image" json:"image"`
    Bio           *string        `gorm:"column:bio" json:"bio"`
    Role          string         `gorm:"column:role;not null;default:user" json:"role"`
    Verified      *bool          `gorm:"column:verified;not null;default:false" json:"verified"`
    Status        string         `gorm:"column:status;not null;default:active" json:"status"`
    CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
    Teams         *[]Teams       `gorm:"many2many:team_members;" json:"teams"`
    Github        *Github        `gorm:"foreignKey:UserID;onDelete:cascade" json:"github"`
    Submissions   *[]Submission  `gorm:"foreignKey:UserID;onDelete:cascade" json:"submissions"`
    Details       *UserDetails   `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"details"`
    Projects      *[]Project     `gorm:"many2many:project_contributors;" json:"projects"`
    LikedProjects *[]Project     `gorm:"many2many:project_likes;" json:"liked_projects"`
    Notifications []Notification `gorm:"foreignKey:UserID;onDelete:cascade" json:"notifications"`
    Location      *Location      `gorm:"foreignKey:UserID;onDelete:cascade" json:"location"`
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

type Location struct {
    ID      uint    `gorm:"primaryKey" json:"id"`
    UserID  uint    `gorm:"column:user_id;not null;unique" json:"user_id"`
    City    *string `gorm:"column:city" json:"city"`
    Country *string `gorm:"column:country" json:"country"`
}

type Github struct {
    ID       uint   `gorm:"column:id;primaryKey"`
    UserID   uint   `gorm:"column:user_id"`
    Username string `gorm:"column:username;unique;not null"`
}
```

### User Relations

| Relation | Type | Description | Cascade Delete |
|----------|------|-------------|----------------|
| Teams | many2many | User team memberships | No |
| Github | 1-to-1 | GitHub profile | ✅ Yes |
| Submissions | 1-to-many | Problem submissions | ✅ Yes |
| Details | 1-to-1 | Extended profile details | ✅ Yes |
| Projects | many2many | Contributed projects | No |
| LikedProjects | many2many | Liked projects | No |
| Notifications | 1-to-many | User notifications | ✅ Yes |
| Location | 1-to-1 | User location | ✅ Yes |

---

## Project Model

```go
type Project struct {
    ID               uint            `gorm:"primaryKey"`
    Title            string          `gorm:"column:title"`
    Image            *string         `gorm:"column:image"`
    Description      string          `gorm:"column:description"`
    GithubLink       string          `gorm:"column:github_link"`
    LiveURL          *string         `gorm:"column:live_url"`
    Likes            int             `gorm:"column:likes;default:0"`
    Views            int             `gorm:"column:views;default:0"`
    Category         string          `gorm:"column:category;default:'General'"`
    Cost             int             `gorm:"column:cost;default:0"`
    Status           string          `gorm:"column:status;default:'active'"`
    Visibility       int             `gorm:"column:visibility;type:tinyint;default:1"`
    Version          string          `gorm:"column:version;default:'0.0.0'"`
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
```

**Status Values:** `active`, `inactive`, `archived`

**Visibility:** `1` = public, `0` = private

### Project Relations

| Relation | Type | Description | Cascade Delete |
|----------|------|-------------|----------------|
| Creator | many-to-1 | Project creator | No |
| Contributors | many2many | Project contributors | No |
| Tags | many2many | Associated tags | No |
| Technologies | many2many | Tech stack | No |
| LikedBy | many2many | Users who liked | No |
| ViewedBy | many2many | Users who viewed | No |
| ForkedByProjects | 1-to-many | Forks of this project | ✅ Yes |
| Comments | 1-to-many | Project comments | ✅ Yes |
| Reviews | 1-to-many | Project reviews | ✅ Yes |

---

## Team Model

```go
type Teams struct {
    ID        int64     `gorm:"primaryKey"`
    Name      string    `gorm:"column:name;uniqueIndex"`
    Image     *string   `gorm:"column:image"`
    GithubUrl *string   `gorm:"column:github_url"`
    Projects  []Project `gorm:"many2many:project_teams;" json:"projects"`
}
```

---

## Comments Model

```go
type Comments struct {
    ID        uint      `gorm:"primaryKey"`
    ProjectID uint      `gorm:"column:project_id;not null"`
    UserID    uint      `gorm:"column:user_id;not null"`
    Content   string    `gorm:"column:content;not null"`
    User      *User     `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"user"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
```

---

## Reviews Model

```go
type Reviews struct {
    ID        uint      `gorm:"primaryKey"`
    ProjectID uint      `gorm:"column:project_id;not null"`
    UserID    uint      `gorm:"column:user_id;not null"`
    Rating    int       `gorm:"column:rating;not null"`
    Comment   string    `gorm:"column:comment"`
    User      *User     `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"user"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
```

---

## Tag Model

```go
type Tag struct {
    ID   uint   `gorm:"primaryKey;autoIncrement"`
    Name string `gorm:"column:name;unique;not null"`
}
```

---

## Technologies Model

```go
type Technologies struct {
    ID   int    `gorm:"column:id;primaryKey;autoIncrement"`
    Name string `gorm:"column:name;unique;not null"`
}
```

---

## Submission Model

```go
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
```

---

## Notification Model

```go
type Notification struct {
    ID          uint   `gorm:"primaryKey"`
    UserID      uint   `gorm:"index"`
    Type        string // like, comment, follow, milestone
    Title       string
    Message     string
    ProjectID   uint `gorm:"index"`
    TriggeredBy uint
    IsRead      bool `gorm:"default:false"`
    CreatedAt   time.Time
    User            *User
    Project         *Project
    TriggeredByUser *User `gorm:"foreignKey:TriggeredBy"`
}
```

---

## ChatBot Model

```go
type ChatBotMessage struct {
    ID        int       `gorm:"primaryKey"`
    AskedBy   *uint     `gorm:"column:asked_by"`
    Role      string    `gorm:"column:role"`
    Category  *string   `gorm:"column:category"` // project, product, etc
    Content   string    `gorm:"column:content"`
    Response  *string   `gorm:"column:response"`
    CreatedAt time.Time `gorm:"column:created_at"`
}
```

---

## Ollama Model

```go
type Ollama struct {
    ID        int    `gorm:"primaryKey"`
    ModelName string `gorm:"column:model_name"`
    Prompt    string `gorm:"column:prompt"`
    Response  string `gorm:"column:response"`
    AskedBy   *uint  `gorm:"column:asked_by"`
}
```

Used for caching Ollama LLM responses to avoid redundant queries.

---

## Order Model

```go
type Order struct {
    ID        uint        `gorm:"primaryKey" json:"id"`
    UserID    uint        `json:"user_id"`
    Total     float64     `json:"total"`
    Status    string      `json:"status"` // pending, completed, cancelled
    Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
}

type OrderItem struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}
```

---

## Post Model

```go
type Post struct {
    ID       int    `gorm:"primaryKey"`
    Title    string `gorm:"not null"`
    Content  string `gorm:"not null"`
    AuthorID int    `gorm:"not null"`
}
```

---

## Analytics Models

### ProjectTemplate

```go
type ProjectTemplate struct {
    ID           uint            `gorm:"primaryKey"`
    ProjectID    uint            `gorm:"index"`
    Name         string
    Description  string
    CreatedBy    uint            `gorm:"index"`
    IsPublic     bool            `gorm:"default:false"`
    UsageCount   int             `gorm:"default:0"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    Creator      *User
    Technologies *[]Technologies `gorm:"many2many:template_technologies;"`
    Tags         *[]Tag          `gorm:"many2many:template_tags;"`
}
```

### ProjectStats

```go
type ProjectStats struct {
    ID            uint    `gorm:"primaryKey"`
    ProjectID     uint    `gorm:"uniqueIndex"`
    ViewCount     int     `gorm:"default:0"`
    LikeCount     int     `gorm:"default:0"`
    CommentCount  int     `gorm:"default:0"`
    ReviewCount   int     `gorm:"default:0"`
    AverageRating float64 `gorm:"default:0"`
    UpdatedAt     time.Time
}
```

### ProjectAnalytics

```go
type ProjectAnalytics struct {
    ID            uint `gorm:"primaryKey"`
    ProjectID     uint `gorm:"uniqueIndex"`
    TotalViews    int
    TotalLikes    int
    TotalComments int
    AverageRating float64
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
```

### Trending Models

```go
type TrendingTech struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string `gorm:"uniqueIndex"`
    UsageCount int
}

type TrendingTag struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string `gorm:"uniqueIndex"`
    UsageCount int
}
```

---

## Relations Summary

### Cascade Delete Behavior

**When User is deleted:**
- ✅ Github profile
- ✅ All Submissions
- ✅ User Details
- ✅ Location information
- ✅ All Notifications

**When Project is deleted:**
- ✅ All Comments
- ✅ All Reviews
- ✅ All Forked projects

### Many-to-Many Join Tables

| Join Table | Models |
|------------|--------|
| `team_members` | User ↔ Teams |
| `project_contributors` | User ↔ Project |
| `project_likes` | User ↔ Project |
| `project_views` | User ↔ Project |
| `project_tags` | Project ↔ Tag |
| `project_technologies` | Project ↔ Technologies |
| `project_teams` | Project ↔ Teams |
| `template_technologies` | ProjectTemplate ↔ Technologies |
| `template_tags` | ProjectTemplate ↔ Tag |

### Database Constraints

- **Unique constraints:** `email`, `username`, `tag.name`, `technologies.name`
- **Foreign keys:** Configured with `onDelete:cascade` where appropriate
- **Not-null constraints:** Applied on essential fields
- **Indexes:** Applied on frequently queried foreign keys

---

## License

[MIT](LICENSE)
