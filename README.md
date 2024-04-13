# Go-based Social Network Web Server

## Table of Contents
- [Overview](#overview)
- [Data Models](#data-models)
- [RESTful API Design](#restful-api-design)
- [High-Level Design](#high-level-design)
- [Handle Methods](#handle-methods)
- [Implementation Considerations](#implementation-considerations)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Config](#config)
- [License](#license)

## Overview
This repository contains the source code for a Go-based Social Network Web Server. 
It focuses on providing a scalable, and secure platform for social networking applications, 
emphasizing users, posts, comments, likes, and user relationships.

## Data Models
### Core Entities:
- **User**: ID, hashed_password, salt, first_name, last_name, email, and user_name.
- **Post**: ID, fk_user_id, content_text, content_image_path, created_at, and visibility.
- **Comment**: ID, fk_post_id, fk_user_id, content, and created_at.
- **Like**: Represents user endorsements of posts, indicating a many-to-many relationship between users and posts.
- **USER_USER**: Facilitates user relationships, follower_id and following_id.

## RESTful API Design
- **Login**: POST v1/sessions {user_name, password} {msg} 
- **Sign up**: POST v1/users {user_name, email, first_name, last_name, birthday, password} {msg} 
- **Edit profile**: PUT v1/users {first_name, last_name, birthday, password} {msg} 
- **See follow list**: GET v1/friends/user_id {} {users} 
- **Follow**: POST v1/friends/user_id {msg} 
- **Unfollow**: DELETE v1/friends/user_id {msg} 
- **See user posts**: GET v1/friends/user_id/posts {posts} 
- **See post**: GET v1/posts/post_id {text, image, comments, likes} 
- **Create post**: POST v1/posts {text, image} {msg} 
- **Edit post**: PUT v1/posts/post_id {text, image} {msg} 
- **Delete post**: DELETE v1/posts/post_id {} {msg} 
- **Comment post**: POST v1/posts/post_id/comments {text} {msg} 
- **Like post**: POST v1/posts/post_id/likes {} {msg} 
- **Newsfeed**: GET v1/newsfeeds {posts} 

## Implement Consider
The system's high-level design outlines interactions between users, web servers, application servers, caches, databases, and a newsfeed generation service. 
- Web Server (Gin): Serves as the entry point for all user actions. It authenticates requests, routes them to the appropriate service, and returns responses to the client.

- Application Servers: Handle the business logic of your application, such as managing posts, users, and feeds. These servers will interact with various caches and the database to fulfill requests.

- Cache (Redis): Used for session management, metadata caching, posts, and media caching, enhancing the performance of frequently accessed data.

- Database (PostgreSQL): Stores user data, posts, comments, likes, and relationships. You might use separate databases for metadata and posts for scalability.

- Newsfeed Generation Service: A specialized service that compiles and organizes newsfeed content based on user actions. This service will utilize caches and the database to generate feeds efficiently.

- Media Storage (AWS S3/Minio): Stores large media files like images and videos. The database stores references (URLs) to these files.

Message Queue (RabbitMQ): Facilitates asynchronous tasks like newsfeed updates. When a user creates a post, a message is sent to a queue, which is processed by a service to update followers' feeds.
## Handle Methods
Methods include session management, caching strategies for authentication and authorization, and password hashing techniques using bcrypt, salt, and pepper for security.

## Implementation Considerations
- Uses Gin for the web server to handle requests and responses efficiently.
- Redis for session caching, username for sign up, list of followers/following users.
- Separates logic into models, repositories, and services for maintainability.
- Utilizes Docker for easy deployment and scaling.

## Project Structure
/social-network/
├── cmd/server/main.go
├── internal/
│ ├── app/
│ ├── model/
│ ├── repo/
│ └── service/
├── pkg/
│ ├── cache/
│ ├── config/
│ └── database/
├── scripts/
├── deployments/
├── .env
├── go.mod
├── go.sum
└── README.md
