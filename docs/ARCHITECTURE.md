Architecture
===

> [!NOTE]
This markdown file serves as an overview of forge architecture.  
It documents the main components of the system and the relationships between them.  
It is intended to provide a quick understanding of how the different parts of forge interact without describing implementation details

## Database Table Overview

Here is an overview of the tables used in the database

***User Table***
| Column | Description |
|--------|-------------|
| id | Unique identifier for the user |
| githubId | Unique identifier of the user's GitHub account |
| username | GitHub username of the user |
| mail | Email address of the user |
| createdAt | Date and time when the user was created |

***Project Table***
| Column | Description |
|--------|-------------|
| id | Unique identifier for the project |
| githubId | Identifier of the user who owns the project |
| name  | Name of the project |
| repositoryUrl | URL of the Git repository associated with the project |
| repositoryId  | Unique identifier of the repository |
| branch | Git branch used for deployments |
| createdAt | Date and time when the project was created |
| lastUpdatedAt | Date and time when the project was last updated |

***Deployment Table***
| Column | Description |
|--------|-------------|
| id | Unique identifier for the deployment |
| projectId | Identifier of the project associated with the deployment |
| commitSHA | SHA identifier of the commit being deployed |
| status | Current status of the deployment |
| imageTag | Docker image tag generated for the deployment |
| createdAt | Date and time when the deployment was created |
| startedAt | Date and time when the deployment started |
| finishedAT | Date and time when the deployment finished |


## Project Architecture
