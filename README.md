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
- **Login**: `POST v1/sessions`
- **Sign Up**: `POST v1/users`
- **Edit Profile**: `PUT v1/users`
- **See Follow List**: `GET v1/friends/user_id`
- **Follow/Unfollow**: `POST/DELETE v1/friends/user_id`
- **See User Posts**: `GET v1/friends/user_id/posts`
- **See Post**: `GET v1/posts/post_id`
- **Create/Edit/Delete Post**: `POST/PUT/DELETE v1/posts`
- **Comment Post**: `POST v1/posts/post_id/comments`
- **Like Post**: `POST v1/posts/post_id/likes`
- **Newsfeed**: `GET v1/newsfeeds`

## High-Level Design
The system's high-level design outlines interactions between users, web servers, application servers, caches, databases, and a newsfeed generation service. 
It emphasizes scalability, where caching is utilized to enhance performance and user experience.

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
