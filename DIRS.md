# Project Structure

I want to keep domain logic in one place. So `domain/user` includes everything that has to do with dealing with users. That includes a sub-router, pages, handlers and templates.

Each domain package can then create a sub-package if necessary.

```
├── app.go | main func
├── internal | technical abstractions
    ├── db
    └── session
└── domain | domain abstractions
    ├── user
    └── todo
```
